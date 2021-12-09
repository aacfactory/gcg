/*
 * Copyright 2021 Wang Min Xiang
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * 	http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package gcg

import "io"

type Render interface {
	Render(w io.Writer) (err error)
}

type Code interface {
	Render
	packages() (ps Packages)
}

type Codes []Code

func (c Codes) Render(w io.Writer) (err error) {
	return
}

func (c Codes) packages() (ps Packages) {
	ps = NewPackages()
	for _, code := range c {
		ps.Merge(code.packages())
	}
	return
}
