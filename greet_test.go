/*
Copyright 2018 The pkg103greet Authors.
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
// pkg103greet returns a greeting in a few different languages.
package pkg103greet

import "testing"

// Table driven test
func TestGreet(t *testing.T) {
	var tests = []struct {
		languageCode string
		greetings    string
	}{
		{"en", "Hello"},
		{"fr", "Bonjour"},
		{"it", "Ciao"},
		{"po", "cześć"},
		{"sp", "Hola"},
		{"ru", "Здравствуйте"},
	}
	for _, test := range tests {
		got := greetMe(test.languageCode)
		if got != test.greetings {
			t.Errorf("Want: %s but got %s", test.greetings, got)
		}
	}
}
