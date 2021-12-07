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
)

func Var2(ident string) (code *Variable) {
	code = &Variable{
		stmt: newStatement(),
	}
	code.stmt.Add(Keyword("var")).Add(Space()).Add(Ident(ident))
	return
}

type Variable struct {
	stmt *Statement
}

func (v *Variable) Type(op Code) *Variable {
	v.stmt.Add(Space()).Add(op)
	return v
}

func (v *Variable) Op(op string) *Variable {
	v.stmt.Add(Space()).Add(newToken(op))
	return v
}

func (v *Variable) Value(code Code) *Variable {
	v.stmt.Add(Space()).Add(code)
	return v
}

func (v Variable) Render(w io.Writer) (err error) {
	v.stmt.Add(Line())
	err = v.stmt.Render(w)
	if err != nil {
		err = fmt.Errorf("render var failed, %v", err)
		return
	}
	return
}

func (v Variable) imports() (imports Imports) {
	imports = v.stmt.imports()
	return
}

func Vars() *Variables {
	return &Variables{}
}

type Variables []*Variable

func (v *Variables) Add(i *Variable) *Variables {
	if i != nil {
		*v = append(*v, i)
	}
	return v
}

func (v Variables) Render(w io.Writer) (err error) {
	stmt := newStatement()
	stmt.Add(newToken("var"), newToken(" ("), Line())
	for _, variable := range v {
		codes := variable.stmt.Codes()[1:]
		stmt.Add(Tab()).Add(codes...).Add(Line())
	}
	stmt.Add(newToken(")"), Line())
	err = stmt.Render(w)
	return
}

func (v Variables) imports() (imports Imports) {
	imports = NewImports()
	for _, variable := range v {
		imports.Merge(variable.imports())
	}
	return
}
