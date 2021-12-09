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

package gcg_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/aacfactory/gcg"
)

func TestNewImports(t *testing.T) {
	ps := gcg.NewPackages()
	ps.Add(gcg.NewPackage("fmt")).Add(gcg.NewPackageWithAlias("github.com/aacfactory/gcg", "a")).Add(gcg.NewPackage("github.com/aacfactory/gcg"))
	ps.Add(gcg.NewPackage("xx/gg"))
	code := ps.MapToCode()
	fmt.Println(code.Render(os.Stdout))
}
