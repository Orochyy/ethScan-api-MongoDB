package common

import "fmt"

const (
	// BlockParameterLatest is a BlockParameter of type Latest.
	BlockParameterLatest BlockParameter = iota
	// BlockParameterEarliest is a BlockParameter of type Earliest.
	BlockParameterEarliest
	// BlockParameterPending is a BlockParameter of type Pending.
	BlockParameterPending
)

const _BlockParameterName = "latestearliestpending"

var _BlockParameterMap = map[BlockParameter]string{
	0: _BlockParameterName[0:6],
	1: _BlockParameterName[6:14],
	2: _BlockParameterName[14:21],
}

// String implements the Stringer interface.
func (x BlockParameter) String() string {
	if str, ok := _BlockParameterMap[x]; ok {
		return str
	}
	return fmt.Sprintf("BlockParameter(%d)", x)
}

var _BlockParameterValue = map[string]BlockParameter{
	_BlockParameterName[0:6]:   0,
	_BlockParameterName[6:14]:  1,
	_BlockParameterName[14:21]: 2,
}

// ParseBlockParameter attempts to convert a string to a BlockParameter
func ParseBlockParameter(name string) (BlockParameter, error) {
	if x, ok := _BlockParameterValue[name]; ok {
		return x, nil
	}
	return BlockParameter(0), fmt.Errorf("%s is not a valid BlockParameter", name)
}

const (
	// SortingPreferenceAsc is a SortingPreference of type Asc.
	SortingPreferenceAsc SortingPreference = iota
	// SortingPreferenceDesc is a SortingPreference of type Desc.
	SortingPreferenceDesc
)

const _SortingPreferenceName = "ascdesc"

var _SortingPreferenceMap = map[SortingPreference]string{
	0: _SortingPreferenceName[0:3],
	1: _SortingPreferenceName[3:7],
}

// String implements the Stringer interface.
func (x SortingPreference) String() string {
	if str, ok := _SortingPreferenceMap[x]; ok {
		return str
	}
	return fmt.Sprintf("SortingPreference(%d)", x)
}

var _SortingPreferenceValue = map[string]SortingPreference{
	_SortingPreferenceName[0:3]: 0,
	_SortingPreferenceName[3:7]: 1,
}

// ParseSortingPreference attempts to convert a string to a SortingPreference
func ParseSortingPreference(name string) (SortingPreference, error) {
	if x, ok := _SortingPreferenceValue[name]; ok {
		return x, nil
	}
	return SortingPreference(0), fmt.Errorf("%s is not a valid SortingPreference", name)
}
