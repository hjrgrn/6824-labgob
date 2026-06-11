package labgob

import (
	"encoding/gob"
	"io"
)

//
// trying to send non-capitalized fields over RPC produces a range of
// misbehavior, including both mysterious incorrect computation and
// outright crashes. so this wrapper around Go's encoding/gob warns
// about non-capitalized field names.
//

type LabEncoder struct {
	gob *gob.Encoder
}

func NewEncoder(w io.Writer) *LabEncoder {
	enc := &LabEncoder{}
	enc.gob = gob.NewEncoder(w)
	return enc
}
