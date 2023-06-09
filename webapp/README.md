## Used Command
```
# initialize module
go mod init github.com/uyutaka/introduction_to_testing_in_go/webapp

# add packages
go get github.com/go-chi/chi/v5
go get github.com/go-chi/chi/v5/middleware
go get github.com/jackc/pgx/v4
```


## Commands
```
cd ./webapp/cmd/web
go run .

cd ./webapp
go test ./...
go test -tags=integration ./... # including dbrepo (integration tests)


# run db
cd webapp
docker compose up

# docker
docker system prune -a                          # remove cache
docker ps                                       # check container id
docker inspect <container id> | grep IPAddress    # get container IP
docker compose run <service name> bash                # go inside container
```
