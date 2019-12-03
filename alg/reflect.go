/**
 * Auth :   liubo
 * Date :   2019/12/3 16:10
 * Comment: 各种反射的应用
 */

package alg

import "reflect"

// 拆箱
func ToSlice(dataList interface{}) []interface{} {
	s := reflect.ValueOf(dataList)

	if s.Kind() != reflect.Slice {
		panic("not slice")
	}

	ret := make([]interface{}, s.Len())
	for i:=0; i<s.Len(); i++ {
		ret[i] = s.Index(i).Interface()
	}
	return ret
}

// 再次拆箱
func SetSliceValue(src []interface{}, dst interface{}) {
	s := reflect.ValueOf(dst)

	if s.Kind() != reflect.Slice {
		panic("not slice")
	}
	if s.Len() != len(src) {
		panic("not equal")
	}

	for i:=0; i<s.Len(); i++ {
		s.Index(i).Set(reflect.ValueOf(src[i]))
	}
}