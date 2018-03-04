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
// pkg103greet library contains greetings in few different languages.
package pkg103greet

// GreetMe takes the country code from en, it, po, sp and ru and returns 'Hello' in that language
func GreetMe(langcode string) string {

	var greet string
	switch langcode {
	case "fr":
		greet = "Bonjour"
	case "it":
		greet = "Ciao"
	case "po":
		greet = "cześć"
	case "sp":
		greet = "Hola"
	case "ru":
		greet = "Здравствуйте"
	default:
		greet = "Hello"
	}
	return greet
}
