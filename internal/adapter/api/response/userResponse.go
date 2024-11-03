package response

import (
	"ecommerce/internal/helper"
)

type UserDataResponse map[string]interface{}
type UserDataArrayResponse []map[string]interface{}

func NewUserResponse(data interface{}, err error) Response {
	jsonResp, _ := helper.ToJson(data)

	return Response{
		Messsage: "successful",
		Data:     UserDataResponse(jsonResp),
	}
}

func NewUserArrayResponse(data interface{}, err error) Response {
	jsonResp, _ := helper.ToArrayJson(data)

	return Response{
		Messsage: "successful",
		Data:     UserDataArrayResponse(jsonResp),
	}
}
