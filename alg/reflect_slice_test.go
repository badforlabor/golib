/**
 * Auth :   liubo
 * Date :   2019/12/6 17:56
 * Comment:
 */

package alg

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSlice1(t *testing.T) {
	var datalist = []int{0,1,2,3,4,5}

	SliceRemoveAt(&datalist, 1)
	assert.Equal(t, len(datalist), 5, "")

	EraseOne(&datalist, IntCompare(0))
	assert.Equal(t, datalist[0], 2, "")

	SliceAddOne(&datalist, 6)

	SliceInsertOne(&datalist, 0, 0)
	SliceInsertOne(&datalist, 1, 1)

	AddUniqueOne(&datalist, 6, IntCompare(6))
	AddUniqueOne(&datalist, 7, IntCompare(7))

	assert.Equal(t, len(datalist), 8, "")

	for i:=len(datalist); i<32; i++ {
		AddUniqueOne(&datalist, i, IntCompare(i))
	}

	for i, v := range datalist {
		assert.Equal(t, v, i, "")
	}
}

