package model

import "errors"

// NAME 名称
var NAME = "Model"

// ErrCallFail 调用失败
var ErrCallFail = errors.New(NAME + ": transfer funcs fail.")

// ErrorLinkFail 数据库链接失败
var ErrorLinkFail = errors.New(NAME + ": Link failed error.")
