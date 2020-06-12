
// Copyright (C) 2017 spica-spc authors
//
// This file is part of the spica-spc library.
//
// the spica-spc library is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// the spica-spc library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with the spica-spc library.  If not, see <http://www.gnu.org/licenses/>.
//

package core

import (
	"github.com/spica-matrix/spica-spc/crypto"
	"github.com/spica-matrix/spica-spc/crypto/keystore"
)

// RecoverSignerFromSignature return address who signs the signature
func RecoverSignerFromSignature(alg keystore.Algorithm, plainText []byte, cipherText []byte) (*Address, error) {
	signature, err := crypto.NewSignature(alg)
	if err != nil {
		return nil, err
	}
	pub, err := signature.RecoverPublic(plainText, cipherText)
	if err != nil {
		return nil, err
	}
	pubdata, err := pub.Encoded()
	if err != nil {
		return nil, err
	}
	addr, err := NewAddressFromPublicKey(pubdata)
	if err != nil {
		return nil, err
	}
	return addr, nil
}
