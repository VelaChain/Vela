package keeper

import (
	"context"

	"github.com/VelaChain/vela/x/market/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) validateSendSharesMsg(ctx sdk.Context, msg *types.MsgSendShares) error {
	if msg.DenomA == msg.DenomB {
		return types.ErrDenomsSame
	}
	if _, err := sdk.AccAddressFromBech32(msg.Creator); err != nil {
		return types.ErrAccAddressFromMsg
	}
	if _, err := sdk.AccAddressFromBech32(msg.Address); err != nil {
		return types.ErrAccAddressFromMsg
	}
	return nil
}

func (k msgServer) SendShares(goCtx context.Context, msg *types.MsgSendShares) (*types.MsgSendSharesResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// validate message
	if err := k.validateSendSharesMsg(ctx, msg); err != nil {
		return &types.MsgSendSharesResponse{}, err
	}
	// check pool exists
	_, found := k.Keeper.GetPool(ctx, msg.DenomA, msg.DenomB)
	if !found {
		return &types.MsgSendSharesResponse{}, types.ErrPoolDNE
	}
	// check provider exists
	poolName := types.NewPoolName(msg.DenomA, msg.DenomB)
	liqProv, found := k.Keeper.GetLiqProv(ctx, poolName, msg.Creator)
	if !found {
		return &types.MsgSendSharesResponse{}, types.ErrLiqProvDNE
	}
	// convert message shares to sdk int
	msgShares, ok := sdk.NewIntFromString(msg.Shares)
	if !ok {
		return &types.MsgSendSharesResponse{}, types.ErrConvertSharesToInt
	} 
	// check msgShares is positive
	if !msgShares.IsPositive() {
		return &types.MsgSendSharesResponse{}, types.ErrShareAmountNotPos
	}
	// check reciever liq prov exists
	var newRecvLiqProv types.LiqProv
	recvLiqProv, found := k.Keeper.GetLiqProv(ctx, poolName, msg.Address) 
	if !found {
		newRecvLiqProv = types.NewLiqProv(msg.Shares, poolName, msg.Address)
	} else {
		recv, ok := sdk.NewIntFromString(recvLiqProv.ShareAmount)
		if !ok {
			return &types.MsgSendSharesResponse{}, types.ErrConvertSharesToInt
		}
		newShares := msgShares.Add(recv)
		newRecvLiqProv = types.NewLiqProv(newShares.String(), poolName, msg.Address) 
	}
	
	liqProvShares, ok := sdk.NewIntFromString(liqProv.ShareAmount)
	if !ok {
		return &types.MsgSendSharesResponse{}, types.ErrConvertSharesToInt
	} 
	newLiqProvShares := liqProvShares.Sub(msgShares)
	// remove shares from creator 
	newLiqProv := types.NewLiqProv(newLiqProvShares.String(), poolName, liqProv.Address)
	// update liq provs in list
	k.Keeper.SetLiqProv(ctx, newLiqProv)
	k.Keeper.SetLiqProv(ctx, newRecvLiqProv)

	return &types.MsgSendSharesResponse{}, nil
}
