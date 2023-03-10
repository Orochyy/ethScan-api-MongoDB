package logs

import "fmt"

const (
	// ComparisonOperatorAnd is a ComparisonOperator of type And.
	ComparisonOperatorAnd ComparisonOperator = iota
	// ComparisonOperatorOr is a ComparisonOperator of type Or.
	ComparisonOperatorOr
)

const _ComparisonOperatorName = "andor"

var _ComparisonOperatorMap = map[ComparisonOperator]string{
	0: _ComparisonOperatorName[0:3],
	1: _ComparisonOperatorName[3:5],
}

// String implements the Stringer interface.
func (x ComparisonOperator) String() string {
	if str, ok := _ComparisonOperatorMap[x]; ok {
		return str
	}
	return fmt.Sprintf("ComparisonOperator(%d)", x)
}

var _ComparisonOperatorValue = map[string]ComparisonOperator{
	_ComparisonOperatorName[0:3]: 0,
	_ComparisonOperatorName[3:5]: 1,
}

// ParseComparisonOperator attempts to convert a string to a ComparisonOperator
func ParseComparisonOperator(name string) (ComparisonOperator, error) {
	if x, ok := _ComparisonOperatorValue[name]; ok {
		return x, nil
	}
	return ComparisonOperator(0), fmt.Errorf("%s is not a valid ComparisonOperator", name)
}
