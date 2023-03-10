package logs

import (
	"context"
	ecommon "ethScan-api-MongoDB/common"
	"ethScan-api-MongoDB/httpapi"
	"ethScan-api-MongoDB/marshallers"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
	"math/big"
	"strconv"
	"time"
)

// LogsClient is the client for logs related actions.
type LogsClient struct {
	API *httpapi.APIClient
}

// LogsRequest contains the request parameters for GetLogs.
type LogsRequest struct {
	FromBlock   LogsBlockParam
	ToBlock     LogsBlockParam
	Address     common.Address
	Topics      []common.Hash
	Comparisons []TopicComparison
}

func (req *LogsRequest) toParams() (map[string]string, error) {
	params := make(map[string]string)

	fromBlock, err := req.FromBlock.toParam()
	if err != nil {
		return nil, err
	}
	params["fromBlock"] = fromBlock

	toBlock, err := req.ToBlock.toParam()
	if err != nil {
		return nil, err
	}
	params["toBlock"] = toBlock

	params["address"] = req.Address.String()

	if err := req.addTopicParams(params); err != nil {
		return nil, err
	}

	if err := req.addCompParams(params); err != nil {
		return nil, err
	}

	return params, nil
}

func (req *LogsRequest) addTopicParams(params map[string]string) error {
	if len(req.Topics) > 4 {
		return errors.New("a maximum of 4 topics is allowed")
	}

	for i := range req.Topics {
		key := fmt.Sprintf("topic%d", i)
		params[key] = req.Topics[i].String()
	}

	return nil
}

func (req *LogsRequest) addCompParams(params map[string]string) error {
	for i := range req.Comparisons {
		k, v, err := req.Comparisons[i].toParam()
		if err != nil {
			return err
		}

		params[k] = v
	}

	return nil
}

// LogsBlockParam contain block-related parameters.
type LogsBlockParam struct {
	Number uint64
	Latest bool
}

func (b LogsBlockParam) toParam() (string, error) {
	if b.Latest {
		if b.Number != 0 {
			return "", errors.New("number must not be specified when latest is true for block")
		}

		return "latest", nil
	}

	return strconv.FormatUint(b.Number, 10), nil
}

// TopicComparison contains parameters for comparing two topics.
type TopicComparison struct {
	Topics   [2]uint8
	Operator ComparisonOperator
}

func (c *TopicComparison) toParam() (string, string, error) {
	if c.Topics[1] <= c.Topics[0] {
		return "", "", errors.New("second topic must be greater than first")
	}

	key := fmt.Sprintf("topic%d_%d_opr", c.Topics[0], c.Topics[1])
	val := c.Operator.String()

	return key, val, nil
}

// ComparisonOperator is an enumeration of topic comparison operators.
// ENUM(and,or)
type ComparisonOperator int32

// LogResponse contains information on an ethereum log.
type LogResponse struct {
	Address          common.Address
	BlockNumber      uint64 `etherscan:"blockNumber,hex"`
	Data             []byte
	GasPrice         *big.Int  `etherscan:"gasPrice,hex"`
	GasUsed          *big.Int  `etherscan:"gasUsed,hex"`
	LogIndex         uint32    `etherscan:"logIndex,hex"`
	Timestamp        time.Time `etherscan:"timeStamp,hex"`
	Topics           []common.Hash
	TransactionHash  common.Hash `etherscan:"transactionHash"`
	TransactionIndex uint32      `etherscan:"transactionIndex,hex"`
}

// GetLogs provides an alternative to the native eth_getLogs.
func (c *LogsClient) GetLogs(ctx context.Context, req *LogsRequest) ([]LogResponse, error) {
	params, err := req.toParams()
	if err != nil {
		return nil, err
	}

	rspData, err := c.API.Get(ctx, &httpapi.RequestParams{
		Module: ecommon.LogsModule,
		Action: "getLogs",
		Other:  params,
	})
	if err != nil {
		return nil, err
	}

	var result []LogResponse
	if err := marshallers.UnmarshalResponse(rspData, &result); err != nil {
		return nil, err
	}

	return result, nil
}
