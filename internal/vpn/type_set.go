package vpn

import "strings"

type ProtocolSet uint8

const (
	TypeOutline ProtocolSet = 1 << iota
	TypeOVC
	TypeWG
	TypeIPSec
)

var (
	type2string = map[ProtocolSet]string{
		TypeOutline: Outline,
		TypeOVC:     OVC,
		TypeWG:      WG,
		TypeIPSec:   IPSec,
	}
	string2type = map[string]ProtocolSet{
		Outline: TypeOutline,
		OVC:     TypeOVC,
		WG:      TypeWG,
		IPSec:   TypeIPSec,
	}
)

func (s ProtocolSet) String() string {
	types := make([]string, 0, len(type2string))
	for k, v := range type2string {
		if s&k != 0 {
			types = append(types, v)
		}
	}
	return strings.Join(types, ",")
}

func (s ProtocolSet) GetSupported(available ProtocolSet) (supported ProtocolSet, unsupported ProtocolSet) {
	supported = s & available
	unsupported = s & ^available
	return
}

func NewTypesFromString(s string) ProtocolSet {
	t := ProtocolSet(0)
	for _, v := range strings.Split(s, ",") {
		t |= string2type[strings.Trim(v, " ")]
	}
	return t
}
