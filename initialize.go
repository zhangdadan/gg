package main

import (
	"dakesolo.com/gg/app/controller/index"
	"net/http"
)


func route(mux * http.ServeMux)  {
	mux.Handle("/index/getIndex", &unit.Context{Action: index.GetIndex})
	mux.Handle("/index/getNav", &unit.Context{Action: index.GetNav})
	mux.Handle("/index/getBanner", &unit.Context{Action: index.GetBanner})
}








