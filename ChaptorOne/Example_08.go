package main

import (
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
	p, err := plot.New()
	if err != nil {
		log.Fatal(err)
	}
	p.Title.Text = "Box plot"
	p.Y.Label.Text = "Values"

	w := vg.Points(50)
	for idx, columnName := range irisDF.Names() {
		if columnName != "species" {
			v := make(plotter.Values, irisDF.Nrow())
			for i, floatVal := range irisDF.Col(columnName).Float() {
				v[i] = floatVal
			}
			b, err := plotter.NewBoxPlot(w, float64(idx), v)
			if err != nil {
				log.Fatal(err)
			}
			p.Add(b)

		}
	}

	p.NominalX("sepal_length", "sepal_width", "petal_length", "petal_width")
	if err := p.Save(6*vg.Inch, 8*vg.Inch, "boxplot.png"); err != nil {
		log.Fatal(err)
	}
}
