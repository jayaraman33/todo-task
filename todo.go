package main
import (
	"io"
	"net/http"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"github.com/jinzhu/gorm"
	"gorm.io/driver/postgres"

)


var db, _ := gorm.Open(postgres.Open("postgres://ecgjxqyi:p0AK5VS0f9j-WtbAeTfwU0AFM1Z3FUmQ@queenie.db.elephantsql.com/ecgjxqyi"),&gorm.Config{})

type TodoItemModel struct{
		 Id int `gorm:"primary_key"`
		 Description string
		 Completed bool
	}




func Healthz(w http.ResponseWriter, r *http.Request) {
	defer db.Close()
	db.Debug().DropTableIfExists(&TodoItemModel{})
    db.Debug().AutoMigrate(&TodoItemModel{})



	log.Info("API Health is OK")
	w.Header().Set("Content-Type", "application/json")
	io.WriteString(w, `{"alive": true}`)
}

func init() {
	log.SetFormatter(&log.TextFormatter{})
	log.SetReportCaller(true)
}

func main() {
	log.Info("Starting Todo API server")
	router := mux.NewRouter()
	router.HandleFunc("/healthz", Healthz).Methods("GET")
	http.ListenAndServe(":8000", router)
}

