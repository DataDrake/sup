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
	"os"
)

var local = "⮞"
var remote = " "

func init() {
	if !term.HasUnicode() {
		local = "H"
		remote = "R "
	}
}

func host() *Piece {
	hostname, _ := os.Hostname()
	// Override for SSH
	if term.IsSSH() {
		return &Piece{
			Content: remote + hostname,
			FG:      "0",
			BG:      "208",
		}
	}
	// localhost
	return &Piece{
		Content: local,
		FG:      "15",
		BG:      "57",
	}
}