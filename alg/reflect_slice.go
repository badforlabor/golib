/**
 * Auth :   liubo
 * Date :   2019/12/3 16:10
 * Comment: 各种反射的应用
 */

package alg


import "reflect"

func isSlice(dataList interface{}) bool {
	s := reflect.ValueOf(dataList)

	if s.Kind() == reflect.Ptr {
		s = s.Elem()
	}

	if s.Kind() == reflect.Slice {
		return true
	}
	return false
}

func reflectSlice(dataList interface{}) reflect.Value {
	s := reflect.ValueOf(dataList)

	if s.Kind() == reflect.Ptr {
		s = s.Elem()
	}

	if s.Kind() != reflect.Slice {
		panic("not slice")
	}

	return s
}

// 拆箱
func ToSlice(dataList interface{}) []interface{} {
	s := reflectSlice(dataList)

	ret := make([]interface{}, s.Len())
	for i:=0; i<s.Len(); i++ {
		ret[i] = s.Index(i).Interface()
	}
	return ret
}

// 再次拆箱
func SetSliceValue(src []interface{}, dst interface{}) {
	s := reflectSlice(dst)

	if s.Len() != len(src) {
		panic("not equal")
	}

	for i:=0; i<s.Len(); i++ {
		s.Index(i).Set(reflect.ValueOf(src[i]))
	}
}

// 删除
func SliceRemoveAt(dataList interface{}, idx int) {
	mustPtr(dataList)

	if idx < 0 {
		panic("invalid index")
	}

	s := reflectSlice(dataList)
	if s.Len() == 0 {
		return
	}

	for i:=idx; i<s.Len()-1; i++ {
		s.Index(i).Set(s.Index(i+1))
	}

	//s.SetLen(s.Len() - 1)
	s.SetLen(s.Len() - 1)
}

func SliceAddOne(dataList, one interface{}) {
	mustPtr(dataList)

	SliceInsertOne(dataList, -1, one)
}
func SliceInsertOne(dataList interface{}, idx int, one interface{}) {
	mustPtr(dataList)

	if idx < 0 && idx != -1 {
		panic("invalid index")
	}

	s := reflectSlice(dataList)

	if s.Len()+1 > s.Cap() {
		dst := reflect.MakeSlice(s.Type(), s.Len() + 1, s.Cap() * 2)
		reflect.Copy(dst, s)
		s.Set(dst)
	} else {
		s.SetLen(s.Len() + 1)
	}

	if idx == -1{
		idx = s.Len() - 1
	} else {
		for i:=s.Len()-1; i>idx; i-- {
			s.Index(i).Set(s.Index(i-1))
		}
	}
	s.Index(idx).Set(reflect.ValueOf(one))
}
func EraseOne(dataList interface{}, compare func(d interface{})bool) {
	idx := Find(dataList, compare)
	if idx >= 0 {
		SliceRemoveAt(dataList, idx)
	}
}
func AddUniqueOne(dataList, data interface{}, compare func(d interface{})bool) {
	if Contain(dataList, compare) {
		return
	}
	SliceAddOne(dataList, data)
}