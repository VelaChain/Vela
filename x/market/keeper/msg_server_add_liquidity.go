package keeper

import (
	"context"

	"github.com/VelaChain/vela/x/market/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) validateAddLiquidityMsg(ctx sdk.Context, msg *types.MsgAddLiquidity) error {
	// check ints from strings ok
	msgAmountA, ok := sdk.NewIntFromString(msg.AmountA)
	if !ok {
		return types.ErrConvertAmountAToInt
	}
	msgAmountB, ok := sdk.NewIntFromString(msg.AmountB)
	if !ok {
		return types.ErrConvertAmountBToInt
	}
	_, ok = sdk.NewIntFromString(msg.MinShares)
	if !ok {
		return types.ErrConvertSharesToInt
	}
	// check different denoms
	if msg.DenomA == msg.DenomB {
		return types.ErrDenomsSame
	}
	// get creator acc address
	accAddr, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return err
	}
	// check account balances
	if !k.bankKeeper.HasBalance(ctx, accAddr, sdk.NewCoin(msg.DenomA, msgAmountA)) {
		return types.ErrInsufficientBalanceA
	}
	if !k.bankKeeper.HasBalance(ctx, accAddr, sdk.NewCoin(msg.DenomB, msgAmountB)) {
		return types.ErrInsufficientBalanceB
	}

	return nil
}
func (k msgServer) AddLiquidity(goCtx context.Context, msg *types.MsgAddLiquidity) (*types.MsgAddLiquidityResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	// validate params and balances are available
	if err := k.validateAddLiquidityMsg(ctx, msg); err != nil {
		return &types.MsgAddLiquidityResponse{}, err
	}
	// check if pool exists
	pool, found := k.Keeper.GetPool(ctx, msg.DenomA, msg.DenomB)
	if !found {
		return &types.MsgAddLiquidityResponse{}, types.ErrPoolDNE
	}
	// check if liq prov exists
	poolName := types.NewPoolName(msg.DenomA, msg.DenomB)
	liqProv, found := k.Keeper.GetLiqProv(ctx, poolName, msg.Creator)
	if !found {
		return &types.MsgAddLiquidityResponse{}, types.ErrLiqProvDNE
	}
	// convert strings to skd.Int
	msgAmountA, ok := sdk.NewIntFromString(msg.AmountA)
	if !ok {
		return &types.MsgAddLiquidityResponse{}, types.ErrConvertAmountAToInt
	}
	msgAmountB, ok := sdk.NewIntFromString(msg.AmountB)
	if !ok {
		return &types.MsgAddLiquidityResponse{}, types.ErrConvertAmountBToInt
	}
	poolAmountA, ok := sdk.NewIntFromString(pool.AmountA)
	if !ok {
		return &types.MsgAddLiquidityResponse{}, types.ErrConvertAmountAToInt
	}
	poolAmountB, ok := sdk.NewIntFromString(pool.AmountB)
	if !ok {
		return &types.MsgAddLiquidityResponse{}, types.ErrConvertAmountBToInt
	}
	poolShares, ok := sdk.NewIntFromString(pool.Shares)
	if !ok {
		return &types.MsgAddLiquidityResponse{}, types.ErrConvertSharesToInt
	}
	msgShares, ok := sdk.NewIntFromString(msg.MinShares)
	if !ok {
		return &types.MsgAddLiquidityResponse{}, types.ErrConvertSharesToInt
	}
	provShares, ok := sdk.NewIntFromString(liqProv.ShareAmount)
	if !ok {
		return &types.MsgAddLiquidityResponse{}, types.ErrConvertSharesToInt
	}
	// check ratios
	if poolAmountA.Mul(msgAmountB) != poolAmountB.Mul(msgAmountA) {
		return &types.MsgAddLiquidityResponse{}, types.ErrInvalidRatio
	}
	// get new shares out
	newShares := msgAmountA.Mul(poolShares).Quo(poolAmountA)
	// check shares are sufficient
	if newShares.LT(msgShares) {
		return &types.MsgAddLiquidityResponse{}, types.ErrNotEnoughSharesOut
	}
	// get creator acc address
	accAddr, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return &types.MsgAddLiquidityResponse{}, types.ErrAccAddressFromMsg
	}
	// create coins to send
	coinA := sdk.NewCoin(msg.DenomA, msgAmountA)
	coinB := sdk.NewCoin(msg.DenomB, msgAmountB)
	// send coins from account to module
	err = k.bankKeeper.SendCoinsFromAccountToModule(ctx, accAddr, types.ModuleName, sdk.NewCoins(coinA, coinB))
	if err != nil {
		return &types.MsgAddLiquidityResponse{}, types.ErrSendCoinToModule
	}
	// get new pool balances
	newAmountA := poolAmountA.Add(msgAmountA)
	newAmountB := poolAmountB.Add(msgAmountB)
	newPoolShares := newShares.Add(poolShares)
	// get new liq prov balances
	newProvShares := newShares.Add(provShares)
	// get new pool and provider
	newPool := types.NewPool(newAmountA.String(), msg.DenomA, newAmountB.String(), msg.DenomB, newPoolShares.String())
	newLiqProv := types.NewLiqProv(newProvShares.String(), poolName, msg.Creator)
	// update pool and provider
	k.Keeper.SetPool(ctx, newPool)
	k.Keeper.SetLiqProv(ctx, newLiqProv)

	return &types.MsgAddLiquidityResponse{}, nil
}
