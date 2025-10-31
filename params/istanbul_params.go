package params

import (
	"sync"
)

// ##CROSS: istanbul param

type (
	istanbulParamEntry struct {
		config    *IstanbulConfig
		timepoint uint64
	}
	istanbulParams struct {
		mu     sync.Mutex
		params map[uint64]*istanbulParamEntry // index => config
	}
)

var istanbulParamsInst = &istanbulParams{
	params: make(map[uint64]*istanbulParamEntry),
}

func (p *istanbulParams) config(timepoint uint64) (config *IstanbulConfig) {
	p.mu.Lock()
	defer p.mu.Unlock()

	var best uint64 = 0

	// find the config with the highest timepoint that is <= given timepoint
	for _, entry := range p.params {
		if entry.timepoint <= timepoint && entry.timepoint >= best {
			best = entry.timepoint
			config = entry.config
		}
	}
	return
}

func (p *istanbulParams) configByIndex(index uint64) (config *IstanbulConfig) {
	p.mu.Lock()
	defer p.mu.Unlock()

	if e := p.params[index]; e != nil {
		return e.config
	}
	return nil
}

func (p *istanbulParams) setConfig(index uint64, config *IstanbulConfig, timepoint uint64) {
	if config == nil {
		return
	}
	p.mu.Lock()
	defer p.mu.Unlock()

	p.params[index] = &istanbulParamEntry{
		config:    config,
		timepoint: timepoint,
	}
}

func (p *istanbulParams) clear() {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.params = make(map[uint64]*istanbulParamEntry)
}

// IstanbulConfigAt returns the IstanbulConfig closest to the given timepoint.
func IstanbulConfigAt(timepoint uint64) *IstanbulConfig {
	return istanbulParamsInst.config(timepoint)
}

// IstanbulConfigByIndex returns the IstanbulConfig at the given index.
func IstanbulConfigByIndex(index uint64) *IstanbulConfig {
	return istanbulParamsInst.configByIndex(index)
}

// SetIstanbulConfig adds the IstanbulConfig of the given index.
func SetIstanbulConfig(index uint64, config *IstanbulConfig, timepoint uint64) {
	istanbulParamsInst.setConfig(index, config, timepoint)
}

// ClearCachedIstanbulConfigs clears the cached Istanbul configs.
func ClearCachedIstanbulConfigs() {
	istanbulParamsInst.clear()
}
