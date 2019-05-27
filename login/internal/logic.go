package internal

func UserLogin(name string, pwd string) (int, *User) {
	if len(name) <= 0 {
		return ERRLOGIN, nil
	}

	if len(pwd) <= 0 {
		return ERRLOGIN, nil
	}

	user := GetUserByName(name)

	if user == nil {
		return ERRNOUSER, nil
	}

	if user.acPwd != pwd {
		return ERRLOGIN, nil
	}

	if user.status == STRUCT_FORBID {
		return ERRFORBID, nil
	}

	return LGOINSUCC, user
}

func UserRegister(name string, pwd string) (int, *User) {
	if len(name) <= 0 {
		return ERRLOGIN, nil
	}

	if len(pwd) <= 0 {
		return ERRLOGIN, nil
	}

	user := GetUserByName(name)
	if user != nil {
		return ERRHAVEUSER, nil
	}

	user = createUser(name, pwd)

	return REGISTERSUCC, user

}
