package main

import (
	cr "crypto/rand"
	"fmt"
	"math"
	"math/big"
	"math/rand"
	"sync"
	"time"
)

// TODO: Algorithms to implement
// [X] Search Problem
// [ ] Scalar Product
// [ ] Selection
// [ ] Merging

func main() {

	searchProblem()
}

func searchProblem() {

	t := time.Now()

	for pow := 0; pow < 31; pow++ {

		res := [20][2]time.Duration{}
		arr := []int{}
		for i := 0; i < 20; i++ {
			pp, _ := cr.Int(cr.Reader, big.NewInt(int64(math.Pow(2, 9))))
			rand.Seed(pp.Int64())
			arr = rand.Perm(int(math.Pow(2, float64(pow))))
			// tA, tB := searcher(arr)
			tB := searcher(arr)
			// res[i][0] = tA
			res[i][1] = tB
		}

		var total1 time.Duration
		for _, v := range res {
			total1 += v[0]
		}
		var total2 time.Duration
		for _, v := range res {
			total2 += v[1]
		}
		var total3 time.Duration
		for _, v := range res {
			total3 += v[0] - v[1]
		}

		// Array Length, Time Sequential, Time Concurrent, Difference Time
		// fmt.Println(len(arr),
		// 	float64(total1.Nanoseconds())/1000,
		// 	float64(total2.Nanoseconds())/1000,
		// 	float64(total3.Nanoseconds())/1000)

		fmt.Println(float64(total2.Nanoseconds()) / 1000)
	}

	tT := time.Since(t)
	fmt.Println(tT)
}

func searcher(arr []int) time.Duration {

	date := 4
	p := 512
	// part := 0
	// pos := 0
	// 16 processors / goroutines
	elems := len(arr) / p // the number of elements per block

	// primitive sequential algorithm
	// t1 := time.Now()
	// _, pResult := findDate(arr, date, part)
	// tA := time.Since(t1)
	// pResult++

	// concurrent algorithm
	var wg sync.WaitGroup

	t2 := time.Now()
	tB := time.Since(t2)
	for i := 0; i < p; i++ {
		wg.Add(1)
		nArr := arr[elems*i : elems*(i+1)]
		i := i

		go func() {
			defer wg.Done()
			partT, posT := findDate(nArr, date, i)
			if partT != -1 && posT != -1 {
				tB = time.Since(t2)
				// part = partT
				// pos = posT
			}
		}()
	}
	wg.Wait()

	// result := (elems * part) + pos
	// fmt.Println("Primitive Result: ", pResult, ", ", tA)
	// fmt.Println("Concurrent Result: ", result, ", ", tB)

	// return tA, tB
	return tB
}

func findDate(nArr []int, date, part int) (int, int) {

	for i := 0; i < len(nArr); i++ {
		if nArr[i] == date {
			return part, i
		}
	}
	return -1, -1
}
