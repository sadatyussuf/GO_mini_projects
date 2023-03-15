package interest

// InterestRate returns the interest rate for the provided balance.
func InterestRate(balance float64) float32 {
	switch{
        case balance < float64(0): return 3.213
        case balance >= float64(0) && balance < float64(1000): return 0.5
        case balance >= float64(1000) && balance < float64(5000): return 1.621
        default: return 2.475
    }
}

// Interest calculates the interest for the provided balance.
func Interest(balance float64) float64 {
	rate := InterestRate(balance)
    interest := (float64(rate) * balance )/ float64(100)
    return interest
}

// AnnualBalanceUpdate calculates the annual balance update, taking into account the interest rate.
func AnnualBalanceUpdate(balance float64) float64 {
	interest := Interest(balance)
    return interest + balance
}

// YearsBeforeDesiredBalance calculates the minimum number of years required to reach the desired balance.
func YearsBeforeDesiredBalance(balance, targetBalance float64) int {
    years := 0
    for (balance < targetBalance) {
        balance = AnnualBalanceUpdate(balance)
        years += 1
    }
    return years
}
