package server

import (
	"github.com/staumann/caluclation/database"
	"github.com/staumann/caluclation/server/ui"
	"github.com/staumann/caluclation/sql"
	"log"
	"net/http"
)

const (
	contentTypeJson = "application/json"
)

var (
	billRepository database.BillRepository
	userRepository database.UserRepository
)

func Start() {
	billRepository = sql.GetBillRepository()
	userRepository = sql.GetUserRepository()

	ui.Prepare(billRepository, userRepository)
	ui.ParseTemplates("frontend/html")

	// ui
	http.HandleFunc("/", ui.HomeHandler)
	http.HandleFunc("/users", ui.UserHandler)
	http.HandleFunc("/users/new", ui.NewUserHandler)
	http.HandleFunc("/users/create", ui.CreateUserHandler)
	http.HandleFunc("/bills", ui.BillHandler)

	// bill handler
	http.HandleFunc("/api/bill/create", createBillHandler)
	http.HandleFunc("/api/bill/get", getBillHandler)
	http.HandleFunc("/api/bill/update", updateBillHandler)
	http.HandleFunc("/api/bill/delete", deleteBillHandler)

	// user handler
	http.HandleFunc("/api/user/create", createUserHandler)
	http.HandleFunc("/api/user/get", getUserHandler)
	http.HandleFunc("/api/user/update", updateUserHandler)

	log.Print(http.ListenAndServe(":8889", nil))
}
