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

package memberlist

import (
	"context"

	clustermodel "github.com/ortuman/jackal/pkg/model/cluster"
)

// MemberList defines cluster memberlist interface.
type MemberList interface {
	// GetMember returns cluster member info associated to an identifier.
	GetMember(instanceID string) (m clustermodel.Member, ok bool)

	// GetMembers returns all cluster registered members.
	GetMembers() map[string]clustermodel.Member

	// Start initializes memberlist.
	Start(ctx context.Context) error

	// Stop releases all underlying memberlist resources.
	Stop(ctx context.Context) error
}
