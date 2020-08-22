package database

import "github.com/staumann/caluclation/model"

type UserRepository interface {
	GetUserByID(int64) *model.User
	GetUsers() []*model.User
	SaveUser(*model.User) error
	UpdateUser(*model.User) error
	DeleteUserByID(int64) error
}

type BillRepository interface {
	GetBillByID(int64) *model.Bill
	SaveBill(*model.Bill) error
	UpdateBill(*model.Bill) error
	DeleteBillByID(int64) error
}

type PositionRepository interface {
	GetPositionByID(int64) *model.Position
	SavePosition(*model.Position) error
	UpdatePosition(*model.Position) error
	DeleteByID(int64) error

	GetByBillID(int64) []*model.Position
}
