## WS-Gin

### Update packages and run
```
go get .
go run .
```

### Get albums
```
curl http://localhost:8080/albums \
    --header "Content-Type: application/json" \
    --request "GET"
```

### Get album by id
```
curl http://localhost:8080/albums/2
```

### Post album
```
curl http://localhost:8080/albums \
    --include \
    --header "Content-Type: application/json" \
    --request "POST" \
    --data '{"id": "4","title": "The Modern Sound of Betty Carter","artist": "Betty Carter","price": 49.99}'
```