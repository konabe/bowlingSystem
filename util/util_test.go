package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnique(t *testing.T) {
	list := []int{0, 1, 2, 3}
	empty_list := []int{}
	assert.ElementsMatch(t, []int{0, 1, 2, 3}, Unique(list))
	assert.ElementsMatch(t, empty_list, Unique(empty_list))
}

func TestContains(t *testing.T) {
	list := []int{0, 1, 2, 3}
	assert.True(t, Contains(list, 0))
	assert.True(t, Contains(list, 1))
	assert.True(t, Contains(list, 2))
	assert.True(t, Contains(list, 3))
	assert.False(t, Contains(list, 4))
}
