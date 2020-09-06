package PaloAlto

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
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

	return &SetttingCofnig{
		IP:      PaloIPmgm,
		Loc:     PaloLocation,
		Vsys:    PaloVsys,
		Key:     PaloKey,
		Version: PaloVersion,
	}

}

// GetZonePalo - Get zone security from Palo Alto
func GetZonePalo() string {

	// Init value for connect API
	var s SetttingCofnig = *InitPalo()

	// Fullpath to Palo Alto API
	var PathURL string = `https://` + s.IP + `/restapi/` + s.Version + `/Network/Zones?location=` + s.Loc + `&vsys=` + s.Vsys

	// Exception SSL
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

	// Create Cleint Object fot HTTP
	client := &http.Client{}

	// Create Request Object for Client
	req, errReq := http.NewRequest("GET", PathURL, nil)

	// Set Header Authentication to API
	req.Header.Set("X-PAN-KEY", s.Key)

	if errReq != nil {
		fmt.Println(errReq)
	}

	resP, errRes := client.Do(req)

	if errRes != nil {
		//sdfsdf
	}

	defer resP.Body.Close()
	body, _ := ioutil.ReadAll(resP.Body)

	return string(body)
}
