package utils

import (
	"crypto/rand"
	"io"
	math "math/rand"
	"strconv"
	"time"
)

var (
	length = 6
)

func EncodeToString() string {
	b := make([]byte, length)
	n, err := io.ReadAtLeast(rand.Reader, b, length)
	if n != length {
		panic(err)
	}
	for i := 0; i < len(b); i++ {
		b[i] = table[int(b[i])%len(table)]
	}
	return string(b)
}

var table = [...]byte{'1', '2', '3', '4', '5', '6', '7', '8', '9', '0'}

func GenerateRandomNumbers(x, max int, ignore []string) []int {
	math.Seed(time.Now().UnixNano())

	result := make([]int, 0, x)
	ignoreMap := make(map[int]bool)

	for _, num := range ignore {
		number, _ := strconv.Atoi(num)
		ignoreMap[number] = true
	}

	for len(result) < x {
		num := math.Intn(max + 1)

		if !ignoreMap[num] {
			result = append(result, num)
		}
	}

	return result
}

func IntSliceToStringSlice(ints []int) []string {
	result := make([]string, len(ints))
	for i, v := range ints {
		result[i] = strconv.Itoa(v)
	}
	return result
}

func StringSliceToIntSlice(strings []string) []int {
	result := make([]int, len(strings))
	for i, v := range strings {
		result[i], _ = strconv.Atoi(v)
	}
	return result
}
