package test1

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSum(t *testing.T) {
	result := Sum()
	fmt.Println(result)
	assert.Equal(t, result, 45)
}

func BenchmarkSum(b *testing.B) {
	for i := 0; i < 100; i++ {
		Sum()
	}
}
