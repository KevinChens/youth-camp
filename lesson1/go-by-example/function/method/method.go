package method

type User struct {
	Name string
	Age int
}

func (u User) ModifyInfo1() User {
	u.Name = "lisi"
	u.Age = 15
	// u修改不成功
	return u
}

func (u *User) ModifyInfo2() *User {
	u.Name = "lisi"
	u.Age = 15
	// u修改成功
	return u
}

