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

package consensuspb

import (
	"fmt"

	"github.com/spica-matrix/spica-spc/util/byteutils"
)

// ToString return a string of consensus root
func (m *ConsensusRoot) ToString() string {
	return fmt.Sprintf(`{"proposer": %s, "timestamp": "%d", "dynasty": "%s"}`,
		byteutils.Hex(m.Proposer),
		m.Timestamp,
		byteutils.Hex(m.DynastyRoot),
	)
}
© 2020 GitHub, Inc.
