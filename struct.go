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
	"sort"
	"strings"
)

func Struct() *StructBuilder {
	return &StructBuilder{
		fields:  make([]Code, 0, 1),
		methods: make([]Code, 0, 1),
	}
}

type StructBuilder struct {
	fields  []Code
	methods []Code
}

func (s *StructBuilder) AddField(field *StructFieldBuilder) {
	s.fields = append(s.fields, field.Build())
}

func (s *StructBuilder) AddMethod(fn *FuncBuilder) {
	s.methods = append(s.methods, fn.Build())
}

func (s *StructBuilder) Build() (code Code) {
	stmt := newStatement()
	stmt.Keyword("struct").Space()
	if len(s.fields) == 0 {
		stmt.Token("{").Token("}")
	} else {
		group := Group("{", "}", "\n")
		for _, field := range s.fields {
			group.Add(field)
		}
		stmt.Add(group).Line()
	}
	if len(s.methods) > 0 {
		for _, method := range s.methods {
			stmt.Add(method).Line()
		}
	}
	code = stmt
	return
}

func StructField(name string) *StructFieldBuilder {
	return &StructFieldBuilder{
		comments: nil,
		name:     strings.TrimSpace(name),
		typ:      nil,
		tags:     make(map[string]string),
	}
}

type StructFieldBuilder struct {
	comments []string
	name     string
	typ      Code
	tags     map[string]string
}

func (b *StructFieldBuilder) Type(code Code) {
	b.typ = code
}

func (b *StructFieldBuilder) Tag(k string, v string) {
	b.tags[strings.TrimSpace(k)] = strings.TrimSpace(v)
}

func (b *StructFieldBuilder) Comments(comments ...string) {
	b.comments = comments
}

func (b *StructFieldBuilder) Build() (code Code) {
	stmt := newStatement()
	if b.comments != nil && len(b.comments) > 0 {
		stmt.Comments(b.name)
		stmt.Comments(b.comments...)
	}
	stmt.Ident(b.name).Space().Add(b.typ)
	if len(b.tags) > 0 {
		content := ""
		tks := make([]string, 0, 1)
		for k := range b.tags {
			tks = append(tks, k)
		}
		sort.Strings(tks)
		for _, k := range tks {
			v := b.tags[k]
			content = content + " " + fmt.Sprintf("%s:\"%s\"", k, v)
		}
		if len(content) > 0 {
			content = content[1:]
		}
		stmt.Space().Token("`").Token(content).Token("`")
	}
	code = stmt
	return
}
