package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"

	"golang.org/x/exp/constraints"
)

type A[T constraints.Ordered, U constraints.Ordered] struct {
	first  T
	second U
	third  int64
}

type B struct {
	first  string
	second int
	third  int64
}

func test_a[T constraints.Ordered, U constraints.Ordered](arr []A[T, U]) time.Duration {
	start := time.Now()
	sort.Slice(arr, func(i, j int) bool {
		if arr[i].first == arr[j].first {
			if arr[i].second == arr[j].second {
				return arr[i].third < arr[j].third
			}
			return arr[i].second < arr[j].second
		}
		return arr[i].first < arr[j].first
	})

	elapsed := time.Since(start)
	return elapsed
}

func generateRandomA() A[string, int] {
	return A[string, int]{
		first:  generateRandomString(),
		second: rand.Intn(100),
		third:  rand.Int63n(100),
	}
}

// Генерирует случайную строку длиной 10 символов.
func generateRandomString() string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	const length = 10
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}

func test_b(arr []B) time.Duration {
	start := time.Now()
	sort.Slice(arr, func(i, j int) bool {
		if arr[i].first == arr[j].first {
			if arr[i].second == arr[j].second {
				return arr[i].third < arr[j].third
			}
			return arr[i].second < arr[j].second
		}
		return arr[i].first < arr[j].first
	})

	elapsed := time.Since(start)
	return elapsed
}

func main() {
	totalSortTimeForA, totalSortTimeForB := time.Duration(0), time.Duration(0)

	nExperiments := 50

	for i := 0; i < nExperiments; i++ {
		vecA := make([]A[string, int], 100000)
		vecB := make([]B, 100000)
		for j := 0; j < 100000; j++ {
			randomA := generateRandomA()
			vecA[j] = randomA
			vecB[j] = B(randomA)
		}

		sortTimeForA := test_a(vecA)
		sortTimeForB := test_b(vecB)

		totalSortTimeForA += sortTimeForA
		totalSortTimeForB += sortTimeForB
	}

	fmt.Printf("Mean elapsed time of sort with generics: %v\n", totalSortTimeForA/time.Duration(nExperiments))
	fmt.Printf("Mean elapsed time of sort without generics: %v\n", totalSortTimeForB/time.Duration(nExperiments))
}
