package create_transaction

import (
	"context"
	"testing"

	"github.com.br/devfullcycle/fc-ms-wallet/internal/entity"
	"github.com.br/devfullcycle/fc-ms-wallet/internal/event"
	"github.com.br/devfullcycle/fc-ms-wallet/internal/usecase/mocks"
	"github.com.br/devfullcycle/fc-ms-wallet/pkg/events"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateTransactionUseCase_Execute(t *testing.T) {
	client1, _ := entity.NewClient("client 1", "j@j.com")
	account1 := entity.NewAccount(client1)
	account1.Credit(1000)

	client2, _ := entity.NewClient("client 2", "j@j2.com")
	account2 := entity.NewAccount(client2)
	account2.Credit(1000)

	mockUwo := &mocks.UowMock{}
	mockUwo.On("Do", mock.Anything, mock.Anything).Return(nil)

	inputDto := CreateTransactionInputDto{
		AccountIDFrom: account1.ID,
		AccountIDTo:   account2.ID,
		Amount:        100,
	}
	dispatcher := events.NewEventDispatcher()
	eventTransaction := event.NewTransactionCreated()
	eventBalance := event.NewBalanceUpdated()
	cxt := context.Background()

	uc := NewCreateTransactionUseCase(mockUwo, dispatcher, eventTransaction, eventBalance)
	output, err := uc.Execute(cxt, inputDto)
	assert.Nil(t, err)
	assert.NotNil(t, output)
	mockUwo.AssertExpectations(t)
	mockUwo.AssertNumberOfCalls(t, "Do", 1)
}
