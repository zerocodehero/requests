/*
 @author: lynn
 @date: 2023/5/3
 @time: 20:04
*/

package requests

import (
	"bytes"
	"encoding/json"
	"github.com/tidwall/gjson"
	"io"
	"log"
	"net/http"
)

/*
*网络请求封装
 */

type ClientOption struct {
	Url     string
	Params  map[string]string
	Headers map[string]string
	Body    map[string]interface{}
}

// Get http Get method
func (option ClientOption) Get() []byte {
	//new request
	req, err := http.NewRequest("GET", option.Url, nil)
	if err != nil {
		log.Println(err)
		return nil
	}
	//add params
	q := req.URL.Query()
	if option.Params != nil {
		for key, val := range option.Params {
			q.Add(key, val)
		}
		req.URL.RawQuery = q.Encode()
	}
	//add headers
	req.Header.Add("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.0.0 Safari/537.36")
	if option.Headers != nil {
		for key, val := range option.Headers {
			req.Header.Add(key, val)
		}
	}
	//http client
	client := &http.Client{}
	log.Printf("Go GET URL : %s \n", req.URL.String())

	//发送请求
	res, err := client.Do(req)
	if err != nil {
		return nil
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Println(err)
		}
	}(res.Body) //一定要关闭res.Body
	//读取body
	resBody, err := io.ReadAll(res.Body) //把  body 内容读入字符串 s
	if err != nil {
		return nil
	}

	return resBody
}

// Post http post method
func (option *ClientOption) Post() []byte {
	//add post body
	var bodyJson []byte
	var req *http.Request
	if option.Body != nil {
		var err error
		bodyJson, err = json.Marshal(option.Body)
		if err != nil {
			log.Println(err)
			return nil
		}
	}
	req, err := http.NewRequest("POST", option.Url, bytes.NewBuffer(bodyJson))
	if err != nil {
		log.Println(err)
		return nil
	}
	req.Header.Set("Content-type", "application/json")
	req.Header.Add("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/102.0.5005.61 Safari/537.36")
	//add params
	q := req.URL.Query()
	if option.Params != nil {
		for key, val := range option.Params {
			q.Add(key, val)
		}
		req.URL.RawQuery = q.Encode()
	}
	//add headers
	if option.Headers != nil {
		for key, val := range option.Headers {
			req.Header.Add(key, val)
		}
	}
	//http client
	client := &http.Client{}
	log.Printf("Go POST URL : %s \n", req.URL.String())

	//发送请求
	res, err := client.Do(req)
	if err != nil {
		return nil
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Println(err)
		}
	}(res.Body) //一定要关闭res.Body
	//读取body
	resBody, err := io.ReadAll(res.Body) //把  body 内容读入字符串 s
	if err != nil {
		return nil
	}

	return resBody
}

func (option *ClientOption) ToJson(resBody []byte) gjson.Result {
	return gjson.Parse(string(resBody))
}
func (option *ClientOption) ToStruct(resBody []byte, structName interface{}) {
	if err := json.Unmarshal(resBody, structName); err != nil {
		log.Println(err)
	}
}
