// Code generated by "stringer -type=TestDummyEnum"; DO NOT EDIT.

package igdb

import "strconv"

const _TestDummyEnum_name = "TestDummyEnum1TestDummyEnum2"

var _TestDummyEnum_index = [...]uint8{0, 14, 28}

func (i TestDummyEnum) String() string {
	i -= 1
	if i < 0 || i >= TestDummyEnum(len(_TestDummyEnum_index)-1) {
		return "TestDummyEnum(" + strconv.FormatInt(int64(i+1), 10) + ")"
	}
	return _TestDummyEnum_name[_TestDummyEnum_index[i]:_TestDummyEnum_index[i+1]]
}
