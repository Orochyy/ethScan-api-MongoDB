package blocks

import "fmt"

const (
	// ClosestAvailableBlockBefore is a ClosestAvailableBlock of type Before.
	ClosestAvailableBlockBefore ClosestAvailableBlock = iota
	// ClosestAvailableBlockAfter is a ClosestAvailableBlock of type After.
	ClosestAvailableBlockAfter
)

const _ClosestAvailableBlockName = "beforeafter"

var _ClosestAvailableBlockMap = map[ClosestAvailableBlock]string{
	0: _ClosestAvailableBlockName[0:6],
	1: _ClosestAvailableBlockName[6:11],
}

// String implements the Stringer interface.
func (x ClosestAvailableBlock) String() string {
	if str, ok := _ClosestAvailableBlockMap[x]; ok {
		return str
	}
	return fmt.Sprintf("ClosestAvailableBlock(%d)", x)
}

var _ClosestAvailableBlockValue = map[string]ClosestAvailableBlock{
	_ClosestAvailableBlockName[0:6]:  0,
	_ClosestAvailableBlockName[6:11]: 1,
}

// ParseClosestAvailableBlock attempts to convert a string to a ClosestAvailableBlock
func ParseClosestAvailableBlock(name string) (ClosestAvailableBlock, error) {
	if x, ok := _ClosestAvailableBlockValue[name]; ok {
		return x, nil
	}
	return ClosestAvailableBlock(0), fmt.Errorf("%s is not a valid ClosestAvailableBlock", name)
}
