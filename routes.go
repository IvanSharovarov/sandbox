package main

import "net/http"

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		Name:        "UserIndex",
		Method:      "GET",
		Pattern:     "/users",
		HandlerFunc: getUsers,
	},
	Route{
		Name:        "UserIndex",
		Method:      "POST",
		Pattern:     "/users",
		HandlerFunc: createUser,
	},
	Route{
		Name:        "UserIndex",
		Method:      "GET",
		Pattern:     "/users/{id}",
		HandlerFunc: getUser,
	},
	Route{
		Name:        "UserIndex",
		Method:      "PUT",
		Pattern:     "/users/{id}",
		HandlerFunc: updateUser,
	},
	Route{
		Name:        "UserIndex",
		Method:      "DELETE",
		Pattern:     "/users/{id}",
		HandlerFunc: deleteUser,
	},
}
