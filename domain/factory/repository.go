package factory

import "tiagoncardoso/imersaofc/domain/repository"

type RepositoryFactory interface {
	CreateTransactionRepository() repository.TransactionRepository
}
