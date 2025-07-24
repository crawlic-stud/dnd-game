package server

import (
	"context"
	"dnd-game/internal/db"
	"dnd-game/internal/util/helper"
	"dnd-game/internal/util/services"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/joho/godotenv/autoload"
)

type Store struct {
	*db.Queries
	Pool *pgxpool.Pool
}

type Server struct {
	*helper.ServerHelper

	Store *Store

	Auth *services.AuthService
}

func (s *Store) Transaction(ctx context.Context, txFunc func(tx *db.Queries) error) error {
	tx, err := s.Pool.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	qTx := s.Queries.WithTx(tx)

	err = txFunc(qTx)
	if err != nil {
		return err
	}

	return tx.Commit(ctx)
}

func NewHTTPServer(handler http.Handler) *http.Server {
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		log.Fatal(err)
	}

	return &http.Server{
		Addr:         fmt.Sprintf(":%d", port),
		Handler:      handler,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}
}
