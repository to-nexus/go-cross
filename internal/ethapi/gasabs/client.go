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

	accesslistLock    sync.RWMutex
	accesslistVersion common.Hash
	whitelistFrom     map[common.Address]struct{}
	whitelistTo       map[common.Address]struct{}
	blacklistFrom     map[common.Address]struct{}
	blacklistTo       map[common.Address]struct{}
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

			accesslistLock:    sync.RWMutex{},
			accesslistVersion: common.Hash{},
			whitelistFrom:     make(map[common.Address]struct{}),
			whitelistTo:       make(map[common.Address]struct{}),
			blacklistFrom:     make(map[common.Address]struct{}),
			blacklistTo:       make(map[common.Address]struct{}),
		}
	)
	go client.sync()
	return client
}

func (ec *Client) sync() {
	ticker := time.NewTicker(syncInterval)
	for {
		ctx, cancel := context.WithTimeout(ec.ctx, syncInterval)
		err := ec.syncAccessList(ctx)
		if err != nil {
			ec.log.Error("Failed to request accesslist", "error", err)
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
	ec.accesslistLock.RLock()
	_, approved = ec.whitelistFrom[address]
	ec.accesslistLock.RUnlock()
	return
}

func (ec *Client) IsApprovedTo(ctx context.Context, address common.Address) (approved bool) {
	ec.accesslistLock.RLock()
	_, approved = ec.whitelistTo[address]
	ec.accesslistLock.RUnlock()
	return
}

func (ec *Client) IsBlacklistedFrom(ctx context.Context, address common.Address) (blacklisted bool) {
	ec.accesslistLock.RLock()
	_, blacklisted = ec.blacklistFrom[address]
	ec.accesslistLock.RUnlock()
	return
}

func (ec *Client) IsBlacklistedTo(ctx context.Context, address common.Address) (blacklisted bool) {
	ec.accesslistLock.RLock()
	_, blacklisted = ec.blacklistTo[address]
	ec.accesslistLock.RUnlock()
	return
}

type accesslistResp struct {
	Version       common.Hash      `json:"version"`
	WhitelistFrom []common.Address `json:"whitelist_from"`
	WhitelistTo   []common.Address `json:"whitelist_to"`
	BlacklistFrom []common.Address `json:"blacklist_from"`
	BlacklistTo   []common.Address `json:"blacklist_to"`
}

func (ec *Client) syncAccessList(ctx context.Context) error {
	var accesslist accesslistResp
	err := ec.c.CallContext(ctx, &accesslist, "gasabs_accesslist", ec.accesslistVersion)
	if err != nil {
		return err
	}

	if accesslist.Version == ec.accesslistVersion {
		return nil
	}

	whitelistFrom := make(map[common.Address]struct{})
	whitelistTo := make(map[common.Address]struct{})
	blacklistFrom := make(map[common.Address]struct{})
	blacklistTo := make(map[common.Address]struct{})
	for _, from := range accesslist.WhitelistFrom {
		whitelistFrom[from] = struct{}{}
	}
	for _, to := range accesslist.WhitelistTo {
		whitelistTo[to] = struct{}{}
	}
	for _, from := range accesslist.BlacklistFrom {
		blacklistFrom[from] = struct{}{}
	}
	for _, to := range accesslist.BlacklistTo {
		blacklistTo[to] = struct{}{}
	}

	ec.accesslistLock.Lock()
	ec.accesslistVersion = accesslist.Version
	ec.whitelistFrom = whitelistFrom
	ec.whitelistTo = whitelistTo
	ec.blacklistFrom = blacklistFrom
	ec.blacklistTo = blacklistTo
	ec.accesslistLock.Unlock()

	ec.log.Info("Gasabs accesslist updated", "version", ec.accesslistVersion, "whitelist_from_length", len(ec.whitelistFrom), "whitelist_to_length", len(ec.whitelistTo), "blacklist_from_length", len(ec.blacklistFrom), "blacklist_to_length", len(ec.blacklistTo))
	return nil
}

// Host
func (ec *Client) Host() string {
	return ec.rawurl
}
