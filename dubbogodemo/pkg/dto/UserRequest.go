package dto

type UserRequest struct {
	UserName string // 用户名
	UserId string // 用户id
}

func(req UserRequest) JavaClassName() string{
	return "com.demo.exp.dto.QueryUserParam"
}