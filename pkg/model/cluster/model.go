// Copyright 2022 The jackal Authors
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

package clustermodel

import (
	"fmt"

	"github.com/ortuman/jackal/pkg/version"
)

// Member represents a cluster instance address and port.
type Member struct {
	InstanceID string
	Host       string
	Port       int
	APIVer     *version.SemanticVersion
}

// String returns Member string representation.
func (m *Member) String() string {
	return fmt.Sprintf("%s:%d", m.Host, m.Port)
}
