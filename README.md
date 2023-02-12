# surgicare-api-gateway
# ni-go-boiler-plate


Client Wrapper for SQL, Kafka, Cache
Test Usage of SQL, Kafka, Cache
error wrapper, instrumentation


export GOPRIVATE=github.com/pharmeasy/surgicare-proto

go get github.com/pharmeasy/surgicare-proto

vim ~/.netrc
machine github.com login <username> password <personal_access_token>

create database ni-api-gateway OR update the db config in app.json
After running once, it will create a user db, add id, name, password

run main.go
http port : 
grpc port : 

To see grpc services
instal grpc_cli & do
/opt/homebrew/Cellar/grpc/1.46.0/bin/grpc_cli ls localhost:8000  -l
