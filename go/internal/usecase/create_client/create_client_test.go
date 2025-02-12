package create_client

import (
	"testing"

	"github.com.br/devfullcycle/fc-ms-wallet/internal/usecase/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateClientUseCase_Execute(t *testing.T) {
	m := &mocks.ClientGatewayMock{}
	m.On("Save", mock.Anything).Return(nil)
	uc := NewCreateClientUseCase(m)

	output, err := uc.Execute(CreateClientInputDto{
		Name:  "John Doe",
		Email: "j@j",
	})
	assert.Nil(t, err)
	assert.NotNil(t, output)
	assert.Equal(t, "John Doe", output.Name)
	assert.Equal(t, "j@j", output.Email)
	m.AssertExpectations(t)
	m.AssertNumberOfCalls(t, "Save", 1)
}
