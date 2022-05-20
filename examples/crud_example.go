/*
Package examples shows how to implement a basic CRUD for two data structures with the api2go server functionality.
To play with this example server you can run some of the following curl requests

In order to demonstrate dynamic baseurl handling for requests, apply the --header="REQUEST_URI:https://www.your.domain.example.com" parameter to any of the commands.

Create a new user:
	curl -X POST http://localhost:31415/v0/users -d '{"data" : {"type" : "users" , "attributes": {"user-name" : "marvin"}}}'

List users:
	curl -X GET http://localhost:31415/v0/users

List paginated users:
	curl -X GET 'http://localhost:31415/v0/users?page\[offset\]=0&page\[limit\]=2'
OR
	curl -X GET 'http://localhost:31415/v0/users?page\[number\]=1&page\[size\]=2'

Update:
	curl -vX PATCH http://localhost:31415/v0/users/1 -d '{ "data" : {"type" : "users", "id": "1", "attributes": {"user-name" : "better marvin"}}}'

Delete:
	curl -vX DELETE http://localhost:31415/v0/users/2

Create a chocolate with the name sweet
	curl -X POST http://localhost:31415/v0/chocolates -d '{"data" : {"type" : "chocolates" , "attributes": {"name" : "Ritter Sport", "taste": "Very Good"}}}'

Create a user with a sweet
	curl -X POST http://localhost:31415/v0/users -d '{"data" : {"type" : "users" , "attributes": {"user-name" : "marvin"}, "relationships": {"sweets": {"data": [{"type": "chocolates", "id": "1"}]}}}}'

List a users sweets
	curl -X GET http://localhost:31415/v0/users/1/sweets

Replace a users sweets
	curl -X PATCH http://localhost:31415/v0/users/1/relationships/sweets -d '{"data" : [{"type": "chocolates", "id": "2"}]}'

Add a sweet
	curl -X POST http://localhost:31415/v0/users/1/relationships/sweets -d '{"data" : [{"type": "chocolates", "id": "2"}]}'

Remove a sweet
	curl -X DELETE http://localhost:31415/v0/users/1/relationships/sweets -d '{"data" : [{"type": "chocolates", "id": "2"}]}'
*/
package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/manyminds/api2go/routing"

	"github.com/manyminds/api2go"
	"github.com/manyminds/api2go/examples/model"
	"github.com/manyminds/api2go/examples/resource"
	"github.com/manyminds/api2go/examples/storage"
)

func main() {
	port := 31415

	r := gin.Default()
	api := api2go.NewAPIWithRouting("v0", api2go.NewStaticResolver(""), routing.Gin(r))
	userStorage := storage.NewUserStorage()
	chocStorage := storage.NewChocolateStorage()
	userStorage.Insert(model.User{
		ID:            "1",
		Username:      "foo",
		PasswordHash:  "foo",
		Chocolates:    nil,
		ChocolatesIDs: nil,
		CreditCard: &model.CreditCard{ID:"foo"},
	})
	api.AddResource(model.User{}, resource.UserResource{ChocStorage: chocStorage, UserStorage: userStorage})
	api.AddResource(model.Chocolate{}, resource.ChocolateResource{ChocStorage: chocStorage, UserStorage: userStorage})
	api.AddResource(model.CreditCard{}, resource.CreditCardResource{UserStorage: userStorage})

	fmt.Printf("Listening on :%d", port)

	// listen and serve on 0.0.0.0:8443 (for windows "localhost:8443")
	r.Run(":8080")
	//r.RunTLS(":8443", "./certs/crt.pem", "./certs/key.pem")
}
