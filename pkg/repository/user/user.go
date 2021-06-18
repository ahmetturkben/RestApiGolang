package userrepo

import(
	"LogPushService/pkg/model"
	"gorm.io/gorm"
)

//Repository
type Repository struct {
	db *gorm.DB
}

//NewRepository
func NewRepository(db *gorm.DB) *Repository{
	return &Repository{
		db: db,
	}
}

//All
func (p *Repository) All() ([]model.User, error){
	users := []model.User{}
	err := p.db.Find(&users).Error
	return users, err
}

//FindById
func (p *Repository) FindByID(id uint) (*model.User, error){
	user := new(model.User)
	err := p.db.Where(`id = ?`, id).First(&user).Error
	return user, err
}

//Insert
func (p *Repository) Create(user *model.User) (*model.User, error) {
	err := p.db.Create(&user).Error
	return user, err
}

//Save
func (p *Repository) Save(user *model.User) (*model.User, error) {
	err := p.db.Save(&user).Error
	return user, err
}

// Delete ...
func (p *Repository) Delete(id uint) error {
	err := p.db.Delete(&model.User{Id: id}).Error
	return err
}

// Migrate ...
func (p *Repository) Migrate() error {
	return p.db.AutoMigrate(&model.User{})
}

