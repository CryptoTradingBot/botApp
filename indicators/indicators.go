package indicators

import "math"

/**
	The technical analysis indicator, designed to smooth the time series by the formula

	RMA = ((RMAt-1×(n-1))+Xt)/n,

	when n – width window, X – price.
	First RMA is calculated as SMA.
	RMA smoothes better than EMA with the same window width.
 */
// rolling moving average индикатор
func RMA(previousRMA float64, period int, price float64) float64 {
	return ((previousRMA * (float64(period) - 1)) + price) / float64(period)
}

/*
	A simple, or arithmetic, moving average is calculated by summing the closing price of the instrument for a certain
	number of single periods (for example, in 12 hours), followed by dividing the sum by the number of periods.

	SMA = SUM (CLOSE (i), N) / N

	when:

	SUM — sum;
	CLOSE (i) — price close current period;
	N — number periods calculating.
*/
func SMA(period int, prices []float64) float64 {
	var sum float64 = 0
	for _, price := range prices {
		sum += price
	}

	return sum / float64(period)
}

// The anti-aliasing effect will be based on what will be transmitted as Prices (RMA, SMA, EMA and eth.)
func RSI(period int, prices []float64) float64 {
	var value = 100 - (100 / (1 + RS(period, prices)))
	return value
}

func RS(period int, prices []float64) float64 {
	var previousPrice = prices[0]

	var averageGain float64 = 0
	var averageLoss float64 = 0
	for i := 1; i < len(prices); i++ {
		changePrice := prices[i] - previousPrice
		if changePrice >= 0 {
			averageLoss += math.Abs(changePrice)
		} else {
			averageGain += math.Abs(changePrice)
		}
		previousPrice = prices[i]
	}

	averageGain /= float64(period)
	averageLoss /= float64(period)

	return averageGain / averageLoss
}
