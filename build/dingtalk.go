package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

func getBody(w http.ResponseWriter,r *http.Request) {
	fmt.Println("hello")

	body, _ := ioutil.ReadAll(r.Body)
	fmt.Println(string(body))  //打印完整参数
	fmt.Println(r.Method)
	s1 := `{"msgtype":"text","text":{"content":`
	s2 := `},"at": {
       "atMobiles": [
           "13916201176",
       ],
       "isAtAll": false
    }`
	s3:= s1+string(body)+s2+"}"

	fmt.Println(s3)
	fmt.Printf("%T\n",s3)
	// 钉钉接口地址
	url := "https://oapi.dingtalk.com/robot/send?access_token=d02cf44281b26ea0780537c8a9562b64b37c353201ab439839eaeda9029fc7b8"
	req,err := http.NewRequest("POST",url,strings.NewReader(s3))
	req.Header.Add("content-type", "application/json")
	if err != nil {
		panic(err)
	}
	defer req.Body.Close()
	client := &http.Client{Timeout: 5 * time.Second}
	resp, error := client.Do(req)
	if error != nil {
		panic(error)
	}
	defer resp.Body.Close()
	result, _ := ioutil.ReadAll(resp.Body)
	content := string(result)
	fmt.Println(content)
	return

}
func main() {
	http.HandleFunc("/register",getBody)
	err:=http.ListenAndServe("0.0.0.0:9000",nil)
	if err != nil{
		fmt.Println("http server start failed!",err)
		return
	}

}
