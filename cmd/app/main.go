package main


import(
	"fmt"
	"net/http"
	"log"
	"github.com/gorilla/mux"
	//"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"gorm.io/driver/postgres"
	"LogPushService/pkg/api"
	"LogPushService/pkg/repository/log"
	"LogPushService/pkg/service"

	// "LogPushService/pkg/api/user"
	"LogPushService/pkg/repository/user"
	// "LogPushService/pkg/service"
)

type App struct {
	Router *mux.Router
	DB 	   *gorm.DB
}

func main() {
	a := App{}

	// Initialize routes
	a.initialize(
		"ahmetturkben/sqlexpress",
		"4096",
		"testuser",
		"12345678",
		"ShoppingCart")
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
func InitLogAPI(db *gorm.DB) api.LogAPI {
	logRepository := logrepo.NewRepository(db)
	logService := service.NewLogService(logRepository)
	logAPI := api.NewLogAPI(logService)
	logAPI.Migrate()
	return logAPI
}

// InitUserAPI
func InitUserAPI(db *gorm.DB) api.UserAPI {
	userRepository := userrepo.NewRepository(db)
	userService := service.NewUserService(userRepository)
	userAPI := api.NewUserAPI(userService)
	userAPI.Migrate()
	return userAPI
}