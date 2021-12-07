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
	"fmt"
	"io"
	"strings"
)

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
		err = code.Render(w)
		if err != nil {
			return
		}
	}
	return
}

func (s Statement) imports() (imports Imports) {
	imports = NewImports()
	for _, code := range s {
		imports.Merge(code.imports())
	}
	return
}

func newStatements() *Statements {
	return &Statements{}
}

type Statements []*Statement

func (s *Statements) Render(w io.Writer) (err error) {
	for _, statement := range *s {
		err = statement.Render(w)
		if err != nil {
			return
		}
	}
	return
}

func (s *Statements) imports() (imports Imports) {
	imports = NewImports()
	for _, statement := range *s {
		imports.Merge(statement.imports())
	}
	return
}

func (s *Statements) Add(v *Statement) *Statements {
	if v == nil {
		return s
	}
	*s = append(*s, v)
	return s
}

func Symbol(s string) (code Code) {
	code = newToken(s)
	return
}

func Space() (code Code) {
	code = newToken(" ")
	return
}

func Line() (code Code) {
	code = newToken("\n")
	return
}

func Tab() (code Code) {
	code = Tabs(1)
	return
}

func Tabs(n int) (code Code) {
	content := ""
	for i := 0; i < n; i++ {
		content = content + "\t"
	}
	code = newToken(content)
	return
}

func newToken(content string) *token {
	return &token{
		content: strings.TrimSpace(content),
	}
}

type token struct {
	content string
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

func (t token) imports() (imports Imports) {
	imports = emptyImports
	return
}
