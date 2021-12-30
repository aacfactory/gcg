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
	"path/filepath"
	"strings"
	"testing"

	"github.com/aacfactory/gcg"
)

func TestNewProject(t *testing.T) {
	gopath, _ := os.LookupEnv("GOPATH")
	if gopath == "" {
		t.Errorf("need $GOPATH")
		return
	}
	gopath = strings.ReplaceAll(gopath, "\\", "/")
	rootDir := filepath.Join(gopath, "github.com/aacfactory/tests")

	p, pErr := gcg.NewProject(
		rootDir,
		gcg.ProjectRemoveBeforeRender(true),
		gcg.ProjectGoProxy("https://goproxy.cn"),
		gcg.ProjectRenderHook(gcg.ProjectGoGenerateHook),
	)
	if pErr != nil {
		t.Errorf("%v", pErr)
		return
	}
	fmt.Println(p)
}
