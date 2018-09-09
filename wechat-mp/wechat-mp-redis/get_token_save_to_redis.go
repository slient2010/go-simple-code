package main

import (
	"encoding/json"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

type AccessTokenAndExpire struct {
	AccessToken string  `json:"access_token"`
	ExpiresTime float64 `json:"expires_in"`
}

const (
	// wxplat info
	appID                = "wxbc77f4a9d733e7ca"
	appSecret            = "6eba0753195214dc5b39582870e7c176"
	accessTokenFetchUrl  = "https://api.weixin.qq.com/cgi-bin/token"
	customServicePostUrl = "https://api.weixin.qq.com/cgi-bin/message/custom/send"

	// redis info
	redisServer = "127.0.0.1"
	redisPort   = "6379"
	redisPass   = "Lj#2017!"
)

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
	// check key
	c, err := redis.Dial("tcp", redisServer+":"+redisPort)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer c.Close()

	for {

		_, err = redis.String(c.Do("GET", "accessToken"))
		if err == nil {
			continue
		} else {
			//获取key和有效时间
			accessToken, expiresIn, e := fetchAccessToken()
			if e != nil {
				log.Println("Get access_token error:", e)
				return
			}

			// save to redis
			v, err := c.Do("SET", "accessToken", accessToken, "EX", expiresIn)
			if err != nil {
				fmt.Println(err)
				return
			}

			// for loop
			v, err = redis.String(c.Do("GET", "accessToken"))
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println(v)

			t, err := redis.Int64(c.Do("PTTL", "accessToken"))
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println(t)
		}
		time.Sleep(time.Second * 5)
	}
}
