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

package shell

import (
	"fmt"
	"github.com/DataDrake/flair"
	"github.com/DataDrake/flair/color"
	"github.com/DataDrake/flair/escape"
	"github.com/DataDrake/sup/pieces"
	"github.com/DataDrake/sup/term"
	"github.com/DataDrake/sup/themes"
	"os"
)

// Shell Definitions
var (
	Bash = Shell{"\001", "\002"}
	Zsh  = Shell{"%{", "%}"}
)

// Current is the most recently set Shell
var Current Shell

// Set changes the Shell by name with Bash as the fallback
func Set(name string) {
	switch name {
	case "bash", "sh", "posix":
		Current = Bash
	case "zsh":
		Current = Zsh
	default:
		fmt.Fprintf(os.Stderr, "unsupported shell '%s', defaulting to bash\n", name)
		Current = Bash
	}
}

// Shell contains the shell-specific escape sequences for marking escape sequences in prompts
type Shell struct {
	Start string
	End   string
}

// First color change
func (s Shell) First(fg, bg color.Color) string {
	return s.Start + escape.Combine(fg.FG(), bg.BG()).String() + s.End
}

// Last color change with carrot, reseting colors after
func (s Shell) Last(fg color.Color, carrot string) string {
	pre := s.Start + escape.Combine(fg.FG(), color.DefaultBG).String() + s.End
	post := s.Start + flair.Reset + s.End
	return pre + carrot + post
}

// Switch the colors between pieces, making sure the carrot looks right
func (s Shell) Switch(bg, nextBG, nextFG color.Color, carrot string) string {
	pre := s.Start + escape.Combine(bg.FG(), nextBG.BG()).String() + s.End
	post := s.Start + nextFG.FG().String() + s.End
	return pre + carrot + post
}

// Full generates a string from each of the varous parts
func (s Shell) Full(ps []pieces.Piece) string {
	carrot := themes.Current["carrot"].Unicode
	carrotSame := themes.Current["carrot-same"].Unicode
	var status string
	for i, curr := range ps {
		if i == 0 {
			// Very first piece needs colors set before-hand
			status += s.First(curr.FG, curr.BG) + " "
		}
		status += curr.Content
		if len(ps) == i+1 {
			// Last piece has nothing after it
			status += " " + s.Last(curr.BG, carrot)
			break
		}
		if next := ps[i+1]; curr.FG != next.FG || curr.BG != next.BG {
			// Deal with color change
			status += " " + s.Switch(curr.BG, next.BG, next.FG, carrot)
		} else {
			// Same color, so just print
			status += carrotSame
		}
	}
	return status
}

// Simple generates a string from each of the varous parts, for Linux console
func (s Shell) Simple(ps []pieces.Piece) string {
	var status string
	for i, curr := range ps {
		status += s.First(curr.FG, curr.BG) + " " + curr.Content + " "
		if len(ps) == i+1 {
			// Last piece has reset and simple prompt
			status += flair.Reset + "$ "
			break
		}
	}
	return status
}

// Prompt writes the configured Prompt to os.Stdout with the specified shell's formatting
func (s Shell) Prompt(statuses []string) {
	// Build each of the requested pieces
	ps := pieces.Build(statuses)
	// Render all the pieces as a single string
	var out string
	if term.HasUnicode() {
		out = s.Full(ps)
	} else {
		out = s.Simple(ps)
	}
	// Print the resulting string to Stdout
	fmt.Fprintf(os.Stdout, out)
}
