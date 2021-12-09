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

package gcg_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/aacfactory/gcg"
)

func TestType(t *testing.T) {
	// struct
	s := gcg.Struct()
	sf1 := gcg.StructField("Name")
	sf1.Comments("name", "string")
	sf1.Type(gcg.Ident("string"))
	sf1.Tag("json", "name")
	sf1.Tag("xml", "name")
	s.AddField(sf1)
	sf2 := gcg.StructField("Age")
	sf2.Comments("age", "int")
	sf2.Type(gcg.Ident("int"))
	sf2.Tag("json", "age")
	sf2.Tag("xml", "age")
	s.AddField(sf2)
	sf3 := gcg.StructField("World")
	sf3.Type(gcg.Star().QualifiedIdent(gcg.NewPackage("foo/bar"), "World"))
	s.AddField(sf3)
	ts := gcg.Type("Hello", s.Build())
	fmt.Println(ts.Render(os.Stdout))
	// func
	fnb := gcg.Func()
	fnb.AddParam("p1", gcg.Ident("string"))
	fnb.AddParam("p2", gcg.Star().QualifiedIdent(gcg.NewPackage("foo/bar"), "G"))
	fnb.AddResult("r1", gcg.Ident("string"))
	fnb.AddResult("err", gcg.Ident("error"))
	tf := gcg.Type("Hello", fnb.Build())
	fmt.Println(tf.Render(os.Stdout))
	// ident
	ti := gcg.Type("Hello", gcg.Ident("string"))
	fmt.Println(ti.Render(os.Stdout))

	// interface
	i := gcg.Interface()
	i.Extend(gcg.QualifiedIdent(gcg.NewPackage("io"), "Writer"))
	i.Extend(gcg.QualifiedIdent(gcg.NewPackage("io"), "Reader"))
	ifn1 := gcg.InterfaceFn("D")
	ifn1.AddParam("a", gcg.Ident("string"))
	ifn1.AddResult("err", gcg.Ident("error"))
	i.AddFn(ifn1, "1", "2")
	tx := gcg.Type("Hello", i.Build())
	fmt.Println(tx.Render(os.Stdout))
}

func TestArray(t *testing.T) {

	x := gcg.Array(gcg.Star().QualifiedIdent(gcg.NewPackage("foo/bar"), "A"))

	fmt.Println(x.MapToType().Render(os.Stdout))
	fmt.Println(x.MapToMakeSlice(0, 10).Render(os.Stdout))
	fmt.Println(x.MapToMakeArray(10).Render(os.Stdout))

}

func TestMap(t *testing.T) {
	key := gcg.Ident("string")
	val := gcg.Ident("string")
	x := gcg.Map(key, val)

	fmt.Println(x.MapToType().Render(os.Stdout))
	fmt.Println(x.MapToMake().Render(os.Stdout))

}

func TestChan(t *testing.T) {

	x := gcg.Chan(gcg.Star().QualifiedIdent(gcg.NewPackage("foo/bar"), "A"))

	fmt.Println(x.MapToType().Render(os.Stdout))
	fmt.Println(x.MapToMake(10).Render(os.Stdout))
	fmt.Println(x.MapToSendType().Render(os.Stdout))
	fmt.Println(x.MapToConsumeType().Render(os.Stdout))

}
