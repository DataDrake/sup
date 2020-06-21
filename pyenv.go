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
	"strings"
)

func pyenv() (pieces []Piece) {
	// Check for Python virtual environemnt
	if env := os.Getenv("VIRTUAL_ENV"); len(env) > 0 {
		// only output the last part of the environment path
		parts := strings.Split(env, "/")
		p := Piece{
			Content: fmt.Sprintf("⟆ %s", parts[len(parts)-1]),
			FG:      30,
			BG:      47,
		}
		pieces = append(pieces, p)
	}
	return
}
