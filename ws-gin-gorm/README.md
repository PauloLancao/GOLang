## WS Gin Gorm

### Get Gin
```
go get -u github.com/gin-gonic/gin
```

### Get Gorm
```
go get -u github.com/jinzhu/gorm
```

### Install mingw via powershell admin
***
Need choco installed
***

```
choco install mingw
```

### Get sqlite3
```
go get github.com/mattn/go-sqlite3
```

### Update packages and run
```
go get .
go run .
```

### Get books
```
curl -s http://localhost:8080/books \
    --header "Content-Type: application/json" \
    --request "GET"
```

### Get book by id
```
curl -s http://localhost:8080/books/1 \
    --header "Content-Type: application/json" \
    --request "GET"
```

### Post book
```
curl -s http://localhost:8080/books \
    --include \
    --header "Content-Type: application/json" \
    --request "POST" \
    --data '{"title": "The Modern Sound of Betty Carter","author": "Betty Carter"}'
```

### Patch book
```
curl -s http://localhost:8080/books/1 \
    --include \
    --header "Content-Type: application/json" \
    --request "PATCH" \
    --data '{"title": "Update The Modern Sound of Betty Carter","author": "Update Betty Carter"}'
```

### Delete book
````
curl -s http://localhost:8080/books/1 \
    --include \
    --header "Content-Type: application/json" \
    --request "DELETE"
```