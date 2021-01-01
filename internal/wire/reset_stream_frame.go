package wire

import (
	"bytes"

	"github.com/lucas-clemente/quic-go/internal/protocol"
	"github.com/lucas-clemente/quic-go/quicvarint"
)

// A ResetStreamFrame is a RESET_STREAM frame in QUIC
type ResetStreamFrame struct {
	StreamID  protocol.StreamID
	ErrorCode protocol.ApplicationErrorCode
	FinalSize protocol.ByteCount
}

func parseResetStreamFrame(r *bytes.Reader, _ protocol.VersionNumber) (*ResetStreamFrame, error) {
	if _, err := r.ReadByte(); err != nil { // read the TypeByte
		return nil, err
	}

	var streamID protocol.StreamID
	var byteOffset protocol.ByteCount
	sid, err := quicvarint.ReadVarInt(r)
	if err != nil {
		return nil, err
	}
	streamID = protocol.StreamID(sid)
	errorCode, err := quicvarint.ReadVarInt(r)
	if err != nil {
		return nil, err
	}
	bo, err := quicvarint.ReadVarInt(r)
	if err != nil {
		return nil, err
	}
	byteOffset = protocol.ByteCount(bo)

	return &ResetStreamFrame{
		StreamID:  streamID,
		ErrorCode: protocol.ApplicationErrorCode(errorCode),
		FinalSize: byteOffset,
	}, nil
}

func (f *ResetStreamFrame) Write(b *bytes.Buffer, _ protocol.VersionNumber) error {
	b.WriteByte(0x4)
	quicvarint.WriteVarInt(b, uint64(f.StreamID))
	quicvarint.WriteVarInt(b, uint64(f.ErrorCode))
	quicvarint.WriteVarInt(b, uint64(f.FinalSize))
	return nil
}

// Length of a written frame
func (f *ResetStreamFrame) Length(version protocol.VersionNumber) protocol.ByteCount {
	return 1 + quicvarint.VarIntLen(uint64(f.StreamID)) + quicvarint.VarIntLen(uint64(f.ErrorCode)) + quicvarint.VarIntLen(uint64(f.FinalSize))
}
