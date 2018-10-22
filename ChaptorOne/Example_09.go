package main

import (
	"fmt"
	"github.com/gonum/stat/distuv"
	"gonum.org/v1/gonum/stat"
)

func main() {
	observed := []float64{260.0, 135.0, 105.0}
	totalObserved := 500.0
	expected := []float64{totalObserved * 0.60, totalObserved * 0.25, totalObserved * 0.15}
	chiSquire := stat.ChiSquare(observed, expected)
	fmt.Printf("\nChi_Squire: %0.2f\n,", chiSquire)

	chiDist := distuv.ChiSquared{
		K:   2.0,
		Src: nil,
	}
	pValue := chiDist.Prob(chiSquire)
	fmt.Printf("p-value: %0.4f\n\n", pValue)

}
