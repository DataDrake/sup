//
// Copyright 2021 Bryan T. Meyers <root@datadrake.com>
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

package pieces

import (
	"github.com/DataDrake/sup/term"
	"github.com/DataDrake/sup/themes"
	"os/exec"
	"strings"
)

func git() string {
	// get the current Git branch
	out, err := exec.Command("git", "rev-parse", "--abbrev-ref", "HEAD").CombinedOutput()
	if err != nil {
		// either not a git repo or fail silently
		return ""
	}
	return "git:" + strings.TrimSpace(string(out))
}

func svn() string {
	// check if this is an SVN repo
	if err := exec.Command("svn", "info", "--show-item", "url").Run(); err != nil {
		return ""
	}
	// TODO: Add special handling for SVN tags/branches
	return "svn"
}

func vcs() *Piece {
	th := themes.Current["vcs"]
	p := Convert(th)
	p.Content = th.ASCII
	if term.HasUnicode() {
		p.Content = th.Unicode
	}
	if content := git(); len(content) > 0 {
		p.Content += content
		return p
	}
	if content := svn(); len(content) > 0 {
		p.Content += content
		return p
	}
	return nil
}
