package dto


type UserRequest struct {
	UserName string // 用户名
	UserId string // 用户id
}

func(req UserRequest) JavaClassName() string{
	return JavaDtoPgkName + "QueryUserParam"
}


type UserResponse struct {
	UserName string // 用户名
	UserId string // 用户id
	BirthDate string // 出生日期
	Age int64 // 年龄
	IsDead bool // 是否已去世
}

func(req UserResponse) JavaClassName() string{
	return JavaDtoPgkName + "QueryUserResponse"
}
