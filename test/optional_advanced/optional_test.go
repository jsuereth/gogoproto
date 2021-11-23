package optional

import (
	fmt "fmt"
	math "math"
	testing "testing"

	proto "github.com/gogo/protobuf/proto"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func TestOptionalProtoSerializes(t *testing.T) {
	param := true
	intp := int32(1)
	longp := int64(2000)
	uintp := uint32(1)
	ulongp := uint64(211300214)
	dp := 1.0
	sp := "Hello"
	p := &OptMessage{
		MyParam:     &param,
		VintParam:   &intp,
		VlongParam:  &longp,
		IntParam:    &uintp,
		LongParam:   &ulongp,
		DoubleParam: &dp,
		StringParam: &sp,
	}
	dAtA, err := p.Marshal()
	if err != nil {
		t.Fatalf("msg = %v, err = %v", p, err)
	}
	msg := &OptMessage{}
	if err := msg.Unmarshal(dAtA); err != nil {
		t.Fatalf("msg = %v, err = %v", p, err)
	}
	if p.MyParam == nil || !*p.MyParam {
		t.Fatalf("Failed to serialize myparam: %v from %v", msg, p)
	}
}
