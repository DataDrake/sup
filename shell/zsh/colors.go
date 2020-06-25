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

package zsh

import (
	"github.com/DataDrake/sup/pieces"
)

// FG prints a color as a foreground sequence
func FG(c pieces.Color) string {
	return "%{\033[38;5;" + string(c) + "m%}"
}

// BG prints a color as a background sequence
func BG(c pieces.Color) string {
	return "%{\033[48;5;" + string(c) + "m%}"
}

// Pair prints two colors as a FG/BG sequence
func Pair(fg, bg pieces.Color) string {
	return "%{\033[38;5;" + string(fg) + ";48;5;" + string(bg) + "m%}"
}
