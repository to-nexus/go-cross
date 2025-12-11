package params

import (
	"slices"
	"sync"
)

// ##CROSS: istanbul param

type (
	istanbulParamEntry struct {
		config    *IstanbulConfig
		timepoint uint64
	}
	istanbulParams struct {
		mu     sync.RWMutex
		params []*istanbulParamEntry // params are sorted by timepoint in ascending order
	}
)

var istanbulParamsInst = &istanbulParams{}

func (p *istanbulParams) length() int {
	p.mu.RLock()
	defer p.mu.RUnlock()

	return len(p.params)
}

func (p *istanbulParams) config(timepoint uint64) *IstanbulConfig {
	p.mu.RLock()
	defer p.mu.RUnlock()

	// find the config with the highest timepoint that is <= given timepoint
	for i := len(p.params) - 1; i >= 0; i-- {
		if e := p.params[i]; e.timepoint <= timepoint {
			return e.config
		}
	}
	return nil
}

func (p *istanbulParams) addConfig(config *IstanbulConfig, timepoint uint64) {
	p.mu.Lock()
	defer p.mu.Unlock()

	p.params = append(p.params, &istanbulParamEntry{
		config:    config,
		timepoint: timepoint,
	})

	slices.SortFunc(p.params, func(a, b *istanbulParamEntry) int {
		return int(a.timepoint - b.timepoint)
	})
}

func (p *istanbulParams) clear() {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.params = nil
}

// IstanbulConfigLength returns the number of cached Istanbul configs.
func IstanbulConfigLength() int {
	return istanbulParamsInst.length()
}

// IstanbulConfigAt returns the IstanbulConfig closest to the given timepoint.
func IstanbulConfigAt(timepoint uint64) *IstanbulConfig {
	return istanbulParamsInst.config(timepoint)
}

// AddIstanbulConfig adds the IstanbulConfig at the given timepoint.
func AddIstanbulConfig(config *IstanbulConfig, timepoint uint64) {
	istanbulParamsInst.addConfig(config, timepoint)
}

// ClearCachedIstanbulConfigs clears the cached Istanbul configs.
func ClearCachedIstanbulConfigs() {
	istanbulParamsInst.clear()
}
