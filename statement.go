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
	"path"
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

func (s *Statement) Symbol(v string) *Statement {
	s.Add(newToken(v))
	return s
}

func Symbol(s string) (stmt *Statement) {
	stmt = newStatement().Symbol(s)
	return
}

func (s *Statement) Space() *Statement {
	s.Add(newToken(" "))
	return s
}

func Space() (stmt *Statement) {
	stmt = newStatement().Space()
	return
}

func (s *Statement) Line() *Statement {
	s.Add(newToken("\n"))
	return s
}

func Line() (stmt *Statement) {
	stmt = newStatement().Line()
	return
}

func (s *Statement) Tab() *Statement {
	s.Add(newToken("\t"))
	return s
}

func Tab() (stmt *Statement) {
	stmt = newStatement().Tab()
	return
}

func (s *Statement) Tabs(n int) *Statement {
	for i := 0; i < n; i++ {
		s.Add(newToken("\t"))
	}
	return s
}

func Tabs(n int) (stmt *Statement) {
	stmt = newStatement().Tabs(n)
	return
}

func (s *Statement) Ident(name string) *Statement {
	s.Add(newToken(name))
	return s
}

func Ident(name string) (stmt *Statement) {
	stmt = newStatement().Ident(name)
	return
}

func (s *Statement) QualifiedIdent(packagePath string, packageAlias, name string) *Statement {
	packagePath = strings.TrimSpace(packagePath)
	_, packageName := path.Split(packagePath)
	packageAlias = strings.TrimSpace(packageAlias)
	if packageAlias != "" {
		packageName = packageAlias
	}
	_import := &Import{
		Name:  packageName,
		Alias: packageAlias,
		Path:  packagePath,
	}
	s.Add(newToken(_import.Name), newToken("."), newToken(name, _import))
	return s
}

func QualifiedIdent(packagePath string, packageAlias, name string) (stmt *Statement) {
	stmt = newStatement().QualifiedIdent(packagePath, packageAlias, name)
	return
}

func (s *Statement) Star() *Statement {
	s.Add(Symbol("*"))
	return s
}

func Star() (stmt *Statement) {
	stmt = newStatement().Star()
	return
}

func (s *Statement) Call(params ...Code) *Statement {
	s.Add(newToken("("))
	if params != nil {
		for i, param := range params {
			s.Add(param)
			if i > 0 && i < len(params)-1 {
				s.Add(newToken(","), Space())
			}
		}
	}
	s.Add(newToken(")"))
	return s
}

func Call(params ...Code) (stmt *Statement) {
	stmt = newStatement().Call(params...)
	return
}

func (s *Statement) Comment(texts ...string) *Statement {
	if texts == nil || len(texts) == 0 {
		return s
	}
	for _, text := range texts {
		s.Add(newToken(fmt.Sprintf("// %s\n", text)))
	}
	return s
}

func Comment(texts ...string) (stmt *Statement) {
	stmt = newStatement().Comment(texts...)
	return
}

func (s *Statement) Var(name string) *Statement {
	if name == "" {
		return s
	}
	s.Add(Keyword("var")).Add(Ident(name)).Add(Space())
	return s
}

func Var(name string) (stmt *Statement) {
	// all expr move into stmt
	stmt = newStatement().Var(name)
	return
}
