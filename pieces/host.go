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
	"github.com/DataDrake/sup/themes"
	"os"
)

func host() *Piece {
	th := themes.Current["host-local"]
	if term.IsSSH() {
		th = themes.Current["host-remote"]
	}
	p := Convert(th)
	p.Content = th.ASCII
	if term.HasUnicode() {
		p.Content = th.Unicode
	}
	// Override for SSH
	if term.IsSSH() {
		hostname, _ := os.Hostname()
		p.Content += hostname
		return p
	}
	// localhost
	return p
}
