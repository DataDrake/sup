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
	"os/exec"
	"strings"
)

func git() *Piece {
	// get the current Git branch
	cmd := exec.Command("git", "rev-parse", "--abbrev-ref", "HEAD")
	out, err := cmd.CombinedOutput()
	if err != nil {
		// either not a git repo or fail silently
		return nil
	}
	return &Piece{
		Content: fmt.Sprintf(" git:%s", strings.TrimSpace(string(out))),
		FG:      30,
		BG:      47,
	}
}

func svn() *Piece {
	// check if this is an SVN repo
	cmd := exec.Command("svn", "info", "--show-item", "url")
	err := cmd.Run()
	if err != nil {
		return nil
	}
	// TODO: Add special handling for SVN tags/branches
	return &Piece{
		Content: " svn",
		FG:      30,
		BG:      47,
	}
}

func vcs() []Piece {
	var pieces []Piece
	if p := git(); p != nil {
		pieces = append(pieces, *p)
	}
	if p := svn(); p != nil {
		pieces = append(pieces, *p)
	}
	return pieces
}
