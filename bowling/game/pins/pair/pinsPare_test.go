package pair

import (
	"bowlingSystem/bowling/game/pins/numbers"
	"github.com/stretchr/testify/assert"
	"testing"
)

// 通常
func TestBowlFirst(t *testing.T) {
	frame := New()
	err := frame.BowlFirst([]int{1, 2, 3})
	assert.Nil(t, err)
	assert.Equal(t, 3, frame.FirstScore)
	assert.ElementsMatch(t, pinNumbers.New(4, 5, 6, 7, 8, 9, 10).Values, frame.firstRestPins.GetRestNumbers().Values)
	assert.Equal(t, 1, frame.bowlCount)
}

func TestBowl_Normal(t *testing.T) {
	cases := []struct {
		firstNumbers      []int
		firstScore        int
		firstRestNumbers  []int
		secondNumbers     []int
		secondScore       int
		secondRestNumbers []int
	}{
		{
			// 一投目がガーターのとき、二投目もガーターのとき
			firstNumbers: []int{}, firstScore: 0, firstRestNumbers: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			secondNumbers: []int{}, secondScore: 0, secondRestNumbers: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		},
		{ // 一投目がガーターのとき、二投目は３本のとき
			firstNumbers: []int{}, firstScore: 0, firstRestNumbers: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			secondNumbers: []int{1, 3, 5}, secondScore: 3, secondRestNumbers: []int{2, 4, 6, 7, 8, 9, 10},
		},
		{ // 一投目が３本のとき、二投目はガーターのとき
			firstNumbers: []int{1, 3, 5}, firstScore: 3, firstRestNumbers: []int{2, 4, 6, 7, 8, 9, 10},
			secondNumbers: []int{}, secondScore: 0, secondRestNumbers: []int{2, 4, 6, 7, 8, 9, 10},
		},
		{ // 一投目が３本のとき、二投目は４本のとき
			firstNumbers: []int{1, 3, 5}, firstScore: 3, firstRestNumbers: []int{2, 4, 6, 7, 8, 9, 10},
			secondNumbers: []int{2, 4, 6, 8}, secondScore: 4, secondRestNumbers: []int{7, 9, 10},
		},
		{ // 一投目が3本のとき、二投目は7本のとき (スペア)
			firstNumbers: []int{1, 3, 5}, firstScore: 3, firstRestNumbers: []int{2, 4, 6, 7, 8, 9, 10},
			secondNumbers: []int{2, 4, 6, 7, 8, 9, 10}, secondScore: 7, secondRestNumbers: nil,
		},
	}
	for _, c := range cases {
		frame := New()
		// 一投目
		err := frame.BowlFirst(c.firstNumbers)
		assert.Nil(t, err)
		assert.Equal(t, c.firstScore, frame.FirstScore)
		assert.ElementsMatch(t, pinNumbers.New(c.firstRestNumbers...).Values, frame.firstRestPins.GetRestNumbers().Values)
		assert.Equal(t, 1, frame.bowlCount)

		// 二投目
		err = frame.BowlSecond(c.secondNumbers)
		assert.Nil(t, err)
		assert.Equal(t, c.firstScore, frame.FirstScore)
		assert.Equal(t, c.secondScore, frame.SecondScore)
		assert.ElementsMatch(t, pinNumbers.New(c.firstRestNumbers...).Values, frame.firstRestPins.GetRestNumbers().Values)
		assert.ElementsMatch(t, pinNumbers.New(c.secondRestNumbers...).Values, frame.secondRestPins.GetRestNumbers().Values)
		assert.Equal(t, 2, frame.bowlCount)
	}
}

func TestBowl_Normal_Strike(t *testing.T) {
	frame := New()
	// ストライク
	err := frame.BowlFirst([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	assert.Nil(t, err)
	// 二投目はエラーになることを確認
	err = frame.BowlSecond([]int{})
	assert.NotNil(t, err)
}

// 異常系: 不正なコールのときにエラーになることをテストする
func TestBowl_Abuseful(t *testing.T) {
	frame := New()
	err := frame.BowlFirst([]int{})
	assert.Nil(t, err)
	err = frame.BowlFirst([]int{})
	assert.NotNil(t, err)
}
