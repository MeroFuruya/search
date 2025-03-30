package main

import (
	"fmt"

	"github.com/alexflint/go-arg"
)

type formatArg struct {
	Format string
}

const (
	FormatArgJson   = "json"
	FormatArgPretty = "pretty"
)

func (f *formatArg) UnmarshalText(text []byte) error {
	f.Format = string(text)

	if f.Format != FormatArgJson && f.Format != FormatArgPretty {
		return fmt.Errorf("invalid format: %s", f.Format)
	}
	return nil
}

type Args struct {
	Format formatArg `arg:"--format" help:"Output format (json, pretty)" default:"json"`
	Port   uint16    `arg:"--port,env:SEARCH_PORT" help:"Port to listen on" default:"3000" `
	Host   string    `arg:"--host,env:SEARCH_HOST" help:"Host to listen on" default:"*"`
	// Postgres
	PostgresHost     string `arg:"--postgres-host,required,env:SEARCH_POSTGRES_HOST" help:"Postgres host"`
	PostgresPort     uint16 `arg:"--postgres-port,env:SEARCH_POSTGRES_PORT" help:"Postgres port" default:"5432"`
	PostgresDatabase string `arg:"--postgres-database,required,env:SEARCH_POSTGRES_DATABASE" help:"Postgres database"`
	PostgresUser     string `arg:"--postgres-user,required,env:SEARCH_POSTGRES_USER" help:"Postgres user"`
	PostgresPassword string `arg:"--postgres-password,required,env:SEARCH_POSTGRES_PASSWORD" help:"Postgres password"`
	PostgresTls      bool   `arg:"--postgres-tls,env:SEARCH_POSTGRES_TLS" help:"Use TLS for Postgres" default:"false"`
	// Redis
	RedisHost     string `arg:"--redis-host,required,env:SEARCH_REDIS_HOST" help:"Redis host"`
	RedisPort     uint16 `arg:"--redis-port,env:SEARCH_REDIS_PORT" help:"Redis port" default:"6379"`
	RedisPassword string `arg:"--redis-password,env:SEARCH_REDIS_PASSWORD" help:"Redis password"`
	RedisTls      bool   `arg:"--redis-tls,env:SEARCH_REDIS_TLS" help:"Use TLS for Redis" default:"false"`
}

var parsedArgs bool
var globalArgs Args

func parseArgs() Args {
	parsedArgs = true
	arg.MustParse(&globalArgs)
	return globalArgs
}

func GetArgs() Args {
	if !parsedArgs {
		parseArgs()
	}
	return globalArgs
}
