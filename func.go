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

func (s *Statement) Call(fn Code, params ...Code) *Statement {
	s.Add(fn)
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

func Call(fn Code, params ...Code) (stmt *Statement) {
	stmt = newStatement().Call(fn, params...)
	return
}

func Func() (builder *FuncBuilder) {
	builder = &FuncBuilder{
		comments: nil,
		receiver: nil,
		name:     "",
		params:   make([]Code, 0, 1),
		results:  make([]Code, 0, 1),
		body:     nil,
	}
	return
}

type FuncBuilder struct {
	comments []string
	receiver Code
	name     string
	params   []Code
	results  []Code
	body     Code
}

func (b *FuncBuilder) Receiver(ident string, typ Code) {
	b.receiver = newStatement().Ident(ident).Space().Add(typ)
}

func (b *FuncBuilder) Name(name string) {
	b.name = name
}

func (b *FuncBuilder) Comments(comments ...string) {
	if comments != nil {
		b.comments = comments
	}
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
	if len(b.comments) > 0 {
		stmt.Comments(b.name)
		stmt.Comments(b.comments...)
	}
	stmt.Keyword("func")
	if b.receiver != nil {
		stmt.Token("(").Add(b.receiver).Token(")")
	}
	stmt.Space()
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
		stmt.Token("(")
		if len(b.results) == 1 {
			stmt.Add(b.results[0])
		} else {
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
	if b.body != nil {
		stmt.Symbol("{").Line().Add(b.body).Line().Symbol("}").Line()
	}
	code = stmt
	return
}
