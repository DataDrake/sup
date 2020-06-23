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

var carrot = " "
var carrotSame = "  "

// Piece makes up a small part of the PS1 line
type Piece struct {
	content string
	fg Color
	bg Color
}

type pieceFn func() *Piece

// Call each function, joining together the results
func build(fns ...pieceFn) (pieces []Piece) {
	for _, fn := range fns {
		if piece := fn(); piece != nil {
			pieces = append(pieces, *piece)
		}
	}
	return
}

// Generate a string from each of the varous parts
func render(pieces []Piece) string {
	var status string
	for i, curr := range pieces {
		if i == 0 {
			// Very first piece needs colors set before-hand
			status += Pair(curr.fg, curr.bg) + " "
		}
		status += curr.content
		if len(pieces) == i+1 {
			// Last piece has nothing after it
			status += " " + FG(curr.bg) + "\\[\\e[49m\\]" + carrot + "\\[\\e[0m\\]"
			break
		}
		if next := pieces[i+1]; curr.fg != next.fg || curr.bg != next.bg {
			// Deal with color change
			status +=  " " + Pair(curr.bg, next.bg) + carrot + FG(next.fg)
		} else {
			// Same color, so just print
			status += carrotSame
		}
	}
	return status
}

// Generate a string from each of the varous parts, for Linux console
func renderSimple(pieces []Piece) string {
	var status string
	for i, curr := range pieces {
		status += Pair(curr.fg, curr.bg) + " " + curr.content + " "
		if len(pieces) == i+1 {
			// Last piece has reset and simple prompt
			status += "\\[\\e[0m\\]$ "
			break
		}
	}
	return status
}
