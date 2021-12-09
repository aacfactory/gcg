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

import (
	"fmt"
)

func (s *Statement) Comments(texts ...string) *Statement {
	if texts == nil || len(texts) == 0 {
		return s
	}
	for _, text := range texts {
		s.Add(newToken(fmt.Sprintf("// %s\n", text)))
	}
	return s
}

func Comments(texts ...string) (code Code) {
	code = newStatement().Comments(texts...)
	return
}
