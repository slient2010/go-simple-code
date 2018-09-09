package main

import (
	"bytes"
	"fmt"
	"net/http"
	"encoding/json"
	"io/ioutil"
)

const (
    APPID                = "wxcc77f4a9d744e7ca"
    APPSECRET            = "6eaa0232295214dc5b39582870e7c176"
)

type AccessTokenResponse struct {
	AccessToken string  `json:"access_token"`
	ExpiresIn   float64 `json:"expires_in"`
}

func main() {
	wxAccessUrl := "https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=" + APPID + "&secret=" + APPSECRET

	resp, err := http.Get(wxAccessUrl)
	if err != nil || resp.StatusCode != http.StatusOK {
		fmt.Println("发送get请求获取 atoken 错误", err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("发送get请求获取 atoken 读取返回body错误", err)
	}

	if bytes.Contains(body, []byte("access_token")) {
		atr := AccessTokenResponse{}
		err = json.Unmarshal(body, &atr)
		if err != nil {
			fmt.Println("发送get请求获取 atoken 返回数据json解析错误", err)
		}
		fmt.Println(atr.AccessToken, atr.ExpiresIn)
	} else {
		fmt.Println("发送get请求获取 微信返回 err")
	}
}
