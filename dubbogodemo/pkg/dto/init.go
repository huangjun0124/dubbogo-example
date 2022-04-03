package dto

type POJO interface {
	JavaClassName() string
}

type WithDto func(...POJO)

func RegisterDtos(registerFunc WithDto){
	registerFunc(
		&UserRequest{},
		&UserResponse{},
		)
}

const JavaDtoPgkName = "com.demo.exp.dto."