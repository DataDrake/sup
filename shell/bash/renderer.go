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

package bash

import (
	"github.com/DataDrake/sup/pieces"
)

var carrot = " "
var carrotSame = "  "

// Bash is a Renderer for the Bash shell
type Bash struct{}

// Full generates a string from each of the varous parts
func (b Bash) Full(ps []pieces.Piece) string {
	var status string
	for i, curr := range ps {
		if i == 0 {
			// Very first piece needs colors set before-hand
			status += Pair(curr.FG, curr.BG) + " "
		}
		status += curr.Content
		if len(ps) == i+1 {
			// Last piece has nothing after it
			status += " " + FG(curr.BG) + "\001\033[49m\002" + carrot + "\001\033[0m\002"
			break
		}
		if next := ps[i+1]; curr.FG != next.FG || curr.BG != next.BG {
			// Deal with color change
			status += " " + Pair(curr.BG, next.BG) + carrot + FG(next.FG)
		} else {
			// Same color, so just print
			status += carrotSame
		}
	}
	return status
}

// Simple generates a string from each of the varous parts, for Linux console
func (b Bash) Simple(ps []pieces.Piece) string {
	var status string
	for i, curr := range ps {
		status += Pair(curr.FG, curr.BG) + " " + curr.Content + " "
		if len(ps) == i+1 {
			// Last piece has reset and simple prompt
			status += "\001\033[0m\002$ "
			break
		}
	}
	return status
}
