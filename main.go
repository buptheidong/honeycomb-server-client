package main

import "honeycomb-server-client/testclient"

func main() {
	//account相关协议
	//testclient.HttpPostFormAccountLogin("test1@test.com", "2", "abc123", "")
	//honeycomb相关协议
	//testclient.HttpPostFormRoleLogin("1", "1100000000", "1")
	testclient.HttpPostFormSubmitPurchase("1100000000", "0", "测试10", "描述10", "1", "2", "2", "1")
	//testclient.HttpPostFormGetPurchase("1100000000", "0", "0")
}
