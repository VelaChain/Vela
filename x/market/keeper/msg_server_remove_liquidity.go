package keeper

import (
	"context"

	"github.com/VelaChain/vela/x/market/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) validateRemoveLiquidityMsg(ctx sdk.Context, msg *types.MsgRemoveLiquidity) (msgShares sdk.Int, accAddr sdk.AccAddress, err error) {
	if msgShares, ok := sdk.NewIntFromString(msg.Shares); !ok {
		return msgShares, accAddr, types.ErrConvertSharesToInt
	}

	if msg.DenomA == msg.DenomB {
		return msgShares, accAddr, types.ErrDenomsSame
	}

	if accAddr, err := sdk.AccAddressFromBech32(msg.Creator); err != nil {
		return msgShares, accAddr, types.ErrAccAddressFromMsg
	}
	return msgShares, accAddr, nil
}

func (k msgServer) RemoveLiquidity(goCtx context.Context, msg *types.MsgRemoveLiquidity) (*types.MsgRemoveLiquidityResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// validate message
	msgShares, accAddr, err := k.validateRemoveLiquidityMsg(ctx, msg)
	if err != nil {
		return &types.MsgRemoveLiquidityResponse{}, err
	}
	// check if pool exists
	pool, found := k.Keeper.GetPool(ctx, msg.DenomA, msg.DenomB)
	if !found {
		return &types.MsgRemoveLiquidityResponse{}, types.ErrPoolDNE
	}
	// check if liq prov exists
	poolName := types.NewPoolName(msg.DenomA, msg.DenomB)
	liqProv, found := k.Keeper.GetLiqProv(ctx, poolName, msg.Creator)
	if !found {
		return &types.MsgRemoveLiquidityResponse{}, types.ErrLiqProvDNE
	}
	// get msg shares as sdk int
	// msgShares, ok := sdk.NewIntFromString(msg.Shares)
	// if !ok {
	// 	return &types.MsgRemoveLiquidityResponse{}, types.ErrConvertSharesToInt
	// }
	// get liq prov shares as sdk int
	provShares, ok := sdk.NewIntFromString(liqProv.ShareAmount)
	if !ok {
		return &types.MsgRemoveLiquidityResponse{}, types.ErrConvertSharesToInt
	}
	// get pool amounts as sdk int
	poolAmountA, ok := sdk.NewIntFromString(pool.AmountA)
	if !ok {
		return &types.MsgRemoveLiquidityResponse{}, types.ErrConvertAmountAToInt
	}
	poolAmountB, ok := sdk.NewIntFromString(pool.AmountB)
	if !ok {
		return &types.MsgRemoveLiquidityResponse{}, types.ErrConvertAmountBToInt
	}
	poolShares, ok := sdk.NewIntFromString(pool.Shares)
	if !ok {
		return &types.MsgRemoveLiquidityResponse{}, types.ErrConvertSharesToInt
	}
	// get assets out amounts
	amountOutA := msgShares.Mul(poolAmountA).Quo(poolShares)
	if amountOutA.IsZero() {
		return &types.MsgRemoveLiquidityResponse{}, types.ErrAmountOutAZero
	}
	amountOutB := msgShares.Mul(poolAmountB).Quo(poolShares)
	if amountOutB.IsZero() {
		return &types.MsgRemoveLiquidityResponse{}, types.ErrAmountOutBZero
	}
	// check liq prov still has shares
	newProvShares := provShares.Sub(msgShares)
	if newProvShares.IsZero() {
		return &types.MsgRemoveLiquidityResponse{}, types.ErrProviderEmptyRemove
	}
	// get creator acc address
	// accAddr, err := sdk.AccAddressFromBech32(msg.Creator)
	// if err != nil {
	// 	return &types.MsgRemoveLiquidityResponse{}, err
	// }
	// create coins to send
	coinA := sdk.NewCoin(msg.DenomA, amountOutA)
	coinB := sdk.NewCoin(msg.DenomB, amountOutB)
	// send coins from module to account
	err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, accAddr, sdk.NewCoins(coinA, coinB))
	if err != nil {
		return &types.MsgRemoveLiquidityResponse{}, err
	}
	// update pool balances
	newAmountA := poolAmountA.Sub(amountOutA)
	newAmountB := poolAmountB.Sub(amountOutB)
	newShares := poolShares.Sub(msgShares)
	// get new pool and liq prov
	newPool := types.NewPool(newAmountA.String(), msg.DenomA, newAmountB.String(), msg.DenomB, newShares.String())
	newLiqProv := types.NewLiqProv(newProvShares.String(), poolName, msg.Creator)
	// set pool and liq prov
	k.Keeper.SetPool(ctx, newPool)
	k.Keeper.SetLiqProv(ctx, newLiqProv)

	return &types.MsgRemoveLiquidityResponse{}, nil
}
