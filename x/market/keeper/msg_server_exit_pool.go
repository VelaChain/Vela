package keeper

import (
	"context"

	"github.com/VelaChain/vela/x/market/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) validateExitPoolMsg(ctx sdk.Context, msg *types.MsgExitPool) (accAddr sdk.AccAddress, err error) {
	if accAddr, err := sdk.AccAddressFromBech32(msg.Creator); err != nil {
		return accAddr, types.ErrAccAddressFromMsg
	}
	return accAddr, nil
}

func (k msgServer) ExitPool(goCtx context.Context, msg *types.MsgExitPool) (*types.MsgExitPoolResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// validate message
	accAddr, err := k.validateExitPoolMsg(ctx, msg)
	if err != nil {
		return &types.MsgExitPoolResponse{}, err
	}
	// check if pool exists
	pool, found := k.Keeper.GetPool(ctx, msg.DenomA, msg.DenomB)
	if !found {
		return &types.MsgExitPoolResponse{}, types.ErrPoolDNE
	}
	// check if liq prov exists
	poolName := types.NewPoolName(msg.DenomA, msg.DenomB)
	liqProv, found := k.Keeper.GetLiqProv(ctx, poolName, msg.Creator)
	if !found {
		return &types.MsgExitPoolResponse{}, types.ErrLiqProvDNE
	}
	// get liq prov shares as sdk int
	provShares, ok := sdk.NewIntFromString(liqProv.ShareAmount)
	if !ok {
		return &types.MsgExitPoolResponse{}, types.ErrConvertSharesToInt
	}
	// get pool amounts as sdk int
	poolAmountA, ok := sdk.NewIntFromString(pool.AmountA)
	if !ok {
		return &types.MsgExitPoolResponse{}, types.ErrConvertAmountAToInt
	}
	poolAmountB, ok := sdk.NewIntFromString(pool.AmountB)
	if !ok {
		return &types.MsgExitPoolResponse{}, types.ErrConvertAmountBToInt
	}
	poolShares, ok := sdk.NewIntFromString(pool.Shares)
	if !ok {
		return &types.MsgExitPoolResponse{}, types.ErrConvertSharesToInt
	}
	// // use provShares - feeAmount for amounts out
	// remShares, err := types.ApplyExitFee(provShares)
	// if err != nil {
	// 	return &types.MsgExitPoolResponse{}, err
	// }
	// get fees for exit
	fees, found := k.Keeper.GetFeeMap(ctx, poolName) 
	if !found {
		return &types.MsgExitPoolResponse{}, types.ErrFeeMapDNE
	}
	// get exit fee as sdk dec
	exitFee, err := sdk.NewDecFromStr(fees.Exit)
	if err != nil {
		return &types.MsgExitPoolResponse{}, err 
	}
	// apply exit fee to provShares
	remShares, err := types.ApplyFee(provShares, exitFee)
	if err != nil {
		return &types.MsgExitPoolResponse{}, err
	}
	// get asset out amounts
	amountOutA := types.AOutGivenShares(poolAmountA, poolShares, remShares)
	if !amountOutA.IsPositive() {
		return &types.MsgExitPoolResponse{}, types.ErrAmountOutANotPos
	}
	amountOutB := types.BOutGivenShares(poolAmountB, poolShares, remShares)
	if !amountOutB.IsPositive() {
		return &types.MsgExitPoolResponse{}, types.ErrAmountOutBNotPos
	}
	// get creator acc address
	accAddr, err = sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return &types.MsgExitPoolResponse{}, err
	}
	// create coins to send
	coinA := sdk.NewCoin(msg.DenomA, amountOutA)
	coinB := sdk.NewCoin(msg.DenomB, amountOutB)
	// send coins from account to module
	err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, accAddr, sdk.NewCoins(coinA, coinB))
	if err != nil {
		return &types.MsgExitPoolResponse{}, err
	}
	// update pool balances
	newAmountA := poolAmountA.Sub(amountOutA)
	newAmountB := poolAmountB.Sub(amountOutB)
	newShares := poolShares.Sub(provShares)
	// remove liq prov
	k.Keeper.RemoveLiqProv(ctx, poolName, msg.Creator)
	// delete pool if empty
	if newAmountA.IsZero() && newAmountB.IsZero() && newShares.IsZero() {
		k.Keeper.RemovePool(ctx, pool.DenomA, pool.DenomB)
	}
	// get new pool
	newPool := types.NewPool(newAmountA.String(), msg.DenomA, newAmountB.String(), msg.DenomB, newShares.String())
	// set pool
	k.Keeper.SetPool(ctx, newPool)

	return &types.MsgExitPoolResponse{}, nil
}
