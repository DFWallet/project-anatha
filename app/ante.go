package app

import (
	sdk "github.com/DFWallet/anatha/types"
	"github.com/DFWallet/anatha/x/auth"
	"github.com/DFWallet/anatha/x/auth/ante"
	"github.com/DFWallet/project-anatha/x/distribution"
	"github.com/DFWallet/project-anatha/x/fee"
)

func (app *AnathaApp) NewAnteHandler() sdk.AnteHandler {
	return sdk.ChainAnteDecorators(
		ante.NewSetUpContextDecorator(), // outermost AnteDecorator. SetUpContext must be called first
		ante.NewMempoolFeeDecorator(),
		ante.NewValidateBasicDecorator(),
		ante.NewValidateMemoDecorator(app.accountKeeper),
		ante.NewConsumeGasForTxSizeDecorator(app.accountKeeper),
		ante.NewSetPubKeyDecorator(app.accountKeeper), // SetPubKeyDecorator must be called before all signature verification decorators
		ante.NewValidateSigCountDecorator(app.accountKeeper),
		ante.NewDeductFeeDecorator(app.accountKeeper, app.supplyKeeper),
		ante.NewSigGasConsumeDecorator(app.accountKeeper, auth.DefaultSigVerificationGasConsumer),
		ante.NewSigVerificationDecorator(app.accountKeeper),
		fee.NewFeeDecorator(app.feeKeeper, app.bankKeeper, app.hraKeeper, app.supplyKeeper, distribution.AmcModuleName),
		ante.NewIncrementSequenceDecorator(app.accountKeeper), // innermost AnteDecorator
	)
}
