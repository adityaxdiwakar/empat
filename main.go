package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log"
	"os"
)

func main() {
	x := []int{34, 494, 733, 237}
	y := []int{34, 293, 292, 129}
	const width, height = 900, 400

	// Create a colored image of the given width and height.
	img := image.NewNRGBA(image.Rect(0, 0, width, height))

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			img.Set(x, y, color.NRGBA{
				R: uint8(255),
				G: uint8(255),
				B: uint8(255),
				A: 255,
			})
		}
	}

	for y := height - 50; y > height-55; y-- {
		for x := 50; x < width-50; x++ {
			img.Set(x, y, color.NRGBA{
				R: uint8(0),
				G: uint8(0),
				B: uint8(0),
				A: 255,
			})
		}
	}

	for x := 50; x < 55; x++ {
		for y := height - 50; y > 50; y-- {
			img.Set(x, y, color.NRGBA{
				R: uint8(0),
				G: uint8(0),
				B: uint8(0),
				A: 255,
			})
		}
	}

	for i := 0; i < len(x); i++ {
		fmt.Println(x[i], y[i])
		img.Set(x[i]+50, y[i]+50, color.NRGBA{
			R: uint8(0),
			G: uint8(0),
			B: uint8(0),
			A: 255,
		})
		img.Set(x[i]+49, y[i]+50, color.NRGBA{
			R: uint8(0),
			G: uint8(0),
			B: uint8(0),
			A: 255,
		})
		img.Set(x[i]+51, y[i]+50, color.NRGBA{
			R: uint8(0),
			G: uint8(0),
			B: uint8(0),
			A: 255,
		})
		img.Set(x[i]+50, y[i]+51, color.NRGBA{
			R: uint8(0),
			G: uint8(0),
			B: uint8(0),
			A: 255,
		})
		img.Set(x[i]+50, y[i]+49, color.NRGBA{
			R: uint8(0),
			G: uint8(0),
			B: uint8(0),
			A: 255,
		})
	}

	f, err := os.Create("image.png")
	if err != nil {
		log.Fatal(err)
	}

	if err := png.Encode(f, img); err != nil {
		f.Close()
		log.Fatal(err)
	}

	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
}
