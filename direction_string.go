// Code generated by "stringer -type=Direction"; DO NOT EDIT

package sirpent

import "fmt"

const _Direction_name = "NorthEastEastSouthEastSouthWestWestNorthWest"

var _Direction_index = [...]uint8{0, 9, 13, 22, 31, 35, 44}

func (i Direction) String() string {
	if i < 0 || i >= Direction(len(_Direction_index)-1) {
		return fmt.Sprintf("Direction(%d)", i)
	}
	return _Direction_name[_Direction_index[i]:_Direction_index[i+1]]
}
