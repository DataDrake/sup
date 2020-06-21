# sup
Go PS1 Status Updater

[![Go Report Card](https://goreportcard.com/badge/github.com/DataDrake/sup)](https://goreportcard.com/report/github.com/DataDrake/cuppa) [![license](https://img.shields.io/github/license/DataDrake/sup.svg)]() 

## Motivation

The prompt is a good place for contextual information, but other tools were slow. So I wrote this instead!

## Goals

 * Spin up and down as fast as reasonable
 * A+ Rating on [Report Card](https://goreportcard.com/report/github.com/DataDrake/cuppa)
 
## Reported Information

* Hostname (SSH)
* Username
* Python Virtual Environment
* Version Control
  * SVN
  * Git (with Branch)
* Working Directory
* Exit codes (including pipelines)

## Requirements

* Go 1.14 (tested)
* Make
* (Optional) Git
* (Optional) Subversion

## Installation

1. Clone repo and enter its directory
2. `make`
3. `sudo make install`

## Usage

Add the following line **exactly** to your `bashrc`:
```
source /usr/share/sup/bash.sh
```

Reload your `bashrc` or open a new terminal. Enjoy!

## License
 
Copyright 2020 Bryan T. Meyers <root@datadrake.com>
 
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
 
http://www.apache.org/licenses/LICENSE-2.0
 
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
