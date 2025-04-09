// Code generated by goctl. DO NOT EDIT.
// goctl 1.8.1

package types

type AddGoodsRequest struct {
	GoodsBase
}

type AddUserRequest struct {
	UserBase
}

type AdminLoginRequest struct {
	Email    string `json:"email" validate:"required,email"`    // 管理员邮箱
	Password string `json:"password" validate:"required,min=6"` // 管理员密码
}

type AdminLoginResponse struct {
	Token string   `json:"token"` // JWT 或其他认证令牌
	User  UserBase `json:"user"`  // 用户信息（可选）
}

type CommonResponse struct {
	Status int         `json:"status"` // 状态码，例如 200 表示成功
	Msg    string      `json:"msg"`    // 返回消息
	Data   interface{} `json:"data"`   // 返回数据（可以是任意类型）
}

type GetUserByIdReq struct {
	ID int `path:"id" validate:"gt=0"`
}

type GoodsBase struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
	Stock int     `json:"stock"`
}

type GoodsBaseOptional struct {
	Name  *string  `json:"name,omitempty"`
	Price *float64 `json:"price,omitempty"`
	Stock *int     `json:"stock,omitempty"`
}

type UpdateGoodsRequest struct {
	ID uint `json:"id"` // 主键，必须传递
	GoodsBaseOptional
}

type UpdateUserRequest struct {
	ID uint `json:"id"` // 主键，必须传递
	UserBaseOptional
}

type UserBase struct {
	Name     string `json:"name"`
	Age      int    `json:"age"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserBaseOptional struct {
	Name     *string `json:"name,omitempty" validate:"omitempty,min=2"`
	Age      *int    `json:"age,omitempty" validate:"omitempty,gte=0,lte=120"`
	Email    *string `json:"email,omitempty" validate:"omitempty,email"`
	Password *string `json:"password,omitempty" validate:"omitempty,min=6"`
}
