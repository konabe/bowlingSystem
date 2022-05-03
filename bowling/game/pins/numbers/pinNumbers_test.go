package pinNumbers

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNew(t *testing.T) {
	cases := []struct {
		numbers []int
		values  []int
		len     int
	}{
		{
			[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			10,
		},
		{
			[]int{1},
			[]int{1},
			1,
		},
		{
			[]int{},
			[]int{},
			0,
		},
		{
			nil,
			[]int{},
			0,
		},
		{
			[]int{1, 1, 2, 2, 3, 3, 3, 3},
			[]int{1, 2, 3},
			3,
		},
		{
			[]int{1, 2, 3, 11},
			[]int{1, 2, 3},
			3,
		},
	}
	for i, c := range cases {
		failedMessage := fmt.Sprintf("case %d is failed", i)
		pn := New(c.numbers...)
		assert.Equal(t, c.values, pn.Values, failedMessage)
		assert.Equal(t, c.len, pn.Len())
	}
}
