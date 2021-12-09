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

func ForRangeArray(ident Code, indexIdent string, elementIdent string, block Code) (code Code) {
	stmt := newStatement()
	stmt.Keyword("for").Space().Ident(indexIdent).Symbol(",").Space().Ident(elementIdent).Space().ColonEqual().Space().Keyword("range").Space().Add(ident).Space().Symbol("{").Line()
	stmt.Tab().Add(block).Line()
	stmt.Symbol("}").Line()
	code = stmt
	return
}

func For(init Code, cond Code, post Code, block Code) (code Code) {
	stmt := newStatement()
	stmt.Keyword("for").Space().Add(init).Symbol(";").Space().Add(cond).Symbol(";").Space().Add(post).Space().Symbol("{").Line()
	stmt.Tab().Add(block).Line()
	stmt.Symbol("}").Line()
	code = stmt
	return
}

func ForCond(cond Code, block Code) (code Code) {
	stmt := newStatement()
	stmt.Keyword("for").Space().Add(cond).Space().Symbol("{").Line()
	stmt.Tab().Add(block).Line()
	stmt.Symbol("}").Line()
	code = stmt
	return
}

func ForDead(block Code) (code Code) {
	stmt := newStatement()
	stmt.Keyword("for").Space().Symbol("{").Line()
	stmt.Tab().Add(block).Line()
	stmt.Symbol("}").Line()
	code = stmt
	return
}
