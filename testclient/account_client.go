package testclient

import (
	"fmt"
	"honeycomb-server-client/constant"
	"honeycomb-server-client/models"
	"honeycomb-server-client/serialize"
	"honeycomb-server-client/tools"
	"io/ioutil"
	"net/url"
)

func HttpPostFormAccountLogin(userAccount string, userRegType string, password string, loginToken string) string {
	resp, err := tools.AccountPostForm(constant.AccountAddress+"/account/user/login",
		url.Values{"userAccount": {userAccount}, "userRegType": {userRegType}, "password": {password}, "loginToken": {loginToken}})

	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	userRet := new(models.UserRet)
	result := serialize.JSONDeserialize(userRet, string(body))
	fmt.Println(userAccount, result)
	return string(body)
}
