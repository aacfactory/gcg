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

func Switch() *SwitchBuilder {
	return &SwitchBuilder{
		expression:   nil,
		cases:        make([]*SwitchCase, 0, 1),
		defaultBlock: nil,
	}
}

type SwitchBuilder struct {
	expression   Code
	cases        []*SwitchCase
	defaultBlock Code
}

func (b *SwitchBuilder) Expression(code Code) {
	b.expression = code
}

func (b *SwitchBuilder) Case(expression Code, block Code) {
	b.cases = append(b.cases, &SwitchCase{
		expression: expression,
		block:      block,
	})
}

func (b *SwitchBuilder) Default(block Code) {
	b.defaultBlock = block
}

func (b *SwitchBuilder) Build() (code Code) {
	stmt := newStatement()
	stmt.Keyword("switch")
	if b.expression != nil {
		stmt.Space().Add(b.expression)
	}
	stmt.Space().Symbol("{").Line()
	for _, sc := range b.cases {
		stmt.Keyword("case").Space().Add(sc.expression).Colon().Line()
		stmt.Tab().Add(sc.block).Line()
	}
	if b.defaultBlock != nil {
		stmt.Keyword("default").Colon().Line()
		stmt.Tab().Add(b.defaultBlock).Line()
	}
	stmt.Symbol("}").Line()
	code = stmt
	return
}

type SwitchCase struct {
	expression Code
	block      Code
}
