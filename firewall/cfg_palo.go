package firewall

import (
	"crypto/tls"
	"fmt"
	"net/http"
)

type (
	// SetttingCofnig -
	SetttingCofnig struct {
		IP      string
		Loc     string
		Vsys    string
		Key     string
		Version string
	}
)

const (
	// PaloIPmgm - IP Management
	PaloIPmgm = "192.168.10.187"
	// PaloLocation -
	PaloLocation = "vsys"
	// PaloVsys -
	PaloVsys = "vsys1"
	// PaloKey -
	PaloKey = "LUFRPT1tdExDSzhTQkdLUmFWUFFiTTBpbnVEeUhFdTQ9QlB4Ny9YZmJLUFZSRVZRazg1MVI2Z1RBTVMzU2hWWnFQL0hrVndwaThlQWZ4c21pS2VvMXZSRjh3Mm9OQ1NvOQ=="
	// PaloVersion -
	PaloVersion = "9.1"
)

// InitPalo - init config
func (s *SetttingCofnig) InitPalo() {
	s.IP = PaloIPmgm
	s.Loc = PaloLocation
	s.Vsys = PaloVsys
	s.Key = PaloKey
	s.Version = PaloVersion
}

// GetZonePalo - Get zone security from Palo Alto
func (s *SetttingCofnig) GetZonePalo() string {

	var PathURL string = `https://` + s.IP + `/restapi/` + s.Version + `/Network/Zones?location=` + s.Loc + `&vsys=` + s.Vsys
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

	client := &http.Client{}
	req, errReq := http.NewRequest("GET", PathURL, nil)
	req.Header.Set("X-PAN-KEY", s.Key)

	if errReq != nil {
		fmt.Println(errReq)
	}

	_, errRes := client.Do(req)

	if errRes != nil {
		//sdfsdf
	}

	return "sdfsdf"
}
