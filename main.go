package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func GetIpAddr(ip string) string {
	client := &http.Client{}
	s := fmt.Sprintf("ip=%s&verifycode=", ip)
	var data = strings.NewReader(s)
	req, err := http.NewRequest("POST", "http://ip.360.cn/IPQuery/ipquery", data)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Accept", "application/json, text/plain, */*")
	req.Header.Set("Accept-Language", "en-US,en;q=0.9,zh-CN;q=0.8,zh;q=0.7")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Cookie", "crypt_code=B02SZv%252B4s1RefYheIKSOy6t5ewxj8aoZhxHpRZiXtDjfQO3td%252FtL8n4LYwhO52y9")
	req.Header.Set("Origin", "http://ip.360.cn")
	req.Header.Set("Referer", "http://ip.360.cn/")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/101.0.4951.67 Safari/537.36")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", bodyText)
	var z map[string]interface{}
	json.Unmarshal(bodyText, &z)
	return z["data"].(string)
}

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		ip := c.Query("ip")
		log.Printf("%s\n", ip)
		c.JSON(200, map[string]interface{}{
			"ip": GetIpAddr(ip),
		})
	})
	r.Run(":5000") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
