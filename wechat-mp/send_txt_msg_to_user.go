//向用户发送文本文件
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

// text msg struct
type CustomServiceMsg struct {
	ToUser  string         `json:"touser"`
	MsgType string         `json:"msgtype"`
	Text    TextMsgContent `json:"text"`
}

// 消息内容struct
type TextMsgContent struct {
	Content string `json:"content"`
}

// access token 及过期时间
type AccessTokenAndExpire struct {
	AccessToken string  `json:"access_token"`
	ExpiresTime float64 `json:"expires_in"`
}

const (
	appID                = "wxcc77f4a9d744e7ca"
	appSecret            = "6eaa0232295214dc5b39582870e7c176"
	accessTokenFetchUrl  = "https://api.weixin.qq.com/cgi-bin/token"
	customServicePostUrl = "https://api.weixin.qq.com/cgi-bin/message/custom/send"
)

//获取token
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

// 发送消息
func pushCustomMsg(accessToken, toUser, msg string) error {

	csMsg := &CustomServiceMsg{
		ToUser:  toUser,
		MsgType: "text",
		Text:    TextMsgContent{Content: msg},
	}

	body, err := json.MarshalIndent(csMsg, " ", "  ")
	if err != nil {
		return err
	}

	postReq, err := http.NewRequest("POST",
		strings.Join([]string{customServicePostUrl, "?access_token=", accessToken}, ""),
		bytes.NewReader(body))
	if err != nil {
		return err
	}

	postReq.Header.Set("Content-Type", "application/json; encoding=utf-8")

	client := &http.Client{}
	resp, err := client.Do(postReq)
	if err != nil {
		return err
	}
	resp.Body.Close()

	return nil
}

func main() {
	// Fetch access_token
	accessToken, expiresIn, err := fetchAccessToken()
	if err != nil {
		log.Println("Get access_token error:", err)
		return
	}
	fmt.Println(accessToken, expiresIn)

	// 用户ID
	userID := "oH64bwpHEzu38qIXoARAqCEcnv3k"
	// Post custom service message
	// 文本消息内容
	msg := "你好, 网站http://www.xxx.com无法响应请求!"
	err = pushCustomMsg(accessToken, userID, msg)
	if err != nil {
		log.Println("Push custom service message err:", err)
		return
	}
}
