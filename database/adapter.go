package database

import "github.com/staumann/caluclation/model"

type Adapter interface {
	GetBillByID(int64) *model.Bill
	SaveBill(*model.Bill) error
	UpdateBill(*model.Bill) error
	DeleteBillByID(int64) error

	GetUserByID(int64) *model.User
	SaveUser(*model.User) error
	UpdateUser(*model.User) error
	DeleteUserByID(int64) error
}
