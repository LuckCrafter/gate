package methods

import (
	"github.com/luckcrafter/gate/pkg/edition/java/proto/state"
	"github.com/luckcrafter/gate/pkg/edition/java/proto/state/states"
	"github.com/luckcrafter/gate/pkg/edition/java/proto/version"
	"github.com/luckcrafter/gate/pkg/gate/proto"
)

// Protocol returns the protocol version of the given subject if provided.
func Protocol(subject any) (proto.Protocol, bool) {
	// this method is implemented by proxy player
	if p, ok := subject.(interface{ Protocol() proto.Protocol }); ok {
		return p.Protocol(), true
	}
	return version.Unknown.Protocol, false
}

// State returns the state of the given subject if provided.
func State(subject any) (states.State, bool) {
	if s, ok := subject.(interface{ State() *state.Registry }); ok {
		return s.State().State, true
	}
	return -1, false
}
