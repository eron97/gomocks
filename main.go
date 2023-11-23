package main

import (
	"github.com/eron97/gomocks.git/api/handlers"
	"github.com/eron97/gomocks.git/api/routes"
	"github.com/eron97/gomocks.git/database/mysql"
	"github.com/eron97/gomocks.git/services"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// Carrega credenciais do banco de dados
	dbConnector := &mysql.MySQLConnector{
		Credentials: "admin:mysql-todolist@tcp(database-1.cpj0eavfzshu.us-east-1.rds.amazonaws.com:3306)/db_todolist",
	}

	// Conecta ao banco de dados
	db, err := dbConnector.Connect()
	if err != nil {
		return
	}

	dbConnector.Connect()

	// Carrega uma instância do MySQLConnector como a implementação da interface DBConnector
	todoListServices := &services.MyTodoListService{
		DBConnector: dbConnector,
	}

	defer db.Close() // Garante que a conexão será fechada no final da função main

	// Inicia o manipulador
	todoListHandler := &handlers.TodoListHandler{
		TodoListService: todoListServices,
	}

	// Configuração do roteador Gin
	router := routes.SetupRouter(todoListHandler)
	// Inicia o servidor
	router.Run(":8080")
}
