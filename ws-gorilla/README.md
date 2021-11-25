## WS Gorilla Mux

### Get Gorilla Mux
```
go get -u github.com/gorilla/mux
```

### Update packages and run
```
go get .
go run .
```

### Get welcome home
```
curl -s http://localhost:8080/ \
    --header "Content-Type: application/json" \
    --request "GET"
```

### Get Events
```
curl -s http://localhost:8080/events \
    --header "Content-Type: application/json" \
    --request "GET"
```

### Get Event by Id
```
curl -s http://localhost:8080/events/1 \
    --header "Content-Type: application/json" \
    --request "GET"
```

### Create Event
```
curl -s http://localhost:8080/event \
    --include \
    --header "Content-Type: application/json" \
    --request "POST" \
    --data '{"id":"2","title":"The Modern Sound of Betty Carter","description": "Betty Carter"}'
```

### Patch Event
```
curl -s http://localhost:8080/events/1 \
    --include \
    --header "Content-Type: application/json" \
    --request "PATCH" \
    --data '{"title": "Update The Modern Sound of Betty Carter","description": "Update Betty Carter"}'
```

### Delete Event
```
curl -s http://localhost:8080/events/1 \
    --include \
    --header "Content-Type: application/json" \
    --request "DELETE"
```