package tools

import (
	"crypto/md5"
	"crypto/tls"
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
	hardwareID    = "hardwareID"
)

func AccountPostForm(address string, forms url.Values) (*http.Response, error) {
	time := strconv.FormatInt(time.Now().Unix(), 10)
	content := time + AccountMD5KEY
	sign := fmt.Sprintf("%x", md5.Sum([]byte(content)))
	log.Println(sign, content)
	forms.Add("sign", sign)
	forms.Add("time", time)
	return generateHttpsClient().PostForm(address, forms)
}

func HoneycombPostForm(address string, forms url.Values) (*http.Response, error) {
	time := strconv.FormatInt(time.Now().Unix(), 10)
	content := time + HoneycombKey
	sign := fmt.Sprintf("%x", md5.Sum([]byte(content)))
	log.Println(sign, content)
	forms.Add("sign", sign)
	forms.Add("time", time)
	forms.Add("debug", DebugSign)
	forms.Add("hwid", hardwareID)
	return generateHttpsClient().PostForm(address, forms)
}

func generateHttpsClient() *http.Client {
	//跳过证书验证
	transport := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	httpClient := &http.Client{
		Timeout:   5 * time.Second,
		Transport: transport,
	}
	return httpClient
}
