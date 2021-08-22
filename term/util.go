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
)

// Has256Color checks for 256 color support
func Has256Color() bool {
	term := os.Getenv("TERM")
	switch {
	case term == "alacritty":
		return true
	case strings.Contains(term, "kitty"):
		return true
	default:
		return strings.Contains(term, "256")
	}
}

// HasUnicode checks for Unicode support
func HasUnicode() bool {
	return !strings.Contains(os.Getenv("TERM"), "linux")
}

// IsSSH checks for the presence of an SSH session
func IsSSH() bool {
	return len(os.Getenv("SSH_CLIENT")) > 0
}
