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

func If(expression Code, block Code) *IfBuilder {
	return &IfBuilder{
		start: &IfElement{
			expression: expression,
			block:      block,
		},
		middens: make([]*IfElement, 0, 1),
		end:     nil,
	}
}

type IfBuilder struct {
	start   *IfElement
	middens []*IfElement
	end     *IfElement
}

func (b *IfBuilder) ElseIf(expression Code, block Code) {
	b.middens = append(b.middens, &IfElement{
		expression: expression,
		block:      block,
	})
	return
}

func (b *IfBuilder) Else(block Code) {
	b.end = &IfElement{
		expression: nil,
		block:      block,
	}
	return
}

func (b *IfBuilder) Build() (code Code) {
	stmt := newStatement()
	stmt.Keyword("if").Space().Add(b.start.expression).Space().Symbol("{").Line()
	stmt.Add(b.start.block).Line()
	stmt.Symbol("}")
	if len(b.middens) > 0 {
		for _, midden := range b.middens {
			stmt.Space().Keyword("else").Space().Keyword("if").Space().Add(midden.expression).Space().Symbol("{").Line()
			stmt.Add(midden.block).Line()
			stmt.Symbol("}")
		}
	}
	if b.end != nil {
		stmt.Space().Keyword("else").Space().Symbol("{").Line()
		stmt.Add(b.end.block).Line()
		stmt.Symbol("}")
	}
	stmt.Line()
	code = stmt
	return
}

type IfElement struct {
	expression Code
	block      Code
}
