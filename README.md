# get-pg-schema
The utility helps to verify that you can connect to a PostgreSQL database and get a list of available databases and tables in that database and public schema.


# Setup
- put the following commands to cmd.
```sh
git clone git@github.com:stevenstr/get-pg-schema.git
cd ./get-pg-schema
go build .
docker-compose up
```

# Usage
- for testing.
```sh
go run .\get-pg-schema.go localhost 5432 stevenstr pass master
```
- for "real" usage.
```sh
get-pg-schema.exe localhost 5432 stevenstr pass master
```

# Acces to docs
- puth to cmd and press enter.
```sh
godoc -http=:6060
```
- go to browser.
```sh
http://localhost:6060/pkg/get-pg-schema/
```