/* Copyright 2018 Bruce Liu.  All rights reserved.
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
	"github.com/kaka89/gorest/router"
	"github.com/kaka89/gorest/filter"
	"github.com/julienschmidt/httprouter"
)

// create a new router
func NewRouter() *router.Router {

	return &router.Router{
		nil,
		nil,
		"/",
		make(map[string](map[string]string)),
		httprouter.New(),
		&filter.DefaultFilterChain{},
	}
}

// open GET function to gorest
func GET(path, name string, handle httprouter.Handle) *router.Router {
	RestApp.Router.GET(path, name, handle)
	return RestApp.Router
}

func SubRouter(path string) *router.Router {
	return RestApp.Router.SubRouter(path)
}

func AddFilterChain(chain filter.FilterChain) *router.Router {
	return RestApp.Router.AddFilterChain(chain)
}

func AddFilter(filter filter.Filter) *router.Router {
	return RestApp.Router.AddFilter(filter)
}
