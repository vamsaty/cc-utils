package cc_utils

import (
	"io"
	"os"
)

// FileExists checks if the file exists.
func FileExists(fileName string) bool {
	if _, err := os.Stat(fileName); err != nil {
		return false
	}
	return true
}

func ReadFile(fileName string) ([]byte, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	return io.ReadAll(file)
}

func Min(a ...int) int {
	ans := a[0]
	for i := 0; i < len(a); i++ {
		if a[i] < ans {
			ans = a[i]
		}
	}
	return ans
}

func Max(a ...int) int {
	ans := a[0]
	for i := 0; i < len(a); i++ {
		if a[i] > ans {
			ans = a[i]
		}
	}
	return ans
}

func PanicIf(err error) {
	if err != nil {
		panic(err)
	}
}

func BoolToInt(x bool) int {
	if x {
		return 1
	}
	return 0
}

func SumOfBools(x ...bool) int {
	sum := 0
	for _, v := range x {
		if v {
			sum++
		}
	}
	return sum
}
