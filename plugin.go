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

	"github.com/Sirupsen/logrus"
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
		logrus.Println(err)
		return err
	}
	return nil
}

func (p Plugin) execSonarRunner() error {
	args := []string{"-jar", "./sonar-scanner-cli-3.2.0.1227.jar"}
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

// func (p Plugin) writePipelineLetter() {
//
// 	f, err := os.OpenFile(".Pipeline-Letter", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0666)
// 	if err != nil {
// 		fmt.Printf("!!> Error creating / appending to .Pipeline-Letter")
// 		return
// 	}
// 	defer f.Close()
//
// 	if _, err := f.WriteString(fmt.Sprintf("*SONAR*: %s/dashboard/index/%s\n", p.Host, strings.Replace(p.Key, "/", ":", -1))); err != nil {
// 		fmt.Printf("!!> Error writing to .Pipeline-Letter")
// 	}
// }

// func (p Plugin) writeRepoSignature() {
// 	expectedContent := fmt.Sprintf("%s/%s/%s\n", p.Repo, p.Branch, time.Now().Format("2006-01-02"))
// 	h := sha256.New()
// 	h.Write([]byte(expectedContent))
// 	expectedSignature := fmt.Sprintf("%x", h.Sum(nil))
//
// 	f, err := os.OpenFile(".SonarSignature", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0666)
// 	if err != nil {
// 		fmt.Printf("!!> Error creating / appending to .SonarSignature")
// 		return
// 	}
// 	defer f.Close()
//
// 	if _, err = f.WriteString(expectedSignature); err != nil {
// 		fmt.Printf("!!> Error writing to .SonarSignature")
// 	}
// }

func printCommand(cmd *exec.Cmd) {
	fmt.Printf("==> Executing: %s\n", strings.Join(cmd.Args, " "))
}

func printOutput(outs []byte) {
	if len(outs) > 0 {
		fmt.Printf("==> Output: %s\n", string(outs))
	}
}
