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

package pieces

import (
	"github.com/DataDrake/sup/term"
	"github.com/DataDrake/sup/themes"
	"os"
	"os/user"
	"strings"
)

func trimHome(wd string) (dirs, parts []string) {
	u, _ := user.Current()
	path := strings.TrimPrefix(wd, u.HomeDir)
	parts = strings.Split(path, "/")
	if len(wd) > len(path) {
		dirs = append(dirs, "~")
		parts = parts[1:]
	}
	return
}

func splitPath() (dirs, parts []string) {
	wd, _ := os.Getwd()
	dirs, parts = trimHome(wd)
	return
}

func startingContext(oldDirs, oldParts []string) (dirs, parts []string) {
	dirs, parts = oldDirs, oldParts
	// Show part of the beginning of the path for context
	if len(parts) > 1 {
		lim := 1
		if parts[0] == "" {
			lim++
		}
		dirs = append(dirs, parts[0:lim]...)
		parts = parts[lim:]
	}
	return
}

func insertDots(th themes.Piece, oldDirs, oldParts []string) (dirs, parts []string) {
	dirs, parts = oldDirs, oldParts
	// Only show the last bit of the path
	if len(parts) > 2 {
		dots := th.ASCII
		if term.HasUnicode() {
			dots = th.Unicode
		}
		dirs = append(dirs, dots)
		parts = parts[len(parts)-2:]
	}
	return
}

func dir() *Piece {
	th := themes.Current["dir"]
	p := Convert(th)
	dirs, parts := splitPath()
	dirs, parts = startingContext(dirs, parts)
	dirs, parts = insertDots(th, dirs, parts)
	// append remaining parts and build string
	dirs = append(dirs, parts...)
	p.Content = strings.Join(dirs, "/")
	return p
}
