package main


import(
	"fmt"
	"net/http"
	"log"
	"github.com/gorilla/mux"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"

	"LogPushService/pkg/api"
	"LogPushService/pkg/repository/logrepo"
	"LogPushService/pkg/service"
)

type App struct {
	Router *mux.Router
	DB 	   *gorm.DB
}

func main() {
	a := App{}

	// Initialize storage
//	a.initialize(
//		os.Getenv("APP_DB_HOST"),
//		os.Getenv("APP_DB_PORT"),
//		os.Getenv("APP_DB_USERNAME"),
//		os.Getenv("APP_DB_PASSWORD"),
//		os.Getenv("APP_DB_NAME"))

a.initialize(
		"ahmetturkben/sqlexpress",
		"4096",
		"testuser",
		"12345678",
		"ShoppingCart")

	// Initialize routes
	a.routes()

	// Start server
	a.run(":8070")
}

func (a *App) initialize(host, port, username, password, dbname string) {
	var err error


	dsn := "sqlserver://username:password@server?database=ShoppingCartv1"
	a.DB, err = gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
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
}

// InitPostAPI ..
func InitLogAPI(db *gorm.DB) api.LogAPI {
	logRepository := logrepo.NewRepository(db)
	logService := service.NewLogService(logRepository)
	logAPI := api.NewLogAPI(logService)
	logAPI.Migrate()
	return logAPI
}
