//http://www.youdao.com/w/go/#keyfrom=dict2.top
package main
import (
"fmt"
"net/http"
"log"
"io/ioutil"
)

func main(){
url:="http://www.1905.com/vod/top/lst/?fr=voddp%20rbb"
resp,respErr:=http.Get(url)
if respErr != nil {
	log.Println(respErr)
	return
}
//fmt.Println(resp,respErr)
defer resp.Body.Close()
respBody,readErr:=ioutil.ReadAll(resp.Body)
if readErr != nil{
	fmt.Println(readErr)
	return
}
fmt.Println(string(respBody))
}

