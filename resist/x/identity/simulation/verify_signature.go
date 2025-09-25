package simulation

import (
	"encoding/hex"
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/client"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"resist/x/identity/keeper"
	"resist/x/identity/types"
)

func SimulateMsgVerifySignature(
	ak types.AuthKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
	txGen client.TxConfig,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)

		// Try to find a challenge for this account
		challengeString, err := k.GetChallenge(ctx, simAccount.Address.String())
		if err != nil || challengeString == "" {
			return simtypes.NoOpMsg(types.ModuleName, sdk.MsgTypeURL(&types.MsgVerifySignature{}), "no challenge found for account"), nil, nil
		}

		// Sign the challenge
		sigBytes, err := simAccount.PrivKey.Sign([]byte(challengeString))
		if err != nil {
			return simtypes.NoOpMsg(types.ModuleName, sdk.MsgTypeURL(&types.MsgVerifySignature{}), "failed to sign challenge"), nil, err
		}
		signatureString := hex.EncodeToString(sigBytes)

		msg := &types.MsgVerifySignature{
			Creator:   simAccount.Address.String(),
			Challenge: challengeString,
			Signature: signatureString,
			Address:   simAccount.Address.String(),
		}

		txCtx := simulation.OperationInput{
			R:               r,
			App:             app,
			TxGen:           txGen,
			Cdc:             nil,
			Msg:             msg,
			Context:         ctx,
			SimAccount:      simAccount,
			ModuleName:      types.ModuleName,
			CoinsSpentInMsg: sdk.NewCoins(),
			AccountKeeper:   ak,
			Bankkeeper:      bk,
		}
		return simulation.GenAndDeliverTxWithRandFees(txCtx)
	}
}
