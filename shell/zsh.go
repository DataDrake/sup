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

// Zsh is a Renderer for the zsh shell
type Zsh struct{}

// FG prints a color as a foreground sequence
func (z Zsh) FG(c string) string {
	return "%{\033[38;5;" + c + "m%}"
}

// BG prints a color as a background sequence
func (z Zsh) BG(c string) string {
	return "%{\033[48;5;" + c + "m%}"
}

// Pair prints two colors as a FG/BG sequence
func (z Zsh) Pair(fg, bg string) string {
	return "%{\033[38;5;" + fg + ";48;5;" + bg + "m%}"
}

// Reset both the FG anf BG colors to default
func (z Zsh) Reset() string {
	return "%{%f%k%}"
}

// ResetBG resets the BG color to default
func (z Zsh) ResetBG() string {
	return "%{%k%}"
}
