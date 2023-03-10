package stats

import "fmt"

const (
	// ETHClientTypeReqGeth is a ETHClientTypeReq of type Geth.
	ETHClientTypeReqGeth ETHClientTypeReq = iota
	// ETHClientTypeReqParity is a ETHClientTypeReq of type Parity.
	ETHClientTypeReqParity
)

const _ETHClientTypeReqName = "gethparity"

var _ETHClientTypeReqMap = map[ETHClientTypeReq]string{
	0: _ETHClientTypeReqName[0:4],
	1: _ETHClientTypeReqName[4:10],
}

// String implements the Stringer interface.
func (x ETHClientTypeReq) String() string {
	if str, ok := _ETHClientTypeReqMap[x]; ok {
		return str
	}
	return fmt.Sprintf("ETHClientTypeReq(%d)", x)
}

var _ETHClientTypeReqValue = map[string]ETHClientTypeReq{
	_ETHClientTypeReqName[0:4]:  0,
	_ETHClientTypeReqName[4:10]: 1,
}

// ParseETHClientTypeReq attempts to convert a string to a ETHClientTypeReq
func ParseETHClientTypeReq(name string) (ETHClientTypeReq, error) {
	if x, ok := _ETHClientTypeReqValue[name]; ok {
		return x, nil
	}
	return ETHClientTypeReq(0), fmt.Errorf("%s is not a valid ETHClientTypeReq", name)
}

// MarshalText implements the text marshaller method
func (x ETHClientTypeReq) MarshalText() ([]byte, error) {
	return []byte(x.String()), nil
}

// UnmarshalText implements the text unmarshaller method
func (x *ETHClientTypeReq) UnmarshalText(text []byte) error {
	name := string(text)
	tmp, err := ParseETHClientTypeReq(name)
	if err != nil {
		return err
	}
	*x = tmp
	return nil
}

const (
	// ETHClientTypeResultGeth is a ETHClientTypeResult of type Geth.
	ETHClientTypeResultGeth ETHClientTypeResult = iota
	// ETHClientTypeResultParity is a ETHClientTypeResult of type Parity.
	ETHClientTypeResultParity
)

const _ETHClientTypeResultName = "GethParity"

var _ETHClientTypeResultMap = map[ETHClientTypeResult]string{
	0: _ETHClientTypeResultName[0:4],
	1: _ETHClientTypeResultName[4:10],
}

// String implements the Stringer interface.
func (x ETHClientTypeResult) String() string {
	if str, ok := _ETHClientTypeResultMap[x]; ok {
		return str
	}
	return fmt.Sprintf("ETHClientTypeResult(%d)", x)
}

var _ETHClientTypeResultValue = map[string]ETHClientTypeResult{
	_ETHClientTypeResultName[0:4]:  0,
	_ETHClientTypeResultName[4:10]: 1,
}

// ParseETHClientTypeResult attempts to convert a string to a ETHClientTypeResult
func ParseETHClientTypeResult(name string) (ETHClientTypeResult, error) {
	if x, ok := _ETHClientTypeResultValue[name]; ok {
		return x, nil
	}
	return ETHClientTypeResult(0), fmt.Errorf("%s is not a valid ETHClientTypeResult", name)
}

// MarshalText implements the text marshaller method
func (x ETHClientTypeResult) MarshalText() ([]byte, error) {
	return []byte(x.String()), nil
}

// UnmarshalText implements the text unmarshaller method
func (x *ETHClientTypeResult) UnmarshalText(text []byte) error {
	name := string(text)
	tmp, err := ParseETHClientTypeResult(name)
	if err != nil {
		return err
	}
	*x = tmp
	return nil
}

const (
	// NodeSyncModeReqDefault is a NodeSyncModeReq of type Default.
	NodeSyncModeReqDefault NodeSyncModeReq = iota
	// NodeSyncModeReqArchive is a NodeSyncModeReq of type Archive.
	NodeSyncModeReqArchive
)

const _NodeSyncModeReqName = "defaultarchive"

var _NodeSyncModeReqMap = map[NodeSyncModeReq]string{
	0: _NodeSyncModeReqName[0:7],
	1: _NodeSyncModeReqName[7:14],
}

// String implements the Stringer interface.
func (x NodeSyncModeReq) String() string {
	if str, ok := _NodeSyncModeReqMap[x]; ok {
		return str
	}
	return fmt.Sprintf("NodeSyncModeReq(%d)", x)
}

var _NodeSyncModeReqValue = map[string]NodeSyncModeReq{
	_NodeSyncModeReqName[0:7]:  0,
	_NodeSyncModeReqName[7:14]: 1,
}

// ParseNodeSyncModeReq attempts to convert a string to a NodeSyncModeReq
func ParseNodeSyncModeReq(name string) (NodeSyncModeReq, error) {
	if x, ok := _NodeSyncModeReqValue[name]; ok {
		return x, nil
	}
	return NodeSyncModeReq(0), fmt.Errorf("%s is not a valid NodeSyncModeReq", name)
}

// MarshalText implements the text marshaller method
func (x NodeSyncModeReq) MarshalText() ([]byte, error) {
	return []byte(x.String()), nil
}

// UnmarshalText implements the text unmarshaller method
func (x *NodeSyncModeReq) UnmarshalText(text []byte) error {
	name := string(text)
	tmp, err := ParseNodeSyncModeReq(name)
	if err != nil {
		return err
	}
	*x = tmp
	return nil
}

const (
	// NodeSyncModeResultDefault is a NodeSyncModeResult of type Default.
	NodeSyncModeResultDefault NodeSyncModeResult = iota
	// NodeSyncModeResultArchive is a NodeSyncModeResult of type Archive.
	NodeSyncModeResultArchive
)

const _NodeSyncModeResultName = "DefaultArchive"

var _NodeSyncModeResultMap = map[NodeSyncModeResult]string{
	0: _NodeSyncModeResultName[0:7],
	1: _NodeSyncModeResultName[7:14],
}

// String implements the Stringer interface.
func (x NodeSyncModeResult) String() string {
	if str, ok := _NodeSyncModeResultMap[x]; ok {
		return str
	}
	return fmt.Sprintf("NodeSyncModeResult(%d)", x)
}

var _NodeSyncModeResultValue = map[string]NodeSyncModeResult{
	_NodeSyncModeResultName[0:7]:  0,
	_NodeSyncModeResultName[7:14]: 1,
}

// ParseNodeSyncModeResult attempts to convert a string to a NodeSyncModeResult
func ParseNodeSyncModeResult(name string) (NodeSyncModeResult, error) {
	if x, ok := _NodeSyncModeResultValue[name]; ok {
		return x, nil
	}
	return NodeSyncModeResult(0), fmt.Errorf("%s is not a valid NodeSyncModeResult", name)
}

// MarshalText implements the text marshaller method
func (x NodeSyncModeResult) MarshalText() ([]byte, error) {
	return []byte(x.String()), nil
}

// UnmarshalText implements the text unmarshaller method
func (x *NodeSyncModeResult) UnmarshalText(text []byte) error {
	name := string(text)
	tmp, err := ParseNodeSyncModeResult(name)
	if err != nil {
		return err
	}
	*x = tmp
	return nil
}
