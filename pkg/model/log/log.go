package model

type Log struct {
	Id uint `gorm:"primary_key" json:"id"`
	Message string `json:"message"`
}

type LogDto struct {
	Id uint `gorm:"primary_key" json:"id"`
	Message string `json:"message"`
}

func ToLog(logDto *LogDto) *Log{
	return &Log{
		Id: logDto.Id,
		Message: logDto.Message,
	}
}

func ToLogDTO(log *Log) *LogDto{
	return &LogDto{
		Id: log.Id,
		Message: log.Message,
	}
}