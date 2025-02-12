package gasabs

import (
	"context"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/node"
	"github.com/ethereum/go-ethereum/rpc"
)

// Client defines typed wrappers for the Ethereum RPC API.
type Client struct {
	Config node.GasAbstraction
	c      *rpc.Client
}

// Dial connects a client to the given URL.
func Dial(cfg node.GasAbstraction) (*Client, error) {
	return DialContext(context.Background(), cfg)
}

// DialContext connects a client to the given URL with context.
func DialContext(ctx context.Context, cfg node.GasAbstraction) (*Client, error) {
	c, err := rpc.DialContext(ctx, cfg.GasAbsURL)
	if err != nil {
		return nil, err
	}
	return NewClient(c, cfg), nil
}

// NewClient creates a client that uses the given RPC client.
func NewClient(c *rpc.Client, cfg node.GasAbstraction) *Client {
	return &Client{cfg, c}
}

// Close closes the underlying RPC connection.
func (ec *Client) Close() {
	ec.c.Close()
}

// SignFeeDelegateTransaction
func (ec *Client) SignFeeDelegateTransaction(ctx context.Context, tx *types.Transaction) (*types.Transaction, error) {
	var resp hexutil.Bytes

	data, err := tx.MarshalBinary()
	if err != nil {
		return tx, err
	}

	if err := ec.c.CallContext(ctx, &resp, "gasabs_signFeeDelegateTransaction", hexutil.Encode(data)); err != nil {
		return tx, err
	} else {
		respTx := new(types.Transaction)
		if errInner := respTx.UnmarshalBinary(resp); errInner != nil {
			return tx, errInner
		}
		return respTx, err
	}
}

func (ec *Client) IsApproved(ctx context.Context, contract common.Address) (bool, error) {
	var resp bool
	if err := ec.c.CallContext(ctx, &resp, "gasabs_isApproved", contract); err != nil {
		return false, err
	}
	return resp, nil
}

func (ec *Client) Host() string {
	return ec.Config.GasAbsURL
}
