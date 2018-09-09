package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

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

const (
	fileCfg = "report.json"
)

func main() {
	data, err := ioutil.ReadFile(fileCfg)
	if err != nil {
		if os.IsNotExist(err) {
			return
		}
		log.Fatal(err)
	}

	var cfg CustomServiceMsg
	err = json.Unmarshal(data, &cfg)
	if err != nil {
		log.Fatal(err)
	}

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

	values, _ := redis.Values(c.Do("lrange", "wxuser", "0", "100"))

	for _, v := range values {
		cfg.ToUser = string(v.([]byte))
		err = pushCustomMsg(accessToken, cfg.ToUser, cfg)
		if err != nil {
			log.Println("Push custom service message err:", err)
			return
		}
	}
	////////fmt.Println(cfg.ToUser)
	////////fmt.Println(cfg.News.ArticlesContent[0].Url)
	////////cfg.News.ArticlesContent[0].Url = "hello"
	////////fmt.Println(cfg.News.ArticlesContent[0].Url)

}

func pushCustomMsg(accessToken, toUser string, msg CustomServiceMsg) error {
	customServicePostUrl := "https://api.weixin.qq.com/cgi-bin/message/custom/send"
	msg.ToUser = toUser
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
