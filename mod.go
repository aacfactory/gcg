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

import (
	"bytes"
	"fmt"
)

type Module struct {
	Name      string
	Path      string
	GoVersion string
	Requires  []Require
}

func (mod Module) String() (s string) {
	b := bytes.NewBufferString("")
	_, _ = b.WriteString(fmt.Sprintf("module %s\n\n", mod.Path))
	_, _ = b.WriteString(fmt.Sprintf("go %s\n\n", mod.GoVersion))
	if len(mod.Requires) > 0 {
		replaces := make([]Require, 0, 1)
		_, _ = b.WriteString("require (\n")
		for _, require := range mod.Requires {
			_, _ = b.WriteString(fmt.Sprintf("\t%s %s", require.Name, require.Version))
			if require.Indirect {
				_, _ = b.WriteString(" // indirect\n")
			} else {
				_, _ = b.WriteString("\n")
			}
			if require.Replace != "" {
				replaces = append(replaces, require)
			}
		}
		_, _ = b.WriteString(")\n\n")

		if len(replaces) > 0 {
			_, _ = b.WriteString("replace (\n")
			for _, require := range replaces {
				_, _ = b.WriteString(fmt.Sprintf("\t%s %s => %v %v \n", require.Name, require.Version, require.Replace, require.ReplaceVersion))
			}
			_, _ = b.WriteString(")\n\n")
		}
	}
	s = b.String()
	return
}

type Require struct {
	Name           string
	Version        string
	Replace        string
	ReplaceVersion string
	Indirect       bool
}
