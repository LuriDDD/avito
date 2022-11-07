package store

import (
	"database/sql"
)

// Store ...
type Store struct {
	db                         *sql.DB
	userRepository             *UserRepository
	reservedFundsRepository    *ReservedFundsRepository
	accountingReportRepository *AccountingReportRepository
}

// New ...
func New(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

// User ...
func (s *Store) User() *UserRepository {
	if s.userRepository == nil {
		s.userRepository = &UserRepository{
			store: s,
		}
	}

	return s.userRepository
}

func (s *Store) ReservedFunds() *ReservedFundsRepository {
	if s.reservedFundsRepository == nil {
		s.reservedFundsRepository = &ReservedFundsRepository{
			store: s,
		}
	}

	return s.reservedFundsRepository
}

func (s *Store) AccountingReport() *AccountingReportRepository {
	if s.accountingReportRepository == nil {
		s.accountingReportRepository = &AccountingReportRepository{
			store: s,
		}
	}

	return s.accountingReportRepository
}
