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
