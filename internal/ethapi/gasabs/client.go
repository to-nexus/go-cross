package gasabs

import (
	"context"
	"time"

	"github.com/ethereum/go-ethereum/log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/node"
	"github.com/ethereum/go-ethereum/rpc"
)

const healthyCheckInterval = 10e9

// Client defines typed wrappers for the Ethereum RPC API.
type Client struct {
	Config    node.GasAbstraction
	c         *rpc.Client
	log       log.Logger
	closed    chan struct{}
	isHealthy bool
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
	cli := NewClient(c, cfg)

	t := time.NewTicker(healthyCheckInterval)
	go func() {
		for {
			_, err = cli.healthy(ctx)
			select {
			case <-cli.closed:
				return
			case <-t.C:
			}
		}
	}()
	return cli, err
}

// NewClient creates a client that uses the given RPC client.
func NewClient(c *rpc.Client, cfg node.GasAbstraction) *Client {
	return &Client{cfg, c, log.New("module", "gasabs"), make(chan struct{}), true}
}

// Close closes the underlying RPC connection.
func (ec *Client) Close() {
	ec.c.Close()
	close(ec.closed)
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

// IsApproved
func (ec *Client) IsApproved(ctx context.Context, contract common.Address) (bool, error) {
	var resp bool
	if err := ec.c.CallContext(ctx, &resp, "gasabs_isApproved", contract); err != nil {
		return false, err
	}
	return resp, nil
}

func (ec *Client) healthy(ctx context.Context) (bool, error) {
	var ok bool
	err := ec.c.CallContext(ctx, &ok, "gasabs_healthy")
	if err != nil {
		ec.log.Error("Failed to request healty check GasAbstraction", "error", err, "ok", ok)
	}
	if ec.isHealthy != ok {
		ec.log.Info("The state of isHealthy has changed", "health", ok)
	}
	ec.isHealthy = ok
	return ok, err
}

// Host
func (ec *Client) Host() string {
	return ec.Config.GasAbsURL
}
