package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/market module sentinel errors
var (
	ErrSample               = sdkerrors.Register(ModuleName, 1100, "sample error")
	ErrInvalidPacketTimeout = sdkerrors.Register(ModuleName, 1500, "invalid packet timeout")
	ErrInvalidVersion       = sdkerrors.Register(ModuleName, 1501, "invalid version")
	ErrConvertAmountAToInt  = sdkerrors.Register(ModuleName, 2000, "cannot convert amount A to int")
	ErrConvertAmountBToInt  = sdkerrors.Register(ModuleName, 2001, "cannot convert amount B to int")
	ErrConvertSharesToInt   = sdkerrors.Register(ModuleName, 2002, "cannot convert shares amount to int")
	ErrDenomsSame           = sdkerrors.Register(ModuleName, 2003, "denoms cannot be the same")
	ErrAccAddressFromMsg    = sdkerrors.Register(ModuleName, 2004, "error getting account address")
	ErrInsufficientBalanceA = sdkerrors.Register(ModuleName, 2005, "insufficient balance for A")
	ErrInsufficientBalanceB = sdkerrors.Register(ModuleName, 2006, "insufficient balance for B")
	ErrNotEnoughSharesOut   = sdkerrors.Register(ModuleName, 2007, "shares out less than requested shares")
	ErrSendCoinToModule     = sdkerrors.Register(ModuleName, 2008, "error sending coins from account to module")
	ErrPoolDNE              = sdkerrors.Register(ModuleName, 2009, "Pool DNE")
	ErrPoolExists           = sdkerrors.Register(ModuleName, 2010, "Pool already exists")
	ErrLiqProvDNE           = sdkerrors.Register(ModuleName, 2011, "Liq Prov DNE")
	ErrLiqProvExists        = sdkerrors.Register(ModuleName, 2012, "Liq Prov already exists")
	ErrInvalidRatio         = sdkerrors.Register(ModuleName, 2013, "Added amounts must be at same ratio as pool")
	ErrProviderNotEmpty		= sdkerrors.Register(ModuleName, 2014, "Cannot remove provider: still has shares")

)
