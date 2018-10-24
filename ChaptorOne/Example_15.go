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
	advertFile, err := os.Open("./CData/Advertising.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer advertFile.Close()
	advDF := dataframe.ReadCSV(advertFile)
	yVals := advDF.Col("Sales").Float()
	fmt.Println(yVals)

	for _, colName := range advDF.Names() {
		pts := make(plotter.XYs, advDF.Nrow())
		for i, floatVal := range advDF.Col(colName).Float() {
			pts[i].X = floatVal
			pts[i].Y = yVals[i]

		}
		p, err := plot.New()
		if err != nil {
			log.Fatal(err)
		}
		p.X.Label.Text = colName
		p.Y.Label.Text = "Y"
		p.Add(plotter.NewGrid())
		s, err := plotter.NewScatter(pts)
		if err != nil {
			log.Fatal(err)
		}
		s.GlyphStyle.Radius = vg.Points(3)
		p.Add(s)
		if err := p.Save(4*vg.Inch, 4*vg.Inch, colName+"_scatter.png"); err != nil {
			log.Fatal(err)
		}
	}

}
