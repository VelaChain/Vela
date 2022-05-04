package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	DefaultExitFee = "0.01"
	DefaultSwapFee = "0.01"
)

func AOutGivenShares(poolAmountA, poolShares, provShares sdk.Int) sdk.Int {
	// get ratio pool amount A / pool shares
	poolRatioAtoShares := sdk.NewDecFromInt(poolAmountA).QuoRoundUp(sdk.NewDecFromInt(poolShares))
	// return amount A
	return poolRatioAtoShares.MulInt(provShares).RoundInt()
}

func BOutGivenShares(poolAmountB, poolShares, provShares sdk.Int) sdk.Int {
	// get ratio pool amount b / pool shares
	poolRatioBtoShares := sdk.NewDecFromInt(poolAmountB).QuoRoundUp(sdk.NewDecFromInt(poolShares))
	// return amount B
	return poolRatioBtoShares.MulInt(provShares).RoundInt()
}

func BOutGivenA(poolAmountA, poolAmountB, msgAmountA sdk.Int) sdk.Int {
	// get ratio poolAmountA / poolAmountB 
	poolRatioAtoB := sdk.NewDecFromInt(poolAmountA).QuoRoundUp(sdk.NewDecFromInt(poolAmountB))
	// return amount B for A
	return poolRatioAtoB.MulInt(msgAmountA).RoundInt()
}

func AOutGivenB(poolAmountA, poolAmountB, msgAmountB sdk.Int) sdk.Int {
	// get ratio poolAmountB / poolAmountA
	poolRatioBtoA := sdk.NewDecFromInt(poolAmountB).QuoRoundUp(sdk.NewDecFromInt(poolAmountA))
	// return amount A for B
	return poolRatioBtoA.MulInt(msgAmountB).RoundInt()
}

func CheckRatios(poolAmountA, poolAmountB, msgAmountA, msgAmountB sdk.Int) error {
	// get amount B needed for input A
	needB := BOutGivenA(poolAmountA, poolAmountB, msgAmountA)
	if !needB.IsPositive() {
		return sdkerrors.Wrapf(ErrAmountBNotPos, "for pool's amounts and %s amount A", msgAmountA.String())
	}
	// get amount A needed for input B
	needA := AOutGivenB(poolAmountA, poolAmountB, msgAmountB)
	if !needA.IsPositive() {
		return sdkerrors.Wrapf(ErrAmountANotPos, "for pool's amounts and %s amount B", msgAmountB.String())
	}
	// check amount b correct for amount a
	if !needB.Equal(msgAmountB) {
		return sdkerrors.Wrapf(ErrInvalidRatio, "For %s amount A you must add %s amount B", msgAmountA.String(), needB.String())
	}
	// check amount a correct for amount b
	if !needA.Equal(msgAmountA) {
		return sdkerrors.Wrapf(ErrInvalidRatio, "For %s beta you must add %s alpha", msgAmountB.String(), needA.String())
	}
	return nil
}

func NewPool(amountA, denomA, amountB, denomB, shares string) Pool {
	return Pool{
		AmountA: amountA,
		DenomA:  denomA,
		AmountB: amountB,
		DenomB:  denomB,
		Shares:  shares,
	}
}

func NewPoolName(denomA, denomB string) string {
	if denomA > denomB {
		return denomB + "-" + denomA
	}
	return denomA + "-" + denomB
}

func NewLiqProv(shareAmount, poolName, address string) LiqProv {
	return LiqProv{
		ShareAmount: shareAmount,
		PoolName:    poolName,
		Address:     address,
	}
}
