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

package hook

import (
	clustermodel "github.com/ortuman/jackal/pkg/model/cluster"
)

const (
	// MemberListUpdated hook runs whenever cluster member list is updated.
	MemberListUpdated = "memberlist.updated"
)

// MemberListInfo contains all info associated to MemberListUpdated event.
type MemberListInfo struct {
	// Registered contains all new registered cluster members.
	Registered []clustermodel.Member

	// UnregisteredKeys contains unregistered cluster members keys.
	UnregisteredKeys []string
}