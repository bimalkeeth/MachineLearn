package main

import (
	"bufio"
	"fmt"
	"github.com/kniren/gota/dataframe"
	"log"
	"os"
)

func main() {
	advertFile, err := os.Open("./CData/Advertising.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer advertFile.Close()
	advDF := dataframe.ReadCSV(advertFile)
	trainingNum := (4 * advDF.Nrow()) / 5
	testNum := advDF.Nrow() / 5

	if trainingNum+testNum < advDF.Nrow() {
		trainingNum++
	}
	trainingIdx := make([]int, trainingNum)
	testIdx := make([]int, testNum)

	for i := 0; i < trainingNum; i++ {
		trainingIdx[i] = i
	}
	fmt.Println(trainingIdx)
	for i := 0; i < testNum; i++ {
		testIdx[i] = trainingNum + i
	}
	fmt.Println(testIdx)

	trainingDF := advDF.Subset(trainingIdx)
	testDF := advDF.Subset(testIdx)

	setMap := map[int]dataframe.DataFrame{
		0: trainingDF,
		1: testDF,
	}
	for idx, setName := range []string{"training.csv", "test.csv"} {

		f, err := os.Create(setName)
		if err != nil {
			log.Fatal(err)
		}
		w := bufio.NewWriter(f)
		if err := setMap[idx].WriteCSV(w); err != nil {
			log.Fatal(err)
		}
	}

}
