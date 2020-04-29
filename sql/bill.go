package sql

import (
	"errors"
	"github.com/staumann/caluclation/model"
	"log"
	"time"
)

func (a *Adapter) GetBillByID(id int64) *model.Bill {
	obj := new(model.Bill)
	err := a.db.QueryRow(a.getScript("get/bill"), id).Scan(&obj.ID, &obj.UserID, &obj.ShopID, &obj.TotalDiscount, &obj.Timestamp)
	if err != nil {
		log.Printf("error getting bill: %s", err.Error())
		return nil
	}
	return obj
}

func (a *Adapter) SaveBill(bill *model.Bill) error {
	smt, err := a.db.Prepare(a.getScript("insert/bill"))
	if err != nil {
		log.Printf("error preparing statement: %s", err.Error())
		return err
	}
	t := time.Now()
	bill.Timestamp = t.Format("2006-01-02 15:04:05")
	if result, e := smt.Exec(bill.UserID, bill.ShopID, bill.TotalDiscount, bill.Timestamp); e != nil {
		log.Printf("error executing sql: %s", e.Error())
	} else {
		bill.ID, _ = result.LastInsertId()
	}

	return err
}

func (a *Adapter) UpdateBill(bill *model.Bill) error {

	_, err := a.db.Exec(a.getScript("update/bill"), bill.UserID, bill.ShopID, bill.TotalDiscount, bill.Timestamp, bill.ID)

	return err
}

func (a *Adapter) DeleteBillByID(id int64) error {
	tx, e := a.db.Begin()
	if e != nil {
		return e
	}
	result, err := tx.Exec(a.getScript("delete/bill"), id)
	if err != nil {
		return err
	}
	if count, _ := result.RowsAffected(); count != 1 {
		log.Printf("error more than one row was affected, rolling back transcation")
		_ = tx.Rollback()
		return errors.New("error more than one row was affected, rolling back transaction")
	}

	if e = tx.Commit(); e != nil {
		return e
	}

	return nil
}
