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
	"github.com/DataDrake/sup/themes"
)

func pipeGood(content string) pieceFn {
	return func() *Piece {
		good := Convert(themes.Current["pipe-good"])
		good.Content = content
		return good
	}
}

func pipeBad(content string) pieceFn {
	return func() *Piece {
		bad := Convert(themes.Current["pipe-bad"])
		bad.Content = content
		return bad
	}
}

func pipeStatus(statuses []string) (fns []pieceFn) {
	failure := false
	// Generate a piece for each return code
	for _, content := range statuses {
		if content == "0" {
			fns = append(fns, pipeGood(content))
		} else {
			failure = true
			fns = append(fns, pipeBad(content))
		}
	}
	// If no failures, don't produce any output
	if !failure {
		fns = make([]pieceFn, 0)
	}
	return
}
