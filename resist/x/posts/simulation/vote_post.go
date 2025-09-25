package simulation

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/client"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"resist/x/posts/keeper"
	"resist/x/posts/types"
)

func SimulateMsgVotePost(
	ak types.AuthKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
	txGen client.TxConfig,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)

		var allPosts []types.SocialPost
		err := k.SocialPost.Walk(ctx, nil, func(key string, value types.SocialPost) (stop bool, err error) {
			allPosts = append(allPosts, value)
			return false, nil
		})
		if err != nil {
			return simtypes.NoOpMsg(types.ModuleName, sdk.MsgTypeURL(&types.MsgVotePost{}), "unable to get posts"), nil, err
		}
		if len(allPosts) == 0 {
			return simtypes.NoOpMsg(types.ModuleName, sdk.MsgTypeURL(&types.MsgVotePost{}), "no posts found"), nil, nil
		}

		post := allPosts[r.Intn(len(allPosts))]

		voteType := "upvote"
		if r.Intn(2) == 0 {
			voteType = "downvote"
		}

		msg := &types.MsgVotePost{
			Creator:   simAccount.Address.String(),
			PostIndex: post.Index,
			VoteType:  voteType,
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
