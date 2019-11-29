/*
 * Copyright 2018 The CovenantSQL Authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package kms

import (
	"strings"

	"github.com/SQLess/SQLess/proto"
)

// AnonymousXXX is only used for DHT.Ping

var (
	// AnonymousNodeID is the anonymous node id
	AnonymousNodeID = proto.NodeID(strings.Repeat("f", 64))
	// AnonymousRawNodeID is the anonymous node id
	AnonymousRawNodeID = AnonymousNodeID.ToRawNodeID()
)