package main

import (
	"fmt"
	"image"
	"image/color"
	"log"
	"math"
	"reflect"

	"github.com/jonas-p/go-shp"
	"github.com/llgcode/draw2d/draw2dimg"
)

// Maximum bounding box of the shapefile to be used for the map
var minLong, maxLong float64
var minLat, maxLat float64

var latDistance float64  // The spread of latitude values represented in the shapefile
var longDistance float64 // The spread of longitude values represented in the shapefile

var imageWidth, imageHeight int

// How much larger/smaller the draw context is relative to shapefile coordinate space
var scaleFactorX, scaleFactorY float64

func distance(min float64, max float64) float64 {
	return math.Abs(min - max)
}

func main() {

	imageWidth = 300
	imageHeight = 300
	// Open the Shapefile
	//shape, err := shp.Open("ne_110m_land.shp")
	shape, err := shp.Open("polygon.shp")
	if err != nil {
		log.Fatal(err)
	}
	defer shape.Close()

	// Create drawing context
	dest := image.NewRGBA(image.Rect(0, 0, imageWidth, imageHeight))
	gc := draw2dimg.NewGraphicContext(dest)
	gc.SetFillColor(color.RGBA{0xFF, 0xFF, 0xFF, 0xFF})
	gc.SetStrokeColor(color.RGBA{0x00, 0x00, 0x00, 0xFF})
	gc.SetLineWidth(3)

	// Work out the mapping between the size of the shapefile coordinate space and
	// the drawing context coordinate space
	bbox := shape.BBox()
	minLong, minLat, maxLong, maxLat = bbox.MinX, bbox.MinY, bbox.MaxX, bbox.MaxY
	latDistance = distance(minLat, maxLat)
	longDistance = distance(minLong, maxLong)

	fmt.Printf("Min Long Value In Shapefile: %f\n", minLong)
	fmt.Printf("Max Long Value In Shapefile: %f\n", maxLong)

	fmt.Printf("Min Lat Value In Shapefile: %f\n", minLat)
	fmt.Printf("Max Lat Value In Shapefile: %f\n", maxLat)

	fmt.Printf("Longitudinal Distance in shapefile: %f\n", longDistance)
	fmt.Printf("Latitudinal Distance in shapefile: %f\n", latDistance)

	// Get the width scale factor (x or longitude)
	xScaleFactor := float64(imageWidth) / longDistance
	yScaleFactor := float64(imageHeight) / latDistance

	fmt.Printf("Scale factor for drawing context X: %f\n", xScaleFactor)
	fmt.Printf("Scale factor for drawing context Y: %f\n", yScaleFactor)

	// Apply scale factor to graphics context

	gc.Scale(xScaleFactor, yScaleFactor)

	// Translate coordinate system for graphics context
	// TODO Maybe?

	fields := shape.Fields()

	// loop through all features in the shapefile
	for shape.Next() {
		n, p := shape.Shape()
		// print feature bounding box

		fmt.Println(reflect.TypeOf(p).Elem(), p.BBox())
		// print attributes
		for k, f := range fields {
			val := shape.ReadAttribute(n, k)
			fmt.Printf("\t%v: %v\n", f, val)

		}
	}
}
