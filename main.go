package main

import "honeycomb-server-client/testclient"

func main() {
	//account相关协议
	//testclient.HttpPostFormAccountLogin("test1@test.com", "2", "abc123", "")
	//honeycomb相关协议
	testclient.HttpPostFormRoleLogin("1", "0", "1", "22")
	testclient.HttpPostFormSubmitPurchase("1200000000", "0", "测试4", "描述4", "1", "2", "2", "1")
}
