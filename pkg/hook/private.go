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

import "github.com/jackal-xmpp/stravaganza"

const (
	// PrivateFetched hook runs when a user private XML is fetched.
	PrivateFetched = "private.fetched"

	// PrivateUpdated hook runs when a user private XML is updated.
	PrivateUpdated = "private.updated"
)

// PrivateInfo contains all information associated to a private event.
type PrivateInfo struct {
	// Username is the name of the user associated to this event.
	Username string

	// Private is the private XML element associated to this event.
	Private stravaganza.Element
}
