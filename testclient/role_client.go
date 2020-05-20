package testclient

import (
	"fmt"
	"honeycomb-server-client/constant"
	"honeycomb-server-client/serialize"
	"honeycomb-server-client/tools"
	"honeycomb/model/returns"
	"io/ioutil"
	"net/url"
)

func HttpPostFormRoleLogin(userID string, roleID string, userType string, HardwareID string) string {
	resp, err := tools.HoneycombPostForm(constant.Address+"/role/login",
		url.Values{"userid": {userID}, "roleid": {roleID}, "userType": {userType}, "hwid": {HardwareID}})

	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	loginReturn := new(returns.LoginReturn)
	result := serialize.JSONDeserialize(loginReturn, string(body))
	fmt.Println(roleID, result)
	return string(body)
}
