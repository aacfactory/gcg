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
	"strings"
)

func Struct() *StructBuilder {
	return &StructBuilder{
		group: Group("{", "}", "\n"),
	}
}

type StructBuilder struct {
	group *StatementGroup
}

func (s *StructBuilder) AddField(ident string, typ Code, tag Code, comments ...string) {
	field := newStatement()
	if comments != nil && len(comments) > 0 {
		field.Comments(ident).Comments(comments...)
	}
	field.Ident(ident).Space().Add(typ)
	if tag != nil {
		field.Space().Add(tag)
	}
	s.group.Add(field)
}

func (s *StructBuilder) Build() (code Code) {
	stmt := newStatement()
	stmt.Keyword("struct").Space().Add(s.group)
	code = stmt
	return
}

func StructFieldTag(ss ...string) (code Code) {
	content := ""
	size := len(ss) / 2
	for i := 0; i < size; i++ {
		key := strings.TrimSpace(ss[i])
		val := strings.TrimSpace(ss[i+1])
		content = content + " " + fmt.Sprintf("%s:\"%s\"", key, val)
	}
	if len(content) > 0 {
		content = content[1:]
	}
	stmt := newStatement()
	stmt.Token("`").Token(content).Token("`")
	code = stmt
	return
}
