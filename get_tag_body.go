package main

import (
        "fmt"
        "io/ioutil"
        "log"
        "net/http"
        "os"
        "strings"

        "github.com/Jeffail/gabs"
        blackfriday "gopkg.in/russross/blackfriday.v2"
)

func main() {
        r, err := http.Get("https://api.github.com/repos/slient2010/webdemo/releases/tags/v0.1.14")
        if err != nil {
                fmt.Println("error: ", err)
        }
        defer r.Body.Close()
        contents, err := ioutil.ReadAll(r.Body)
        if err != nil {
                fmt.Printf("%s", err)
                os.Exit(1)
        }

        // 匹配body中是否包含"body"
        if strings.Contains(string(contents), "\"body\"") {
                jsonData, e := gabs.ParseJSON([]byte(contents))
                if e != nil {
                        log.Println(e)
                }

                // fmt.Println(jsonData.Path("body").Data().(string))
                git := jsonData.Path("body").Data().(string)

                output := blackfriday.Run([]byte(git))
                fmt.Println(string(output))
        }

}
