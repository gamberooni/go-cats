# go-cats
Super simple CRUD app in Go about cats

## How it works
Client makes a request to the web server > the corresponding handler handles the request > handler requests for data from db via `store` > `store` fetches data from db via `gorm` > handler gets the requested data from `store` > handler sends the data back to the client

## Routes
1. GET `/api/cats`
2. POST `/api/cats`
3. GET `/api/cats/:id`
4. PUT `/api/cats/:id`
5. DELETE `/api/cats/:id`

## Starting the server
In the project root directory, run `go run main.go`
