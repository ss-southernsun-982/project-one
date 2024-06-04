package shop_find_all_test

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestFindAll(t *testing.T) {

	url := "http://localhost:8088/v1/shops"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		t.Fatalf(`Hello("Gladys") = %q`, err)
	}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		t.Fatalf(`Hello("Gladys") = %q`, err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		t.Fatalf(`Hello("Gladys") = %q, %v`, string(body), err)
	}
}
