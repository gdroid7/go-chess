# go-chess

Install GO

https://golang.org/doc/install

Clone this repo to your workspace and cd to this project

To run program

````
go run main.go <piece> <position>
````

Piece options : `PAWN` / `KING` / `QUEEN`

Position format : `A1` / `B1` / `H1` etc

To Run Tests

````
go test ./tests/
````

To Run Specific Tests

````
go test ./tests/<fileName>
````

To Check code coverage

first run

````
go test ./tests/ -coverprofile cover.out -coverpkg=./...
````

then 
````
go tool cover -html=cover.out
````
