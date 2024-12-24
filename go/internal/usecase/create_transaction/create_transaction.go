package create_transaction

import (
	"context"

	"github.com.br/devfullcycle/fc-ms-wallet/internal/entity"
	"github.com.br/devfullcycle/fc-ms-wallet/internal/gateway"
	"github.com.br/devfullcycle/fc-ms-wallet/pkg/events"
	"github.com.br/devfullcycle/fc-ms-wallet/pkg/uow"
)

type CreateTransactionInputDto struct {
	AccountIDFrom string `json:"account_id_from"`
	AccountIDTo   string `json:"account_id_to"`
	Amount        float64
}

type CreateTransactionOutputID struct {
	ID            string  `json:"id"`
	AccountIDFrom string  `json:"account_id_from"`
	AccountIDTo   string  `json:"account_id_to"`
	Amount        float64 `json:"amount"`
}

type BalanceUpdatedOutputDTO struct {
	AccountIDFrom         string  `json:"account_id_from"`
	AccountIDTo           string  `json:"account_id_to"`
	BallanceAccountIDFrom float64 `json:"balance_account_id_from"`
	BallanceAccountIDTo   float64 `json:"balance_account_id_to"`
}

type CreateTransactionUseCase struct {
	Uow                uow.UowInterface
	EventDispatcher    events.EventDispatcherInterface
	TransactionCreated events.EventInterface
	BalanceUpdated     events.EventInterface
}

func NewCreateTransactionUseCase(
	uow uow.UowInterface,
	eventDispatcher events.EventDispatcherInterface,
	transactionCreated events.EventInterface,
	balanceUpdated events.EventInterface) *CreateTransactionUseCase {
	return &CreateTransactionUseCase{
		Uow:                uow,
		EventDispatcher:    eventDispatcher,
		TransactionCreated: transactionCreated,
		BalanceUpdated:     balanceUpdated,
	}
}

func (uc *CreateTransactionUseCase) Execute(cxt context.Context, input CreateTransactionInputDto) (*CreateTransactionOutputID, error) {
	output := &CreateTransactionOutputID{}
	balanceUpdatedOutput := &BalanceUpdatedOutputDTO{}
	err := uc.Uow.Do(cxt, func(_ *uow.Uow) error {
		accountRepository := uc.getAccountRepository(cxt)
		transactionRepository := uc.getTransactionRepository(cxt)
		accountFrom, err := accountRepository.FindById(input.AccountIDFrom)
		if err != nil {
			return err
		}
		accountTo, err := accountRepository.FindById(input.AccountIDTo)
		if err != nil {
			return err
		}
		transaction, err := entity.NewTransaction(accountFrom, accountTo, input.Amount)
		if err != nil {
			return err
		}
		// err = accountRepository.UpdateBalance(accountFrom)

		// if err != nil {
		// 	return err
		// }

		// err = accountRepository.UpdateBalance(accountTo)

		// if err != nil {
		// 	return err
		// }

		err = transactionRepository.Create(transaction)
		if err != nil {
			return err
		}

		output.ID = transaction.ID
		output.AccountIDFrom = input.AccountIDFrom
		output.AccountIDTo = input.AccountIDTo
		output.Amount = input.Amount

		balanceUpdatedOutput.AccountIDFrom = input.AccountIDFrom
		balanceUpdatedOutput.AccountIDTo = input.AccountIDTo
		balanceUpdatedOutput.BallanceAccountIDFrom = accountFrom.Balance
		balanceUpdatedOutput.BallanceAccountIDTo = accountTo.Balance
		return nil
	})

	if err != nil {
		return nil, err
	}

	uc.TransactionCreated.SetPayload(output)
	uc.EventDispatcher.Dispatch(uc.TransactionCreated)

	uc.BalanceUpdated.SetPayload(balanceUpdatedOutput)
	uc.EventDispatcher.Dispatch(uc.BalanceUpdated)

	return output, nil
}

func (uc *CreateTransactionUseCase) getAccountRepository(cxt context.Context) gateway.AccountGateway {
	repo, err := uc.Uow.GetRepository(cxt, "AccountDB")
	if err != nil {
		panic(err)
	}
	return repo.(gateway.AccountGateway)

}

func (uc *CreateTransactionUseCase) getTransactionRepository(cxt context.Context) gateway.TransactionGateway {
	repo, err := uc.Uow.GetRepository(cxt, "TransactionDB")
	if err != nil {
		panic(err)
	}
	return repo.(gateway.TransactionGateway)
}
