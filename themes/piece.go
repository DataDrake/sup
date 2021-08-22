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

package themes

import (
	"github.com/DataDrake/flair/color"
)

// Piece is the configuration for a Piece when rendering
type Piece struct {
	// ASCII symbol for non-Unicode terminals (Linux console)
	ASCII string `json:"ascii"`
	// Unicode symbol for Unicode terminals
	Unicode string `json:"unicode"`
	// 4-bit color for Linux Console
	C8 struct {
		FG color.Color `json:"fg"`
		BG color.Color `json:"bg"`
	} `json:"8"`
	// 8-bit color for full terminals
	C256 struct {
		FG color.Color `json:"fg"`
		BG color.Color `json:"bg"`
	} `json:"256"`
}
