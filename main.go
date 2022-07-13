package main

import (
	"io"
	"net/http"

	"encoding/json"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	log "github.com/sirupsen/logrus"
	"strconv"

)

var db, err = gorm.Open("mysql", "root:root123@tcp(localhost:33061)/todolist?charset=utf8&parseTime=True&loc=Local")

type TodoItemModel struct {
	Id          int `gorm:"primary_key"`
	Description string
	Completed   bool
}

func CreateItem(w http.ResponseWriter, r *http.Request) {
	description := r.FormValue("description")
	log.WithFields(log.Fields{"description": description}).Info("Add new TodoItem. Saving to database")
	todo := &TodoItemModel{Description: description, Completed: false}
	db.Create(&todo)
	result := db.Last(&todo)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result.Value)

}

// getItemByID

func GetItemByID(Id int ) bool {
	todo := &TodoItemModel{}
	result := db.First(&todo, Id)
	if result.Error !=nil {
		log.Warn("TodoItem not fount DB")
		return false 
	}
	return true
}


// update

func UpdateItem(w http.ResponseWriter, r *http.Request) {

vars := mux.Vars(r)
id, _ := strconv.Atoi(vars["id"])

err := GetItemByID(id)
if err ==false {
	w.Header().Set("Content-type", "Application/json")
	io.WriteString(w, `{deleted": false, "error": "not found"}`)
} else {
	completed, _ := strconv.ParseBool(r.FormValue("completed"))
	log.WithFields(log.Fields{"Id" : id, "Completed": completed}). Info("Updating TodoItem")
	todo := TodoItemModel{}
	db.First(&todo, id)
	todo.Completed = completed 
	db.Save(&todo)
	w.Header().Set("Content-type", "Application/json")
	io.WriteString(w, `{"updated" : true}`)
	
}
}

// delete

func DeleteItem(w http.ResponseWriter, r *http.Request) {

	Vars:= mux.Vars(r)
	id, _ := strconv.Atoi(Vars["id"])

	err := GetItemByID(id)
	if err == false {
	w.Header().Set("Content-type", "Application/json")
	io.WriteString(w, `{"deleted" : false, "error" : "not found"}`)

} else {
	log.WithFields(log.Fields{"Id" : id,}). Info("Deleting TodoItem")
	todo := TodoItemModel{}
	db.First(&todo, id)
	db.Delete(&todo)
	w.Header().Set("Content-type", "Application/json")
	io.WriteString(w, `{"Deleted" : true}`)
}
}


	


func Healthz(w http.ResponseWriter, r *http.Request) {
	log.Info("API Health is OK")
	w.Header().Set("Content-Type", "application/json")
	io.WriteString(w, `{"alive": true}`)
}

func main() {

	defer db.Close()
	if err != nil {
		log.Error(err)
	}
	db.Debug().DropTableIfExists(&TodoItemModel{})
	db.Debug().AutoMigrate(&TodoItemModel{})

	log.Info("Starting Todolist API server")
	router := mux.NewRouter()
	router.HandleFunc("/healthz", Healthz).Methods("GET")
	router.HandleFunc("/todo", CreateItem).Methods("POST")
	router.HandleFunc("/todo/{id}", UpdateItem).Methods("POST")
	router.HandleFunc("/todo/{id}", DeleteItem).Methods("DELETE")
	http.ListenAndServe(":8000", router)
}
