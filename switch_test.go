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

func TestSwitch(t *testing.T) {

	x := gcg.Switch()

	x.Expression(gcg.Ident("x"))
	x.Case(gcg.Ident("x").Space().GT().Space().Literal(10), gcg.Return())
	x.Case(gcg.Ident("x").Space().GT().Space().Literal(20), gcg.Return())
	x.Case(gcg.Ident("x").Space().GT().Space().Literal(30), gcg.Return())
	x.Default(gcg.Return())

	fmt.Println(x.Build().Render(os.Stdout))

}
