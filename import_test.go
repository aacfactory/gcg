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
	"github.com/aacfactory/gcg"
	"testing"
)

func TestNewImports(t *testing.T) {
	imports := gcg.NewImports()
	imports.Add(&gcg.Import{
		Name:  "foo",
		Alias: "bar",
		Path:  "x/foo",
	})
	imports.Add(&gcg.Import{
		Name:  "foo",
		Alias: "bar",
		Path:  "x/foo",
	})
	imports.Add(&gcg.Import{
		Name:  "boo",
		Alias: "bar",
		Path:  "x/boo",
	})
	for _, i := range imports {
		fmt.Println(i)
	}
}
