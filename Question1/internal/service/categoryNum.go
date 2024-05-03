package service

import (
	"fmt"
	"net/http"
)
func FetchData(cat string, n string) (*http.Response, error) {
	requestURL := fmt.Sprintf("http://20.244.56.144/test/companies/AMZ/categories/%s/products?top=%s&minPrice=1&maxPrice=10000", cat, n)

	req, err := http.NewRequest("GET", requestURL, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating HTTP request: %s", err)
	}
	req.Header.Add("Authorization", Token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error making HTTP request: %s", err)
	}
	return resp, nil
}
