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

package router

import (
	"github.com/julienschmidt/httprouter"
	"github.com/kaka89/gorest/filter"
	"net/http"
)

// the router for gorest app
type Router struct {
	// use tree to record the router structure, just like a namespace
	Parent   *Router
	Children []*Router

	// prefix of this router, default value is "/"
	Prefix string

	// The index, maps names to urls, the real urls, the value structure will be: [method][name][url], we use this to generate the api
	Urls map[string](map[string]string)

	// the http router used, and each gorest app instance only has one instance of httprouter
	R *httprouter.Router

	// the filter chain which will execute before controller, the filter chain of parent will execute first before execute sub router's filter chain
	FilterChain filter.FilterChain
}

// create a sub router of current router
func (r *Router) SubRouter(subPrefix string) *Router {
	subRouter := &Router{
		Urls:   make(map[string](map[string]string)),
		Prefix: r.fullPath(subPrefix),
		R:      r.R,
	}
	// Init relationships
	r.Children = append(r.Children, subRouter)
	subRouter.Parent = r

	// release all parent filters to sub router. TODO: 2018/2/13 Bruce we only use default filter here, change it, add filter chain before add sub router
	subRouter.FilterChain = &filter.DefaultFilterChain{}
	subRouter.FilterChain.AddFilter(r.FilterChain.GetAllFilters()...)

	return subRouter

}

// add an router to the current router
func (r *Router) addPath(method, path, name string, fn httprouter.Handle) {

	fullPath := r.fullPath(path)

	// record the full path of this router, and if name is empty, means do not export the api file
	if len(name) > 0 {
		if r.Urls[method] == nil {
			r.Urls[method] = make(map[string]string)
		}
		r.Urls[method][name] = fullPath
	}

	// Wrapper function to bypass the parameter problem
	wf := func(response http.ResponseWriter, request *http.Request, params httprouter.Params) {
		// do filter first
		r.FilterChain.DoFilter(response, request)
		// go to the controller
		fn(response, request, params)
	}

	r.R.Handle(method, fullPath, wf)
}

// open GET function to gorest
func (r *Router) GET(path, name string, handle httprouter.Handle) *Router {
	r.addPath("GET", path, name, handle)
	return r
}

func (r *Router) AddFilter(filter ...filter.Filter) *Router {
	r.FilterChain.AddFilter(filter...)
	return r
}

func (r *Router) SetFilterChain(filterChain filter.FilterChain) *Router {
	r.FilterChain = filterChain
	return r
}

func (r *Router) AddFilterChain(filterChain filter.FilterChain) *Router {
	r.FilterChain.AddFilter(filterChain.GetAllFilters()...)
	return r
}

// subPath returns the full path of the current router
func (r *Router) fullPath(path string) string {

	pre := r.Prefix

	//TODO: 2018/2/13 Bruce fix the duplicate slash bug
	if (pre == "/" || pre[:len(pre)-1] == "/") && path[:1] == "/" {
		pre = pre[:len(pre)-1]
	}

	return pre + path
}
