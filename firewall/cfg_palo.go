package firewall

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
	PaloVersion string = "9.1"
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
func (s *SetttingCofnig) GetZonePalo() (string, error) {

	return `1`, nil
}

/*
func GetZonePalo() string {

	var PathURL string = `https://` + PaloIPmgm + `/restapi/` + PaloVersion + `/Network/Zones?location=` + PaloLocation + `&vsys=` + PaloVsys
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

	client := &http.Client{}
	req, errReq := http.NewRequest("GET", PathURL, nil)
	req.Header.Set("X-PAN-KEY", PaloKey)

	if errReq != nil {
		fmt.Println(errReq)
	}

	resP, errRes := client.Do(req)

	if errRes != nil {
		//sdfsdf
	}
	resP.
	return `1`
}
*/
