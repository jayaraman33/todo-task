// package main

// import (
// 	"io"
// 	"net/http"

// 	"github.com/gorilla/mux"
// 	log "github.com/sirupsen/logrus"
// 	//"github.com/jinzhu/gorm"
// 	//"gorm.io/driver/postgres"
// 	// "github.com/go-sql-driver/mysql"
// 	//  "github.com/jinzhu/gorm"
// 	//  "github.com/jinzhu/gorm/dialects/mysql"
// )

// //var db, _ := gorm.Open(postgres.Open("postgres://ecgjxqyi:p0AK5VS0f9j-WtbAeTfwU0AFM1Z3FUmQ@queenie.db.elephantsql.com/ecgjxqyi"),&gorm.Config{})

// //var db, _ := gorm.Open("mysql", "root:root@/todolist?charset=utf8&parseTime=True&loc=Local")

// // type TodoItemModel struct{
// // 		 Id int `gorm:"primary_key"`
// // 		 Description string
// // 		 Completed bool
// // 	}

// func Healthz(w http.ResponseWriter, r *http.Request) {
// 	//  defer db.Close()
// 	//  db.Debug().DropTableIfExists(&TodoItemModel{})
// 	//  db.Debug().AutoMigrate(&TodoItemModel{})

// 	log.Info("API Health is OK")
// 	w.Header().Set("Content-Type", "application/json")
// 	io.WriteString(w, `{"alive": true}`)
// }

// // func init() {
// // 	log.SetFormatter(&log.TextFormatter{})
// // 	log.SetReportCaller(true)
// // }

// func main() {
// 	log.Info("Starting Todo API server")
// 	router := mux.NewRouter()
// 	router.HandleFunc("/healthz", Healthz).Methods("GET")
// 	http.ListenAndServe(":8000", router)
// }


package main

import (

	//"io"
	"net/http"
	//"strconv"

	"github.com/gorilla/mux"
)

// Healthz

func Healthz(W http.ResponseWriter, r *http.Request) {

}

// GetCompletedItems

func GetCompletedItems(w http.ResponseWriter, r *http.Request) {

}

// GetIncompleteItems

func GetIncompleteItems(W http.ResponseWriter, r *http.Request) {

}

// CreateItem

func CreateItem(W http.ResponseWriter, r *http.Request) {

}

// UpdateItem

func UpdateItem(W http.ResponseWriter, r *http.Request) {

}

// DeleteItem

func DeleteItem(W http.ResponseWriter, r *http.Request){

}

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/healthz", Healthz).Methods("GET")
	router.HandleFunc("/todo-completed", GetCompletedItems).Methods("GET")
	router.HandleFunc("/todo-incomplete", GetIncompleteItems).Methods("GET")
	router.HandleFunc("/todo", CreateItem).Methods("POST")
	router.HandleFunc("/todo/{id}", UpdateItem).Methods("POST")
	router.HandleFunc("/todo/{id}", DeleteItem).Methods("DELETE")
	http.ListenAndServe(":8000", router)
}

