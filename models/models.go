package models

//通用返回信息
type CommonRet struct {
	ErrCode string `json:"errCode"`
	ErrMsg  string `json:"errMsg"`
}

//UserRet 返回给客户端的用户注册信息
type UserRet struct {
	UserServerInfo *UserServerInfo
	CommonRet
}

//UserServerInfo 返回给客户端的用户信息
type UserServerInfo struct {
	UID         int64  `json:"uid"`        //用户ID
	RID         int64  `json:"rid"`        //用户角色ID
	LoginToken  string `json:"loginToken"` //TOKEN
	UserRegType int    `json:"userRegType"`
	UserType    int    `json:"userType"`
}
