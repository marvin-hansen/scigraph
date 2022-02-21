// Copyright (c) 2021. Marvin Friedrich Lars Hansen. All Rights Reserved. Contact: marvin.hansen@gmail.com

package collection_types

import (
	"femto/src/types/imx/imx_types"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
)

func getBar() imx_types.BarData {

	return imx_types.BarData{
		Ticker:   "TEST",
		Date:     "TODAY",
		Open:     rand.Float64(),
		High:     rand.Float64(),
		Low:      rand.Float64(),
		Close:    rand.Float64(),
		Volume:   0,
		BarCount: 0,
		Average:  0,
	}
}

func TestFilled(t *testing.T) {

	w := NewRollingWindow[imx_types.BarData](3, 4)

	b1 := getBar()
	b2 := getBar()
	w.PushBack(b1)
	w.PushBack(b2)

	var actual = w.Filled()
	var expected = false
	assert.Equal(t, expected, actual, "Should be equal")

	b3 := getBar()
	b4 := getBar()

	w.PushBack(b3)
	w.PushBack(b4)
	actual = w.Filled()
	expected = true
	assert.Equal(t, expected, actual, "Should be equal")

}
