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
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
)

func init() {
	Dirs()
	if name := os.Getenv("SUP_THEME"); len(name) > 0 {
		if Load(name) == nil {
			return
		}
	}
	Load("default")
}

// ErrThemeMissing occurs when a configured theme cannot be found
var ErrThemeMissing = errors.New("unable to find specified theme")

var (
	// SystemDir is the Location of installed themes
	SystemDir string
	// Current is the theme to use when rendering
	Current Theme
)

// Theme is a theme for sup
type Theme map[string]Piece

// Load reads in a theme
func Load(name string) error {
	for _, dir := range dirs {
		path := filepath.Join(dir, name+".json")
		if _, err := os.Stat(path); err != nil {
			continue
		}
		f, err := os.Open(path)
		if err != nil {
			return err
		}
		defer f.Close()
		dec := json.NewDecoder(f)
		return dec.Decode(&Current)
	}
	return ErrThemeMissing
}

var dirs []string

// Dirs builds a list of theme directories for searching
func Dirs() {
	home, err := os.UserHomeDir()
	if err == nil {
		dirs = append(dirs, filepath.Join(home, ".config", "sup", "themes"))
	}
	dirs = append(dirs, SystemDir)
}
