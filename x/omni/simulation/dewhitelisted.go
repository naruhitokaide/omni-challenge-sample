package simulation

import (
	"math/rand"

	"omni/x/omni/keeper"
	"omni/x/omni/types"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
)

func SimulateMsgDewhitelisted(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgDewhitelisted{
			Authority: simAccount.Address.String(),
		}

		// TODO: Handling the Dewhitelisted simulation

		return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "Dewhitelisted simulation not implemented"), nil, nil
	}
}
