# efishery

The service has 5 endpoints:

1. `localhost:8080/api/v1/efishery/user` HTTP POST endpoint for user to register.
2. `localhost:8080/api/v1/efishery/login` HTTP POST endpoint to get token.
3. `localhost:8080/api/v1/efishery/validate` HTTP POST endpoint to validate token.
4. `localhost:8080/api/v1/efishery/storages` HTTP GET endpoint to get data from storage with converted currency.
5. `localhost:8080/api/v1/efishery/aggregate` HTTP GET endpoint to get aggregated data from storage.

To start using the app first:
#### Clone the project
```bash
git clone https://github.com/nurroy/efishery.git
```

#### Go to directory
```bash
cd efishery
```

#### Set environtment
You can set environtment in `.env` file.

or you can directly run the app

#### Running App
For running the app type
```bash
# run
go run main.go
```

You can use `eFishery.postman_collection.json` to try

#### Using step
1. register your account using `register endpoint` and you will get password for login.
2. Use password from register to `login endpoint` for get jwt token.
3. You can validate token using `validate endpoint` endpoint
4. use your token to access `storages endpoint`, can be accessed by any role.
5. You can also access `aggregate endpoint` if you have admin role.
