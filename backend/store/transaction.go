package store

import "gorm.io/gorm"

// GORM Postgres Implementation of TransactionStore

type PostgreTransactionStore struct {
	db *gorm.DB
}

func (pts *PostgreTransactionStore) GetTransaction() {
}

func (pts *PostgreTransactionStore) GetTransactions() {
}

func (pts *PostgreTransactionStore) CreateTransaction() {
}

func (pts *PostgreTransactionStore) UpdateTransaction() {}

func (pts *PostgreTransactionStore) DeleteTransaction() {}
