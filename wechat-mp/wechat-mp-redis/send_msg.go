package main

import (
	"bytes"
	"encoding/json"
	//	"io/ioutil"
	"fmt"
	"github.com/garyburd/redigo/redis"
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

func pushCustomMsg(accessToken, toUser string, msg CustomServiceMsg) error {
// func pushCustomMsg(accessToken string, msg CustomServiceMsg) error {
	customServicePostUrl := "https://api.weixin.qq.com/cgi-bin/message/custom/send"
	body, _ := json.Marshal(msg)
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

const (
	fileCfg = "warn.json"
)

func main() {
	// Connect to redis and get the key.
	c, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer c.Close()
	accessToken, err := redis.String(c.Do("GET", "accessToken"))
	if err != nil {
		fmt.Println(err)
		return
	}

	jsonStr := []byte(`{
    "touser": "oH64bwkyegY9EMbBS_9ET_I4Xmv0", 
    "msgtype": "news", 
    "news": {
        "articles": [
            {
                "title": "今日[2017-02-06]长垣数据中心系统报告", 
                "description": "长垣数据中心健康报告！内容包括，出入口流量统计，每日用户PV，系统日志情况，系统安全等等相关内容！", 
                "url": "https://www.os4u.xyz", 
                "picurl": "http://www.ue.cn/style/front/images/lianjiaologo.png"
            }, 
            {
                "title": "今日[2017-02-06]0校园网络带宽报告", 
                "description": "每所学校使用带宽统计相关内容！", 
                "url": "https://www.os4u.xyz", 
                "picurl": "http://www.ue.cn/style/front/images/lianjiaologo.png"
            }, 
            {
                "title": "今日[2017-02-06]1校园网络带宽报告", 
                "description": "每所学校使用带宽统计相关内容！", 
                "url": "https://www.os4u.xyz", 
                "picurl": "http://www.ue.cn/style/front/images/lianjiaologo.png"
            }, 
            {
                "title": "今日[2017-02-06]4校园网络带宽报告", 
                "description": "每所学校使用带宽统计相关内容！", 
                "url": "https://www.os4u.xyz/ping", 
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

	err = pushCustomMsg(accessToken, toUser, csMsg)
	if err != nil {
		log.Println("Push custom service message err:", err)
		return
	}
}
