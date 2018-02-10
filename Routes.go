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
		"Index",
		"GET",
		"/",
		Index,
	},
	Route{
		"User Count",
		"GET",
		"/usercount",
		APIusercount,
	},
	Route{
		"Most Links",
		"GET",
		"/mostlinks",
		APImostlinks,
	},
	Route{
		"Most Mentions",
		"GET",
		"/mentioned",
		APImentioned,
	},
	Route{
		"Most Active",
		"GET",
		"/mostactive",
		APImostactive,
	},
	Route{
		"Most Grumpy",
		"GET",
		"/mostgrumpy",
		APImostgrumpy,
	},
	Route{
		"Most Happy",
		"GET",
		"/mosthappy",
		APImosthappy,
	},

}