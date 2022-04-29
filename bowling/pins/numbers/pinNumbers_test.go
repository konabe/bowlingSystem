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
		isNil   bool
	}{
		{
			[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			false,
		},
		{
			[]int{1},
			[]int{1},
			false,
		},
		{
			[]int{},
			[]int{},
			false,
		},
		{
			nil,
			nil,
			false,
		},
		{
			[]int{1, 1, 2, 2, 3, 3, 3, 3},
			[]int{1, 2, 3},
			false,
		},
		{
			[]int{1, 2, 3, 11},
			nil,
			true,
		},
	}
	for i, c := range cases {
		failedMessage := fmt.Sprintf("case %d is failed", i)
		pn := New(c.numbers...)
		assert.Equal(t, c.isNil, pn == nil, failedMessage)
		if pn != nil {
			assert.ElementsMatch(t, c.values, pn.Values, failedMessage)
		}
	}
}
