package paloalto

import (
	"io/ioutil"
	"log"
	"net/http"
)

// GetZonePalo - Get zone security from Palo Alto
func GetZonePalo() string {

	// Init value for connect API
	var s SetttingCofnig = *InitPalo()

	// Fullpath to Palo Alto API
	var PathURL string = `https://` + s.IP + `/restapi/` + s.Version + `/Network/Zones?location=` + s.Loc + `&vsys=` + s.Vsys

	// Create Cleint Object fot HTTP
	client := &http.Client{}

	// Create Request to API Server
	req, errReq := http.NewRequest("GET", PathURL, nil)

	// Set Header Authentication to API
	req.Header.Set("X-PAN-KEY", s.Key)

	if errReq != nil {
		log.Println(errReq)
	}

	resP, errRes := client.Do(req)

	if errRes != nil {
		//Handle Error Response
		log.Println(errRes)
	}

	defer resP.Body.Close()
	body, errBodyRes := ioutil.ReadAll(resP.Body)

	if errBodyRes != nil {
		log.Println(errBodyRes)
	}

	return string(body)
}
