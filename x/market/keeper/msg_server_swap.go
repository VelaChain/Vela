package keeper

import (
	"context"

	"github.com/VelaChain/vela/x/market/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) validateSwapMsg(ctx sdk.Context, msg *types.MsgSwap) (amountIn sdk.Int, amountOut sdk.Int, accAddr sdk.AccAddress, err error) {
	amountIn, ok := sdk.NewIntFromString(msg.AmountIn)
	if !ok {
		return amountIn, amountOut, accAddr, types.ErrConvertAmountIn
	}

	amountOut, ok = sdk.NewIntFromString(msg.MinAmountOut)
	if !ok {
		return amountIn, amountOut, accAddr, types.ErrConvertAmountOut
	}

	if msg.DenomIn == msg.DenomOut {
		return amountIn, amountOut, accAddr, types.ErrDenomsSame
	}

	accAddr, err = sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return amountIn, amountOut, accAddr, types.ErrAccAddressFromMsg
	}

	if !k.bankKeeper.HasBalance(ctx, accAddr, sdk.NewCoin(msg.DenomIn, amountIn)) {
		return amountIn, amountOut, accAddr, types.ErrInsufficientBalanceIn
	}

	return amountIn, amountOut, accAddr,  nil
}

func (k msgServer) Swap(goCtx context.Context, msg *types.MsgSwap) (*types.MsgSwapResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// validate message
	msgAmountIn, minAmountOut, accAddr, err := k.validateSwapMsg(ctx, msg)
	if err != nil {
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
	// msgAmountIn, ok := sdk.NewIntFromString(msg.AmountIn)
	// if !ok {
	// 	return &types.MsgSwapResponse{}, types.ErrConvertAmountIn
	// }
	// minAmountOut, ok := sdk.NewIntFromString(msg.MinAmountOut)
	// if !ok {
	// 	return &types.MsgSwapResponse{}, types.ErrConvertAmountOut
	// }
	// get pool amounts as sdk int
	poolAmountA, ok := sdk.NewIntFromString(pool.AmountA)
	if !ok {
		return &types.MsgSwapResponse{}, types.ErrConvertAmountAToInt
	}
	poolAmountB, ok := sdk.NewIntFromString(pool.AmountB)
	if !ok {
		return &types.MsgSwapResponse{}, types.ErrConvertAmountBToInt
	}
	// get fee as sdk dec
	swapFee, err := sdk.NewDecFromStr(types.DefaultSwapFee)
	if err != nil {
		return &types.MsgSwapResponse{}, err
	}
	// get fee for msg amount in
	feeAmount := swapFee.MulInt(msgAmountIn)
	if !feeAmount.IsPositive() {
		return &types.MsgSwapResponse{}, types.ErrSwapFeeAmountNotPos
	}
	// only swap with msg amount in - fee amount
	swapAmount := msgAmountIn.SubRaw(feeAmount.RoundInt64())
	// get asset out amount and pool balances
	var amountOut, newAmountA, newAmountB sdk.Int
	if pool.DenomA == msg.DenomIn {
		// amount out b denom
		amountOut = types.BOutGivenA(poolAmountA, poolAmountB, swapAmount)
		newAmountA = poolAmountA.Add(msgAmountIn)
		newAmountB = poolAmountB.Sub(amountOut)
	} else {
		// amount out a denom
		amountOut = types.AOutGivenB(poolAmountA, poolAmountB, swapAmount)
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
	// accAddr, err := sdk.AccAddressFromBech32(msg.Creator)
	// if err != nil {
	// 	return &types.MsgSwapResponse{}, types.ErrAccAddressFromMsg
	// }
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
	if err != nil {
		return &types.MsgSwapResponse{}, err
	}
	// get new pool
	newPool := types.NewPool(newAmountA.String(), pool.DenomA, newAmountB.String(), pool.DenomB, pool.Shares)
	// set pool
	k.Keeper.SetPool(ctx, newPool)

	return &types.MsgSwapResponse{}, nil
}
