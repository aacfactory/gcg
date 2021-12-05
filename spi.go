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

type Generator interface {
	Generate() (err error)
}

type Spec interface {
	Build() (imports Imports, code string, err error)
}

type Statement interface {
	Code() (code string, err error)
	Imports() (imports Imports)
}

type Expr interface {
	Imports() (imports Imports)
}

type TypeExpr interface {
	Expr
	Type() (typ *Type)
}
