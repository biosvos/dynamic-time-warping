package dynamic_time_warping

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDynamicTimeWarping(t *testing.T) {
	ret := DynamicTimeWarping([]int{1, 2, 3}, []int{1, 2, 3})
	assert.Equal(t, uint64(0), ret)
}
