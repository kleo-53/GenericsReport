package main

import (
	"fmt"
	"math/rand"
	"reflect"
	"sort"
	"time"
)

type A[T comparable, U comparable] struct {
	first  T
	second U
	third  int64
}

type B struct {
	first  string
	second int
	third  int64
}

type Comparable[T any] interface {
	Less(other T) bool
}

func test[T comparable](arr []T) time.Duration {
	start := time.Now()
	sort.Slice(arr, func(i, j int) bool {
		a := arr[i]
		b := arr[j]

		// Проверяем, реализует ли тип T интерфейс Comparable.
		if isComparable(a) {
			// Получаем метод Less и вызываем его для сравнения элементов.
			lessMethod := reflect.ValueOf(a).MethodByName("Less")
			args := []reflect.Value{reflect.ValueOf(b)}
			result := lessMethod.Call(args)
			return result[0].Bool()
		}

		// Если тип не реализует интерфейс Comparable, возвращаем false.
		return false
	})
	elapsed := time.Since(start)
	return elapsed
}

// Функция isComparable проверяет, реализует ли тип T интерфейс Comparable.
func isComparable[T comparable](a T) bool {
	compType := reflect.TypeOf((*Comparable[T])(nil)).Elem()
	return reflect.TypeOf(a).Implements(compType)
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

		sortTimeForA := test(vecA)
		sortTimeForB := test_b(vecB)

		totalSortTimeForA += sortTimeForA
		totalSortTimeForB += sortTimeForB
	}

	fmt.Printf("Mean elapsed time of sort with generics: %v\n", totalSortTimeForA/time.Duration(nExperiments))
	fmt.Printf("Mean elapsed time of sort without generics: %v\n", totalSortTimeForB/time.Duration(nExperiments))
}
