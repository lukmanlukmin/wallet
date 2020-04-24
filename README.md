
# BE PRETEST DCI (Mini e-wallet)

__Asumption__<br>
[task & issue](https://www.evernote.com/shard/s355/client/snv?noteGuid=3ce40a4f-d15f-4f35-b4f8-9cfe6c0cc0ec&noteKey=c73db594218f81a5c5ea32b02274461b&sn=https%3A%2F%2Fwww.evernote.com%2Fshard%2Fs355%2Fsh%2F3ce40a4f-d15f-4f35-b4f8-9cfe6c0cc0ec%2Fc73db594218f81a5c5ea32b02274461b&title=BE%2BPRETEST%2BDCI%2B%2528Mini%2Be-wallet%2529)
please read this [document](https://github.com/lukmanlukmin/wallet/blob/master/assumption.md) for additional assumption

__Introduction__<br>
we can run this application by just run with [docker configuration](https://github.com/lukmanlukmin/wallet/blob/master/Dockerfile) without any problem or even with [docker-compose](https://github.com/lukmanlukmin/wallet/blob/master/docker-compose.yml) if we want database attached already by just simply run below command
```bash
docker-compose up --build
```
noted: database data will always reset every build, because not have persistence data on docker configuration.
but if we want to run without docker, you need continue follow below instruction:

__Manager__<br>
this project managed using [go-module](https://blog.golang.org/using-go-modules)

__Database__<br>
this application will perfectly run with postgreSQL-12

__Migration__<br>
this project use [gorm](https://gorm.io/) as ORM Library, but for some reason migration tool that to be used is using  [golang-migrate](https://github.com/golang-migrate/migrate)
* if this is first time usage, we need to build CLI-migrate to our environment as per [this documentation](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate)
```bash
$ go get -tags 'postgres' -u github.com/golang-migrate/migrate/v4/cmd/migrat
```
* migration file will be on path /migration/{domain}
* to run migration, follow below guide: 
```bash
migrate -verbose -source file://migration/postgresql -database postgres:'//username_database:password_databasehost:port/database_name?option=option_value' up
```
* example to make migration file
```bash
migrate create -ext sql -dir path/absolute_path scheme_name_anything dbdriver://username:password@host:port/dbname?option1=option_value&option2=option_value2
```

__Test__<br>
Library that used to test on this project using [testify](https://github.com/stretchr/testify)
* to run our test we can start with below comman
```bash
go test -v -cover ./...
or
go test -coverprofile=coverage.out ./...
```
* if we ned more information about our test result, we have option to run below command
```bash
go tool cover -func=coverage.out
```