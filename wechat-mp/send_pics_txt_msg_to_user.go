//发送图文信息
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
	ToUser  string   `json:"touser"`
	MsgType string   `json:"msgtype"`
	News    Articles `json:"news"`
}

type Articles struct {
	ArticlesContent []TextMsgContent `json:"articles"`
}

type TextMsgContent struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Url         string `json:"url"`
	Picurl      string `json:"picurl"`
}

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

// 发送消息
func pushCustomMsg(accessToken string, msg CustomServiceMsg) error {
        body, _:= json.Marshal(msg)
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
	//Fetch access_token
	accessToken, expiresIn, err := fetchAccessToken()
	if err != nil {
		log.Println("Get access_token error:", err)
		return
	}
	fmt.Println(expiresIn)
	// 消息内容
	jsonStr := []byte(`{
    "touser": "oH64bwpHEzu38qIXoARAqCEcnv3k",
    "msgtype": "news", 
    "news": {
        "articles": [
            {
                "title": "今日[2017-01-08]长垣数据中心系统报告", 
                "description": "长垣数据中心健康报告！内容包括，出入口流量统计，每日用户PV，系统日志情况，系统安全等等相关内容！", 
                "url": "http://www.os4u.xyz", 
                "picurl": "http://www.ue.cn/style/front/images/lianjiaologo.png"
            }, 
            {
                "title": "今日[2017-01-08]0校园网络带宽报告", 
                "description": "每所学校使用带宽统计相关内容！", 
                "url": "http://www.os4u.xyz", 
                "picurl": "http://www.ue.cn/style/front/images/lianjiaologo.png"
            }, 
            {
                "title": "今日[2017-01-08]1校园网络带宽报告", 
                "description": "每所学校使用带宽统计相关内容！", 
                "url": "http://www.os4u.xyz", 
                "picurl": "http://www.ue.cn/style/front/images/lianjiaologo.png"
            }, 
            {
                "title": "今日[2017-01-08]2校园网络带宽报告", 
                "description": "每所学校使用带宽统计相关内容！", 
                "url": "http://www.os4u.xyz", 
                "picurl": "http://www.ue.cn/style/front/images/lianjiaologo.png"
            }, 
            {
                "title": "今日[2017-01-08]3校园网络带宽报告", 
                "description": "每所学校使用带宽统计相关内容！", 
                "url": "http://www.os4u.xyz", 
                "picurl": "http://www.ue.cn/style/front/images/lianjiaologo.png"
            }, 
            {
                "title": "今日[2017-01-08]4校园网络带宽报告", 
                "description": "每所学校使用带宽统计相关内容！", 
                "url": "http://www.os4u.xyz", 
                "picurl": "http://www.ue.cn/style/front/images/lianjiaologo.png"
            }
        ]
      }
    }`)

	var csMsg CustomServiceMsg

	if err := json.Unmarshal(jsonStr, &csMsg); err != nil {
		fmt.Println(err)
		return
	}

	err := pushCustomMsg(accessToken, csMsg)
	if err != nil {
		log.Println("Push custom service message err:", err)
		return
	}
}
