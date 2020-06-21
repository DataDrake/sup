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
	"os"
)

func pipeStatus() (pieces []Piece) {
	failure := false
	// Generate a piece for each return code
	for _, content := range os.Args[1:] {
		// Assume success
		piece := Piece{
			Content: content,
			FG:      30,
			BG:      42,
		}
		// override for failure
		if content != "0" {
			failure = true
			piece.FG = 97
			piece.BG = 41
		}
		pieces = append(pieces, piece)
	}
	// If no failures, don't produce any output
	if !failure {
		pieces = make([]Piece, 0)
	}
	return
}
