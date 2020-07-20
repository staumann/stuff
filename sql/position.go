package sql

import (
	"github.com/staumann/caluclation/model"
	"log"
)

type positionRepo struct {
	adapter *Adapter
}

func (p *positionRepo) GetPositionByID(id int64) *model.Position {
	pos := new(model.Position)
	row := p.adapter.db.QueryRow(p.adapter.getScript("get/position"), id)
	err := row.Scan(&pos.ID, &pos.Amount, &pos.Description, &pos.SinglePrice, &pos.Discount, &pos.BillID, &pos.Type)
	if err != nil {
		log.Printf("error getting position: %s", err.Error())
		return nil
	}
	return pos
}

func (p *positionRepo) SavePosition(pos *model.Position) error {
	smt, err := p.adapter.db.Prepare(p.adapter.getScript("insert/position"))
	if err != nil {
		log.Printf("error preparing statement: %s", err.Error())
		return err
	}

	if result, e := smt.Exec(pos.Amount, pos.Description, pos.SinglePrice, pos.Discount, pos.BillID, pos.Type); e != nil {
		log.Printf("error executing sql: %s", e.Error())
	} else {
		pos.ID, _ = result.LastInsertId()
	}

	return err
}

func (p *positionRepo) UpdatePosition(pos *model.Position) error {
	_, err := p.adapter.db.Exec(p.adapter.getScript("update/position"), pos.Amount, pos.Description, pos.SinglePrice, pos.Discount, pos.BillID, pos.Type)

	return err
}

func (p *positionRepo) DeleteByID(id int64) error {
	_, err := p.adapter.db.Exec(p.adapter.getScript("delete/position"), id)

	return err
}

func (p *positionRepo) GetByBillID(id int64) []*model.Position {
	rows, err := p.adapter.db.Query(p.adapter.getScript("get/positions"), id)

	if err != nil {
		log.Printf("error getting positions by bill id: %s", err.Error())
		return nil
	}
	positions := make([]*model.Position, 0)
	for rows.Next() {
		pos := new(model.Position)
		if err := rows.Scan(&pos.ID, &pos.Amount, &pos.Description, &pos.SinglePrice, &pos.Discount, &pos.BillID, &pos.Type); err != nil {
			log.Printf("error getting data for bill id %d: %s", id, err.Error())
		} else {
			positions = append(positions, pos)
		}
	}
	return positions
}
