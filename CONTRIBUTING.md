# Get started
### Git clone
```
$ git clone https://github.com/tomoish/github-persona.git
```
### Writing the .env File
Create your GitHub token (starting with ghp_) and enter it into the .env file located at ./src/.env. Use the following format:
```
GITHUB_TOKEN=<your github token>
GITHUB_TOKEN1=<your github token>
GITHUB_TOKEN2=<your github token>
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

3. Test it out at http://localhost:8080/create?username= (username)

#### Using Docker
1. docker compose up
    ```
    # ./
    $ docker compose up -d
    ```

2. Test it out at http://localhost:8080/create?username= (username)

### CI
#### golangci-lint
1. install golangci-lint
    ```
    $ berw install golangci-lint
    ```
2. run golanci-lint in ./src
    the setting file is ./.golangci.yml
    ```
    # ./src
    $ golangci-lint run --fix
    ```

#### pre-commit
1. install pre-commit
    ```
    $ brew install pre-commit
    ```
2. setting for pre-commit in root directory
    the setting file is ./.pre-commit-config.yaml
    ```
    $ pre-commit install -t pre-commit
    ```
