package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/market module sentinel errors
var (
	ErrSample                = sdkerrors.Register(ModuleName, 1100, "sample error")
	ErrInvalidPacketTimeout  = sdkerrors.Register(ModuleName, 1500, "invalid packet timeout")
	ErrInvalidVersion        = sdkerrors.Register(ModuleName, 1501, "invalid version")
	ErrConvertAmountAToInt   = sdkerrors.Register(ModuleName, 2000, "cannot convert amount A to int")
	ErrConvertAmountBToInt   = sdkerrors.Register(ModuleName, 2001, "cannot convert amount B to int")
	ErrConvertSharesToInt    = sdkerrors.Register(ModuleName, 2002, "cannot convert shares amount to int")
	ErrDenomsSame            = sdkerrors.Register(ModuleName, 2003, "denoms cannot be the same")
	ErrAccAddressFromMsg     = sdkerrors.Register(ModuleName, 2004, "error getting account address")
	ErrInsufficientBalanceA  = sdkerrors.Register(ModuleName, 2005, "insufficient balance for A")
	ErrInsufficientBalanceB  = sdkerrors.Register(ModuleName, 2006, "insufficient balance for B")
	ErrNotEnoughSharesOut    = sdkerrors.Register(ModuleName, 2007, "shares out less than requested shares")
	ErrSendCoinToModule      = sdkerrors.Register(ModuleName, 2008, "error sending coins from account to module")
	ErrPoolDNE               = sdkerrors.Register(ModuleName, 2009, "Pool DNE")
	ErrPoolExists            = sdkerrors.Register(ModuleName, 2010, "Pool already exists")
	ErrLiqProvDNE            = sdkerrors.Register(ModuleName, 2011, "Liq Prov DNE")
	ErrLiqProvExists         = sdkerrors.Register(ModuleName, 2012, "Liq Prov already exists")
	ErrInvalidRatio          = sdkerrors.Register(ModuleName, 2013, "Added amounts must be at same ratio as pool")
	ErrProviderNotEmpty      = sdkerrors.Register(ModuleName, 2014, "Cannot remove provider: still has shares")
	ErrProviderEmptyRemove   = sdkerrors.Register(ModuleName, 2015, "Cannot remove liquidity: use exit pool to remove all liquidity")
	ErrAmountOutAZero        = sdkerrors.Register(ModuleName, 2016, "Amount out A is zero")
	ErrAmountOutBZero        = sdkerrors.Register(ModuleName, 2017, "Amount out B is zero")
	ErrConvertAmountIn       = sdkerrors.Register(ModuleName, 2018, "cannot convert amount in to int")
	ErrConvertAmountOut      = sdkerrors.Register(ModuleName, 2019, "cannot convert amount out to int")
	ErrInsufficientBalanceIn = sdkerrors.Register(ModuleName, 2020, "insufficient balance for amount in")
	ErrNotEnoughAmountOut    = sdkerrors.Register(ModuleName, 2021, "amount out less than requested amount out")
	ErrPoolAmountAZero       = sdkerrors.Register(ModuleName, 2022, "pool amount A zero")
	ErrPoolAmountBZero       = sdkerrors.Register(ModuleName, 2023, "pool amount b zero")
	ErrShareAmountNotPos	 = sdkerrors.Register(ModuleName, 2024, "shares amount not positive")
	ErrExitfeeAmountNotPos	 = sdkerrors.Register(ModuleName, 2025, "exit fee amount not positive")
	ErrConvertFeeToDec		 = sdkerrors.Register(ModuleName, 2026, "cannot convert fee to dec")
	ErrSwapFeeAmountNotPos	 = sdkerrors.Register(ModuleName, 2027, "swap fee amount not positive")
)
