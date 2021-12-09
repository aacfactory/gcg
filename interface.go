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

import "strings"

func Interface() *InterfaceBuilder {
	return &InterfaceBuilder{
		extends:   make([]Code, 0, 1),
		functions: make([]Code, 0, 1),
	}
}

type InterfaceBuilder struct {
	extends   []Code
	functions []Code
}

func (b *InterfaceBuilder) Extend(code Code) {
	b.extends = append(b.extends)
	return
}

func (b *InterfaceBuilder) AddFn(fn *InterfaceFnBuilder, comments ...string) {
	b.functions = append(b.functions, fn.Build(comments...))
	return
}

func (b *InterfaceBuilder) Build() (code Code) {
	group := Group("{", "}", "\n")
	for _, extend := range b.extends {
		group.Add(extend)
	}
	for _, function := range b.functions {
		group.Add(function)
	}
	code = newStatement().Keyword("interface").Space().Add(group)
	return
}

func InterfaceFn(name string) *InterfaceFnBuilder {
	return &InterfaceFnBuilder{
		name:    strings.TrimSpace(name),
		params:  make([]Code, 0, 1),
		results: make([]Code, 0, 1),
	}
}

type InterfaceFnBuilder struct {
	name    string
	params  []Code
	results []Code
}

func (b *InterfaceFnBuilder) AddParam(name string, typ Code) {
	stmt := newStatement()
	stmt.Ident(name).Space().Add(typ)
	b.params = append(b.params, stmt)
}

func (b *InterfaceFnBuilder) AddResult(name string, typ Code) {
	stmt := newStatement()
	if name != "" {
		stmt.Ident(name).Space()
	}
	stmt.Add(typ)
	b.results = append(b.results, stmt)
}

func (b *InterfaceFnBuilder) Build(comments ...string) (code Code) {
	stmt := newStatement()
	if comments != nil && len(comments) > 0 {
		stmt.Comments(b.name)
		stmt.Comments(comments...)
	}
	stmt.Ident(b.name)
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
		stmt.Token("(")
		if len(b.results) == 1 {
			stmt.Add(b.results[0])
		} else {
			stmt.Token("(")
			for i, result := range b.results {
				if i > 0 {
					stmt.Symbol(",").Space()
				}
				stmt.Add(result)
			}
		}
		stmt.Token(")")
		stmt.Space()
	}
	stmt.Line()
	code = stmt
	return
}
