package services_test

import (
	"context"
	"database/sql"
	"testing"

	"github.com/eron97/gomocks.git/mocks"
	"github.com/eron97/gomocks.git/services"
	"github.com/stretchr/testify/assert"
	gomock "go.uber.org/mock/gomock"
)

func TestMyTodoListService_GetTodoList(t *testing.T) {
	// Criar um novo controlador de mock
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Criar uma nova instância do mock para a interface mysql.DBConnector
	mockDBConnector := mocks.NewMockDBConnector(ctrl)

	// Configurar o comportamento esperado do mock
	mockDBConnector.EXPECT().Connect().Return(&sql.DB{}, nil)

	// Criar uma nova instância do serviço com o mock
	todoListService := &services.MyTodoListService{
		DBConnector: mockDBConnector,
	}

	// Chamar o método GetTodoList no serviço
	todoItems, err := todoListService.GetTodoList(context.Background())

	// Fazer as afirmações
	assert.NoError(t, err)
	assert.Len(t, todoItems, 2)
}
