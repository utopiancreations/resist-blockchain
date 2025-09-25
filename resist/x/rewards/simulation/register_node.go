package simulation

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/client"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"resist/x/rewards/keeper"
	"resist/x/rewards/types"
)

func SimulateMsgRegisterNode(
	ak types.AuthKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
	txGen client.TxConfig,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)

		nodeId := simtypes.RandStringOfLength(r, 10)

		// Check if node already exists
		_, err := k.Nodes.Get(ctx, nodeId)
		if err == nil {
			return simtypes.NoOpMsg(types.ModuleName, sdk.MsgTypeURL(&types.MsgRegisterNode{}), "node already exists"), nil, nil
		}

		nodeTypes := []string{"validator", "full", "light", "archive"}
		nodeType := nodeTypes[r.Intn(len(nodeTypes))]

		stakeAmount := r.Uint64()

		msg := &types.MsgRegisterNode{
			Creator:     simAccount.Address.String(),
			NodeId:      nodeId,
			NodeType:    nodeType,
			StakeAmount: stakeAmount,
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
