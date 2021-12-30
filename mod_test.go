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
	"runtime"
	"strings"
	"testing"

	"github.com/aacfactory/gcg"
)

func TestModule_String(t *testing.T) {
	version := runtime.Version()
	versionItems := strings.Split(version, ".")
	version = versionItems[0][2:] + "." + versionItems[1]
	mod := gcg.Module{
		Name:      "gcg",
		Path:      "github.com/aacfactory/gcg",
		GoVersion: version,
		Requires:  make([]gcg.Require, 0, 1),
	}
	mod.Requires = append(mod.Requires,
		gcg.Require{
			Name:           "golang.org/x/mod",
			Version:        "v0.5.1",
			Replace:        "",
			ReplaceVersion: "",
			Indirect:       false,
		},
		gcg.Require{
			Name:           "golang.org/x/sys",
			Version:        "v0.0.0-20211019181941-9d821ace8654",
			Replace:        "",
			ReplaceVersion: "",
			Indirect:       true,
		},
		gcg.Require{
			Name:           "golang.org/x/tools",
			Version:        "v0.0.0-20211019181941-9d821ace8654",
			Replace:        "golang.org/x/tools",
			ReplaceVersion: "v0.5.1",
			Indirect:       false,
		},
	)

	fmt.Println(mod)
}
