package tools

import (
	"crypto/md5"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

const (
	AccountMD5KEY = "cKdG41L6HLzsDJFDF!@#210392l;k12hUNh"
	HoneycombKey  = "ASBA@e1h2e982eFlk123e09193j"        //作为debug或工具使用的md5 key
	DebugSign     = "KKD&(@#324j12087:LJF)2332j4j209Lf0" //作为以debug模式发协议的特征字符串
)

func AccountPostForm(address string, forms url.Values) (*http.Response, error) {
	time := strconv.FormatInt(time.Now().Unix(), 10)
	content := time + AccountMD5KEY
	sign := fmt.Sprintf("%x", md5.Sum([]byte(content)))
	log.Println(sign, content)
	forms.Add("sign", sign)
	forms.Add("time", time)
	return http.PostForm(address, forms)
}

func HoneycombPostForm(address string, forms url.Values) (*http.Response, error) {
	time := strconv.FormatInt(time.Now().Unix(), 10)
	content := time + HoneycombKey
	sign := fmt.Sprintf("%x", md5.Sum([]byte(content)))
	log.Println(sign, content)
	forms.Add("sign", sign)
	forms.Add("time", time)
	forms.Add("debug", DebugSign)
	return http.PostForm(address, forms)
}
