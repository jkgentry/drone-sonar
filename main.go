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
	"log"
	"os"
)

func main() {
	var args []Argument
	var isPR bool

	args, _ = addIfExist("PLUGIN_PROJECT_NAME", "-Dsonar.projectName", args, "DRONE_REPO_NAME")
	args, _ = addIfExist("PLUGIN_PROJECT_KEY", "-Dsonar.projectKey", args, "DRONE_REPO_NAME")
	args, _ = addIfExist("PLUGIN_LOGIN", "-Dsonar.login", args, "")
	args, _ = addIfExist("LOGIN", "-Dsonar.login", args, "")
	args, isPR = addIfExist("PLUGIN_PR_KEY", "-Dsonar.pullrequest.key", args, "DRONE_PULL_REQUEST")

	if isPR {
		args, _ = addIfExist("PLUGIN_PR_BRANCH", "-Dsonar.pullrequest.branch", args, "DRONE_SOURCE_BRANCH")
	}

	args, _ = addIfExist("GITHUB_OAUTH", "-Dsonar.github.oauth", args, "")
	args, _ = addIfExist("PLUGIN_GITHUB_REPOSITORY", "-Dsonar.github.repository", args, "DRONE_REPO")

	s := Plugin{Args: args}
	val, ok := os.LookupEnv("PLUGIN_CERTIFICATE_AUTHORITY_URL")
	if ok {
		s.CertificateAuthorityUrl = val
	}
	err := s.execSonarRunner()
	if err != nil {
		log.Fatal(err)
	}
}

func addIfExist(envVariable string, argument string, args []Argument, defaultEnv string) ([]Argument, bool) {
	val, ok := os.LookupEnv(envVariable)
	if ok {
		return append(args, Argument{Value: val, Argument: argument}), true
	} else if defaultEnv != "" {
		val, ok := os.LookupEnv(defaultEnv)
		if ok {
			return append(args, Argument{Value: val, Argument: argument}), true
		}
	}
	return args, false
}
