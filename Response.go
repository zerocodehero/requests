/*
 @author: lynn
 @date: 2023/5/3
 @time: 20:55
*/

package requests

import (
	"github.com/tidwall/gjson"
	"io"
	"log"
	"net/http"
)

type Response struct {
	RESPONSE *http.Response
	ERR      error
}

func (R *Response) JSON() gjson.Result {
	return gjson.Parse(string(R.RAW()))
}

func (R *Response) RAW() []byte {
	//读取body
	resBody, err := io.ReadAll(R.RESPONSE.Body) //把  body 内容读入字符串
	if err != nil {
		log.Println(err)
		return []byte(err.Error())
	}
	return resBody
}
