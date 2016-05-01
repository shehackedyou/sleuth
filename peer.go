// Copyright 2016 Afshin Darian. All rights reserved.
// Use of this source code is governed by The MIT License
// that can be found in the LICENSE file.

package sleuth

type Peer struct {
	// Name is the short public name attached to all events/peers.
	Name string
	// Node is the full peer node name used for whispering.
	Node string
	// Service is the name of the service being offered by a peer.
	Service string
	// Version is the optional service version running on a peer.
	Version string
}
