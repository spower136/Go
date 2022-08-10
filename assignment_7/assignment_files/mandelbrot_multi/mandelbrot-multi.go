package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
	"runtime"
	"time"
)

type triple struct {
	ix, iy int64
	c      int64
}

// Like the numpy linspace function
func linspace(start, end float64, num int64) []float64 {
	result := make([]float64, num)
	step := (end - start) / float64(num-1)
	for i := range result {
		result[i] = start + float64(i)*step
	}
	return result

}

// For each c, the function returns the num of iterations until abs(z)>2
// If c is in Mandelbrot set, the max iteration (threshold) is returned
func countIterationsUntilDivergent(c complex128, threshold int64) int64 {
	z := complex(0, 0)
	var iter int64
	for i := int64(0); i < threshold; i++ {
		iter = i
		z = (z * z) + c
		if cmplx.Abs(z) > 2 {
			return i
		}
	}
	return iter // iter = threshold
}

func calcRow(ix, iy int64, c complex128, threshold int64) triple {
	return triple{ix, iy, countIterationsUntilDivergent(c, threshold)}
}

func mandelbrot_multi(threshold, density int64) [][]int64 {
	realAxis := linspace(-1.8, 1, density)
	imaginaryAxis := linspace(-1.4, 1.4, density)

	// Create an atlas
	atlas := make([][]int64, len(realAxis))
	for i := range atlas {
		atlas[i] = make([]int64, len(imaginaryAxis))
	}

	// Make a buffered channel
	ch := make(chan triple, int64(len(realAxis))*int64(len(imaginaryAxis)))

	for ix, _ := range realAxis {
		go func(ix int) {
			for iy, _ := range imaginaryAxis {
				cx := realAxis[ix]
				cy := imaginaryAxis[iy]
				c := complex(cx, cy)
				res := calcRow(int64(ix), int64(iy), c, threshold)
				ch <- res
			}
		}(ix)
	}

	for i := int64(0); i < int64(len(realAxis))*int64(len(imaginaryAxis)); i++ {
		select {
		case res := <-ch:
			atlas[res.ix][res.iy] = res.c
		}
	}
	return atlas
}

func main() {
	// n := runtime.NumCPU()

	for n := 1; n < 5; n++ {
		runtime_avg := []float64{}

		// fmt.Printf("Number of cores: %d \n", n)

		for i := 1; i < n+1; i++ {
			runtime.GOMAXPROCS(i)

			start := time.Now()
			m := mandelbrot_multi(100, 1000)
			duration := time.Since(start)
			duration_time_microseconds := duration.Round(time.Microsecond)
			fmt.Printf("%d CPUs: %s\n", i, duration_time_microseconds)
			runtime_avg = append(runtime_avg, float64(duration_time_microseconds))
			saveImage(m)
			// fmt.Println(runtime_avg)
		}
		runtime_avg_sum := 0.0
		for _, v := range runtime_avg {
			runtime_avg_sum += v
		}
		runtime_avg_sum /= float64(len(runtime_avg))

	}
}

// Plot the data and save the image to a PNG file
func saveImage(data [][]int64) {
	const contrast = 10
	img := image.NewRGBA(image.Rect(0, 0, 1000, 1000))

	for ix, _ := range data {
		for iy, _ := range data[0] {
			n := data[ix][iy]

			r := 100 - contrast*n
			g := contrast * n
			b := g
			img.Set(ix, iy, color.RGBA{R: uint8(r), G: uint8(g), B: uint8(b), A: 255})

		}
	}
	f, _ := os.Create("mandelbrot-multi.png") // Encode as PNG
	png.Encode(f, img)
}
