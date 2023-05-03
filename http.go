/*
 @author: lynn
 @date: 2023/5/3
 @time: 20:04
*/

package requests

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

func Get(config Config) *Response {
	config.Method = "GET"
	return Send(config)
}

func Post(config Config) *Response {
	config.Method = "POST"
	return Send(config)
}

func Send(config Config) (R *Response) {
	//add post body
	var bodyJson []byte
	if len(config.Body) != 0 {
		var err error
		bodyJson, err = json.Marshal(config.Body)
		if err != nil {
			log.Println(err)
		}
	} else {
		bodyJson = nil
	}

	req, err := http.NewRequest(config.Method, config.Url, bytes.NewBuffer(bodyJson))
	if err != nil {
		log.Println(err)
	}
	req.Header.Set("Content-type", "application/json")
	req.Header.Add("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/102.0.5005.61 Safari/537.36")
	//add params
	q := req.URL.Query()
	for key, val := range config.Params {
		q.Add(key, val)
	}
	req.URL.RawQuery = q.Encode()
	//add headers
	for key, val := range config.Headers {
		req.Header.Add(key, val)
	}
	//http client
	client := &http.Client{}
	log.Printf("Go POST URL : %s \n", req.URL.String())

	//发送请求
	R.RESPONSE, err = client.Do(req)

	if err != nil {
		log.Println("request error")
		R.ERR = err
		return
	}

	return

}
