package main

import (
	"flag"
	"fmt"
	"net"
	"net/http"
	"os"
	"time"

	"github.com/default23/crypto-api/application/transaction/transport/rpc"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"

	transactionendpoint "github.com/default23/crypto-api/application/transaction/endpoint"
	"github.com/default23/crypto-api/application/transaction/service"
	transactionHTTP "github.com/default23/crypto-api/application/transaction/transport/http"
	"github.com/default23/crypto-api/domain"
	"github.com/default23/crypto-api/infrastructure/config"
	"github.com/default23/crypto-api/infrastructure/logging"
)

func main() {
	var configPath string
	logger := logging.New()

	flag.StringVar(&configPath, "config", "", "path to configuration file")
	flag.Parse()

	if configPath == "" {
		panic("config path is not specified. Use '--config=/path/to/config' parameter to run application")
	}

	configFile, err := os.OpenFile(configPath, os.O_RDONLY, 0)
	if err != nil {
		logger.Fatal(err)
		return
	}

	cfg, err := config.Parse(configFile)
	if err != nil {
		logger.Fatal(err)
		return
	}

	seed, err := domain.NewSeed(cfg.Seed)
	if err != nil {
		panic(fmt.Sprintf("seed is invalid: %s", err))
	}

	go startJSONRPC(*cfg, logger, seed)

	router := configureRouter(*cfg, logger, seed)

	srv := &http.Server{
		Handler:      router,
		Addr:         cfg.Server.Addr(),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	logger.Fatal(srv.ListenAndServe())
}

func startJSONRPC(c config.Config, logger *logrus.Entry, seed domain.Seed) {
	if c.Server.JsonRPCPort == "" {
		return
	}

	httpListener, err := net.Listen("tcp", fmt.Sprintf(":%s", c.Server.JsonRPCPort))
	if err != nil {
		logger.WithError(err).Error("failed to start listening on address: %s", c.Server.JsonRPCPort)
		return
	}

	transactionService := service.NewTransactionService(seed)
	transactionEndpoints := transactionendpoint.NewEndpoints(transactionService)
	jsonrpcHandler := rpc.NewJSONRPCHandler(transactionEndpoints, logger)

	logger.Fatal(http.Serve(httpListener, jsonrpcHandler))
}

func configureRouter(c config.Config, logger *logrus.Entry, seed domain.Seed) *mux.Router {
	router := mux.NewRouter()
	router.Use(logging.WithLoggerMiddleware(logger))

	v1Router := router.PathPrefix("/api/v1").Subrouter()

	// configure transaction layer
	{
		transactionService := service.NewTransactionService(seed)
		transactionEndpoints := transactionendpoint.NewEndpoints(transactionService)
		transactionHTTPHandlers := transactionHTTP.NewHTTPHandler(transactionEndpoints)

		transactionHTTPHandlers.Register(v1Router)
	}

	return router
}
