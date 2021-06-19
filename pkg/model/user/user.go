package model

type User struct {
	Id uint `gorm:"primary_key" json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
	Password string `json:"password"`
}

type UserDto struct {
	Id uint `gorm:"primary_key" json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
	Password string `json:"password"`
}

func ToUser(userDto *UserDto) *User{
	return &User{
		Id: userDto.Id,
		Name: userDto.Name,
		Email: userDto.Email,
		Password: userDto.Password,	
	}
}

func ToUserDTO(user *User) *UserDto{
	return &UserDto{
		Id: user.Id,
		Name: user.Name,
		Email: user.Email,
		Password: user.Password,
	}
}