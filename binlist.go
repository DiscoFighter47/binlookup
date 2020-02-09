package binlookup

import (
	"fmt"
	"net/http"
)

// LookUpResponse ...
type LookUpResponse struct {
	Number struct {
		Length int  `json:"length"`
		Luhn   bool `json:"luhn"`
	} `json:"number"`
	Scheme  string `json:"scheme"`
	Type    string `json:"type"`
	Brand   string `json:"brand"`
	Prepaid bool   `json:"prepaid"`
	Country struct {
		Numeric   string `json:"numeric"`
		Alpha2    string `json:"alpha2"`
		Name      string `json:"name"`
		Emoji     string `json:"emoji"`
		Currency  string `json:"currency"`
		Latitude  int    `json:"latitude"`
		Longitude int    `json:"longitude"`
	} `json:"country"`
	Bank struct {
		Name  string `json:"name"`
		URL   string `json:"url"`
		Phone string `json:"phone"`
		City  string `json:"city"`
	} `json:"bank"`
}

// LookUp ...
func (c *Client) LookUp(bin string) (*LookUpResponse, error) {
	r, err := c.do(bin)
	if err != nil {
		return nil, fmt.Errorf("unable to connect server: %v", err)
	}
	if r.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("http status is: %v", r.StatusCode)
	}
	resp := &LookUpResponse{}
	if err = c.parseBody(r, resp); err != nil {
		return nil, fmt.Errorf("unable to parse body: %v", err)
	}
	return resp, nil
}
