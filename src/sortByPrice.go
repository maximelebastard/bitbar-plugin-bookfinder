package main

import (
	"strconv"
)


// ByPrice implements sort.Interface for strings based on
// the prices
type SortByPrice []string

func (p SortByPrice) Len() int { 
	return len(p) 
}
func (p SortByPrice) Swap(i, j int) { 
	p[i], p[j] = p[j], p[i] 
}
func (p SortByPrice) Less(i, j int) bool { 
	var priceA, errA = strconv.ParseFloat(replaceAll(p[i], "[^\\d.]", ""), 32)
	var priceB, errB = strconv.ParseFloat(replaceAll(p[j], "[^\\d.]", ""), 32)

	deal(errA)
	deal(errB)

	return priceA < priceB 
}