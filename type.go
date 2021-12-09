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

func Type(ident string, element Code, comments ...string) (code Code) {
	stmt := newStatement()
	if comments != nil && len(comments) > 0 {
		stmt.Comments(ident)
		stmt.Comments(comments...)
	}
	stmt.Keyword("type").Space().Ident(ident).Space().Add(element).Line()
	code = stmt
	return
}

func Array(elm Code) *ArrayBuilder {
	return &ArrayBuilder{
		elm: elm,
	}
}

type ArrayBuilder struct {
	elm Code
}

func (b *ArrayBuilder) MapToType() (code Code) {
	stmt := newStatement().Token("[]").Add(b.elm)
	code = stmt
	return
}

func (b *ArrayBuilder) MapToMakeSlice(length int, capacity int) (code Code) {
	stmt := newStatement().Token("make").Symbol("(").Token("[]").Add(b.elm).Symbol(",").Space().Literal(length).Symbol(",").Space().Literal(capacity).Symbol(")")
	code = stmt
	return
}

func (b *ArrayBuilder) MapToMakeArray(capacity int) (code Code) {
	stmt := newStatement().Token("make").Symbol("(").Token("[]").Add(b.elm).Symbol(",").Space().Literal(capacity).Symbol(")")
	code = stmt
	return
}

func Map(key Code, val Code) *MapBuilder {
	return &MapBuilder{
		key: key,
		val: val,
	}
}

type MapBuilder struct {
	key Code
	val Code
}

func (b *MapBuilder) MapToType() (code Code) {
	stmt := newStatement().Ident("map").Symbol("[").Add(b.key).Symbol("]").Add(b.val)
	code = stmt
	return
}

func (b *MapBuilder) MapToMake() (code Code) {
	stmt := newStatement().Ident("make").Symbol("(").Add(b.MapToType()).Symbol(")")
	code = stmt
	return
}

func Chan(elm Code) *ChanBuilder {
	return &ChanBuilder{
		elm: elm,
	}
}

type ChanBuilder struct {
	elm Code
}

func (b *ChanBuilder) MapToType() (code Code) {
	stmt := newStatement().Ident("chan").Space().Add(b.elm)
	code = stmt
	return
}

func (b *ChanBuilder) MapToSendType() (code Code) {
	stmt := newStatement().Ident("chan<-").Space().Add(b.elm)
	code = stmt
	return
}

func (b *ChanBuilder) MapToConsumeType() (code Code) {
	stmt := newStatement().Ident("<-chan").Space().Add(b.elm)
	code = stmt
	return
}

func (b *ChanBuilder) MapToMake(length int) (code Code) {
	stmt := newStatement().Ident("make").Symbol("(").Ident("chan").Space().Add(b.elm).Symbol(",").Space().Literal(length).Symbol(")")
	code = stmt
	return
}
