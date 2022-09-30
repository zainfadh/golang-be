#!/usr/bin/env bash

#main.go
export SERVER_PORT=":40001"

#routers/router.go
export SWAGGER_HOST="localhost"

#config/database.go
export DATABASE_USER="postgres"
export DATABASE_PASSWORD="password"
export DATABASE_NAME="postgres"
export DATABASE_HOST="localhost"
export DATABASE_PORT="8080"
export DATABASE_SSLMODE="disable"
export DATABASE_TIMEOUT="5"

#routers/router.go
export APPS_DEBUG="debug"

# For Running Local Development
go run main.go

# For Running at Development / Production Server

# default_nohup_file_name=nohup.out
# current_time=$(date "+%Y-%m-%d_%H.%M.%S")
# nohup_backup_file_name=$default_nohup_file_name.$current_time
# cp $default_nohup_file_name $nohup_backup_file_name

# go build
# nohup ./golang-be > $default_nohup_file_name 2>&1 &