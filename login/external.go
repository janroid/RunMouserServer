package login

import (
	"server/login/internal"
)

var (
	Module       = new(internal.Module)
	ChanRPC      = internal.ChanRPC
	UserLogin    = internal.UserLogin
	UserRegister = internal.UserRegister
	GetUserByUID = internal.GetUserByUID
	LoginByName  = internal.LoginByName
)
