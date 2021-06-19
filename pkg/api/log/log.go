package logapi

import(
	"net/http"
	"RestApiGolang/pkg/service/log"
	"RestApiGolang/pkg/model/log"
	"github.com/gorilla/mux"
	"log"
	"fmt"
	"strconv"
	"encoding/json"
)

// PostAPI ...
type LogAPI struct {
	LogService logservice.LogService
}

// NewPostAPI ...
func NewLogAPI(l logservice.LogService) LogAPI{
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

// FindByID ...
func (l LogAPI) FindByID() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// Check if id is integer
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			RespondWithError(w, http.StatusBadRequest, err.Error())
			return
		}

		// Find login by id from db
		log, err := l.LogService.FindByID(uint(id))
		if err != nil {
			RespondWithError(w, http.StatusNotFound, err.Error())
			return
		}

		RespondWithJSON(w, http.StatusOK, model.ToLogDTO(log))
	}
}

// CreatePost ...
func (l LogAPI) CreateLog() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var logDTO model.LogDto
		log.Print("createLog")

		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&logDTO); err != nil {
			RespondWithError(w, http.StatusBadRequest, err.Error())
			return
		}
		defer r.Body.Close()

		createdLog, err := l.LogService.Create(model.ToLog(&logDTO))
		if err != nil {
			RespondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}

		RespondWithJSON(w, http.StatusOK, model.ToLogDTO(createdLog))
	}
}

// UpdateLog ...
func (l LogAPI) UpdateLog() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			RespondWithError(w, http.StatusBadRequest, err.Error())
			return
		}

		var logDTO model.LogDto
		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&logDTO); err != nil {
			RespondWithError(w, http.StatusBadRequest, err.Error())
			return
		}
		defer r.Body.Close()

		log, err := l.LogService.FindByID(uint(id))
		if err != nil {
			RespondWithError(w, http.StatusNotFound, err.Error())
			return
		}

		log.Message = logDTO.Message
		updatedLog, err := l.LogService.Save(log)
		if err != nil {
			RespondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}

		RespondWithJSON(w, http.StatusOK, model.ToLogDTO(updatedLog))
	}
}

// DeleteLog ...
func (l LogAPI) DeleteLog() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			RespondWithError(w, http.StatusBadRequest, err.Error())
			return
		}

		log, err := l.LogService.FindByID(uint(id))
		if err != nil {
			RespondWithError(w, http.StatusNotFound, err.Error())
			return
		}

		err = l.LogService.Delete(log.Id)
		if err != nil {
			RespondWithError(w, http.StatusNotFound, err.Error())
			return
		}

		type Response struct {
			Message string
		}

		response := Response{
			Message: "Log deleted successfully!",
		}
		RespondWithJSON(w, http.StatusOK, response)
	}
}

func (p LogAPI) Migrate() {
	err := p.LogService.Migrate()
	if err != nil {
		log.Println(err)
	}
}
