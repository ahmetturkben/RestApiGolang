package service

import(
	"LogPushService/pkg/model"
	"LogPushService/pkg/repository/user"
)

//userService
type UserService struct{
	UserRepository *userrepo.Repository
}

// NewUserService ...
func NewUserService(l *userrepo.Repository) UserService{
	return UserService{UserRepository: l}
}

//All
func (u *UserService) All() ([]model.User, error) {
	return u.UserRepository.All()
}

// FindByID ...
func (u *UserService) FindByID(id uint) (*model.User, error) {
	return u.UserRepository.FindByID(id)
}

// Insert 
func (u *UserService) Create(user *model.User) (*model.User, error) {
	return u.UserRepository.Create(user)
}

// Save ...
func (u *UserService) Save(user *model.User) (*model.User, error) {
	return u.UserRepository.Save(user)
}

// Delete ...
func (u *UserService) Delete(id uint) error {
	return u.UserRepository.Delete(id)
}

// Migrate ...
func (u *UserService) Migrate() error {
	return u.UserRepository.Migrate()
}
