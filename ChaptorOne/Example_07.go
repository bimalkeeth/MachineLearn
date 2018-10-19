package main

import (
	"fmt"
	"github.com/kniren/gota/dataframe"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
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
	for _, columnName := range irisDF.Names() {

		if columnName != "species" {
			v := make(plotter.Values, irisDF.Nrow())
			for i, floatVal := range irisDF.Col(columnName).Float() {
				v[i] = floatVal
			}
			p, err := plot.New()
			if err != nil {
				log.Fatal(err)
			}
			p.Title.Text = fmt.Sprintf("Histogram of a %s", columnName)
			h, err := plotter.NewHist(v, 16)
			if err != nil {
				log.Fatal(err)
			}
			h.Normalize(1)
			p.Add(h)
			if err := p.Save(4*vg.Inch, 4*vg.Inch, columnName+"_hist.png"); err != nil {
				log.Fatal(err)
			}
		}
	}

}
