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

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func FileExist(path string) (ok bool) {
	_, lstatErr := os.Lstat(path)
	ok = !os.IsNotExist(lstatErr)
	return
}

func FileRender(path string, cleanBeforeWrite bool) (w io.WriteCloser) {
	path = strings.TrimSpace(path)
	if path == "" {
		panic(fmt.Errorf("gcg: new file render failed, path is empty"))
	}
	if FileExist(path) {
		if !cleanBeforeWrite {
			panic(fmt.Errorf("gcg: new file render failed, %s exist", path))
		}
		rmErr := os.Remove(path)
		if rmErr != nil {
			panic(fmt.Errorf("gcg: new file render failed, remove %s failed, %v", path, rmErr))
		}
	}

	f, createErr := os.Create(path)
	if createErr != nil {
		panic(fmt.Errorf("gcg: new file render failed, create %s failed, %v", path, createErr))
	}
	w = f
	return
}
