# API Specification
https://www.postman.com/final-project-engineering-5/workspace/github-com-sinulingga23-voucher

# Migration
```migrate -path migration/ -database "mysql://<username>:<password>@tcp(<host>:<port>)/<db-name>" up 3```

# Environment Keys
- APP_ENV=<test|production>

## If the APP_ENV=test. The env below should be set:
- MYSQL_DB_TEST_USER
- MYSQL_DB_TEST_PASSWORD
- MYSQL_DB_TEST_HOST
- MYSQL_DB_TEST_PORT
- MYSQL_DB_TEST_NAME

## If the APP_ENV=production. The env below should be set:
- MYSQL_DB_USER
- MYSQL_DB_PASSWORD
- MYSQL_DB_HOST
- MYSQL_DB_PORT
- MYSQL_DB_NAME

# Integration Test
After setting the env keys for DB. You can run the Integration Test using the command: ```go test -v ./repository/ ./service/```

Output:
```=== RUN   TestBrandRepository_Create_Success
--- PASS: TestBrandRepository_Create_Success (0.01s)
=== RUN   TestBrandRepository_IsExistsById_Exists
--- PASS: TestBrandRepository_IsExistsById_Exists (0.01s)
=== RUN   TestBrandRepository_IsExistsById_NotExists
--- PASS: TestBrandRepository_IsExistsById_NotExists (0.00s)
=== RUN   TestVoucherRepository_Create_Success
--- PASS: TestVoucherRepository_Create_Success (0.02s)
=== RUN   TestVoucherRepository_IsExistsById_Exists
--- PASS: TestVoucherRepository_IsExistsById_Exists (0.02s)
=== RUN   TestVoucherRepository_IsExistsById_NotExists
--- PASS: TestVoucherRepository_IsExistsById_NotExists (0.00s)
=== RUN   TestVoucherRepository_FindById_Success
--- PASS: TestVoucherRepository_FindById_Success (0.01s)
=== RUN   TestVoucherRepository_FindAllByBrandId
--- PASS: TestVoucherRepository_FindAllByBrandId (0.04s)
PASS
ok      github.com/sinulingga23/voucher/repository      0.113s
=== RUN   TestBrandService_Create_Success
--- PASS: TestBrandService_Create_Success (0.01s)
=== RUN   TestVoucherService_Create_Success
--- PASS: TestVoucherService_Create_Success (0.02s)
=== RUN   TestVoucherService_FindById_Found
--- PASS: TestVoucherService_FindById_Found (0.02s)
PASS
ok      github.com/sinulingga23/voucher/service 0.052s
```


# Tech Stack
- Go 1.19
- MySQL 8.0
- Gin

# Run
After setting the env keys for DB. You can run the services using the command:
```
go run main.go
```

Output:
```
[GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.

[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:   export GIN_MODE=release
 - using code:  gin.SetMode(gin.ReleaseMode)

[GIN-debug] POST   /brand                    --> github.com/sinulingga23/voucher/handler.Handler.CreateBrand-fm (3 handlers)
[GIN-debug] POST   /voucher                  --> github.com/sinulingga23/voucher/handler.Handler.CreateVoucher-fm (3 handlers)
[GIN-debug] POST   /transaction/redemption   --> github.com/sinulingga23/voucher/handler.Handler.CreateRedemption-fm (3 handlers)
[GIN-debug] GET    /voucher                  --> github.com/sinulingga23/voucher/handler.Handler.FindVoucherById-fm (3 handlers)
[GIN-debug] GET    /voucher/brand            --> github.com/sinulingga23/voucher/handler.Handler.FindAllVoucherByBrandId-fm (3 handlers)
[GIN-debug] GET    /transaction/redemption   --> github.com/sinulingga23/voucher/handler.Handler.FindRedemptionById-fm (3 handlers)
```

The services will be served on:8080.
