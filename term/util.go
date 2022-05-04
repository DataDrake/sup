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

package term

import (
	"os"
	"strings"

	"github.com/xo/terminfo"
)

// Has256Color checks for 256 color support
func Has256Color() bool {
	// Load the current terminfo
	ti, err := terminfo.LoadFromEnv()
	if err != nil {
		return false
	}

	// Get the colors from the terminfo if set
	var colors int
	if colors = ti.Num(terminfo.MaxColors); colors <= 0 {
		// Colors don't seem to be set, assume basic colors only
		colors = int(terminfo.ColorLevelBasic)
	}

	return colors >= 256
}

// HasUnicode checks for Unicode support
func HasUnicode() bool {
	return !strings.Contains(os.Getenv("TERM"), "linux")
}

// IsSSH checks for the presence of an SSH session
func IsSSH() bool {
	return len(os.Getenv("SSH_CLIENT")) > 0
}
