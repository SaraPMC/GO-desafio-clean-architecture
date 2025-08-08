package main

import (
	"database/sql"
	"fmt"
	"log"

	// mysql
	"github.com/devfullcycle/20-CleanArch/configs"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	log.Println("Inicializando banco de dados MySQL...")

	configs, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	db, err := sql.Open(configs.DBDriver, fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", configs.DBUser, configs.DBPassword, configs.DBHost, configs.DBPort, configs.DBName))
	if err != nil {
		log.Fatalf("Erro ao conectar com o banco: %v", err)
	}
	defer db.Close()

	// Testa a conexão
	if err := db.Ping(); err != nil {
		log.Fatalf("Erro ao fazer ping no banco: %v", err)
	}

	// Cria tabela
	createOrdersTable := `CREATE TABLE IF NOT EXISTS orders 
	(
		id varchar(255) NOT NULL, 
		price float NOT NULL, 
		tax float NOT NULL, 
		final_price float NOT NULL, 
		PRIMARY KEY (id)
	);`

	if _, err := db.Exec(createOrdersTable); err != nil {
		log.Fatalf("Erro ao criar tabela orders: %v", err)
	}

	log.Println("✓ Banco de dados inicializado com sucesso!")
}

//para rodar e criar a tabela orders:
//cd cmd/initdb; go run main.go

//Para acessar o BD no docker:
//docker ps (para ver o que esta rodando)
//docker exec -it <container_id> bash (para acessar o mysql dentro do container)
//docker exec -it 7db5c70e0059 bash
//mysql -u root -p orders
//senha root
//show databases;
