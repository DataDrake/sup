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

package shell

// Bash is a Renderer for the Bash shell
type Bash struct{}

// FG prints a color as a foreground sequence
func (b Bash) FG(c string) string {
	return "\001\033[38;5;" + c + "m\002"
}

// BG prints a color as a background sequence
func (b Bash) BG(c string) string {
	return "\001\033[48;5;" + c + "m\002"
}

// Pair prints two colors as a FG/BG sequence
func (b Bash) Pair(fg, bg string) string {
	return "\001\033[38;5;" + fg + ";48;5;" + bg + "m\002"
}

// Reset both the FG anf BG colors to default
func (b Bash) Reset() string {
	return "\001\033[0m\002"
}

// ResetBG resets the BG color to default
func (b Bash) ResetBG() string {
	return "\001\033[49m\002"
}
