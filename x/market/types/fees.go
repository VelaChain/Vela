package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	DefaultExitFee = "0.01"
	DefaultSwapFee = "0.01"
)

func ApplySwapFee(msgAmountIn sdk.Int) (sdk.Int, error) {
	// get fee as sdk dec
	swapFee, err := sdk.NewDecFromStr(DefaultSwapFee)
	if err != nil {
		return sdk.ZeroInt(), err
	}
	// get fee for msg amount in
	feeAmount := swapFee.MulInt(msgAmountIn)
	if !feeAmount.IsPositive() {
		return sdk.ZeroInt(), ErrSwapFeeAmountNotPos
	}
	// only swap with msg amount in - fee amount
	return msgAmountIn.SubRaw(feeAmount.RoundInt64()), nil
}

func ApplyExitFee(provShares sdk.Int) (sdk.Int, error) {
	// get fee as sdk dec
	exitFee, err := sdk.NewDecFromStr(DefaultExitFee)
	if err != nil {
		return sdk.ZeroInt(), err
	}
	// get fee amount
	feeAmount := exitFee.MulInt(provShares)
	if !feeAmount.IsPositive(){
		return sdk.ZeroInt(), ErrExitfeeAmountNotPos
	}
	// use provShares - feeAmount for amounts out
	return provShares.SubRaw(feeAmount.RoundInt64()), nil
}
