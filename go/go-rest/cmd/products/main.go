package main

import (
	"context"
	"os"
	"os/signal"
	"time"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"go-tdd/cmd/products/http_handler"
	"go-tdd/internal/products"
)

func main() {
	// Inicialización de Echo
	e := echo.New()
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())

	// Se crea mapa para 'persistencia' de datos
	productMap := make(map[int32]*products.Product)
	// Se crea repositorio con su mapa de dependencia
	repo := products.NewRepository(productMap)
	// Se crea service que tiene como dependencia su repo
	service := products.NewService(repo)
	// Se inicializan las URI de echo
	http_handler.NewServerHandler(e, service)

	go func() {
		if err := e.Start(":8080"); err != nil {
			e.Logger.Info("shutting down the server: " + err.Error())
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}
