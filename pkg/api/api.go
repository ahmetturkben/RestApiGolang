package api

import(
	"net/http"
	"LogPushService/pkg/service"
	"log"
	"fmt"
)

// PostAPI ...
type LogAPI struct {
	LogService service.LogService
}

// NewPostAPI ...
func NewLogAPI(l service.LogService) LogAPI{
	return LogAPI{LogService: l}
}

// FindAllLogs ...
func (l LogAPI) FindAllLogs() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logs, err := l.LogService.All()
		if err != nil {
			RespondWithError(w, http.StatusNotFound, err.Error())
			fmt.Println(err)
			return
		}

		RespondWithJSON(w, http.StatusOK, logs)
	}
}

func (p LogAPI) Migrate() {
	err := p.LogService.Migrate()
	if err != nil {
		log.Println(err)
	}
}
