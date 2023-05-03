/*
 @author: lynn
 @date: 2023/5/3
 @time: 20:55
*/

package requests

type Config struct {
	Url     string
	Method  string
	Params  map[string]string
	Headers map[string]string
	Body    map[string]interface{}
}
