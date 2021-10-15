package optional

import (
	fmt "fmt"
	math "math"
	testing "testing"

	github_com_gogo_protobuf_proto "github.com/gogo/protobuf/proto"
	proto "github.com/gogo/protobuf/proto"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func TestOptionalProtoSerializes(t *testing.T) {
	param := true
	p := &OptMessage{
		MyParam: &param,
	}
	dAtA, err := github_com_gogo_protobuf_proto.Marshal(p)
	if err != nil {
		t.Fatalf("msg = %v, err = %v", p, err)
	}
	msg := &OptMessage{}
	if err := github_com_gogo_protobuf_proto.Unmarshal(dAtA, msg); err != nil {
		t.Fatalf("msg = %v, err = %v", p, err)
	}
	if p.MyParam == nil || !*p.MyParam {
		t.Fatalf("Failed to serialize myparam: %v from %v", msg, p)
	}
}
