package vpnapi

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/netip"
	"net/url"
	"os"
)

// WGStats - wg_stats endpoint-API call.
type WGStats struct {
	Code         string `json:"code"`
	Traffic      string `json:"traffic"`
	LastActivity string `json:"last-seen"`
	Endpoints    string `json:"endpoints"`
	Timestamp    string `json:"timestamp"`
}

// WgPeerAdd - peer_add endpoint-API call.
func WgPeerAdd(
	actualAddrPort,
	calculatedAddrPort netip.AddrPort,
	wgPub, wgIfacePub,
	wgPSK []byte,
	ipv4,
	ipv6,
	keydesk netip.Addr,
	ovcCertRequest string,
) ([]byte, error) {
	query := fmt.Sprintf("peer_add=%s&wg-public-key=%s&wg-psk-key=%s&allowed-ips=%s",
		url.QueryEscape(base64.StdEncoding.WithPadding(base64.StdPadding).EncodeToString(wgPub)),
		url.QueryEscape(base64.StdEncoding.WithPadding(base64.StdPadding).EncodeToString(wgIfacePub)),
		url.QueryEscape(base64.StdEncoding.WithPadding(base64.StdPadding).EncodeToString(wgPSK)),
		url.QueryEscape(ipv4.String()+","+ipv6.String()),
	)

	if ovcCertRequest != "" {
		query += fmt.Sprintf("&openvpn-client-csr=%s", url.QueryEscape(ovcCertRequest))
	}

	if keydesk.IsValid() {
		query += fmt.Sprintf("&control-host=%s", url.QueryEscape(keydesk.String()))
	}

	body, err := getAPIRequest(actualAddrPort, calculatedAddrPort, query)
	if err != nil {
		return nil, fmt.Errorf("api: %w", err)
	}

	return body, nil
}

// WgPeerDel - peer_del endpoint-API call.
func WgPeerDel(actualAddrPort, calculatedAddrPort netip.AddrPort, wgPub, wgIfacePub []byte) error {
	query := fmt.Sprintf("peer_del=%s&wg-public-key=%s",
		url.QueryEscape(base64.StdEncoding.WithPadding(base64.StdPadding).EncodeToString(wgPub)),
		url.QueryEscape(base64.StdEncoding.WithPadding(base64.StdPadding).EncodeToString(wgIfacePub)),
	)

	_, err := getAPIRequest(actualAddrPort, calculatedAddrPort, query)
	if err != nil {
		return fmt.Errorf("api: %w", err)
	}

	return nil
}

// WgAdd - wg_add endpoint-API call.
func WgAdd(
	actualAddrPort,
	calculatedAddrPort netip.AddrPort,
	wgPriv []byte,
	endpointIPv4 netip.Addr,
	endpointPort uint16,
	IPv4CGNAT,
	IPv6ULA netip.Prefix,
	ovcUID string,
	ovcFakeDomain string,
	ovcCACert string,
	ovcRouterCAKey string,
) error {
	fmt.Fprintf(os.Stderr, "WgAdd: %d\n", len(wgPriv))

	query := fmt.Sprintf("wg_add=%s&external-ip=%s&wireguard-port=%s&internal-nets=%s",
		url.QueryEscape(base64.StdEncoding.WithPadding(base64.StdPadding).EncodeToString(wgPriv)),
		url.QueryEscape(endpointIPv4.String()),
		url.QueryEscape(fmt.Sprintf("%d", endpointPort)),
		url.QueryEscape(IPv4CGNAT.String()+","+IPv6ULA.String()),
	)

	if ovcUID != "" && ovcCACert != "" && len(ovcRouterCAKey) > 0 {
		query += fmt.Sprintf("&cloak-bypass-uid=%s&openvpn-ca-crt=%s&openvpn-ca-key=%s&cloak-domain=%s",
			url.QueryEscape(ovcUID),
			url.QueryEscape(ovcCACert),
			url.QueryEscape(ovcRouterCAKey),
			url.QueryEscape(ovcFakeDomain),
		)
	}

	_, err := getAPIRequest(actualAddrPort, calculatedAddrPort, query)
	if err != nil {
		return fmt.Errorf("api: %w", err)
	}

	return nil
}

// WgDel - wg_del endpoint API call.
func WgDel(actualAddrPort, calculatedAddrPort netip.AddrPort, wgPriv []byte) error {
	query := fmt.Sprintf("wg_del=%s",
		url.QueryEscape(base64.StdEncoding.WithPadding(base64.StdPadding).EncodeToString(wgPriv)),
	)

	_, err := getAPIRequest(actualAddrPort, calculatedAddrPort, query)
	if err != nil {
		return fmt.Errorf("api: %w", err)
	}

	return nil
}

// WgStat - stat endpoint API call.
func WgStat(actualAddrPort, calculatedAddrPort netip.AddrPort, wgPub []byte) (*WGStats, error) {
	query := fmt.Sprintf("stat=%s",
		url.QueryEscape(base64.StdEncoding.WithPadding(base64.StdPadding).EncodeToString(wgPub)),
	)

	body, err := getAPIRequest(actualAddrPort, calculatedAddrPort, query)
	if err != nil {
		return nil, fmt.Errorf("api: %w", err)
	}

	if body == nil {
		return nil, nil
	}

	data := &WGStats{}
	if err := json.Unmarshal(body, data); err != nil {
		return nil, fmt.Errorf("api payload: %w", err)
	}

	return data, nil
}
