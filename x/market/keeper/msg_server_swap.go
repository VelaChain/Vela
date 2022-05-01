package keeper

import (
	"context"

	"github.com/VelaChain/vela/x/market/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) validateSwapMsg(ctx sdk.Context, msg *types.MsgSwap) error {
	amountIn, ok := sdk.NewIntFromString(msg.AmountIn)
	if !ok {
		return types.ErrConvertAmountIn
	} 	

	_, ok = sdk.NewIntFromString(msg.MinAmountOut)
	if !ok {
		return types.ErrConvertAmountOut
	}

	if msg.DenomIn == msg.DenomOut {
		return types.ErrDenomsSame
	}

	accAddr, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return types.ErrAccAddressFromMsg
	}

	if !k.bankKeeper.HasBalance(ctx, accAddr, sdk.NewCoin(msg.DenomIn, amountIn)){
		return types.ErrInsufficientBalanceIn
	}

	return nil
}

func (k msgServer) Swap(goCtx context.Context, msg *types.MsgSwap) (*types.MsgSwapResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// validate message
	if err := k.validateSwapMsg(ctx, msg); err != nil {
		return &types.MsgSwapResponse{}, err
	}
	// get denoms in order to get pool
	var denomA, denomB string
	if msg.DenomIn > msg.DenomOut {
		denomA, denomB = msg.DenomOut, msg.DenomIn
	} else {
		denomA, denomB = msg.DenomIn, msg.DenomOut
	}
	// check if pool exists
	pool, found := k.Keeper.GetPool(ctx, denomA, denomB)
	if !found {
		return &types.MsgSwapResponse{}, types.ErrPoolDNE
	}
	// get msg amounts as sdk int
	msgAmountIn, ok := sdk.NewIntFromString(msg.AmountIn)
	if !ok {
		return &types.MsgSwapResponse{}, types.ErrConvertAmountIn
	} 
	minAmountOut, ok := sdk.NewIntFromString(msg.MinAmountOut)
	if !ok {
		return &types.MsgSwapResponse{}, types.ErrConvertAmountOut
	}
	// get pool amounts as sdk int
	poolAmountA, ok := sdk.NewIntFromString(pool.AmountA)
	if !ok {
		return &types.MsgSwapResponse{}, types.ErrConvertAmountAToInt
	}
	poolAmountB, ok := sdk.NewIntFromString(pool.AmountB)
	if !ok {
		return &types.MsgSwapResponse{}, types.ErrConvertAmountBToInt
	}
	// get asset out amount and pool balances
	var amountOut, newAmountA, newAmountB sdk.Int
	if (pool.DenomA == msg.DenomIn){
		// amount out b denom
		amountOut = msgAmountIn.Mul(poolAmountB).Quo(poolAmountA)
		newAmountA = poolAmountA.Add(msgAmountIn)
		newAmountB = poolAmountB.Sub(amountOut)
	} else {
		// amount out a denom
		amountOut = msgAmountIn.Mul(poolAmountA).Quo(poolAmountB)
		newAmountA = poolAmountA.Sub(amountOut)
		newAmountB = poolAmountB.Add(msgAmountIn)
	}
	// check asset out amount is >= requested amount out
	if amountOut.LT(minAmountOut) {
		return &types.MsgSwapResponse{}, types.ErrNotEnoughAmountOut
	}
	if newAmountA.IsZero() {
		return &types.MsgSwapResponse{}, types.ErrPoolAmountAZero
	}
	if newAmountB.IsZero() {
		return &types.MsgSwapResponse{}, types.ErrPoolAmountBZero
	}
	// get creator acc address
	accAddr, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return &types.MsgSwapResponse{}, types.ErrAccAddressFromMsg
	}
	// create coins to send
	coinIn := sdk.NewCoin(msg.DenomIn, msgAmountIn)
	coinOut := sdk.NewCoin(msg.DenomOut, amountOut)
	// send coins in from account to module
	err = k.bankKeeper.SendCoinsFromAccountToModule(ctx, accAddr, types.ModuleName, sdk.NewCoins(coinIn))
	if err != nil {
		return &types.MsgSwapResponse{}, err
	}
	// send coins out from module to account
	err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, accAddr, sdk.NewCoins(coinOut))
	if err != nil{
		return &types.MsgSwapResponse{}, err
	}
	// get new pool
	newPool := types.NewPool(newAmountA.String(), pool.DenomA, newAmountB.String(), pool.DenomB, pool.Shares)
	// set pool
	k.Keeper.SetPool(ctx, newPool)

	return &types.MsgSwapResponse{}, nil
}
