package model

import "errors"
//根据业务逻辑的需要，自定义一一些错误
var (
	ERROR_USER_NOTEXISTS = errors.New("用户不存在")
	ERROR_USER_PWD = errors.New("密码不正确")
	ERROR_USER_TEXISTS = errors.New("用户已存在")
)