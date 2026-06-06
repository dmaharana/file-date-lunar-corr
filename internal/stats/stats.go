package stats

import (
	"gonum.org/v1/gonum/stat/distuv"
)

// ChiSquareResult holds the results of the Chi-squared test
type ChiSquareResult struct {
	Statistic float64
	PValue    float64
	DF        int
}

// PerformChiSquareTest calculates the Chi-squared statistic and p-value
func PerformChiSquareTest(observed []int) ChiSquareResult {
	if len(observed) == 0 {
		return ChiSquareResult{}
	}

	total := 0
	for _, count := range observed {
		total += count
	}

	if total == 0 {
		return ChiSquareResult{}
	}

	expected := float64(total) / float64(len(observed))
	var chi2 float64
	for _, count := range observed {
		diff := float64(count) - expected
		chi2 += (diff * diff) / expected
	}

	df := len(observed) - 1
	chiDist := distuv.ChiSquared{K: float64(df), Src: nil}
	pValue := 1.0 - chiDist.CDF(chi2)

	return ChiSquareResult{
		Statistic: chi2,
		PValue:    pValue,
		DF:        df,
	}
}
