package main

import (
	"fmt"
	"log"
	"reflect"

	"github.com/jonas-p/go-shp"
)

func main() {
	shape, err := shp.Open("ne_110m_land.shp")

	if err != nil {
		log.Fatal(err)
	}
	defer shape.Close()

	bbox := shape.BBox()
	fmt.Printf("Bbox of shapefile: %v\n", bbox)

	fields := shape.Fields()

	// loop through all features in the shapefile
	for shape.Next() {
		n, p := shape.Shape()

		// print feature
		fmt.Println(reflect.TypeOf(p).Elem(), p.BBox())

		// print attributes
		for k, f := range fields {
			val := shape.ReadAttribute(n, k)
			fmt.Printf("\t%v: %v\n", f, val)
		}
		fmt.Println()
	}
}
