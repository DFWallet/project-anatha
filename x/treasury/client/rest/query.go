package rest

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/DFWallet/anatha/client/context"
	"github.com/DFWallet/anatha/types/rest"
	"github.com/DFWallet/project-anatha/x/treasury/internal/types"
)

func registerQueryRoutes(cliCtx context.CLIContext, r *mux.Router) {
	// TODO: Define your GET REST endpoints
	r.HandleFunc(
		"/treasury/parameters",
		queryParamsHandlerFn(cliCtx),
	).Methods("GET")
}

func queryParamsHandlerFn(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cliCtx, ok := rest.ParseQueryHeightOrReturnBadRequest(w, cliCtx, r)
		if !ok {
			return
		}

		route := fmt.Sprintf("custom/%s/parameters", types.QuerierRoute)

		res, height, err := cliCtx.QueryWithData(route, nil)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		cliCtx = cliCtx.WithHeight(height)
		rest.PostProcessResponse(w, cliCtx, res)
	}
}
