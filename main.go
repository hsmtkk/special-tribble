package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/zap"
)

func main() {
	httpPortStr := os.Getenv("HTTP_PORT")
	if httpPortStr == "" {
		log.Fatal("HTTP_PORT must be defined")
	}
	httpPort, err := strconv.Atoi(httpPortStr)
	if err != nil {
		log.Fatalf("failed to parse as int; %s; %v", httpPortStr, err)
	}
	addr := fmt.Sprintf("0.0.0.0:%d", httpPort)

	logConfig := zap.NewDevelopmentConfig()
	logConfig.OutputPaths = []string{"stdout", "/tmp/log"}
	logger, err := logConfig.Build()
	if err != nil {
		log.Fatalf("failed to build logger; %v", err)
	}
	defer logger.Sync()
	sugar := logger.Sugar()

	sugar.Info("starting server at %s", addr)

	var handler handler = &myHandler{sugar}

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.GET("/", handler.Handle)
	e.Logger.Fatal(e.Start(addr))
}

type handler interface {
	Handle(ctx echo.Context) error
}

type myHandler struct {
	sugar *zap.SugaredLogger
}

func (h *myHandler) Handle(ctx echo.Context) error {
	h.sugar.Info("handle begin")
	h.sugar.Info("handle end")
	return ctx.String(http.StatusOK, "Hello, World!")
}
