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

var carrot = " "
var carrotSame = "  "

// Piece makes up a small part of the PS1 line
type Piece struct {
	Content string
	FG      Color
	BG      Color
}

type pieceFn func() *Piece

var all = []pieceFn{
	host,
	username,
	pyenv,
	vcs,
	dir,
}

// Build calls each function, joining together the results
func Build(statuses []string) (pieces []Piece) {
	all = append(all, pipeStatus(statuses)...)
	for _, fn := range all {
		if piece := fn(); piece != nil {
			pieces = append(pieces, *piece)
		}
	}
	return
}
