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
	"github.com/DataDrake/sup/pieces"
	"github.com/DataDrake/sup/themes"
)

// Renderer presents a common interface for supporting any shell syntax
type Renderer interface {
	FG(fg string) string
	BG(bg string) string
	Pair(fg, bg string) string
	Reset() string
	ResetBG() string
}

// Full generates a string from each of the varous parts
func Full(r Renderer, ps []pieces.Piece) string {
	carrot := themes.Current["carrot"].Unicode
	carrotSame := themes.Current["carrot-same"].Unicode
	var status string
	for i, curr := range ps {
		if i == 0 {
			// Very first piece needs colors set before-hand
			status += r.Pair(curr.FG, curr.BG) + " "
		}
		status += curr.Content
		if len(ps) == i+1 {
			// Last piece has nothing after it
			status += " " + r.FG(curr.BG) + r.ResetBG() + carrot + r.Reset()
			break
		}
		if next := ps[i+1]; curr.FG != next.FG || curr.BG != next.BG {
			// Deal with color change
			status += " " + r.Pair(curr.BG, next.BG) + carrot + r.FG(next.FG)
		} else {
			// Same color, so just print
			status += carrotSame
		}
	}
	return status
}

// Simple generates a string from each of the varous parts, for Linux console
func Simple(r Renderer, ps []pieces.Piece) string {
	var status string
	for i, curr := range ps {
		status += r.Pair(curr.FG, curr.BG) + " " + curr.Content + " "
		if len(ps) == i+1 {
			// Last piece has reset and simple prompt
			status += r.Reset() + "$ "
			break
		}
	}
	return status
}
