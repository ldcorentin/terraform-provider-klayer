package klayers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

const baseLatestURL = "https://api.klayers.cloud/api/v2"

func getKlayerArns(name string, region string, python_version string) interface{} {
	// Default timeout 1h, set to 10s for production
	client := &http.Client{Timeout: 10 * time.Second}

	// Declare outputs map
	var mBody []map[string]string
	// Build url for api call
	url := fmt.Sprintf("%s/p%s/%s/%s/%s", baseLatestURL, python_version, "layers", region, name)
	fmt.Println(url)

	// Call to API
	resp, err := client.Get(url)
	if err != nil {
		fmt.Println(err)
		// return mBody
		return mBody
	}
	defer resp.Body.Close()

	// Response body from io to []byte
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		err = fmt.Errorf("Cannot read API response: %v", err)
		fmt.Println(err)
		return mBody
	}

	// Unmarshal or Decode the JSON ([]bytes) to the interface.
	err = json.Unmarshal(body, &mBody)

	return mBody
}