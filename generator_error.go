/*
Copyright 2014 Zachary Klippenstein

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

   http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package regen

import (
	"fmt"
)

type GeneratorError struct {
	ErrorStr string
	Cause    error
}

func generatorError(cause error, format string, args ...interface{}) error {
	return GeneratorError{fmt.Sprintf(format, args...), cause}
}

func (err GeneratorError) Error() string {
	return fmt.Sprintf("%s\ncaused by %s", err.ErrorStr, err.Cause.Error())
}
