package vpnapi

import (
	"errors"
	"fmt"
	"net/netip"
	"strconv"
	"strings"
	"time"
)

// WgStatTimestamp - VPN stat timestamp.
type WgStatTimestamp struct {
	Timestamp int64
	Time      time.Time
}

// WgStatTraffic - VPN stat traffic.
type WgStatTraffic struct {
	Rx uint64
	Tx uint64
}

// WgStatTrafficMap - VPN stat traffic map, key is User wg_public_key.
// Dedicated map objects for wg and ipsec.
type WgStatTrafficMap struct {
	Wg    map[string]*WgStatTraffic
	IPSec map[string]*WgStatTraffic
}

// WgStatLastActivityMap - VPN stat last activity map, key is User wg_public_key.
// Dedicated map objects for wg and ipsec.
type WgStatLastActivityMap struct {
	Wg    map[string]time.Time
	IPSec map[string]time.Time
}

// WgStatEndpointMap - VPN stat endpoint map, key is User wg_public_key.
// Dedicated map objects for wg and ipsec.
type WgStatEndpointMap struct {
	Wg    map[string]netip.Prefix
	IPSec map[string]netip.Prefix
}

// ErrInvalidStatFormat - invalid stat format.
var ErrInvalidStatFormat = errors.New("invalid stat")

// NewWgStatTrafficMap - create new WgStatTrafficMap.
func NewWgStatTrafficMap() *WgStatTrafficMap {
	return &WgStatTrafficMap{
		Wg:    make(map[string]*WgStatTraffic),
		IPSec: make(map[string]*WgStatTraffic),
	}
}

// NewWgStatLastActivityMap - create new WgStatLastActivityMap.
func NewWgStatLastActivityMap() *WgStatLastActivityMap {
	return &WgStatLastActivityMap{
		Wg:    make(map[string]time.Time),
		IPSec: make(map[string]time.Time),
	}
}

// NewWgStatEndpointMap - create new WgStatEndpointMap.
func NewWgStatEndpointMap() *WgStatEndpointMap {
	return &WgStatEndpointMap{
		Wg:    make(map[string]netip.Prefix),
		IPSec: make(map[string]netip.Prefix),
	}
}

// WgStatParseTimestamp - parse timestamp value.
func WgStatParseTimestamp(timestamp string) (*WgStatTimestamp, error) {
	ts, err := strconv.ParseInt(timestamp, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("parse: %w", err)
	}

	return &WgStatTimestamp{
		Timestamp: ts,
		Time:      time.Unix(ts, 0).UTC(),
	}, nil
}

// WgStatParseTraffic - parse traffic from text.
func WgStatParseTraffic(traffic string) (*WgStatTrafficMap, error) {
	m := NewWgStatTrafficMap()

	for _, line := range strings.Split(traffic, "\n") {
		if line == "" {
			continue
		}

		clmns := strings.Split(line, "\t")
		if len(clmns) < 3 {
			return nil, fmt.Errorf("traffic: %w", ErrInvalidStatFormat)
		}

		rx, err := strconv.ParseUint(clmns[1], 10, 64)
		if err != nil {
			continue
		}

		tx, err := strconv.ParseUint(clmns[2], 10, 64)
		if err != nil {
			continue
		}

		m.Wg[clmns[0]] = &WgStatTraffic{
			Rx: rx,
			Tx: tx,
		}

		if len(clmns) >= 5 {
			rx, err := strconv.ParseUint(clmns[3], 10, 64)
			if err != nil {
				continue
			}

			tx, err := strconv.ParseUint(clmns[4], 10, 64)
			if err != nil {
				continue
			}

			m.IPSec[clmns[0]] = &WgStatTraffic{
				Rx: rx,
				Tx: tx,
			}
		}
	}

	return m, nil
}

// WgStatParseLastActivity - parse last activity time from text.
func WgStatParseLastActivity(lastSeen string) (*WgStatLastActivityMap, error) {
	m := NewWgStatLastActivityMap()

	for _, line := range strings.Split(lastSeen, "\n") {
		if line == "" {
			continue
		}

		clmns := strings.Split(line, "\t")
		if len(clmns) < 2 {
			return nil, fmt.Errorf("last seen: %w", ErrInvalidStatFormat)
		}

		ts, err := strconv.ParseInt(clmns[1], 10, 64)
		if err != nil {
			continue
		}

		if ts != 0 {
			m.Wg[clmns[0]] = time.Unix(ts, 0).UTC()
		}

		if len(clmns) >= 3 {
			ts, err := strconv.ParseInt(clmns[2], 10, 64)
			if err != nil {
				continue
			}

			if ts != 0 {
				m.IPSec[clmns[0]] = time.Unix(ts, 0).UTC()
			}
		}
	}

	return m, nil
}

// WgStatParseEndpoints - parse last seen endpoints from text.
func WgStatParseEndpoints(lastSeen string) (*WgStatEndpointMap, error) {
	m := NewWgStatEndpointMap()

	for _, line := range strings.Split(lastSeen, "\n") {
		if line == "" {
			continue
		}

		clmns := strings.Split(line, "\t")
		if len(clmns) < 2 {
			return nil, fmt.Errorf("endpoints: %w", ErrInvalidStatFormat)
		}

		prefix, err := netip.ParsePrefix(clmns[1])
		if err != nil {
			continue
		}

		if prefix.IsValid() {
			m.Wg[clmns[0]] = prefix
		}

		if len(clmns) >= 3 {
			prefix, err := netip.ParsePrefix(clmns[2])
			if err != nil {
				continue
			}

			if prefix.IsValid() {
				m.IPSec[clmns[0]] = prefix
			}
		}
	}

	return m, nil
}

// WgStatParse - parse stats from parsed response.
// Most of fileds have a text format, so we need to parse them.
func WgStatParse(resp *WGStats) (*WgStatTimestamp, *WgStatTrafficMap, *WgStatLastActivityMap, *WgStatEndpointMap, error) {
	ts, err := WgStatParseTimestamp(resp.Timestamp)
	if err != nil {
		return nil, nil, nil, nil, fmt.Errorf("parse: %w", err)
	}

	trafficMap, err := WgStatParseTraffic(resp.Traffic)
	if err != nil {
		return nil, nil, nil, nil, fmt.Errorf("parse: %w", err)
	}

	lastActivityMap, err := WgStatParseLastActivity(resp.LastActivity)
	if err != nil {
		return nil, nil, nil, nil, fmt.Errorf("parse: %w", err)
	}

	endpointsMap, err := WgStatParseEndpoints(resp.Endpoints)
	if err != nil {
		return nil, nil, nil, nil, fmt.Errorf("parse: %w", err)
	}

	return ts, trafficMap, lastActivityMap, endpointsMap, nil
}
