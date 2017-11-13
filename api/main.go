package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/expenses-tracker/api/models"
	"github.com/gorilla/mux"
)

var (
	database *mgo.Database
)

func init() {
	session, err := mgo.Dial("mongodb://localhost:27017/expenses_tracker")
	if err != nil {
		panic(err)
	}

	database = session.DB("expenses_tracker")
}

// GetStatus retrieves a status code and server message
func GetStatus(w http.ResponseWriter, r *http.Request) {
	type Status struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	}

	var status Status

	status.Code = 200
	status.Message = "Server is up and running"

	json.NewEncoder(w).Encode(status)
}

// CreateExpense creates a single document at expenses collection
func CreateExpense(w http.ResponseWriter, r *http.Request) {
	expensesCollection := database.C("expenses")

	var e models.Expense

	e.ID = bson.NewObjectId()
	e.Date = time.Now()
	e.Name = "Trocar óleo do carro"
	e.Value = 127.53

	err := expensesCollection.Insert(e)

	if err != nil {
		panic(err)
	}

	json.NewEncoder(w).Encode(e)
}

// UpdateExpense updates a single document in expenses collection
func UpdateExpense(w http.ResponseWriter, r *http.Request) {
	expensesCollection := database.C("expenses")

	var expense models.Expense

	if err := json.NewDecoder(r.Body).Decode(&expense); err != nil {
		panic(err)
	}

	if err := expensesCollection.UpdateId(expense.ID, &expense); err != nil {
		panic(err)
	}

	json.NewEncoder(w).Encode(expense)
}

// GetExpenses returns a list of expenses
func GetExpenses(w http.ResponseWriter, r *http.Request) {
	expensesCollection := database.C("expenses")

	var expenses []models.Expense

	err := expensesCollection.Find(nil).All(&expenses)

	if err != nil {
		panic(err)
	}

	json.NewEncoder(w).Encode(expenses)
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/status", GetStatus).Methods("GET")
	router.HandleFunc("/expenses", GetExpenses).Methods("GET")
	router.HandleFunc("/expenses", CreateExpense).Methods("POST")
	router.HandleFunc("/expenses", UpdateExpense).Methods("PUT")
	log.Fatal(http.ListenAndServe(":8000", router))
}
