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

func TestNewFile(t *testing.T) {

	f := gcg.NewFile("foo")

	f.FileComments("file comment1", "file comment2")

	addConst(f)

	addFunc1(f)

	addTypeStruct(f)
	addStructFunc1(f)
	addStructFunc2(f)

	renderErr := f.Render(os.Stdout)
	if renderErr != nil {
		t.Errorf("%v", renderErr)
	}

}

func addConst(f *gcg.File) {

	constCode := gcg.Constants()
	constCode.Add("Namespace", "campaigns")
	constCode.Add("QueryFn", `query`)
	constCode.Add("GetFn", `get`)

	f.AddCode(constCode.Build())

}

func addFunc1(f *gcg.File) {

	code := gcg.Func()
	code.Name("Service")
	code.AddResult("svc", gcg.QualifiedIdent(gcg.NewPackage("github.com/aacfactory/fns"), "Service"))
	body := gcg.Ident("svc").Symbol(":=").Symbol("&").Token("service{}").Line()
	body.Keyword("return")
	code.Body(body)

	f.AddCode(code.Build())

}

func addTypeStruct(f *gcg.File) {

	s := gcg.Struct()
	code := gcg.Type("service", s.Build())

	f.AddCode(code)

}

func addStructFunc1(f *gcg.File) {

	code := gcg.Func()

	code.Receiver("s", gcg.Token("*service"))
	code.Name("Namespace")
	code.AddResult("namespace", gcg.String())

	body := gcg.Ident("namespace").Symbol("=").Ident("Namespace").Line().Keyword("return")

	code.Body(body)

	f.AddCode(code.Build())

}

func addStructFunc2(f *gcg.File) {

	code := gcg.Func()
	code.Receiver("s", gcg.Token("*service"))
	code.Name("Handle")
	code.AddParam("ctx", gcg.QualifiedIdent(gcg.NewPackage("github.com/aacfactory/fns"), "Context"))
	code.AddParam("fn", gcg.String())
	code.AddParam("argument", gcg.QualifiedIdent(gcg.NewPackage("github.com/aacfactory/fns"), "Argument"))
	code.AddResult("result", gcg.Token("interface{}"))
	code.AddResult("err", gcg.QualifiedIdent(gcg.NewPackage("github.com/aacfactory/errors"), "CodeError"))

	body := gcg.Statements()

	body.Keyword("switch").Space().Ident("fn").Space().Symbol("{").Line()

	body.Keyword("case").Space().Ident("QueryFn").Symbol(":").Line()
	body.Ident("result").Symbol(",").Ident("err").Symbol("=").Ident("s").Token("invokeQueryFn(ctx, argument)").Line()

	body.Keyword("case").Space().Ident("GetFn").Symbol(":").Line()
	body.Ident("result").Symbol(",").Ident("err").Symbol("=").Ident("s").Token("svc.invokeGetFn(ctx, argument)").Line()
	body.Keyword("default").Symbol(":").Line()
	body.Ident("err").Symbol("=").Token("errors.NotFound(fmt.Sprintf(\"%s/%s 的句柄未找到\", Namespace, fn))", gcg.NewPackage("github.com/aacfactory/errors"), gcg.FmtPackage()).Line()

	body.Keyword("}").Line()
	body.Keyword("return")

	code.Body(body)

	f.AddCode(code.Build())
}

func TestFileExist(t *testing.T) {
	fmt.Println("ok", gcg.FileExist("/Users/wangminxiang/Game/sss"), gcg.FileExist("/Users/wangminxiang/Game/sss/a.txt"))
	fmt.Println("ko", gcg.FileExist("/Users/wangminxiang/Game/ss"), gcg.FileExist("/Users/wangminxiang/Game/sss/b.txt"))
}

func TestFileRender(t *testing.T) {
	x := "/Users/wangminxiang/Game/sss/a.txt"
	w := gcg.FileRender(x, true)
	w.Write([]byte("x"))
	w.Close()
}
