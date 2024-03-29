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

const (
	// RosterRequested hook runs whenever a user requests the roster.
	RosterRequested = "roster.requested"

	// RosterItemUpdated hook runs whenever a roster item subscription is updated.
	RosterItemUpdated = "roster.item.updated"
)

// RosterInfo contains all information associated to a roster event.
type RosterInfo struct {
	// Username is the name of the roster owner.
	Username string

	// JID is the event contact JID.
	JID string

	// Subscription is the roster event subscription value.
	Subscription string
}
