package services_test

import (
	"testing"

	"github.com/eron97/gomocks.git/services"
	"github.com/golang/mock/gomock"
)

func TestMyTodoListService_GetTodoList(t *testing.T) {
	// controlador mock
	ctrl := gomock.NewController
	defer ctrl.Finish()

	// instância do mock
	mockTodoListService := mock_services.NewMockTodoListService(ctrl)

	// config. comportamento esperado do teste

	mockTodoListService.EXPECT().GetTodoList(gomock.Any()).Return([]services.TodoItem{
		{ID: 1, Task_Name: "Task 1", Priority: "High"},
		{ID: 2, Task_Name: "Task 2", Priority: "Medium"},
	}, nil)

}
