package gasabs

import (
	"context"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rpc"
)

const syncInterval = 30 * time.Second

// Client defines typed wrappers for the Ethereum RPC API.
type Client struct {
	ctx    context.Context
	cancel context.CancelFunc
	rawurl string
	c      *rpc.Client
	log    log.Logger

	whitelistLock    sync.RWMutex
	whitelistVersion common.Hash
	whitelistFrom    map[common.Address]struct{}
	whitelistTo      map[common.Address]struct{}
}

// Dial connects a client to the given URL.
func Dial(rawurl string) (*Client, error) {
	return DialContext(context.Background(), rawurl)
}

// DialContext connects a client to the given URL with context.
func DialContext(ctx context.Context, rawurl string) (*Client, error) {
	c, err := rpc.DialContext(ctx, rawurl)
	if err != nil {
		return nil, err
	}
	cli := NewClient(ctx, c, rawurl)

	return cli, err
}

// NewClient creates a client that uses the given RPC client.
func NewClient(ctx context.Context, c *rpc.Client, rawurl string) *Client {
	var (
		chlid, cancel = context.WithCancel(ctx)
		client        = &Client{
			ctx:    chlid,
			cancel: cancel,
			rawurl: rawurl,
			c:      c,
			log:    log.New("module", "gasabs"),

			whitelistLock:    sync.RWMutex{},
			whitelistVersion: common.Hash{},
			whitelistFrom:    make(map[common.Address]struct{}),
			whitelistTo:      make(map[common.Address]struct{}),
		}
	)
	go client.sync()
	return client
}

func (ec *Client) sync() {
	ticker := time.NewTicker(syncInterval)
	for {
		ctx, cancel := context.WithTimeout(ec.ctx, syncInterval)
		err := ec.syncWhitelist(ctx)
		if err != nil {
			ec.log.Error("Failed to request whitelist", "error", err)
		}
		cancel()

		select {
		case <-ec.ctx.Done():
			return
		case <-ticker.C:
		}
	}
}

// Close closes the underlying RPC connection.
func (ec *Client) Close() {
	ec.cancel()
	ec.c.Close()
}

// ##CROSS: fee delegation
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
func (ec *Client) IsApprovedFrom(ctx context.Context, address common.Address) (approved bool) {
	ec.whitelistLock.RLock()
	_, approved = ec.whitelistFrom[address]
	ec.whitelistLock.RUnlock()
	return
}

func (ec *Client) IsApprovedTo(ctx context.Context, address common.Address) (approved bool) {
	ec.whitelistLock.RLock()
	_, approved = ec.whitelistTo[address]
	ec.whitelistLock.RUnlock()
	return
}

type whitelistResp struct {
	Version common.Hash
	From    []common.Address
	To      []common.Address
}

func (ec *Client) syncWhitelist(ctx context.Context) error {
	var whitelist whitelistResp
	err := ec.c.CallContext(ctx, &whitelist, "gasabs_whitelist", ec.whitelistVersion)
	if err != nil {
		return err
	}

	if whitelist.Version == ec.whitelistVersion {
		return nil
	}

	whitelistFrom := make(map[common.Address]struct{})
	whitelistTo := make(map[common.Address]struct{})
	for _, from := range whitelist.From {
		whitelistFrom[from] = struct{}{}
	}
	for _, to := range whitelist.To {
		whitelistTo[to] = struct{}{}
	}

	ec.whitelistLock.Lock()
	ec.whitelistVersion = whitelist.Version
	ec.whitelistFrom = whitelistFrom
	ec.whitelistTo = whitelistTo
	ec.whitelistLock.Unlock()

	ec.log.Info("Gasabs whitelist updated", "version", ec.whitelistVersion, "from_length", len(ec.whitelistFrom), "to_length", len(ec.whitelistTo))
	return nil
}

// Host
func (ec *Client) Host() string {
	return ec.rawurl
}
