package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/jackc/pgx/v5"
	// import postgresql driver
	_ "github.com/jackc/pgx/v5/stdlib"
)

type application struct {
	infoLog  *log.Logger
	errorLog *log.Logger
}

func main() {
	// config
	addr := flag.String("addr", ":4000", "HTTP web server address")
	dsn := flag.String("dsn", "postgres://snippetbox:password@localhost:5432/snippetbox", "PostgreSQL data source name")

	flag.Parse()

	// log config
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	app := &application{
		infoLog:  infoLog,
		errorLog: errorLog,
	}

	// db
	conn, err := connDB(context.Background(), *dsn)
	if err != nil {
		errorLog.Fatal(err)
	}

	var id, title string
	err = conn.QueryRow(context.Background(), "SELECT id, title FROM public.snippets WHERE id=$1", 1).Scan(&id, &title)
	if err != nil {
		errorLog.Fatal(err)
	}
	fmt.Println(id, title)

	defer func() {
		conn.Close(context.Background())
	}()

	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  app.routes(),
	}

	// Start a new web server
	infoLog.Printf("Starting server on %s", *addr)
	err = srv.ListenAndServe()
	if err != nil {
		errorLog.Fatal(err)
	}
}

func connDB(ctx context.Context, dsn string) (*pgx.Conn, error) {
	conn, err := pgx.Connect(ctx, dsn)
	if err != nil {
		return nil, fmt.Errorf("postgresql: unable to connect to db: %w", err)
	}

	if err = conn.Ping(ctx); err != nil {
		return nil, fmt.Errorf("postgresql: unable to ping db: %w", err)
	}
	return conn, nil
}
