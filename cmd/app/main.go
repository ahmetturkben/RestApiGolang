package main


import(
	"fmt"
	"net/http"
	"log"
	"github.com/gorilla/mux"
	//"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"gorm.io/driver/postgres"
	"RestApiGolang/pkg/api/log"
	"RestApiGolang/pkg/repository/log"
	"RestApiGolang/pkg/service/log"

	"RestApiGolang/pkg/api/user"
	"RestApiGolang/pkg/repository/user"
	"RestApiGolang/pkg/service/user"
)

type App struct {
	Router *mux.Router
	DB 	   *gorm.DB
}

func main() {
	a := App{}

	// Initialize routes
	a.initialize(
		"",
		"",
		"",
		"",
		"")
		a.routes()

	// Start server
	a.run(":8070")
}

func (a *App) initialize(host, port, username, password, dbname string) {
	var err error

	//dsn := "sqlserver://aturkben_SQLLogin_1:cq7tpjbi66@ShoppingCartv1.mssql.somee.com?database=ShoppingCartv1"
	dsn := "host=127.0.0.1 user=postgres password=admin dbname=UserApp port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	a.DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
		fmt.Printf("%s", err)
	}

	a.Router = mux.NewRouter()
}

func (a *App) run(addr string) {
	fmt.Printf("Server started at %s\n", addr)
	log.Fatal(http.ListenAndServe(addr, a.Router))
}

func (a *App) routes() {
	logAPI := InitLogAPI(a.DB)
	a.Router.HandleFunc("/logs", logAPI.FindAllLogs()).Methods("GET")
	a.Router.HandleFunc("/logs", logAPI.CreateLog()).Methods("POST")
	a.Router.HandleFunc("/logs/{id:[0-9]+}", logAPI.FindByID()).Methods("GET")
	a.Router.HandleFunc("/logs/{id:[0-9]+}", logAPI.UpdateLog()).Methods("PUT")
	a.Router.HandleFunc("/logs/{id:[0-9]+}", logAPI.DeleteLog()).Methods("DELETE")
	
	userAPI := InitUserAPI(a.DB)
	a.Router.HandleFunc("/users", userAPI.FindAllUsers()).Methods("GET")
	a.Router.HandleFunc("/users", userAPI.CreateUser()).Methods("POST")
	a.Router.HandleFunc("/users/{id:[0-9]+}", userAPI.FindByID()).Methods("GET")
	a.Router.HandleFunc("/users/{id:[0-9]+}", userAPI.UpdateUser()).Methods("PUT")
	a.Router.HandleFunc("/users/{id:[0-9]+}", userAPI.UpdateUser()).Methods("PATCH")
	a.Router.HandleFunc("/users/{id:[0-9]+}", userAPI.DeleteUser()).Methods("DELETE")
}

// InitLogAPI ..
func InitLogAPI(db *gorm.DB) logapi.LogAPI {
	logRepository := logrepo.NewRepository(db)
	logService := logservice.NewLogService(logRepository)
	logAPI := logapi.NewLogAPI(logService)
	logAPI.Migrate()
	return logAPI
}

// InitUserAPI
func InitUserAPI(db *gorm.DB) userapi.UserAPI {
	userRepository := userrepo.NewRepository(db)
	userService := userservice.NewUserService(userRepository)
	userAPI := userapi.NewUserAPI(userService)
	userAPI.Migrate()
	return userAPI
}