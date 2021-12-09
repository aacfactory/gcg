/*
 * Copyright 2021 Wang Min Xiang
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *  	http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package gcg

import (
	"fmt"
	"io"
)

func (s *Statement) Token(name string, ps ...*Package) *Statement {
	s.Add(newToken(name, ps...))
	return s
}

func Token(name string, ps ...*Package) (stmt *Statement) {
	stmt = newStatement().Token(name, ps...)
	return
}

func newToken(content string, ps ...*Package) *token {
	t := &token{
		content: content,
		ps:      NewPackages(),
	}
	if ps != nil {
		for _, i := range ps {
			t.ps.Add(i)
		}
	}
	return t
}

type token struct {
	content string
	ps      Packages
}

func (t token) Render(w io.Writer) (err error) {
	if t.content == "" {
		return
	}
	_, err = w.Write([]byte(t.content))
	if err != nil {
		err = fmt.Errorf("render token %s failed, %v", t.content, err)
		return
	}
	return
}

func (t token) packages() (ps Packages) {
	ps = t.ps
	return
}

func (t token) addPackage(i *Package) {
	t.ps.Add(i)
	return
}
