// Copyright (c) 2021-2022. Marvin Hansen | marvin.hansen@gmail.com

package math_utils

import "errors"

func GetMinValue(values []float64) (min float64, minIdx int, err error) {
	if len(values) == 0 {
		return 0, 0, errors.New("cannot detect a minimum value in an empty slice")
	}
	min = values[0]
	for i, v := range values {
		if v < min {
			min = v
			minIdx = i
		}
	}
	return min, minIdx, nil
}

func GetMaxValue(values []float64) (max float64, maxIdx int, err error) {
	if len(values) == 0 {
		return 0, 0, errors.New("cannot detect a minimum value in an empty slice")
	}
	max = values[0]
	for i, v := range values {
		if v > max {
			max = v
			maxIdx = i
		}
	}
	return max, maxIdx, nil
}
