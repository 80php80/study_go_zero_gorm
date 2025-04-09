package handler

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"gorm/internal/logic"
	"gorm/internal/svc"
	"gorm/internal/types"
)

func CreateUserHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AddUserRequest
		// 1. 解析请求体
		if err := httpx.Parse(r, &req); err != nil {
			// 返回通用错误响应
			httpx.OkJsonCtx(r.Context(), w, types.CommonResponse{
				Status: 400,
				Msg:    "Invalid request body: " + err.Error(),
				Data:   nil,
			})
			return
		}
		// 2. 校验请求数据
		validate := validator.New()
		if err := validate.Struct(req); err != nil {
			// 提取校验错误信息并返回通用错误响应
			var errMsg string
			if validationErrs, ok := err.(validator.ValidationErrors); ok {
				// 提取第一个错误并格式化为用户友好的提示
				field := validationErrs[0].Field() // 字段名
				tag := validationErrs[0].Tag()     // 校验规则
				errMsg = formatValidationError(field, tag)
			} else {
				// 如果不是 ValidationErrors 类型，返回通用错误
				errMsg = "Validation failed: " + err.Error()
			}
			httpx.OkJsonCtx(r.Context(), w, types.CommonResponse{
				Status: 400,
				Msg:    errMsg,
				Data:   nil,
			})
			return
		}

		l := logic.NewCreateUserLogic(r.Context(), svcCtx)
		resp, err := l.CreateUser(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}

// 辅助函数：格式化校验错误信息为用户友好的提示
func formatValidationError(field, tag string) string {
	switch tag {
	case "required":
		return fmt.Sprintf("Field '%s' is required", field)
	case "email":
		return fmt.Sprintf("Field '%s' must be a valid email address", field)
	case "min":
		return fmt.Sprintf("Field '%s' value is too short", field)
	case "max":
		return fmt.Sprintf("Field '%s' value is too long", field)
	default:
		return fmt.Sprintf("Field '%s' is invalid (%s)", field, tag)
	}
}
