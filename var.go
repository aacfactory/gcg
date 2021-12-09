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

func Vars() *VarsBuilder {
	return &VarsBuilder{
		items: make([]*VarBuilder, 0, 1),
	}
}

type VarsBuilder struct {
	items []*VarBuilder
}

func (b *VarsBuilder) Add(v *VarBuilder) {
	b.items = append(b.items, v)
	return
}

func (b *VarsBuilder) Build() (code Code) {
	stmt := newStatement()
	if len(b.items) == 0 {
		code = stmt
		return
	}
	group := Group("(", ")", "\n")
	for _, item := range b.items {
		group.Add(item.BuildVarsElement())
	}
	stmt.Keyword("var").Space().Add(group).Line()
	code = stmt
	return
}

func Var(name string, typ Code) *VarBuilder {
	return &VarBuilder{
		ident: name,
		typ:   typ,
		value: nil,
	}
}

type VarBuilder struct {
	ident    string
	typ      Code
	value    Code
	comments []string
}

func (b *VarBuilder) Comments(text ...string) {
	b.comments = text
}

func (b *VarBuilder) Value(v Code) {
	b.value = v
}

func (b *VarBuilder) Build() (code Code) {
	stmt := newStatement()
	if b.comments != nil && len(b.comments) > 0 {
		stmt.Comments(b.ident)
		stmt.Comments(b.comments...)
	}
	stmt.Keyword("var").Space().Ident(b.ident).Space().Add(b.typ)
	if b.value != nil {
		stmt.Space().Equal().Space().Add(b.value)
	}
	stmt.Line()
	code = stmt
	return
}

func (b *VarBuilder) BuildVarsElement() (code Code) {
	stmt := newStatement()
	if b.comments != nil && len(b.comments) > 0 {
		stmt.Comments(b.ident)
		stmt.Comments(b.comments...)
	}
	stmt.Ident(b.ident).Space().Add(b.typ)
	if b.value != nil {
		stmt.Space().Equal().Space().Add(b.value)
	}
	stmt.Line()
	code = stmt
	return
}
