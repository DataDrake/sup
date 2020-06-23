//
// Copyright 2020 Bryan T. Meyers <root@datadrake.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package main

import (
	"fmt"
	"os"
	"os/user"
	"strings"
)

var (
	// Hostname is the name of this machine
	Hostname string
	// Username is the login of the current user
	Username string

	// HomeDir is the path of the current user's HOME
	HomeDir string
	// WorkDir is the current working directory
	WorkDir string
)

// Has256Color chechs for 256 color support
func Has256Color() bool {
	return strings.Contains(os.Getenv("TERM"), "256")
}

// HasUnicode checks for Unicode support
func HasUnicode() bool {
	return !strings.Contains(os.Getenv("TERM"), "linux")
}

// IsSSH checks for the presence of an SSH session
func IsSSH() bool {
	return len(os.Getenv("SSH_CLIENT")) > 0
}

func init() {
	u, _ := user.Current()

	Hostname, _ = os.Hostname()
	Username = u.Username

	HomeDir = u.HomeDir
	WorkDir, _ = os.Getwd()
}

func main() {
	// Build each of the requested pieces
	fns := []pieceFn{host, username, pyenv, vcs, dir}
	fns = append(fns, pipeStatus()...)
	pieces := build(fns...)
	// Render all the pieces as a single string
	var out string
	if HasUnicode() {
		out = render(pieces)
	} else {
		out = renderSimple(pieces)
	}
	// Print the resulting string to Stdout
	fmt.Print(out)
}
