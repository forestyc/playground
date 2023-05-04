package main

import (
	"fmt"
	"github.com/go-playground/validator/v10"
)

type Json struct {
	Id string `json:"id" validate:"eq=1|eq=2"`
}

var validate = validator.New()

func main() {
	// 结构体验证
	j := Json{
		Id: "1",
	}
	errs := validate.Struct(j)
	if errs != nil {
		fmt.Println(errs.Error())
	}

}
