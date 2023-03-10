package accounts

import (
	"fmt"
)

const (
	// BlockTypeBlocks is a BlockType of type Blocks.
	BlockTypeBlocks BlockType = iota
	// BlockTypeUncles is a BlockType of type Uncles.
	BlockTypeUncles
)

const _BlockTypeName = "blocksuncles"

var _BlockTypeMap = map[BlockType]string{
	0: _BlockTypeName[0:6],
	1: _BlockTypeName[6:12],
}

// String implements the Stringer interface.
func (x BlockType) String() string {
	if str, ok := _BlockTypeMap[x]; ok {
		return str
	}
	return fmt.Sprintf("BlockType(%d)", x)
}

var _BlockTypeValue = map[string]BlockType{
	_BlockTypeName[0:6]:  0,
	_BlockTypeName[6:12]: 1,
}

// ParseBlockType attempts to convert a string to a BlockType
func ParseBlockType(name string) (BlockType, error) {
	if x, ok := _BlockTypeValue[name]; ok {
		return x, nil
	}
	return BlockType(0), fmt.Errorf("%s is not a valid BlockType", name)
}
