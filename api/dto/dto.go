package dto

import (
	"fmt"
	"github.com/go-playground/validator"
	"unicode/utf8"
)

type RegisterDto struct {
	UserName string `validate:"required,checkName" json:"username"`
	Password string `validate:"required" json:"password"`
}

var validate *validator.Validate

func init() {
	validate = validator.New()
	_ = validate.RegisterValidation("checkName", CheckNameFunc)
}

// CheckNameFunc 自定义校验器校验用户名
func CheckNameFunc(f validator.FieldLevel) bool {
	count := utf8.RuneCountInString(f.Field().String())
	if count >= 2 && count <= 12 {
		return true
	} else {
		return false
	}
}

// ValidatorRegister 定义校验数据的方法
func ValidatorRegister(account RegisterDto) error {
	err := validate.Struct(account)
	if err != nil {
		// 输出校验错误 .(validator.ValidationErrors)是断言
		for _, e := range err.(validator.ValidationErrors)[:1] {
			fmt.Println("错误字段:", e.Field())
			fmt.Println("错误的值:", e.Value())
			fmt.Println("错误的tag:", e.Tag())
		}
		return err
	} else {
		return nil
	}
}
