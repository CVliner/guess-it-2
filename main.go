package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	var numbers []float64
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		x, err := strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			fmt.Println("Invalid input:", err)
			continue
		}

		numbers = append(numbers, x)

		if len(numbers) > 1 {
			a, b := Range(numbers[:len(numbers)-1])
			fmt.Printf("%d %d\n", a, b)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
	}
}

func Range(numbers []float64) (int, int) {
	x := make([]float64, len(numbers))
	y := numbers

	for i := range x {
		x[i] = float64(i)
	}

	a, b := LinearRegression(x, y)
	next := float64(len(numbers))
	pred := a + b*next

	res := make([]float64, len(numbers))
	for i := range res {
		res[i] = y[i] - (a + b*x[i])
	}

	sd := StandardDeviation(res)

	lower := int(pred - sd)
	if lower < 1 {
		lower = 1
	}
	upper := int(pred + sd)

	return lower, upper
}

func Average(numbers []float64) float64 {
	var sum float64
	for _, value := range numbers {
		sum = value + sum
	}
	return sum / float64(len(numbers))
}

func StandardDeviation(res []float64) float64 {
	avg := Average(res)
	var sqrt float64
	for _, r := range res {
		sqrt += (r - avg) * (r - avg)
	}
	return math.Sqrt(sqrt / float64(len(res)-1))
}

func LinearRegression(x []float64, y []float64) (float64, float64) {
	length := float64(len(x))
	var X, Y, XY, XX float64
	for i, yi := range y {
		xi := x[i]
		X += xi
		Y += yi
		XY += xi * yi
		XX += xi * xi
	}
	a := (length*XY - X*Y) / (length*XX - X*X)
	b := (Y - a*X) / length
	return b, a
}
