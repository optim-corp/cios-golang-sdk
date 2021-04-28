package convert

import (
	"encoding/binary"
	"math"
	"time"
	"unsafe"

	. "github.com/optim-corp/cios-golang-sdk/util/go_advance_type/util"
)

func Time2Str(x time.Time) string {
	return time.Time.Format(x, "2006-01-02 15:04:05")
}
func Long2Bytes(i int64) []byte {
	var buf = make([]byte, 8)
	binary.BigEndian.PutUint64(buf, uint64(i))
	return buf
}
func Float2Bytes(f float32) []byte {
	b := make([]byte, 8)
	binary.LittleEndian.PutUint32(b, math.Float32bits(f))
	return b
}
func Double2Bytes(f float64) []byte {
	b := make([]byte, 8)
	binary.LittleEndian.PutUint64(b, math.Float64bits(f))
	return b
}
func Bytes2Long(buf []byte) int64 {
	return int64(binary.BigEndian.Uint64(buf))
}
func Bytes2Float(b []byte) float32 {
	return math.Float32frombits(binary.LittleEndian.Uint32(b))
}

func BytesAsULong(b []byte) uint64 {
	return binary.LittleEndian.Uint64(b)
}

func Int16AsUint16(i int16) uint16 {
	return Is(0 < i).
		Ok(uint16(i)).
		No(^uint16(-i) + 1).Value().AsUint16()
}
func Uint16AsInt16(ui uint16) int16 {
	return Is(ui > math.MaxInt16).
		Ok(int16(ui - math.MaxUint16 - 1)).
		No(int16(ui)).Value().AsInt16()

}
func Int32AsUint32(i int32) uint32 {
	return Is(0 < i).
		Ok(uint32(i)).
		No(^uint32(-i) + 1).Value().AsUint32()
}

func Uint32AsInt32(ui uint32) int32 {
	return Is(ui > math.MaxInt32).
		Ok(int32(ui - math.MaxUint32 - 1)).
		No(int32(ui)).Value().AsInt32()
}

func Int64AsUint64(i int64) uint64 {
	return Is(0 < i).
		Ok(uint64(i)).
		No(^uint64(-i) + 1).Value().AsUint64()
}

func Uint64AsInt64(ui uint64) int64 {
	return Is(ui > math.MaxInt64).
		Ok(int64(ui - math.MaxUint64 - 1)).
		No(int64(ui)).Value().AsInt64()
}
func Str2Bytes(s string) []byte {
	return *(*[]byte)(unsafe.Pointer(&s))
}
