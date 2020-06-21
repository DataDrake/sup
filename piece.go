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
)

// Piece makes up a small part of the PS1 line
type Piece struct {
	Content string
	FG      int
	BG      int
}

type pieceFn func() []Piece

// Call each function, joining together the results
func build(fns ...pieceFn) (pieces []Piece) {
	for _, fn := range fns {
		pieces = append(pieces, fn()...)
	}
	return
}

// Generate a string from each of the varous parts
func render(pieces []Piece) string {
	var status string
	for i, curr := range pieces {
		if i == 0 {
			// Very first piece needs colors set before-hand
			status += fmt.Sprintf("\\[\\e[%d;%dm\\]", curr.FG, curr.BG)
		}
		if len(pieces) == i+1 {
			// Last piece has nothing after it
			status += fmt.Sprintf(" %s \\[\\e[%d;49m\\]", curr.Content, curr.BG-10)
			break
		}
		if next := pieces[i+1]; curr.FG != next.FG || curr.BG != next.BG {
			// Deal with color change
			status += fmt.Sprintf(" %s \\[\\e[%d;%dm\\]\\[\\e[%dm\\]", curr.Content, curr.BG-10, next.BG, next.FG)
		} else {
			// Same color, so just print
			status += fmt.Sprintf(" %s ", curr.Content)
		}
	}
	return fmt.Sprintf("%s \\[\\e[0m\\]", status) // make sure to reset styling after
}
