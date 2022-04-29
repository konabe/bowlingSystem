package frame

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// 通常
func TestBowlFirst(t *testing.T) {
	frame := New()
	err := frame.BowlFirst([]int{1, 2, 3})
	assert.Nil(t, err)
	assert.Equal(t, 3, frame.FirstScore)
	assert.Equal(t, []int{4, 5, 6, 7, 8, 9, 10}, frame.FirstRestPins.GetRestNumbers())
	assert.Equal(t, 1, frame.BowlCount)
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
		assert.Equal(t, c.firstRestNumbers, frame.FirstRestPins.GetRestNumbers())
		assert.Equal(t, 1, frame.BowlCount)

		// 二投目
		err = frame.BowlSecond(c.secondNumbers)
		assert.Nil(t, err)
		assert.Equal(t, c.firstScore, frame.FirstScore)
		assert.Equal(t, c.secondScore, frame.SecondScore)
		assert.Equal(t, c.firstRestNumbers, frame.FirstRestPins.GetRestNumbers())
		assert.Equal(t, c.secondRestNumbers, frame.SecondRestPins.GetRestNumbers())
		assert.Equal(t, 2, frame.BowlCount)
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

// Frame

// 通常
func TestBowlFirst_TenFrame(t *testing.T) {
	frame := New()
	err := frame.BowlFirst([]int{1, 2, 3})
	assert.Nil(t, err)
	assert.Equal(t, 3, frame.FirstScore)
	assert.Equal(t, []int{4, 5, 6, 7, 8, 9, 10}, frame.FirstRestPins.GetRestNumbers())
	assert.Equal(t, 1, frame.BowlCount)
}

func TestBowl_Normal_TenFrame(t *testing.T) {
	cases := []struct {
		firstNumbers      []int
		firstScore        int
		firstRestNumbers  []int
		secondNumbers     []int
		secondScore       int
		secondRestNumbers []int
		thirdNumbers      []int
		thirdScore        int
		thirdRestNumbers  []int
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
		{ // 一投目が3本のとき、二投目は7本のとき (スペア)、三投目は3本のとき
			firstNumbers: []int{1, 3, 5}, firstScore: 3, firstRestNumbers: []int{2, 4, 6, 7, 8, 9, 10},
			secondNumbers: []int{2, 4, 6, 7, 8, 9, 10}, secondScore: 7, secondRestNumbers: nil,
			thirdNumbers: []int{1, 2, 3}, thirdScore: 3, thirdRestNumbers: []int{4, 5, 6, 7, 8, 9, 10},
		},
		{ // ダブル 1,2
			firstNumbers: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, firstScore: 10, firstRestNumbers: nil,
			secondNumbers: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, secondScore: 10, secondRestNumbers: nil,
			thirdNumbers: []int{1, 2, 3}, thirdScore: 3, thirdRestNumbers: []int{4, 5, 6, 7, 8, 9, 10},
		},
		{ // スペア、ストライク
			firstNumbers: []int{1, 2, 3, 4, 5}, firstScore: 5, firstRestNumbers: []int{6, 7, 8, 9, 10},
			secondNumbers: []int{6, 7, 8, 9, 10}, secondScore: 5, secondRestNumbers: nil,
			thirdNumbers: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, thirdScore: 10, thirdRestNumbers: nil,
		},
		{ // ストライク、スペア
			firstNumbers: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, firstScore: 10, firstRestNumbers: nil,
			secondNumbers: []int{1, 2, 3, 4, 5}, secondScore: 5, secondRestNumbers: []int{6, 7, 8, 9, 10},
			thirdNumbers: []int{6, 7, 8, 9, 10}, thirdScore: 5, thirdRestNumbers: nil,
		},
		{ // ターキー
			firstNumbers: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, firstScore: 10, firstRestNumbers: nil,
			secondNumbers: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, secondScore: 10, secondRestNumbers: nil,
			thirdNumbers: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, thirdScore: 10, thirdRestNumbers: nil,
		},
	}
	for _, c := range cases {
		frame := New()
		frame.IsForTenFrame = true
		// 一投目
		err := frame.BowlFirst(c.firstNumbers)
		assert.Nil(t, err)
		assert.Equal(t, c.firstScore, frame.FirstScore)
		assert.Equal(t, c.firstRestNumbers, frame.FirstRestPins.GetRestNumbers())
		assert.Equal(t, 1, frame.BowlCount)

		// 二投目
		err = frame.BowlSecond(c.secondNumbers)
		assert.Nil(t, err)
		assert.Equal(t, c.firstScore, frame.FirstScore)
		assert.Equal(t, c.secondScore, frame.SecondScore)
		assert.Equal(t, c.firstRestNumbers, frame.FirstRestPins.GetRestNumbers())
		assert.Equal(t, c.secondRestNumbers, frame.SecondRestPins.GetRestNumbers())
		assert.Equal(t, 2, frame.BowlCount)

		// 三投目
		err = frame.BowlThird(c.thirdNumbers)
		if c.thirdNumbers == nil {
			assert.NotNil(t, err)
			continue
		}
		assert.Nil(t, err)
		assert.Equal(t, c.firstScore, frame.FirstScore)
		assert.Equal(t, c.secondScore, frame.SecondScore)
		assert.Equal(t, c.thirdScore, frame.ThirdScore)
		assert.Equal(t, c.firstRestNumbers, frame.FirstRestPins.GetRestNumbers())
		assert.Equal(t, c.secondRestNumbers, frame.SecondRestPins.GetRestNumbers())
		assert.Equal(t, c.thirdRestNumbers, frame.ThirdRestPins.GetRestNumbers())
		assert.Equal(t, 3, frame.BowlCount)
	}
}

// 異常系: 不正なコールのときにエラーになることをテストする
func TestBowl_Abuseful_TenFrame(t *testing.T) {
	frame := New()
	frame.IsForTenFrame = true
	err := frame.BowlFirst([]int{})
	assert.Nil(t, err)
	err = frame.BowlFirst([]int{})
	assert.NotNil(t, err)
	err = frame.BowlSecond([]int{})
	assert.Nil(t, err)
	err = frame.BowlSecond([]int{})
	assert.NotNil(t, err)
}
