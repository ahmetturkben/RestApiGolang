package userapi

import(
	"net/http"
	"RestApiGolang/pkg/service/user"
	"RestApiGolang/pkg/model/user"
	"github.com/gorilla/mux"
	"log"
	"fmt"
	"strconv"
	"encoding/json"
)

// PostAPI ...
type UserAPI struct {
	UserService userservice.UserService
}

// NewPostAPI ...
func NewUserAPI(u userservice.UserService) UserAPI{
	return UserAPI{UserService: u}
}

// FindAllUsers ...
func (u UserAPI) FindAllUsers() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		users, err := u.UserService.All()
		if err != nil {
			RespondWithError(w, http.StatusNotFound, err.Error())
			fmt.Println(err)
			return
		}

		RespondWithJSON(w, http.StatusOK, users)
	}
}

// FindByID ...
func (u UserAPI) FindByID() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// Check if id is integer
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			RespondWithError(w, http.StatusBadRequest, err.Error())
			return
		}

		// Find uogin by id from db
		user, err := u.UserService.FindByID(uint(id))
		if err != nil {
			RespondWithError(w, http.StatusNotFound, err.Error())
			return
		}

		RespondWithJSON(w, http.StatusOK, model.ToUserDTO(user))
	}
}

// CreatePost ...
func (u UserAPI) CreateUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var UserDTO model.UserDto
		log.Print("createUser")

		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&UserDTO); err != nil {
			RespondWithError(w, http.StatusBadRequest, err.Error())
			return
		}
		defer r.Body.Close()

		createdUser, err := u.UserService.Create(model.ToUser(&UserDTO))
		if err != nil {
			RespondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}

		RespondWithJSON(w, http.StatusOK, model.ToUserDTO(createdUser))
	}
}

// UpdateUser ...
func (u UserAPI) UpdateUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			RespondWithError(w, http.StatusBadRequest, err.Error())
			return
		}

		var UserDTO model.UserDto
		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&UserDTO); err != nil {
			RespondWithError(w, http.StatusBadRequest, err.Error())
			return
		}
		defer r.Body.Close()

		user, err := u.UserService.FindByID(uint(id))
		if err != nil {
			RespondWithError(w, http.StatusNotFound, err.Error())
			return
		}

		user.Name = UserDTO.Name
		user.Email = UserDTO.Email
		user.Password = UserDTO.Password
		updatedUser, err := u.UserService.Save(user)
		if err != nil {
			RespondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}

		RespondWithJSON(w, http.StatusOK, model.ToUserDTO(updatedUser))
	}
}

// DeleteUser ...
func (u UserAPI) DeleteUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			RespondWithError(w, http.StatusBadRequest, err.Error())
			return
		}

		user, err := u.UserService.FindByID(uint(id))
		if err != nil {
			RespondWithError(w, http.StatusNotFound, err.Error())
			return
		}

		err = u.UserService.Delete(user.Id)
		if err != nil {
			RespondWithError(w, http.StatusNotFound, err.Error())
			return
		}

		type Response struct {
			Message string
		}

		response := Response{
			Message: "User deleted successfully!",
		}
		RespondWithJSON(w, http.StatusOK, response)
	}
}

func (u UserAPI) Migrate() {
	err := u.UserService.Migrate()
	if err != nil {
		log.Println(err)
	}
}
