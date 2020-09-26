package models

type ResponseResult struct {
	Cod int`json:"code"`
	Message string`json:"message"`
	Data interface{}`json:"data"`
}
