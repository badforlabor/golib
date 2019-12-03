/**
 * Auth :   liubo
 * Date :   2019/12/3 15:51
 * Comment: 查找
 */

package alg

import "reflect"

func Find(dataList interface{}, compare func(d interface{})bool) int {
	s := reflect.ValueOf(dataList)
	if s.Kind() != reflect.Slice {
		return -2
	}
	return FindList(ToSlice(dataList), compare)
}
func FindList(dataList []interface{}, compare func(d interface{})bool) int {
	for i:=0; i<len(dataList); i++ {
		if compare(dataList[i]) {
			return i
		}
	}
	return -1
}

func FindInt(a int) func(d interface{})bool {
	var function = func(d interface{})bool {
		var aa = d.(int)
		return a == aa
	}
	return function
}
