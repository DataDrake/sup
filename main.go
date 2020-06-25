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
	"flag"
	"fmt"
	"github.com/DataDrake/sup/pieces"
	"github.com/DataDrake/sup/shell"
	"github.com/DataDrake/sup/shell/bash"
	"github.com/DataDrake/sup/shell/zsh"
	"github.com/DataDrake/sup/term"
	"os"
)

func main() {
	var r shell.Renderer
	// Deal with flags
	var sh = flag.String("sh", "bash", "select shell to use")
	flag.Parse()
	switch *sh {
	case "bash", "sh", "posix":
		r = bash.Bash{}
	case "zsh":
		r = zsh.Zsh{}
	default:
		fmt.Fprintf(os.Stderr, "unsupported shell '%s', defaulting to bash\n", *sh)
		r = bash.Bash{}
	}
	// Build each of the requested pieces
	ps := pieces.Build(flag.Args())
	// Render all the pieces as a single string
	var out string
	if term.HasUnicode() {
		out = r.Full(ps)
	} else {
		out = r.Simple(ps)
	}
	// Print the resulting string to Stdout
	fmt.Print(out)
}
