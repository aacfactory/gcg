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

var emptyImports = NewImports()

func NewImports() Imports {
	return make([]*Import, 0, 1)
}

type Imports []*Import

func (n Imports) Empty(v bool) {
	v = len(n) == 0
	return
}

func (n *Imports) Add(v *Import) {
	added := false
	for _, import0 := range *n {
		if import0.Path == v.Path {
			added = true
			break
		}
	}
	if added {
		return
	}
	*n = append(*n, v)
}

func (n *Imports) Merge(x Imports) {
	if x == nil || len(x) == 0 {
		return
	}
	for _, i := range x {
		n.Add(i)
	}
}

type Import struct {
	Name  string
	Alias string
	Path  string
}
