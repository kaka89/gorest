/*
* Copyright 2018 Bruce Liu.  All rights reserved.
*
* Licensed under the Apache License, Version 2.0 (the "License");
* you may not use this file except in compliance with the License.
* You may obtain a copy of the License at
*
*      http://www.apache.org/licenses/LICENSE-2.0
*
* Unless required by applicable law or agreed to in writing, software
* distributed under the License is distributed on an "AS IS" BASIS,
* WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
* See the License for the specific language governing permissions and
* limitations under the License.
*/
package gorest

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"fmt"
	"time"
	"github.com/kaka89/gorest/router"
)

// app default instance, each gorest application only has one instance
var RestApp *App

func init() {
	// init the default instance
	RestApp = NewApp()

	// give a default health check api
	GET("/", "", ShowAPI)
	GET("/health", "health check url", Health)
}

// returns an new app instance
func NewApp() *App {

	app := &App{}
	app.Router = NewRouter()

	return app
}

type App struct {
	// the http router used
	Router *router.Router
}

// start the http server
func (app *App) Run() {

	http.ListenAndServe(":8080", app.Router.R)
}

func Health(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintln(w, time.Now())
}

// show all api whose name is not empty
func ShowAPI(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintln(w, time.Now())
}
