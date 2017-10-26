package fibonacci

import (
	"fmt"
	"testing"
)

const fibLen = 64

func TestFib1(t *testing.T) {
	for i := 0; i < fibLen; i++ {
		fmt.Printf("fib(%d) = %d\n", i, Fib1(i))
	}
}

func TestFib2(t *testing.T) {
	for i := 0; i < fibLen; i++ {
		fmt.Printf("fib(%d) = %d\n", i, Fib2(i))
	}
}
