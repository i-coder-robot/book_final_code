package test1

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSum(t *testing.T) {
	result:=Sum()
	fmt.Println(result)
	assert.Equal(t,result,45)
}

