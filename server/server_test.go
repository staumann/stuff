package server

import "github.com/staumann/caluclation/model"

type AdapterSpy struct {
	getHandler        func(int64) *model.Bill
	saveHandler       func(*model.Bill) error
	updateHandler     func(*model.Bill) error
	deleteHandlerById func(int64) error

	getUserHandler  func(int64) *model.User
	saveUserHandler func(*model.User) error
}

func (as *AdapterSpy) GetBillByID(id int64) *model.Bill {
	if as.getHandler != nil {
		return as.getHandler(id)
	}
	return nil
}

func (as *AdapterSpy) SaveBill(bill *model.Bill) error {
	if as.saveHandler != nil {
		return as.saveHandler(bill)
	}
	return nil
}

func (as *AdapterSpy) UpdateBill(bill *model.Bill) error {
	if as.updateHandler != nil {
		return as.updateHandler(bill)
	}
	return nil
}

func (as *AdapterSpy) SaveUser(user *model.User) error {
	if as.saveUserHandler != nil {
		return as.saveUserHandler(user)
	}
	return nil
}

func (as *AdapterSpy) DeleteBillByID(id int64) error {
	if as.deleteHandlerById != nil {
		return as.deleteHandlerById(id)
	}
	return nil
}

func (as *AdapterSpy) GetUserByID(id int64) *model.User {
	if as.getUserHandler != nil {
		return as.getUserHandler(id)
	}
	return nil
}
