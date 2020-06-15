package testclient

import (
	"fmt"
	"honeycomb-server-client/constant"
	"honeycomb-server-client/serialize"
	"honeycomb-server-client/tools"
	"honeycomb/model/global"
	"honeycomb/model/returns"
	"io/ioutil"
	"net/url"
)

func HttpPostFormRoleLogin(userID string, roleID string, userType string) string {
	resp, err := tools.HoneycombPostForm(constant.HoneycombAddress+"/role/login",
		url.Values{"userid": {userID}, "roleid": {roleID}, "usertype": {userType}})

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
	fmt.Println(roleID, result.(*returns.LoginReturn).ErrCode+result.(*returns.LoginReturn).ErrMsg)
	switch result.(*returns.LoginReturn).UserType {
	case global.USER_TYPE_SUPER, global.USER_TYPE_ADMIN:
		fmt.Println(result.(*returns.LoginReturn).AdminInstance)
	case global.USER_TYPE_BEEKEEPER:
		fmt.Println(result.(*returns.LoginReturn).BeekeeperInstance)
	case global.USER_TYPE_CONSUMER:
		fmt.Println(result.(*returns.LoginReturn).ConsumerInstance)
	}
	return string(body)
}
