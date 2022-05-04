package keeper

import (
	"context"

	"github.com/VelaChain/vela/x/market/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	//	sdkerror "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) validateJoinPoolMsg(ctx sdk.Context, msg *types.MsgJoinPool) (msgAmountA sdk.Int, msgAmountB sdk.Int, msgShares sdk.Int, accAddr sdk.AccAddress, err error) {
	// check ints from strings ok
	msgAmountA, ok := sdk.NewIntFromString(msg.AmountA)
	if !ok {
		return msgAmountA, msgAmountB, msgShares, accAddr, types.ErrConvertAmountAToInt
	}
	msgAmountB, ok = sdk.NewIntFromString(msg.AmountB)
	if !ok {
		return msgAmountA, msgAmountB, msgShares, accAddr, types.ErrConvertAmountBToInt
	}
	msgShares, ok = sdk.NewIntFromString(msg.MinShares)
	if !ok {
		return msgAmountA, msgAmountB, msgShares, accAddr, types.ErrConvertSharesToInt
	}
	// check different denoms
	if msg.DenomA == msg.DenomB {
		return msgAmountA, msgAmountB, msgShares, accAddr, types.ErrDenomsSame
	}
	// get creator acc address
	accAddr, err = sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return msgAmountA, msgAmountB, msgShares, accAddr, err
	}
	// check account balances
	if !k.bankKeeper.HasBalance(ctx, accAddr, sdk.NewCoin(msg.DenomA, msgAmountA)) {
		return msgAmountA, msgAmountB, msgShares, accAddr, types.ErrInsufficientBalanceA
	}
	if !k.bankKeeper.HasBalance(ctx, accAddr, sdk.NewCoin(msg.DenomB, msgAmountB)) {
		return msgAmountA, msgAmountB, msgShares, accAddr, types.ErrInsufficientBalanceB
	}

	return msgAmountA, msgAmountB, msgShares, accAddr, nil
}

// JoinPool creates a new liquidity provider for a pool
// errors if msg creator already providing liquidity to pool
// use AddLiquidity to add liquidity to existing provider
func (k msgServer) JoinPool(goCtx context.Context, msg *types.MsgJoinPool) (*types.MsgJoinPoolResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// validate params and balances are available
	msgAmountA, msgAmountB, msgShares, accAddr, err := k.validateJoinPoolMsg(ctx, msg)
	if err != nil {
		return &types.MsgJoinPoolResponse{}, err
	}
	// check if pool exists
	pool, found := k.Keeper.GetPool(ctx, msg.DenomA, msg.DenomB)
	if !found {
		return &types.MsgJoinPoolResponse{}, types.ErrPoolDNE
	}
	// check if liq prov exists
	poolName := types.NewPoolName(msg.DenomA, msg.DenomB)
	_, found = k.Keeper.GetLiqProv(ctx, poolName, msg.Creator)
	if found {
		return &types.MsgJoinPoolResponse{}, types.ErrLiqProvExists
	}
	// convert strings to skd.Int
	// msgAmountA, ok := sdk.NewIntFromString(msg.AmountA)
	// if !ok {
	// 	return &types.MsgJoinPoolResponse{}, types.ErrConvertAmountAToInt
	// }
	// msgAmountB, ok := sdk.NewIntFromString(msg.AmountB)
	// if !ok {
	// 	return &types.MsgJoinPoolResponse{}, types.ErrConvertAmountBToInt
	// }
	poolAmountA, ok := sdk.NewIntFromString(pool.AmountA)
	if !ok {
		return &types.MsgJoinPoolResponse{}, types.ErrConvertAmountAToInt
	}
	poolAmountB, ok := sdk.NewIntFromString(pool.AmountB)
	if !ok {
		return &types.MsgJoinPoolResponse{}, types.ErrConvertAmountBToInt
	}
	poolShares, ok := sdk.NewIntFromString(pool.Shares)
	if !ok {
		return &types.MsgJoinPoolResponse{}, types.ErrConvertSharesToInt
	}
	// msgShares, ok := sdk.NewIntFromString(msg.MinShares)
	// if !ok {
	// 	return &types.MsgJoinPoolResponse{}, types.ErrConvertSharesToInt
	// }
	// check ratios
	// pool ratio rounds up to nearest int
	// poolRatioAtoB := sdk.NewDecFromInt(poolAmountA).QuoRoundUp(sdk.NewDecFromInt(poolAmountB))
	// poolRatioBtoA := sdk.NewDecFromInt(poolAmountB).QuoRoundUp(sdk.NewDecFromInt(poolAmountA))
	// // get amount B needed for input A
	// needB := poolRatioAtoB.MulInt(msgAmountA).RoundInt()
	// // get amount A needed for input B
	// needA := poolRatioBtoA.MulInt(msgAmountB).RoundInt()
	// // check amount b correct for amount a
	// if !needB.Equal(msgAmountB) {
	// 	return &types.MsgJoinPoolResponse{}, sdkerror.Wrapf(types.ErrInvalidRatio, "For %s alpha you must add %s beta", msgAmountA.String(), needB.String())
	// }
	// // check amount a correct for amount b
	// if !needA.Equal(msgAmountA) {
	// 	return &types.MsgJoinPoolResponse{}, sdkerror.Wrapf(types.ErrInvalidRatio, "For %s beta you must add %s alpha", msgAmountB.String(), needA.String())
	// }
	// msg ratio truncated
	// msgRatio := sdk.NewDecFromInt(msgAmountA).Quo(sdk.NewDecFromInt(msgAmountB))
	// if !poolRatio.Equal(msgRatio) {
	// 	return &types.MsgJoinPoolResponse{}, sdkerror.Wrapf(types.ErrInvalidRatio, "Pool Ratio %s must equal msg ratio %s, add %s beta for %s alpha", poolRatio.String(), msgRatio.String(), poolRatio.MulInt(msgAmountA).RoundInt().String(), msgAmountA.String())
	// }

	// if !poolAmountA.Quo(poolAmountB).Equal(msgAmountA.Quo(msgAmountB)) {
	// 	amtB := poolAmountB.Mul(msgAmountA).Quo(poolAmountA)
	// 	return &types.MsgJoinPoolResponse{}, sdkerror.Wrapf(types.ErrInvalidRatio, "For %s alpha, add %s beta", msgAmountA.String(), amtB.String())
	// }

	// if !poolAmountA.Mul(msgAmountB).RoundInt64().Equal(poolAmountB.Mul(msgAmountA).RoundInt64()) {
	// 	amtB := poolAmountB.Mul(msgAmountA).Quo(poolAmountA)
	// 	return &types.MsgJoinPoolResponse{}, sdkerror.Wrapf(types.ErrInvalidRatio, "For %s alpha, add %s beta", msgAmountA.String(), amtB.String())
	// }

	// check pool ratios
	if err := types.CheckRatios(poolAmountA, poolAmountB, msgAmountA, msgAmountB); err != nil {
		return &types.MsgJoinPoolResponse{}, err
	}
	// get new shares out
	var newShares sdk.Int
	if poolShares.IsZero() {
		newShares = msgShares
	} else {
		newShares = msgAmountA.Mul(poolShares).Quo(poolAmountA)
	}
	// check shares are sufficient
	if newShares.LT(msgShares) {
		return &types.MsgJoinPoolResponse{}, types.ErrNotEnoughSharesOut
	}
	// get creator acc address
	// accAddr, err := sdk.AccAddressFromBech32(msg.Creator)
	// if err != nil {
	// 	return &types.MsgJoinPoolResponse{}, types.ErrAccAddressFromMsg
	// }
	// create coins to send
	coinA := sdk.NewCoin(msg.DenomA, msgAmountA)
	coinB := sdk.NewCoin(msg.DenomB, msgAmountB)
	// send coins from account to module
	err = k.bankKeeper.SendCoinsFromAccountToModule(ctx, accAddr, types.ModuleName, sdk.NewCoins(coinA, coinB))
	if err != nil {
		return &types.MsgJoinPoolResponse{}, types.ErrSendCoinToModule
	}
	// get new pool balances
	newAmountA := poolAmountA.Add(msgAmountA)
	newAmountB := poolAmountB.Add(msgAmountB)
	newPoolShares := newShares.Add(poolShares)
	// get new pool and provider
	newPool := types.NewPool(newAmountA.String(), msg.DenomA, newAmountB.String(), msg.DenomB, newPoolShares.String())
	newLiqProv := types.NewLiqProv(newShares.String(), poolName, msg.Creator)
	// update pool and provider
	k.Keeper.SetPool(ctx, newPool)
	k.Keeper.SetLiqProv(ctx, newLiqProv)

	return &types.MsgJoinPoolResponse{}, nil
}
