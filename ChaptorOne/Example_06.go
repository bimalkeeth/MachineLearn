package main

import (
	"fmt"
	"github.com/gonum/floats"
	"github.com/kniren/gota/dataframe"
	"gonum.org/v1/gonum/stat"
	"log"
	"os"
)

func main() {

	irisFile, err := os.Open("./CData/iris_labeled.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer irisFile.Close()
	irisDF := dataframe.ReadCSV(irisFile)

	// Get the float values from the "sepal_length" column as
	// we will be looking at the measures for this variable.

	sepalLength := irisDF.Col("petal_length").Float()

	// Calculate the Max of the variable.
	minVal := floats.Min(sepalLength)

	// Calculate the Max of the variable.
	maxVal := floats.Max(sepalLength)

	// Calculate the Median of the variable.
	reangeVal := maxVal - minVal

	// Calculate the variance of the variable.
	varianceVal := stat.Variance(sepalLength, nil)

	// Calculate the standard deviation of the variable.
	stadVal := stat.StdDev(sepalLength, nil)

	// Sort the values.
	inds := make([]int, len(sepalLength))
	floats.Argsort(sepalLength, inds)

	// Get the Quantiles
	quant25 := stat.Quantile(0.25, stat.Empirical, sepalLength, nil)
	quant50 := stat.Quantile(0.50, stat.Empirical, sepalLength, nil)
	quant75 := stat.Quantile(0.75, stat.Empirical, sepalLength, nil)

	// Output the results to standard out.
	fmt.Printf("\nSepal Length Summary Statistics:\n")
	fmt.Printf("Max value: %0.2f\n", maxVal)
	fmt.Printf("Min value: %0.2f\n", minVal)
	fmt.Printf("Range value: %0.2f\n", reangeVal)
	fmt.Printf("Variance value: %0.2f\n", varianceVal)
	fmt.Printf("Std Dev value: %0.2f\n", stadVal)
	fmt.Printf("25 Quantile: %0.2f\n", quant25)
	fmt.Printf("50 Quantile: %0.2f\n", quant50)
	fmt.Printf("75 Quantile: %0.2f\n\n", quant75)
}
