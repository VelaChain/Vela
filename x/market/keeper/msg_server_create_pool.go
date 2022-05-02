package keeper

import (
	"context"

	"github.com/VelaChain/vela/x/market/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) validateCreatePoolMsg(ctx sdk.Context, msg *types.MsgCreatePool) (amtA sdk.Int, amtB sdk.Int, addr sdk.AccAddress, err error) {
	// check ints from strings ok
	msgAmountA, ok := sdk.NewIntFromString(msg.AmountA)
	if !ok {
		return amtA, amtB, addr, types.ErrConvertAmountAToInt
	}
	msgAmountB, ok := sdk.NewIntFromString(msg.AmountB)
	if !ok {
		return amtA, amtB, addr, types.ErrConvertAmountBToInt
	}
	if _, ok = sdk.NewIntFromString(msg.MinShares); !ok {
		return amtA, amtB, addr, types.ErrConvertSharesToInt
	}
	// check different denoms
	if msg.DenomA == msg.DenomB {
		return amtA, amtB, addr, types.ErrDenomsSame
	}
	// get creator account address
	accAddr, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return amtA, amtB, addr, types.ErrAccAddressFromMsg
	}
	// check account balances
	if !k.bankKeeper.HasBalance(ctx, accAddr, sdk.NewCoin(msg.DenomA, msgAmountA)) {
		return amtA, amtB, addr, types.ErrInsufficientBalanceA
	}
	if !k.bankKeeper.HasBalance(ctx, accAddr, sdk.NewCoin(msg.DenomB, msgAmountB)) {
		return amtA, amtB, addr, types.ErrInsufficientBalanceB
	}

	return msgAmountA, msgAmountB, accAddr, nil
}

func (k msgServer) CreatePool(goCtx context.Context, msg *types.MsgCreatePool) (*types.MsgCreatePoolResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// validate params and balances are available
	msgAmountA, msgAmountB, accAddr, err := k.validateCreatePoolMsg(ctx, msg)
	if err != nil {
		return &types.MsgCreatePoolResponse{}, err
	}
	// check if pool exists
	if _, found := k.GetPool(ctx, msg.DenomA, msg.DenomB); found {
		return &types.MsgCreatePoolResponse{}, types.ErrPoolExists
	}
	// check if liq prov exists
	poolName := types.NewPoolName(msg.DenomA, msg.DenomB)
	if _, found := k.GetLiqProv(ctx, poolName, msg.Creator); found {
		return &types.MsgCreatePoolResponse{}, types.ErrLiqProvExists
	}
	// convert amounts to strings	
	// msgAmountA, ok := sdk.NewIntFromString(msg.AmountA)
	// if !ok {
	// 	return &types.MsgCreatePoolResponse{}, types.ErrConvertAmountAToInt
	// }
	// msgAmountB, ok := sdk.NewIntFromString(msg.AmountB)
	// if !ok {
	// 	return &types.MsgCreatePoolResponse{}, types.ErrConvertAmountBToInt
	// }
	// create pool
	pool := types.NewPool(msg.AmountA, msg.DenomA, msg.AmountB, msg.DenomB, msg.MinShares)
	// get creator acc address
	// accAddr, err := sdk.AccAddressFromBech32(msg.Creator)
	// if err != nil {
	// 	return &types.MsgCreatePoolResponse{}, types.ErrAccAddressFromMsg
	// }
	// create coins to send
	coinA := sdk.NewCoin(msg.DenomA, msgAmountA)
	coinB := sdk.NewCoin(msg.DenomB, msgAmountB)
	// send coins from account to module
	if err = k.bankKeeper.SendCoinsFromAccountToModule(ctx, accAddr, types.ModuleName, sdk.NewCoins(coinA, coinB)); err != nil {
		return &types.MsgCreatePoolResponse{}, types.ErrSendCoinToModule
	}
	// create liquidity provider
	liqProv := types.NewLiqProv(msg.MinShares, poolName, msg.Creator)
	// update pool and liq prov
	k.Keeper.SetPool(ctx, pool)
	k.Keeper.SetLiqProv(ctx, liqProv)

	return &types.MsgCreatePoolResponse{}, nil
}
