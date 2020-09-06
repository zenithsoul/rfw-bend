package paloalto

import (
	"crypto/tls"
	"net/http"
)

// SetttingCofnig -
type SetttingCofnig struct {
	IP      string
	Loc     string
	Vsys    string
	Key     string
	Version string
}

const (
	// PaloIPmgm - IP Management
	PaloIPmgm string = "192.168.10.187"
	// PaloLocation -
	PaloLocation string = "vsys"
	// PaloVsys -
	PaloVsys string = "vsys1"
	// PaloKey -
	PaloKey string = "LUFRPT1tdExDSzhTQkdLUmFWUFFiTTBpbnVEeUhFdTQ9QlB4Ny9YZmJLUFZSRVZRazg1MVI2Z1RBTVMzU2hWWnFQL0hrVndwaThlQWZ4c21pS2VvMXZSRjh3Mm9OQ1NvOQ=="
	// PaloVersion -
	PaloVersion string = "v9.1"
)

// InitPalo - init config
func InitPalo() *SetttingCofnig {

	// Exception SSL
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

	return &SetttingCofnig{
		IP:      PaloIPmgm,
		Loc:     PaloLocation,
		Vsys:    PaloVsys,
		Key:     PaloKey,
		Version: PaloVersion,
	}

}
