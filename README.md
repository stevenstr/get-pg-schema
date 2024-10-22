# get-pg-schema
The utility helps to verify that you can connect to a PostgreSQL database and get a list of available databases and tables in that database and public schema.


# Setup
```sh
git clone git@github.com:stevenstr/get-pg-schema.git
cd ./get-pg-schema
go build .
docker-compose up
```

# Usage
for testing
```sh
go run .\get-pg-schema.go localhost 5432 stevenstr pass master
```
for "real" usage
```sh
get-pg-schema.exe localhost 5432 stevenstr pass master
```