// Copyright 2017 gf Author(https://gitee.com/johng/gf). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://gitee.com/johng/gf.

package gconv

import "fmt"

// 任意类型转换为[]int类型
func Ints(i interface{}) []int {
	if i == nil {
		return nil
	}
	if r, ok := i.([]int); ok {
		return r
	} else {
		array := make([]int, 0)
		switch i.(type) {
		case []string:
			for _, v := range i.([]string) {
				array = append(array, Int(v))
			}
		case []int8:
			for _, v := range i.([]int8) {
				array = append(array, Int(v))
			}
		case []int16:
			for _, v := range i.([]int16) {
				array = append(array, Int(v))
			}
		case []int32:
			for _, v := range i.([]int32) {
				array = append(array, Int(v))
			}
		case []int64:
			for _, v := range i.([]int64) {
				array = append(array, Int(v))
			}
		case []uint:
			for _, v := range i.([]uint) {
				array = append(array, Int(v))
			}
		case []uint8:
			for _, v := range i.([]uint8) {
				array = append(array, Int(v))
			}
		case []uint16:
			for _, v := range i.([]uint16) {
				array = append(array, Int(v))
			}
		case []uint32:
			for _, v := range i.([]uint32) {
				array = append(array, Int(v))
			}
		case []uint64:
			for _, v := range i.([]uint64) {
				array = append(array, Int(v))
			}
		case []bool:
			for _, v := range i.([]bool) {
				array = append(array, Int(v))
			}
		case []float32:
			for _, v := range i.([]float32) {
				array = append(array, Int(v))
			}
		case []float64:
			for _, v := range i.([]float64) {
				array = append(array, Int(v))
			}
		case []interface{}:
			for _, v := range i.([]interface{}) {
				array = append(array, Int(v))
			}
		}
		if len(array) > 0 {
			return array
		}
	}
	return []int{Int(i)}
}

// 任意类型转换为[]string类型
func Strings(i interface{}) []string {
	if i == nil {
		return nil
	}
	if r, ok := i.([]string); ok {
		return r
	} else {
		array := make([]string, 0)
		switch i.(type) {
		case []int:
			for _, v := range i.([]int) {
				array = append(array, String(v))
			}
		case []int8:
			for _, v := range i.([]int8) {
				array = append(array, String(v))
			}
		case []int16:
			for _, v := range i.([]int16) {
				array = append(array, String(v))
			}
		case []int32:
			for _, v := range i.([]int32) {
				array = append(array, String(v))
			}
		case []int64:
			for _, v := range i.([]int64) {
				array = append(array, String(v))
			}
		case []uint:
			for _, v := range i.([]uint) {
				array = append(array, String(v))
			}
		case []uint8:
			for _, v := range i.([]uint8) {
				array = append(array, String(v))
			}
		case []uint16:
			for _, v := range i.([]uint16) {
				array = append(array, String(v))
			}
		case []uint32:
			for _, v := range i.([]uint32) {
				array = append(array, String(v))
			}
		case []uint64:
			for _, v := range i.([]uint64) {
				array = append(array, String(v))
			}
		case []bool:
			for _, v := range i.([]bool) {
				array = append(array, String(v))
			}
		case []float32:
			for _, v := range i.([]float32) {
				array = append(array, String(v))
			}
		case []float64:
			for _, v := range i.([]float64) {
				array = append(array, String(v))
			}
		case []interface{}:
			for _, v := range i.([]interface{}) {
				array = append(array, String(v))
			}
		}
		if len(array) > 0 {
			return array
		}
	}
	return []string{fmt.Sprintf("%v", i)}
}

// 任意类型转换为[]float64类型
func Floats(i interface{}) []float64 {
	if i == nil {
		return nil
	}
	if r, ok := i.([]float64); ok {
		return r
	} else {
		array := make([]float64, 0)
		switch i.(type) {
		case []string:
			for _, v := range i.([]string) {
				array = append(array, Float64(v))
			}
		case []int:
			for _, v := range i.([]int) {
				array = append(array, Float64(v))
			}
		case []int8:
			for _, v := range i.([]int8) {
				array = append(array, Float64(v))
			}
		case []int16:
			for _, v := range i.([]int16) {
				array = append(array, Float64(v))
			}
		case []int32:
			for _, v := range i.([]int32) {
				array = append(array, Float64(v))
			}
		case []int64:
			for _, v := range i.([]int64) {
				array = append(array, Float64(v))
			}
		case []uint:
			for _, v := range i.([]uint) {
				array = append(array, Float64(v))
			}
		case []uint8:
			for _, v := range i.([]uint8) {
				array = append(array, Float64(v))
			}
		case []uint16:
			for _, v := range i.([]uint16) {
				array = append(array, Float64(v))
			}
		case []uint32:
			for _, v := range i.([]uint32) {
				array = append(array, Float64(v))
			}
		case []uint64:
			for _, v := range i.([]uint64) {
				array = append(array, Float64(v))
			}
		case []bool:
			for _, v := range i.([]bool) {
				array = append(array, Float64(v))
			}
		case []float32:
			for _, v := range i.([]float32) {
				array = append(array, Float64(v))
			}
		case []interface{}:
			for _, v := range i.([]interface{}) {
				array = append(array, Float64(v))
			}
		}
		if len(array) > 0 {
			return array
		}
	}
	return []float64{Float64(i)}
}

// 任意类型转换为[]interface{}类型
func Interfaces(i interface{}) []interface{} {
	if i == nil {
		return nil
	}
	if r, ok := i.([]interface{}); ok {
		return r
	} else {
		array := make([]interface{}, 0)
		switch i.(type) {
		case []string:
			for _, v := range i.([]string) {
				array = append(array, v)
			}
		case []int:
			for _, v := range i.([]int) {
				array = append(array, v)
			}
		case []int8:
			for _, v := range i.([]int8) {
				array = append(array, v)
			}
		case []int16:
			for _, v := range i.([]int16) {
				array = append(array, v)
			}
		case []int32:
			for _, v := range i.([]int32) {
				array = append(array, v)
			}
		case []int64:
			for _, v := range i.([]int64) {
				array = append(array, v)
			}
		case []uint:
			for _, v := range i.([]uint) {
				array = append(array, v)
			}
		case []uint8:
			for _, v := range i.([]uint8) {
				array = append(array, v)
			}
		case []uint16:
			for _, v := range i.([]uint16) {
				array = append(array, v)
			}
		case []uint32:
			for _, v := range i.([]uint32) {
				array = append(array, v)
			}
		case []uint64:
			for _, v := range i.([]uint64) {
				array = append(array, v)
			}
		case []bool:
			for _, v := range i.([]bool) {
				array = append(array, v)
			}
		case []float32:
			for _, v := range i.([]float32) {
				array = append(array, v)
			}
		case []float64:
			for _, v := range i.([]float64) {
				array = append(array, v)
			}
		}
		if len(array) > 0 {
			return array
		}
	}
	return []interface{}{i}
}
