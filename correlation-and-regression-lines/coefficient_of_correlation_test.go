package correlation_and_regression_lines

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"strconv"
	"strings"
	"testing"
)

// Given the test scores of 10 students in Physics and History,
// compute Karl Pearson’s coefficient of correlation between these scores.
// Round the result to three decimal places.
// Using Pearson's formula: r = (n∑XY - ∑X∑Y) / √[(n∑X² - (∑X)²)(n∑Y² - (∑Y)²)]
func TestCoefficientOfCorrelation(t *testing.T) {
	//Enter your code here. Read input from STDIN. Print output to STDOUT
	//reader := bufio.NewReaderSize(os.Stdin, 1024)
	//
	//physicsScores, err := readInts(reader)
	//if err != nil {
	//	// fmt.Printf("Error: %+v")
	//}
	//// fmt.Println(physicsScores)
	//
	//historyScores, err := readInts(reader)
	//if err != nil {
	//	// fmt.Printf("Error: %+v")
	//}
	//// fmt.Println(historyScores)
	//
	//n := len(physicsScores)
	//if n != len(historyScores) || n == 0 {
	//	fmt.Printf("input slices must be non-empty and equal length")
	//}

	physicsScores := []float64{15.0, 12.0, 8.0, 8.0, 7.0, 7.0, 7.0, 6.0, 5.0, 3.0}
	historyScores := []float64{10.0, 25.0, 17.0, 11.0, 13.0, 17.0, 20.0, 13.0, 9.0, 15.0}

	n := len(physicsScores)

	// Compute sums
	sumX, sumY, sumXY, sumX2, sumY2 := 0.0, 0.0, 0.0, 0.0, 0.0
	for i := 0; i < n; i++ {
		sumX += physicsScores[i]
		sumY += historyScores[i]
		sumXY += physicsScores[i] * historyScores[i]
		sumX2 += physicsScores[i] * physicsScores[i]
		sumY2 += historyScores[i] * historyScores[i]
	}

	numerator := float64(n)*sumXY - sumX*sumY
	denomX := float64(n)*sumX2 - sumX*sumX
	denomY := float64(n)*sumY2 - sumY*sumY
	denominator := math.Sqrt(denomX * denomY)

	fmt.Printf("%.3f", numerator/denominator)
}

func readInts(reader *bufio.Reader) ([]int, error) {
	fields := strings.Split(strings.TrimSpace(readLine(reader)), " ")
	nums := make([]int, len(fields))
	for i, f := range fields {
		n, err := strconv.Atoi(f) // convert string -> int
		if err != nil {
			return nil, err
		}
		nums[i] = n
	}
	return nums, nil
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}
