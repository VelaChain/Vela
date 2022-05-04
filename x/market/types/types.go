package types

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

// TODO add fees to pool type 
func NewFee(swapFee, exitFee, poolName string ) FeeMap {
	return FeeMap{
		Swap:		swapFee,
		Exit:		exitFee, 
		PoolName:	poolName,
	}
}

func NewLiqProv(shareAmount, poolName, address string) LiqProv {
	return LiqProv{
		ShareAmount: shareAmount,
		PoolName:    poolName,
		Address:     address,
	}
}
