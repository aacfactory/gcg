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

import "strings"

func (s *Statement) Call(fnName string, params ...Code) *Statement {
	fnName = strings.TrimSpace(fnName)
	if fnName != "" {
		s.Token(fnName)
	}
	s.Token("(")
	if params != nil {
		for i, param := range params {
			s.Add(param)
			if i > 0 && i < len(params)-1 {
				s.Symbol(",").Space()
			}
		}
	}
	s.Token(")")
	return s
}

func Call(fnName string, params ...Code) (stmt *Statement) {
	stmt = newStatement().Call(fnName, params...)
	return
}

func Func() (builder *FuncBuilder) {
	builder = &FuncBuilder{
		name:    "",
		params:  make([]Code, 0, 1),
		results: make([]Code, 0, 1),
		body:    nil,
	}
	return
}

type FuncBuilder struct {
	name    string
	params  []Code
	results []Code
	body    Code
}

func (b *FuncBuilder) Name(name string) {
	b.name = name
}

func (b *FuncBuilder) AddParam(name string, typ Code) {
	stmt := newStatement()
	stmt.Ident(name).Space().Add(typ)
	b.params = append(b.params, stmt)
}

func (b *FuncBuilder) AddResult(name string, typ Code) {
	stmt := newStatement()
	if name != "" {
		stmt.Ident(name).Space()
	}
	stmt.Add(typ)
	b.results = append(b.results, stmt)
}

func (b *FuncBuilder) Body(body Code) {
	b.body = body
}

func (b *FuncBuilder) Build() (code Code) {
	stmt := newStatement()
	stmt.Keyword("func").Space()
	if b.name != "" {
		stmt.Ident(b.name)
	}
	stmt.Token("(")
	if len(b.params) > 0 {
		for i, param := range b.params {
			if i > 0 {
				stmt.Symbol(",").Space()
			}
			stmt.Add(param)
		}
	}
	stmt.Token(")").Space()
	if len(b.results) > 0 {
		if len(b.results) == 1 {
			stmt.Add(b.results[0])
		} else {
			// todo
			stmt.Token("(")
			for i, result := range b.results {
				if i > 0 {
					stmt.Symbol(",").Space()
				}
				stmt.Add(result)
			}
			stmt.Token(")")
		}
		stmt.Space()
	}
	if b.body != nil {
		stmt.Symbol("{").Line().Add(b.body).Line().Symbol("}").Line()
	}
	code = stmt
	return
}
