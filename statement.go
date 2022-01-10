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
	"io"
	"strings"
)

func Statements() *Statement {
	return newStatement()
}

func newStatement() *Statement {
	return &Statement{}
}

type Statement []Code

func (s *Statement) Add(code ...Code) *Statement {
	if code == nil {
		return s
	}
	*s = append(*s, code...)
	return s
}

func (s Statement) Codes() (codes []Code) {
	codes = s
	return
}

func (s Statement) Render(w io.Writer) (err error) {
	if len(s) == 0 {
		return
	}
	for _, code := range s {
		if code == nil {
			continue
		}
		err = code.Render(w)
		if err != nil {
			return
		}
	}
	return
}

func (s Statement) packages() (ps Packages) {
	ps = NewPackages()
	for _, code := range s {
		if code == nil {
			continue
		}
		ps.Merge(code.packages())
	}
	return
}

func Group(open string, close string, separator string) *StatementGroup {
	return &StatementGroup{
		items:     make([]Code, 0, 1),
		open:      strings.TrimSpace(open),
		close:     strings.TrimSpace(close),
		separator: separator,
	}
}

type StatementGroup struct {
	items     []Code
	open      string
	close     string
	separator string
}

func (s *StatementGroup) Add(code Code) *StatementGroup {
	if code == nil {
		return s
	}
	s.items = append(s.items, code)
	return s
}

func (s StatementGroup) Render(w io.Writer) (err error) {
	stmt := newStatement()
	stmt.Add(newToken(s.open)).Line()
	for _, item := range s.items {
		if item == nil {
			continue
		}
		stmt.Tab().Add(item).Add(newToken(s.separator))
	}
	stmt.Add(newToken(s.close))
	err = stmt.Render(w)
	return
}

func (s StatementGroup) packages() (ps Packages) {
	ps = NewPackages()
	for _, item := range s.items {
		if item == nil {
			continue
		}
		ps.Merge(item.packages())
	}
	return
}
