// Code generated by protoc-gen-go. DO NOT EDIT.
// source: discover.proto

package v0

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Discover struct {
	// Types that are valid to be assigned to Info:
	//	*Discover_Sensor
	Info isDiscover_Info `protobuf_oneof:"info"`
}

func (m *Discover) Reset()                    { *m = Discover{} }
func (m *Discover) String() string            { return proto.CompactTextString(m) }
func (*Discover) ProtoMessage()               {}
func (*Discover) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{0} }

type isDiscover_Info interface {
	isDiscover_Info()
}

type Discover_Sensor struct {
	Sensor *Sensor `protobuf:"bytes,3,opt,name=sensor,oneof"`
}

func (*Discover_Sensor) isDiscover_Info() {}

func (m *Discover) GetInfo() isDiscover_Info {
	if m != nil {
		return m.Info
	}
	return nil
}

func (m *Discover) GetSensor() *Sensor {
	if x, ok := m.GetInfo().(*Discover_Sensor); ok {
		return x.Sensor
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Discover) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Discover_OneofMarshaler, _Discover_OneofUnmarshaler, _Discover_OneofSizer, []interface{}{
		(*Discover_Sensor)(nil),
	}
}

func _Discover_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Discover)
	// info
	switch x := m.Info.(type) {
	case *Discover_Sensor:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Sensor); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("Discover.Info has unexpected type %T", x)
	}
	return nil
}

func _Discover_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Discover)
	switch tag {
	case 3: // info.sensor
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Sensor)
		err := b.DecodeMessage(msg)
		m.Info = &Discover_Sensor{msg}
		return true, err
	default:
		return false, nil
	}
}

func _Discover_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Discover)
	// info
	switch x := m.Info.(type) {
	case *Discover_Sensor:
		s := proto.Size(x.Sensor)
		n += proto.SizeVarint(3<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

func init() {
	proto.RegisterType((*Discover)(nil), "capsule8.v0.Discover")
}

func init() { proto.RegisterFile("discover.proto", fileDescriptor4) }

var fileDescriptor4 = []byte{
	// 114 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4b, 0xc9, 0x2c, 0x4e,
	0xce, 0x2f, 0x4b, 0x2d, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x4e, 0x4e, 0x2c, 0x28,
	0x2e, 0xcd, 0x49, 0xb5, 0xd0, 0x2b, 0x33, 0x90, 0xe2, 0x29, 0x4e, 0xcd, 0x2b, 0xce, 0x87, 0x4a,
	0x29, 0x39, 0x72, 0x71, 0xb8, 0x40, 0x15, 0x0b, 0xe9, 0x72, 0xb1, 0x41, 0xe4, 0x24, 0x98, 0x15,
	0x18, 0x35, 0xb8, 0x8d, 0x84, 0xf5, 0x90, 0xf4, 0xe9, 0x05, 0x83, 0xa5, 0x3c, 0x18, 0x82, 0xa0,
	0x8a, 0x9c, 0xd8, 0xb8, 0x58, 0x32, 0xf3, 0xd2, 0xf2, 0x9d, 0x58, 0xa2, 0x98, 0xca, 0x0c, 0x92,
	0xd8, 0xc0, 0xe6, 0x19, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x91, 0x28, 0xaa, 0x48, 0x7c, 0x00,
	0x00, 0x00,
}