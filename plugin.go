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
	"fmt"
	"os/exec"
	"strings"
)

type Argument struct {
	Value    string
	Argument string
}

type Plugin struct {
	Args []Argument
}

func (p *Plugin) Exec() error {
	err := p.execSonarRunner()
	if err != nil {
		return err
	}
	return nil
}

func (p Plugin) execSonarRunner() error {
	args := []string{"-jar", "/bin/sonar-runner.jar"}
	for _, arg := range p.Args {
		args = append(args, arg.Argument+"="+arg.Value)
	}
	cmd := exec.Command("java", args...)
	printCommand(cmd)
	output, err := cmd.CombinedOutput()
	printOutput(output)

	if err != nil {
		return err
	}

	return nil
}

func printCommand(cmd *exec.Cmd) {
	fmt.Printf("==> Executing: %s\n", strings.Join(cmd.Args, " "))
}

func printOutput(outs []byte) {
	if len(outs) > 0 {
		fmt.Printf("==> Output: %s\n", string(outs))
	}
}
