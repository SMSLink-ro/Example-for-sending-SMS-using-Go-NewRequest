package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
)

func main() {

  //
  //  Get your SMSLink / SMS Gateway Connection ID and Password from 
  //  https://www.smslink.ro/get-api-key/
  //

	url := "https://secure.smslink.ro/sms/gateway/communicate/index.php?connection_id=MyConnectionID&password=MyConnectionPassword&to=07xyzzzzzz&message=TestMessage"

	req, _ := http.NewRequest("GET", url, nil)

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))

}