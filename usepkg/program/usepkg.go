package main

import (
	"fmt"
	"goproject/usepkg/custompkg"
	"goproject/usepkg/exinit"

	"github.com/guptarohit/asciigraph"
	"github.com/tuckersGo/musthaveGo/ch16/expkg"
)

func main() {
	custompkg.PrintCustom()
	expkg.PrintSample()

	data := []float64{3, 4, 5, 6, 7, 8, 9, 10, 11, 2}
	graph := asciigraph.Plot(data)
	fmt.Println(graph)
	exinit.PrintD()

}
