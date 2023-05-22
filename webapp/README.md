## Used Command
```
# initialize module
go mod init github.com/uyutaka/introduction_to_testing_in_go/webapp

# add packages
go get github.com/go-chi/chi/v5
go get github.com/go-chi/chi/v5/middleware
```


## Commands
```
cd ./webapp/cmd/web
go run .

cd ./webapp
go test ./...

# run db
cd webapp
docker compose up

# docker
docker system prune -a                          # remove cache
docker ps                                       # check container ip
docker inspect ad2ca99be62d | grep IPAddress    # get container IP
docker compose run postgres bash                # go inside container
```
