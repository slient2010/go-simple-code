package main

import (
	"bytes"
	//	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
         _ "encoding/json"
)
// get wxplat token

const (
    appID                = "wxcc77f4a9d744e7ca"
    appSecret            = "6eaa0232295214dc5b39582870e7c176"
    accessTokenFetchUrl  = "https://api.weixin.qq.com/cgi-bin/token"
    customServicePostUrl = "https://api.weixin.qq.com/cgi-bin/message/custom/send"
)

//获取 accessToken
func fetchAccessToken() (string, float64, error) {
    requestLine := strings.Join([]string{accessTokenFetchUrl,
        "?grant_type=client_credential&appid=",
        appID,
        "&secret=",
        appSecret}, "")

    resp, err := http.Get(requestLine)
    if err != nil || resp.StatusCode != http.StatusOK {
        return "", 0.0, err
    }

    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return "", 0.0, err
    }
    accessToken := AccessTokenAndExpire{}
    err = json.Unmarshal(body, &accessToken)
    return accessToken.AccessToken, accessToken.ExpiresTime, err
}

func main() {
    accessToken, _, _ := fetchAccessToken()

	//菜单JSON字符串
	menuStr := `{
            "button": [
            {
                "name": "联教官网",
                "type": "view",
                "url": "http://www.ue.cn/"
            },
            {

                "name":"管理中心",
                 "sub_button":[
                        {
                        "name": "用户中心",
                        "type": "click",
                        "key": "molan_user_center"
                        },
                        {
                        "name": "公告",
                        "type": "click",
                        "key": "molan_institution"
                        }]
            },
            {
                "name": "资料修改",
                "type": "view",
                "url": "http://www.os4u.xyz/ping"
            }
            ]
        }`

	urlPost := "https://api.weixin.qq.com/cgi-bin/menu/create?access_token=" + accessToken
	bodyType := "application/json;charset=utf-8"
	bytesPost := bytes.NewBuffer([]byte(menuStr))

	res_post, err := http.Post(urlPost, bodyType, bytesPost)

	if err == nil {
		body_post, _ := ioutil.ReadAll(res_post.Body)
		defer res_post.Body.Close()
		fmt.Println(string(body_post))
	}
}
