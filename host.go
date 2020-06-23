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
	"os"
)

func host() []Piece {
	// Assume localhost
	p := Piece{
		content: "⮞",
		fg:      15,
		bg:      57,
	}
	// Override for SSH
	if len(os.Getenv("SSH_CLIENT")) > 0 {
		p = Piece {
			content: fmt.Sprintf(" %s", Hostname),
			fg: 0,
			bg: 208,
		}
	}
	return []Piece{p}
}
