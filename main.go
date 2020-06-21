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

func init() {
	u, _ := user.Current()

	Hostname, _ = os.Hostname()
	Username = u.Username

	HomeDir = u.HomeDir
	WorkDir, _ = os.Getwd()
}

func main() {
	// Build each of the requested pieces
	pieces := build(host, username, pyenv, vcs, dir, pipeStatus)
	// Render all the pieces as a single string
	out := render(pieces)
	// Print the resulting string to Stdout
	fmt.Print(out)
}
