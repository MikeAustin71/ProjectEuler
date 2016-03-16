package main

import (
	"fmt"
	"math"
)

// Brute force approach - This is slowest and worst performing solution.
// Note: at upper limits <= 20, performance is satisfactory. Beginning at
// upper limit = 30 speed slows dramatically. Note: Analysis is based on
// uint64.
//
// See setup info at bottom of this file.

func main() {
	series := getSeries()

	smallestDividend, success := findSmallestDividend(series)

	if success {
		fmt.Println("Success. The smallest dividend evenly divisible by the series 1 to 20 is - ", smallestDividend)

	} else {
		fmt.Println("Failure - Did Not Locate smallest dividend.")
	}

}

func getSeries() []uint64 {

	var MAXSERIESNUM uint64 = 20

	series := make([]uint64, MAXSERIESNUM-1)

	for i := uint64(2); i <= MAXSERIESNUM; i++ {
		series[i-2] = i
	}

	return series
}

func findSmallestDividend(series []uint64) (uint64, bool) {

	for i := series[len(series)-1] + 1; i < math.MaxUint64; i++ {

		if isTestNumDivisbleBySeries(series, i) {
			return i, true
		}
	}

	return 0, false
}

func isTestNumDivisbleBySeries(series []uint64, testNum uint64) bool {

	for _, v := range series {

		if testNum%v != 0 {
			return false
		}

	}

	return true
}

/*

* Source: https://projecteuler.net/problem=5

* Title: Smallest multiple - Problem 5

* Problem Description

-----------------------------------------------------------------------
2520 is the smallest number that can be divided by each of the numbers from 1 to 10 without any remainder. What is the smallest positive number that is evenly divisible by all of the numbers from 1 to 20?
-----------------------------------------------------------------------
Success, Correct Answer = 232792560 - Confirmed by Project Euler

*/
