# First of all I really appreciate Victor Steven craft
## This project base on https://github.com/victorsteven/Go-JWT-Postgres-Mysql-Restful-API 
I just fixed some bug in gist and update range loop there, so if someone want to fork my project please credit Victor Steven project

### For testing 
`go test -cover -coverprofile controllertestresult.out ./tests/controllertests/`
`go test -cover -coverprofile modelstestresult.out ./tests/modeltests/`
#### and for reading test coverage 
`go tool cover -html=controllertestresult.out`
`go tool cover -html=cmodelstestresult.out`
