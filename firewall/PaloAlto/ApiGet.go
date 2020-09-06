package paloalto

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
)

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

	// Create Request to API Server
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
