package config

import (
	"backend/utils"
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5"
)

var Conn *pgx.Conn

func Init() {
	dsn := utils.GetEnv("DB_PATH")
	if dsn == "" {
		return
	}

	var err error
	Conn, err = pgx.Connect(context.Background(), dsn)
	if err != nil {
		log.Fatalf("Erro ao conectar no banco de dados: %v", err)
	}

	fmt.Println("Conectado ao banco")
}

func Close() {
	if Conn != nil {
		Conn.Close(context.Background())
		fmt.Println("Desconectado ao banco")
	}
}
