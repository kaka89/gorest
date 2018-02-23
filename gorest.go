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
	"strings"
)

const (
	// VERSION of the gorest framework
	VERSION = "0.1"

	// two basic running model
	// DEV is for develop
	DEV = "dev"

	// PROD is for production, this is the default model
	PROD = "prod"
)

// start the rest server
func Run(args ...string) {

	if len(args) > 0 && args[0] != "" {
		// first argument should be host and port
		hp := strings.Split(args[0], ":")
		if len(hp) > 0 && hp[0] != "" {
			// TODO: 2018/2/12 Bruce set the config of host
		}
		if len(hp) > 1 && hp[1] != "" {
			// TODO: 2018/2/12 Bruce set the config of port
		}
	}

	// TODO: 2018/2/12 liuquan Read the config asa

	RestApp.Run()
}
