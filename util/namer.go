// Copyright © 2018 Joel Kratz joel@kratz.nu
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package util

import (
	"crypto/rand"
	"encoding/hex"
)

// Descriptive adjectives
var prefix = []string{
	"Adorable", "Attractive", "Alluring",
	"Beautiful", "Bewildered", "Boorish",
	"Bright", "Confident", "Cheerful",
	"Cultured", "Clumsy", "Drab",
	"Dull", "Dynamic", "Disillusioned",
	"Elegant", "Energetic", "Fair",
	"Fancy", "Filthy", "Gentle",
	"Glamorous", "Handsome", "Homely",
	"Hurt", "Jolly",
	"Lovely", "Magnificent", "Neat",
	"Nervous", "Pleasant", "Perfect",
	"Plucky", "Prim", "Smiling",
	"Splendid", "Snobbish",
	"Thoughtful", "Tense", "Timid",
	"Upset", "Vivacious", "Wonderful",
	"Worried", "Wild", "Zaftig",
}

// GenerateName with a given length
func GenerateName(n int) (string, error) {
	bytes := make([]byte, n)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}
