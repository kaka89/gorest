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

package filter

import (
	"net/http"
	"sort"
)

// filter chain that execute the filter
type FilterChain interface {
	// execute the filter chain
	DoFilter(http.ResponseWriter, *http.Request)

	// add new filter to the filter chain
	AddFilter(filter ...Filter)

	// get all filters of this filter chain
	GetAllFilters() []Filter
}

// implements FilterChain
type DefaultFilterChain struct {
	sorted          bool
	currentPosition int
	Filters         []Filter
}

// execute the filter chain
func (this *DefaultFilterChain) DoFilter(response http.ResponseWriter, request *http.Request) {
	if !this.sorted {
		this.sort()
	}

	if this.currentPosition < len(this.Filters) {
		// TODO: 2018/2/13 Bruce 测试以下并行模式是否会有问题
		this.currentPosition ++
		nextFilter := this.Filters[this.currentPosition-1]
		nextFilter.DoFilter(response, request, this)
	}
}

// add new filter to this filter chain
func (this *DefaultFilterChain) AddFilter(filter ...Filter) {
	this.Filters = append(this.Filters, filter...)
	this.sorted = false
	this.sort()
}

// get all filters of this filter chain
func (this *DefaultFilterChain) GetAllFilters() []Filter {
	return this.Filters
}

// sort the filters, sort by order asc
func (this *DefaultFilterChain) sort() {
	if !this.sorted {
		this.sorted = true
		sort.Sort(ByOrder(this.Filters))
	}
}
