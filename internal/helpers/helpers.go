package helpers

import "sort"

func Unique(arr []int) []int {
	var m = make(map[int]bool)
	var newArr []int
	for _, valueArr := range arr {
		if _, value := m[valueArr]; !value {
			m[valueArr] = true
			newArr = append(newArr, valueArr)
		}
	}
	sort.Ints(newArr)
	return newArr
}
func UniqueStringArray(arr []string) []string {
	var m = make(map[string]bool)
	var newArr []string
	for _, valueArr := range arr {
		if _, value := m[valueArr]; !value {
			m[valueArr] = true
			newArr = append(newArr, valueArr)
		}
	}
	return newArr
}
