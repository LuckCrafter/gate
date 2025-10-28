package proto

import (
	"io"

	"github.com/go-logr/logr"
	"github.com/luckcrafter/gate/pkg/gate/proto"
	"github.com/sandertv/gophertunnel/minecraft/protocol/packet"
)

type Encoder struct {
	dec       *packet.Encoder
	log       logr.Logger
	direction proto.Direction
}

func NewEncoder(w io.Writer, direction proto.Direction, log logr.Logger) *Encoder {
	return &Encoder{
		dec:       packet.NewEncoder(w),
		log:       log,
		direction: direction,
	}
}
