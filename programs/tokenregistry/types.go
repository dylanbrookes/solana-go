// Copyright 2020 dfuse Platform Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package tokenregistry

import (
	"fmt"

	"github.com/dfuse-io/solana-go"
)

type TokenMeta struct {
	IsInitialized         bool
	Reg                   [3]byte `text:"-"`
	DataType              byte
	MintAddress           *solana.PublicKey
	RegistrationAuthority *solana.PublicKey
	Logo                  Logo
	Name                  Name
	Symbol                Symbol
}

type Logo [32]byte

func LogoFromString(logo string) (Logo, error) {
	data := []byte(logo)
	if len(data) > 32 {
		return Logo{}, fmt.Errorf("logo data to long expected 32 got %d", len(data))
	}
	l := Logo{}
	copy(l[:], data)
	return l, nil
}
func (l Logo) String() string {
	return AsciiString(l[:])
}

type Name [32]byte

func NameFromString(name string) (Name, error) {
	data := []byte(name)
	if len(data) > 32 {
		return Name{}, fmt.Errorf("name data to long expected 32 got %d", len(data))
	}
	n := Name{}
	copy(n[:], data)
	return n, nil
}

func (n Name) String() string {
	return AsciiString(n[:])
}

type Symbol [12]byte

func SymbolFromString(symbol string) (Symbol, error) {
	data := []byte(symbol)
	if len(data) > 32 {
		return Symbol{}, fmt.Errorf("symbol data to long expected 12 got %d", len(data))
	}
	s := Symbol{}
	copy(s[:], data)
	return s, nil
}

func (s Symbol) String() string {
	return AsciiString(s[:])
}

func AsciiString(data []byte) string {
	var trimmed []byte
	for _, b := range data {
		if b > 0 {
			trimmed = append(trimmed, b)
		}
	}
	return string(trimmed)
}