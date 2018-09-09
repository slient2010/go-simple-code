package main

import (
	"encoding/json"
	"fmt"
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

func main() {
	body := `{
    "touser": "oH64bwkyegY9EMbBS_9ET_IAxxdea",
    "msgtype": "news",
    "news": {
        "articles": [
            {
                "title": "今日[2017-01-08]数据中心系统报告",
                "description": "数据中心健康报告！内容包括，出入口流量统计，每日用户PV，系统日志情况，系统安全等等相关内容！",
                "url": "http://www.demo.com",
                "picurl": "http://www.demo.cn/style/front/images/logo.png"
            },
            {
                "title": "今日[2017-01-08]0校园网络带宽报告",
                "description": "每所学校使用带宽统计相关内容！",
                "url": "http://www.demo.com",
                "picurl": "http://www.demo.cn/style/front/images/logo.png"
            }
        ]
      }
    }`
	var r CustomServiceMsg
	err := json.Unmarshal([]byte(body), &r)
	if err != nil {
		fmt.Printf("err was %v", err)
	}
	fmt.Println(r)
	fmt.Println(r.ToUser)
	fmt.Println(r.News.ArticlesContent[0].Url)
	r.News.ArticlesContent[0].Url = "hello"
	fmt.Println(r.News.ArticlesContent[0].Url)
	fmt.Println(r)

}
