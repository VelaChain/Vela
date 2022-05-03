package types

const (
	DefaultExitFee = "0.01"
	DefaultSwapFee = "0.01"
)

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
