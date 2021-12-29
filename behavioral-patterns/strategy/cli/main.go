package main

import (
	"flag"
	"go-design-patterns/behavioral-patterns/strategy/shapes"
	"log"
	"os"
)

var output = flag.String("output", "text", "The output to use between 'console' and 'image' file")

func main() {
	flag.Parse()

	activeStrategy, err := shapes.NewPrinter(*output)
	if err != nil {
		log.Fatal(err)
	}
	switch *output {
	case shapes.TEXT_STRATEGY:
		activeStrategy.SetWriter(os.Stdout)
	case shapes.IMAGE_STRATEGY:
		w, err := os.Create("./results/red-square.jpg")
		if err != nil {
			log.Fatal("Error opening image")
		}
		defer w.Close()
		activeStrategy.SetWriter(w)
	}
	err = activeStrategy.Print()
	if err != nil {
		log.Fatal(err)
	}

}
