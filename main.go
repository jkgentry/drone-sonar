// Copyright 2018 Jake Gentry
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// 		http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"os"
)

func main() {
	var args []Argument
	args = addIfExist("PLUGIN_PROJECT_NAME", "-Dsonar.projectName", args, "DRONE_REPO_NAME")
	args = addIfExist("PLUGIN_LOGIN", "-Dsonar.login", args, "")
	args = addIfExist("PLUGIN_GITHUB_PULL_REQUEST", "-Dsonar.github.pullRequest", args, "DRONE_PULL_REQUEST")
	args = addIfExist("PLUGIN_ANALYSIS_MODE", "-Dsonar.analysis.mode", args, "")
	args = addIfExist("PLUGIN_GITHUB_OAUTH", "-Dsonar.github.oauth", args, "")
	args = addIfExist("PLUGIN_GITHUB_REPOSITORY", "-Dsonar.github.repository", args, "DRONE_REPO")

	s := Plugin{Args: args}
	s.execSonarRunner()
}

func addIfExist(envVariable string, argument string, args []Argument, defaultEnv string) []Argument {
	val, ok := os.LookupEnv(envVariable)
	if ok {
		return append(args, Argument{Value: val, Argument: "-Dsonar.projectName"})
	} else if defaultEnv != "" {
		val, ok := os.LookupEnv(defaultEnv)
		if ok {
			return append(args, Argument{Value: val, Argument: "-Dsonar.projectName"})
		}
	}
	return args
}
