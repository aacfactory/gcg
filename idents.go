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

func Error() (stmt *Statement) {
	stmt = newStatement().Ident("error")
	return
}

func String() (stmt *Statement) {
	stmt = newStatement().Ident("string")
	return
}

func Int() (stmt *Statement) {
	stmt = newStatement().Ident("int")
	return
}

func Int64() (stmt *Statement) {
	stmt = newStatement().Ident("int64")
	return
}

func UInt() (stmt *Statement) {
	stmt = newStatement().Ident("uint")
	return
}

func UInt64() (stmt *Statement) {
	stmt = newStatement().Ident("uint64")
	return
}

func Float32() (stmt *Statement) {
	stmt = newStatement().Ident("float32")
	return
}

func Float64() (stmt *Statement) {
	stmt = newStatement().Ident("float64")
	return
}

func True() (stmt *Statement) {
	stmt = newStatement().Ident("true")
	return
}

func False() (stmt *Statement) {
	stmt = newStatement().Ident("false")
	return
}

func Byte() (stmt *Statement) {
	stmt = newStatement().Ident("byte")
	return
}

func Context() (stmt *Statement) {
	stmt = newStatement().QualifiedIdent(NewPackage("context"), "Context")
	return
}
