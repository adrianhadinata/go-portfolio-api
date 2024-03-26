package repository

import (
	"database/sql"
	"fmt"
	"portfolio-api/entity"
	"time"
)

type VisitsRepository interface {
	Insert(visitor entity.Visitor) (entity.Visitor, error)
}

type visitsRepository struct {
	db *sql.DB
}

func (v *visitsRepository) Insert(visitor entity.Visitor) (entity.Visitor, error) {
	fmt.Println(v.db)
	err := v.db.QueryRow(
		"INSERT INTO tx_visits (Visitor_IP, Date_Created) VALUES (?, ?) RETURNING Visit_ID, Visitor_IP, Date_Created",
		visitor.Visitor_IP,
		time.Now(),
	).Scan(&visitor.Visit_ID, &visitor.Visitor_IP, &visitor.Date_Created)

	if err != nil {
		return entity.Visitor{}, err
	}

	return visitor, nil
}

func NewVisitsRepository(db *sql.DB) VisitsRepository {
	return &visitsRepository{db: db}
}