package mongoproto
import (
	"github.com/10gen/llmgo"
	"io"
)

// OpMsg sends a diagnostic message to the database. The database sends back a fixed response.
// OpMsg is Deprecated
// http://docs.mongodb.org/meta-driver/latest/legacy/mongodb-wire-protocol/#op-msg
type OpMsg struct {
	Header  MsgHeader
	Message string
}

func(op *OpMsg) OpCode() OpCode {
	return OpCodeMessage
}

func(op *OpMsg) FromReader(r io.Reader) error {
	return nil
}

func(op *OpMsg) Execute(session *mgo.Session) (*mgo.ReplyOp, error) {
	return nil, nil
}