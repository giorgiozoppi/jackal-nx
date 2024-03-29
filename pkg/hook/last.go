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

import "github.com/jackal-xmpp/stravaganza/jid"

const (
	// LastActivityFetched hook runs when a user last activity is fetched.
	LastActivityFetched = "last.fetched"
)

// LastActivityInfo contains all information associated to a last activity event.
type LastActivityInfo struct {
	// Username is the name of the user associated to this event.
	Username string

	// JID represents the event associated JID.
	JID *jid.JID
}
