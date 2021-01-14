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

// Colors defines FG/BG color pairs for themes
type Colors struct {
	// 4-bit color for Linux Console
	C8 struct {
		FG string
		BG string
	}
	// 8-bit color for full terminals
	C256 struct {
		FG string
		BG string
	}
}
