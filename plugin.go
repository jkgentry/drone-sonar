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
	"io"
	"net/http"
	"os"
	"os/exec"
	"strings"
)

type Argument struct {
	Value    string
	Argument string
}

type Plugin struct {
	Args                    []Argument
	CertificateAuthorityUrl string
}

func (p *Plugin) Exec() error {
	err := p.execSonarRunner()
	if err != nil {
		return err
	}
	return nil
}

func (p Plugin) execSonarRunner() error {
	if p.CertificateAuthorityUrl != "" {
		err := downloadFile(p.CertificateAuthorityUrl, "/usr/local/share/ca-certificates/authority.pem")
		if err != nil {
			return err
		}
		cmd := exec.Command("update-ca-certificates")
		printCommand(cmd)
		output, err := cmd.CombinedOutput()
		printOutput(output)
	}

	args := []string{"-Dsonar.userHome=/sonar/"}
	for _, arg := range p.Args {
		args = append(args, arg.Argument+"="+arg.Value)
	}
	cmd := exec.Command("sonar-scanner", args...)
	// printCommand(cmd)
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

// Download data from url and write it to specified file
func downloadFile(url string, file string) error {
	// create the file
	out, err := os.Create(file)
	if err != nil {
		return err
	}
	defer out.Close()

	// get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// writer the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}
	return nil
}
