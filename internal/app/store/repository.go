package store

import (
	"database/sql"

	"http-rest-api/internal/app/model"
)

// UserRepository ...
type UserRepository struct {
	store *Store
}

// Create ...
func (r *UserRepository) ReplenishOrCreate(u *model.User) error {

	user, err := r.GetBalance(u.ID)
	if err == ErrRecordNotFound {
		r.store.db.QueryRow(
			"INSERT INTO users (id, balance) VALUES ($1, $2)",
			u.ID,
			u.Balance,
		)
		return nil
	}

	if err == nil {
		if user.Balance+u.Balance > 0 {

			r.store.db.QueryRow(
				"UPDATE users SET balance = $1 WHERE id = $2",
				user.Balance+u.Balance,
				u.ID,
			)
			return nil
		} else {
			return ErrLowBalance
		}
	}
	return ErrUnknown
}

// Find ...
func (r *UserRepository) GetBalance(id int) (*model.User, error) {
	u := &model.User{ID: id}
	if err := r.store.db.QueryRow(
		"SELECT balance FROM users WHERE id = $1",
		id,
	).Scan(
		&u.Balance,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrRecordNotFound
		}

		return nil, err
	}

	return u, nil
}

type ReservedFundsRepository struct {
	store *Store
}

func (r *ReservedFundsRepository) Create(f *model.ReservedFund) error {
	userRepository := r.store.User()
	user := &model.User{ID: f.ID, Balance: -f.Price}
	err := userRepository.ReplenishOrCreate(user)
	if err != nil {
		return err
	}
	var id_order int
	err = r.store.db.QueryRow(
		"SELECT id_order FROM reserved_funds WHERE id_order = $1",
		f.IDOrder,
	).Scan(&id_order)
	if err == nil {
		return ErrRowExist
	}

	if err == sql.ErrNoRows {
		r.store.db.QueryRow(
			"INSERT INTO reserved_funds (id_user, id_service, id_order, price) VALUES ($1, $2, $3, $4)",
			f.ID,
			f.IDService,
			f.IDOrder,
			f.Price,
		)
		return nil
	}
	return err
}

func (r *ReservedFundsRepository) Recognize(f *model.ReservedFund) error {
	var id_order int
	if err := r.store.db.QueryRow(
		"DELETE FROM reserved_funds WHERE id_order = $1 RETURNING id_order",
		f.IDOrder,
	).Scan(
		&id_order,
	); err != nil {
		if err == sql.ErrNoRows {
			return ErrRowNotExist
		}
		return err
	}
	// add report
	return nil
}

type AccountingReportRepository struct {
	store *Store
}

func (r *AccountingReportRepository) AddOrCreate(a *model.AccountingReport) error {
	var funds int
	if err := r.store.db.QueryRow(
		"SELECT funds FROM accounting_report WHERE id_service = $1 AND month = $2 AND year = $3",
		a.IDService,
		a.Month,
		a.Year,
	).Scan(
		&funds,
	); err != nil {
		r.store.db.QueryRow(
			"UPDATE accounting_report SET funds = $1 WHERE id_service = $2 AND month = $3 AND year = $4",
			a.Funds+funds,
			a.IDService,
			a.Month,
			a.Year,
		)
	} else {
		r.store.db.QueryRow(
			"INSERT INTO users (id_service, month, year, funds) VALUES ($1, $2, $3, $4)",
			a.IDService,
			a.Month,
			a.Year,
			a.Funds+funds,
		)
	}
	return nil
}

func (r *AccountingReportRepository) GetReport(u *model.User) error {
	return nil
}
