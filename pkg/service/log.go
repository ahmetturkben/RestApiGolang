package service

import(
	"LogPushService/pkg/model"
	"LogPushService/pkg/repository/log"
)

//logService
type LogService struct{
	LogRepository *logrepo.Repository
}

// NewLogService ...
func NewLogService(l *logrepo.Repository) LogService{
	return LogService{LogRepository: l}
}

//All
func (l *LogService) All() ([]model.Log, error) {
	return l.LogRepository.All()
}

// FindByID ...
func (l *LogService) FindByID(id uint) (*model.Log, error) {
	return l.LogRepository.FindByID(id)
}

// Insert 
func (l *LogService) Create(log *model.Log) (*model.Log, error) {
	return l.LogRepository.Create(log)
}

// Save ...
func (l *LogService) Save(log *model.Log) (*model.Log, error) {
	return l.LogRepository.Save(log)
}

// Delete ...
func (l *LogService) Delete(id uint) error {
	return l.LogRepository.Delete(id)
}

// Migrate ...
func (l *LogService) Migrate() error {
	return l.LogRepository.Migrate()
}
