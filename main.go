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

package main

import (
	"flag"
	"fmt"
	"github.com/DataDrake/sup/pieces"
	"github.com/DataDrake/sup/shell"
	"github.com/DataDrake/sup/term"
	"github.com/DataDrake/sup/themes"
	"os"
)

func main() {
	if theme := os.Getenv("SUP_THEME"); len(theme) > 0 {
		if err := themes.Load(theme); err != nil {
			themes.Load("default")
		}
	}
	var s shell.Shell
	// Deal with flags
	var sh = flag.String("sh", "bash", "select shell to use")
	flag.Parse()
	switch *sh {
	case "bash", "sh", "posix":
		s = shell.Bash
	case "zsh":
		s = shell.Zsh
	default:
		fmt.Fprintf(os.Stderr, "unsupported shell '%s', defaulting to bash\n", *sh)
		s = shell.Bash
	}
	// Build each of the requested pieces
	ps := pieces.Build(flag.Args())
	// Render all the pieces as a single string
	var out string
	if term.HasUnicode() {
		out = s.Full(ps)
	} else {
		out = s.Simple(ps)
	}
	// Print the resulting string to Stdout
	fmt.Print(out)
}
