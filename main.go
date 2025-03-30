package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/merofuruya/search/core/cache"
	"github.com/merofuruya/search/core/database"
	"github.com/merofuruya/search/core/http/middleware"
	httputil "github.com/merofuruya/search/core/http/util"
	"github.com/merofuruya/search/core/logging"
)

func initDatabase() {
	args := GetArgs()
	databaseConfig := database.DatabaseConfig{
		Host:     args.PostgresHost,
		Port:     args.PostgresPort,
		Database: args.PostgresDatabase,
		User:     args.PostgresUser,
		Password: args.PostgresPassword,
		Tls:      args.PostgresTls,
	}
	err := database.InitDatabase(context.Background(), databaseConfig)
	if err != nil {
		panic(err)
	}
}

func initRedis() {
	args := GetArgs()
	redisConfig := redis.Options{
		Addr:     fmt.Sprintf("%s:%d", args.RedisHost, args.RedisPort),
		Password: args.RedisPassword,
		DB:       0,
	}
	err := cache.InitRedis(context.Background(), redisConfig)
	if err != nil {
		panic(err)
	}
}

func initZerolog() {
	args := GetArgs()
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	if args.Format.Format == FormatArgPretty {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Warn().Err(err).Msg("Failed to load .env file")
	}

	initZerolog()
	logger := logging.GetLogger("server")

	initDatabase()
	initRedis()

	router := chi.NewRouter()
	router.Use(middleware.CorsMiddlewareFactory())
	router.Mount("/api", ApiRouter())
	router.Group(HealthRouter)
	router.Group(StandaloneRouter)

	// debug route
	httputil.DebugRoute("", router)

	args := GetArgs()

	host := httputil.ParseHost(args.Host, args.Port)
	logger.Info().Str("host", host).Msg("Starting server")
	err = http.ListenAndServe(host, router)
	if err != nil {
		logger.Fatal().Err(err).Msg("Failed to start server")
	}
}
