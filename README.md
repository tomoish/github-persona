# read-me

# Get started
### Git clone
```
$ git clone https://github.com/tomoish/read-me.git
```

### run http server
1. run main.go
    ```
    # ./src
    $ go run main.go
    ```
    or
    ```
    $ docker-compose exec app go run main.go
    ```

3. Test it out at http://localhost:8080/ (username)

### Using Docker
1. docker compose up
    ```
    # ./
    $ docker compose up -d
    ```

2. Test it out at http://localhost:8080/ (username)

### Format code
1. install golangci-lint and pre-commit
    ```
    $ berw install golangci-lint
    $ brew install pre-commit
    ```
2. setting for pre-commit
    ```
    $ pre-commit install -t pre-commit
    ```
