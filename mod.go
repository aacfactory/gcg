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

type Module struct {
	Name      string
	Path      string
	GoVersion string
	Requires  []Require
}

func (mod Module) String() (s string) {
	s = mod.Name + " " + mod.Path + "\n"
	s = s + mod.GoVersion + "\n"
	for _, require := range mod.Requires {
		s = s + require.Name + " " + require.Version
		if require.Replace != "" {
			s = s + " => " + require.Replace
		}
		s = s + "\n"
	}
	return
}

type Require struct {
	Name           string
	Version        string
	Replace        string
	ReplaceVersion string
	Indirect       bool
}
