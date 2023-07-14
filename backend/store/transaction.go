package store

import "gorm.io/gorm"

// Really only using interface here incase I change my mind on DB later
type TransactionStore interface {
	GetTransaction()
	GetTransactions()
	CreateTransaction()
	UpdateTransaction()
	DeleteTransaction()
}

type PostgreTransactionStore struct {
	db *gorm.DB
}

func NewPostgreTransactionStore(db *gorm.DB) *PostgreTransactionStore {
	return &PostgreTransactionStore{
		db: db,
	}
}

func (pts *PostgreTransactionStore) GetTransaction() {
}

func (pts *PostgreTransactionStore) GetTransactions() {
}

func (pts *PostgreTransactionStore) CreateTransaction() {
}

func (pts *PostgreTransactionStore) UpdateTransaction() {}

func (pts *PostgreTransactionStore) DeleteTransaction() {}
