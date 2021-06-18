package logrepo

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
func (p *Repository) All() ([]model.Log, error){
	logs := []model.Log{}
	err := p.db.Find(&logs).Error
	return logs, err
}

//FindById
func (p *Repository) FindByID(id uint) (*model.Log, error){
	log := new(model.Log)
	err := p.db.Where(`id = ?`, id).First(&log).Error
	return log, err
}

//Insert
func (p *Repository) Create(log *model.Log) (*model.Log, error) {
	err := p.db.Create(&log).Error
	return log, err
}

//Save
func (p *Repository) Save(log *model.Log) (*model.Log, error) {
	err := p.db.Save(&log).Error
	return log, err
}

// Delete ...
func (p *Repository) Delete(id uint) error {
	err := p.db.Delete(&model.Log{Id: id}).Error
	return err
}

// Migrate ...
func (p *Repository) Migrate() error {
	return p.db.AutoMigrate(&model.Log{})
}

