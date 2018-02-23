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

import "net/http"

// the real filter to execute
type Filter interface {
	Order() int
	//Init()
	DoFilter(http.ResponseWriter, *http.Request, FilterChain)
	//Destroy()
}

// ByOrder implements sort.Interface for []Person based on
// the Age field.
type ByOrder []Filter

func (a ByOrder) Len() int           { return len(a) }
func (a ByOrder) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByOrder) Less(i, j int) bool { return a[i].Order() > a[j].Order() }
