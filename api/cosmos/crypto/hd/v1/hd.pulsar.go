// Code generated by protoc-gen-go-pulsar. DO NOT EDIT.
package hdv1

import (
	_ "cosmossdk.io/api/amino"
	fmt "fmt"
	runtime "github.com/cosmos/cosmos-proto/runtime"
	_ "github.com/cosmos/gogoproto/gogoproto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoiface "google.golang.org/protobuf/runtime/protoiface"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	io "io"
	reflect "reflect"
	sync "sync"
)

var (
	md_BIP44Params               protoreflect.MessageDescriptor
	fd_BIP44Params_purpose       protoreflect.FieldDescriptor
	fd_BIP44Params_coin_type     protoreflect.FieldDescriptor
	fd_BIP44Params_account       protoreflect.FieldDescriptor
	fd_BIP44Params_change        protoreflect.FieldDescriptor
	fd_BIP44Params_address_index protoreflect.FieldDescriptor
)

func init() {
	file_cosmos_crypto_hd_v1_hd_proto_init()
	md_BIP44Params = File_cosmos_crypto_hd_v1_hd_proto.Messages().ByName("BIP44Params")
	fd_BIP44Params_purpose = md_BIP44Params.Fields().ByName("purpose")
	fd_BIP44Params_coin_type = md_BIP44Params.Fields().ByName("coin_type")
	fd_BIP44Params_account = md_BIP44Params.Fields().ByName("account")
	fd_BIP44Params_change = md_BIP44Params.Fields().ByName("change")
	fd_BIP44Params_address_index = md_BIP44Params.Fields().ByName("address_index")
}

var _ protoreflect.Message = (*fastReflection_BIP44Params)(nil)

type fastReflection_BIP44Params BIP44Params

func (x *BIP44Params) ProtoReflect() protoreflect.Message {
	return (*fastReflection_BIP44Params)(x)
}

func (x *BIP44Params) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_crypto_hd_v1_hd_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_BIP44Params_messageType fastReflection_BIP44Params_messageType
var _ protoreflect.MessageType = fastReflection_BIP44Params_messageType{}

type fastReflection_BIP44Params_messageType struct{}

func (x fastReflection_BIP44Params_messageType) Zero() protoreflect.Message {
	return (*fastReflection_BIP44Params)(nil)
}
func (x fastReflection_BIP44Params_messageType) New() protoreflect.Message {
	return new(fastReflection_BIP44Params)
}
func (x fastReflection_BIP44Params_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_BIP44Params
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_BIP44Params) Descriptor() protoreflect.MessageDescriptor {
	return md_BIP44Params
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_BIP44Params) Type() protoreflect.MessageType {
	return _fastReflection_BIP44Params_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_BIP44Params) New() protoreflect.Message {
	return new(fastReflection_BIP44Params)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_BIP44Params) Interface() protoreflect.ProtoMessage {
	return (*BIP44Params)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_BIP44Params) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.Purpose != uint32(0) {
		value := protoreflect.ValueOfUint32(x.Purpose)
		if !f(fd_BIP44Params_purpose, value) {
			return
		}
	}
	if x.CoinType != uint32(0) {
		value := protoreflect.ValueOfUint32(x.CoinType)
		if !f(fd_BIP44Params_coin_type, value) {
			return
		}
	}
	if x.Account != uint32(0) {
		value := protoreflect.ValueOfUint32(x.Account)
		if !f(fd_BIP44Params_account, value) {
			return
		}
	}
	if x.Change != false {
		value := protoreflect.ValueOfBool(x.Change)
		if !f(fd_BIP44Params_change, value) {
			return
		}
	}
	if x.AddressIndex != uint32(0) {
		value := protoreflect.ValueOfUint32(x.AddressIndex)
		if !f(fd_BIP44Params_address_index, value) {
			return
		}
	}
}

// Has reports whether a field is populated.
//
// Some fields have the property of nullability where it is possible to
// distinguish between the default value of a field and whether the field
// was explicitly populated with the default value. Singular message fields,
// member fields of a oneof, and proto2 scalar fields are nullable. Such
// fields are populated only if explicitly set.
//
// In other cases (aside from the nullable cases above),
// a proto3 scalar field is populated if it contains a non-zero value, and
// a repeated field is populated if it is non-empty.
func (x *fastReflection_BIP44Params) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "cosmos.crypto.hd.v1.BIP44Params.purpose":
		return x.Purpose != uint32(0)
	case "cosmos.crypto.hd.v1.BIP44Params.coin_type":
		return x.CoinType != uint32(0)
	case "cosmos.crypto.hd.v1.BIP44Params.account":
		return x.Account != uint32(0)
	case "cosmos.crypto.hd.v1.BIP44Params.change":
		return x.Change != false
	case "cosmos.crypto.hd.v1.BIP44Params.address_index":
		return x.AddressIndex != uint32(0)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.crypto.hd.v1.BIP44Params"))
		}
		panic(fmt.Errorf("message cosmos.crypto.hd.v1.BIP44Params does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_BIP44Params) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "cosmos.crypto.hd.v1.BIP44Params.purpose":
		x.Purpose = uint32(0)
	case "cosmos.crypto.hd.v1.BIP44Params.coin_type":
		x.CoinType = uint32(0)
	case "cosmos.crypto.hd.v1.BIP44Params.account":
		x.Account = uint32(0)
	case "cosmos.crypto.hd.v1.BIP44Params.change":
		x.Change = false
	case "cosmos.crypto.hd.v1.BIP44Params.address_index":
		x.AddressIndex = uint32(0)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.crypto.hd.v1.BIP44Params"))
		}
		panic(fmt.Errorf("message cosmos.crypto.hd.v1.BIP44Params does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_BIP44Params) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "cosmos.crypto.hd.v1.BIP44Params.purpose":
		value := x.Purpose
		return protoreflect.ValueOfUint32(value)
	case "cosmos.crypto.hd.v1.BIP44Params.coin_type":
		value := x.CoinType
		return protoreflect.ValueOfUint32(value)
	case "cosmos.crypto.hd.v1.BIP44Params.account":
		value := x.Account
		return protoreflect.ValueOfUint32(value)
	case "cosmos.crypto.hd.v1.BIP44Params.change":
		value := x.Change
		return protoreflect.ValueOfBool(value)
	case "cosmos.crypto.hd.v1.BIP44Params.address_index":
		value := x.AddressIndex
		return protoreflect.ValueOfUint32(value)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.crypto.hd.v1.BIP44Params"))
		}
		panic(fmt.Errorf("message cosmos.crypto.hd.v1.BIP44Params does not contain field %s", descriptor.FullName()))
	}
}

// Set stores the value for a field.
//
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType.
// When setting a composite type, it is unspecified whether the stored value
// aliases the source's memory in any way. If the composite value is an
// empty, read-only value, then it panics.
//
// Set is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_BIP44Params) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "cosmos.crypto.hd.v1.BIP44Params.purpose":
		x.Purpose = uint32(value.Uint())
	case "cosmos.crypto.hd.v1.BIP44Params.coin_type":
		x.CoinType = uint32(value.Uint())
	case "cosmos.crypto.hd.v1.BIP44Params.account":
		x.Account = uint32(value.Uint())
	case "cosmos.crypto.hd.v1.BIP44Params.change":
		x.Change = value.Bool()
	case "cosmos.crypto.hd.v1.BIP44Params.address_index":
		x.AddressIndex = uint32(value.Uint())
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.crypto.hd.v1.BIP44Params"))
		}
		panic(fmt.Errorf("message cosmos.crypto.hd.v1.BIP44Params does not contain field %s", fd.FullName()))
	}
}

// Mutable returns a mutable reference to a composite type.
//
// If the field is unpopulated, it may allocate a composite value.
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType
// if not already stored.
// It panics if the field does not contain a composite type.
//
// Mutable is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_BIP44Params) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.crypto.hd.v1.BIP44Params.purpose":
		panic(fmt.Errorf("field purpose of message cosmos.crypto.hd.v1.BIP44Params is not mutable"))
	case "cosmos.crypto.hd.v1.BIP44Params.coin_type":
		panic(fmt.Errorf("field coin_type of message cosmos.crypto.hd.v1.BIP44Params is not mutable"))
	case "cosmos.crypto.hd.v1.BIP44Params.account":
		panic(fmt.Errorf("field account of message cosmos.crypto.hd.v1.BIP44Params is not mutable"))
	case "cosmos.crypto.hd.v1.BIP44Params.change":
		panic(fmt.Errorf("field change of message cosmos.crypto.hd.v1.BIP44Params is not mutable"))
	case "cosmos.crypto.hd.v1.BIP44Params.address_index":
		panic(fmt.Errorf("field address_index of message cosmos.crypto.hd.v1.BIP44Params is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.crypto.hd.v1.BIP44Params"))
		}
		panic(fmt.Errorf("message cosmos.crypto.hd.v1.BIP44Params does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_BIP44Params) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.crypto.hd.v1.BIP44Params.purpose":
		return protoreflect.ValueOfUint32(uint32(0))
	case "cosmos.crypto.hd.v1.BIP44Params.coin_type":
		return protoreflect.ValueOfUint32(uint32(0))
	case "cosmos.crypto.hd.v1.BIP44Params.account":
		return protoreflect.ValueOfUint32(uint32(0))
	case "cosmos.crypto.hd.v1.BIP44Params.change":
		return protoreflect.ValueOfBool(false)
	case "cosmos.crypto.hd.v1.BIP44Params.address_index":
		return protoreflect.ValueOfUint32(uint32(0))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.crypto.hd.v1.BIP44Params"))
		}
		panic(fmt.Errorf("message cosmos.crypto.hd.v1.BIP44Params does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_BIP44Params) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos.crypto.hd.v1.BIP44Params", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_BIP44Params) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_BIP44Params) SetUnknown(fields protoreflect.RawFields) {
	x.unknownFields = fields
}

// IsValid reports whether the message is valid.
//
// An invalid message is an empty, read-only value.
//
// An invalid message often corresponds to a nil pointer of the concrete
// message type, but the details are implementation dependent.
// Validity is not part of the protobuf data model, and may not
// be preserved in marshaling or other operations.
func (x *fastReflection_BIP44Params) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_BIP44Params) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*BIP44Params)
		if x == nil {
			return protoiface.SizeOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Size:              0,
			}
		}
		options := runtime.SizeInputToOptions(input)
		_ = options
		var n int
		var l int
		_ = l
		if x.Purpose != 0 {
			n += 1 + runtime.Sov(uint64(x.Purpose))
		}
		if x.CoinType != 0 {
			n += 1 + runtime.Sov(uint64(x.CoinType))
		}
		if x.Account != 0 {
			n += 1 + runtime.Sov(uint64(x.Account))
		}
		if x.Change {
			n += 2
		}
		if x.AddressIndex != 0 {
			n += 1 + runtime.Sov(uint64(x.AddressIndex))
		}
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*BIP44Params)
		if x == nil {
			return protoiface.MarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Buf:               input.Buf,
			}, nil
		}
		options := runtime.MarshalInputToOptions(input)
		_ = options
		size := options.Size(x)
		dAtA := make([]byte, size)
		i := len(dAtA)
		_ = i
		var l int
		_ = l
		if x.unknownFields != nil {
			i -= len(x.unknownFields)
			copy(dAtA[i:], x.unknownFields)
		}
		if x.AddressIndex != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.AddressIndex))
			i--
			dAtA[i] = 0x28
		}
		if x.Change {
			i--
			if x.Change {
				dAtA[i] = 1
			} else {
				dAtA[i] = 0
			}
			i--
			dAtA[i] = 0x20
		}
		if x.Account != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.Account))
			i--
			dAtA[i] = 0x18
		}
		if x.CoinType != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.CoinType))
			i--
			dAtA[i] = 0x10
		}
		if x.Purpose != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.Purpose))
			i--
			dAtA[i] = 0x8
		}
		if input.Buf != nil {
			input.Buf = append(input.Buf, dAtA...)
		} else {
			input.Buf = dAtA
		}
		return protoiface.MarshalOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Buf:               input.Buf,
		}, nil
	}
	unmarshal := func(input protoiface.UnmarshalInput) (protoiface.UnmarshalOutput, error) {
		x := input.Message.Interface().(*BIP44Params)
		if x == nil {
			return protoiface.UnmarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Flags:             input.Flags,
			}, nil
		}
		options := runtime.UnmarshalInputToOptions(input)
		_ = options
		dAtA := input.Buf
		l := len(dAtA)
		iNdEx := 0
		for iNdEx < l {
			preIndex := iNdEx
			var wire uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
				}
				if iNdEx >= l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				wire |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			fieldNum := int32(wire >> 3)
			wireType := int(wire & 0x7)
			if wireType == 4 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: BIP44Params: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: BIP44Params: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Purpose", wireType)
				}
				x.Purpose = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.Purpose |= uint32(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			case 2:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field CoinType", wireType)
				}
				x.CoinType = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.CoinType |= uint32(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			case 3:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Account", wireType)
				}
				x.Account = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.Account |= uint32(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			case 4:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Change", wireType)
				}
				var v int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				x.Change = bool(v != 0)
			case 5:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field AddressIndex", wireType)
				}
				x.AddressIndex = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.AddressIndex |= uint32(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			default:
				iNdEx = preIndex
				skippy, err := runtime.Skip(dAtA[iNdEx:])
				if err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				if (skippy < 0) || (iNdEx+skippy) < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if (iNdEx + skippy) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if !options.DiscardUnknown {
					x.unknownFields = append(x.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
				}
				iNdEx += skippy
			}
		}

		if iNdEx > l {
			return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
		}
		return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, nil
	}
	return &protoiface.Methods{
		NoUnkeyedLiterals: struct{}{},
		Flags:             protoiface.SupportMarshalDeterministic | protoiface.SupportUnmarshalDiscardUnknown,
		Size:              size,
		Marshal:           marshal,
		Unmarshal:         unmarshal,
		Merge:             nil,
		CheckInitialized:  nil,
	}
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.0
// 	protoc        (unknown)
// source: cosmos/crypto/hd/v1/hd.proto

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// BIP44Params is used as path field in ledger item in Record.
type BIP44Params struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// purpose is a constant set to 44' (or 0x8000002C) following the BIP43 recommendation
	Purpose uint32 `protobuf:"varint,1,opt,name=purpose,proto3" json:"purpose,omitempty"`
	// coin_type is a constant that improves privacy
	CoinType uint32 `protobuf:"varint,2,opt,name=coin_type,json=coinType,proto3" json:"coin_type,omitempty"`
	// account splits the key space into independent user identities
	Account uint32 `protobuf:"varint,3,opt,name=account,proto3" json:"account,omitempty"`
	// change is a constant used for public derivation. Constant 0 is used for external chain and constant 1 for internal
	// chain.
	Change bool `protobuf:"varint,4,opt,name=change,proto3" json:"change,omitempty"`
	// address_index is used as child index in BIP32 derivation
	AddressIndex uint32 `protobuf:"varint,5,opt,name=address_index,json=addressIndex,proto3" json:"address_index,omitempty"`
}

func (x *BIP44Params) Reset() {
	*x = BIP44Params{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_crypto_hd_v1_hd_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BIP44Params) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BIP44Params) ProtoMessage() {}

// Deprecated: Use BIP44Params.ProtoReflect.Descriptor instead.
func (*BIP44Params) Descriptor() ([]byte, []int) {
	return file_cosmos_crypto_hd_v1_hd_proto_rawDescGZIP(), []int{0}
}

func (x *BIP44Params) GetPurpose() uint32 {
	if x != nil {
		return x.Purpose
	}
	return 0
}

func (x *BIP44Params) GetCoinType() uint32 {
	if x != nil {
		return x.CoinType
	}
	return 0
}

func (x *BIP44Params) GetAccount() uint32 {
	if x != nil {
		return x.Account
	}
	return 0
}

func (x *BIP44Params) GetChange() bool {
	if x != nil {
		return x.Change
	}
	return false
}

func (x *BIP44Params) GetAddressIndex() uint32 {
	if x != nil {
		return x.AddressIndex
	}
	return 0
}

var File_cosmos_crypto_hd_v1_hd_proto protoreflect.FileDescriptor

var file_cosmos_crypto_hd_v1_hd_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x2f,
	0x68, 0x64, 0x2f, 0x76, 0x31, 0x2f, 0x68, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13,
	0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x2e, 0x68, 0x64,
	0x2e, 0x76, 0x31, 0x1a, 0x11, 0x61, 0x6d, 0x69, 0x6e, 0x6f, 0x2f, 0x61, 0x6d, 0x69, 0x6e, 0x6f,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x14, 0x67, 0x6f, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x67, 0x6f, 0x67, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc0, 0x01, 0x0a,
	0x0b, 0x42, 0x49, 0x50, 0x34, 0x34, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x18, 0x0a, 0x07,
	0x70, 0x75, 0x72, 0x70, 0x6f, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x70,
	0x75, 0x72, 0x70, 0x6f, 0x73, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6f, 0x69, 0x6e, 0x5f, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x63, 0x6f, 0x69, 0x6e, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x16, 0x0a,
	0x06, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x63,
	0x68, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c, 0x61, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x3a, 0x23, 0x98, 0xa0, 0x1f, 0x00,
	0x8a, 0xe7, 0xb0, 0x2a, 0x1a, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x2f, 0x6b, 0x65, 0x79, 0x73,
	0x2f, 0x68, 0x64, 0x2f, 0x42, 0x49, 0x50, 0x34, 0x34, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x42,
	0xc0, 0x01, 0xc8, 0xe1, 0x1e, 0x00, 0x0a, 0x17, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x73, 0x6d,
	0x6f, 0x73, 0x2e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x2e, 0x68, 0x64, 0x2e, 0x76, 0x31, 0x42,
	0x07, 0x48, 0x64, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x29, 0x63, 0x6f, 0x73, 0x6d,
	0x6f, 0x73, 0x73, 0x64, 0x6b, 0x2e, 0x69, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x73,
	0x6d, 0x6f, 0x73, 0x2f, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x2f, 0x68, 0x64, 0x2f, 0x76, 0x31,
	0x3b, 0x68, 0x64, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x43, 0x43, 0x48, 0xaa, 0x02, 0x13, 0x43, 0x6f,
	0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x43, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x2e, 0x48, 0x64, 0x2e, 0x56,
	0x31, 0xca, 0x02, 0x13, 0x43, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x5c, 0x43, 0x72, 0x79, 0x70, 0x74,
	0x6f, 0x5c, 0x48, 0x64, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x1f, 0x43, 0x6f, 0x73, 0x6d, 0x6f, 0x73,
	0x5c, 0x43, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x5c, 0x48, 0x64, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50,
	0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x16, 0x43, 0x6f, 0x73, 0x6d,
	0x6f, 0x73, 0x3a, 0x3a, 0x43, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x3a, 0x3a, 0x48, 0x64, 0x3a, 0x3a,
	0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cosmos_crypto_hd_v1_hd_proto_rawDescOnce sync.Once
	file_cosmos_crypto_hd_v1_hd_proto_rawDescData = file_cosmos_crypto_hd_v1_hd_proto_rawDesc
)

func file_cosmos_crypto_hd_v1_hd_proto_rawDescGZIP() []byte {
	file_cosmos_crypto_hd_v1_hd_proto_rawDescOnce.Do(func() {
		file_cosmos_crypto_hd_v1_hd_proto_rawDescData = protoimpl.X.CompressGZIP(file_cosmos_crypto_hd_v1_hd_proto_rawDescData)
	})
	return file_cosmos_crypto_hd_v1_hd_proto_rawDescData
}

var file_cosmos_crypto_hd_v1_hd_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_cosmos_crypto_hd_v1_hd_proto_goTypes = []interface{}{
	(*BIP44Params)(nil), // 0: cosmos.crypto.hd.v1.BIP44Params
}
var file_cosmos_crypto_hd_v1_hd_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_cosmos_crypto_hd_v1_hd_proto_init() }
func file_cosmos_crypto_hd_v1_hd_proto_init() {
	if File_cosmos_crypto_hd_v1_hd_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cosmos_crypto_hd_v1_hd_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BIP44Params); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_cosmos_crypto_hd_v1_hd_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cosmos_crypto_hd_v1_hd_proto_goTypes,
		DependencyIndexes: file_cosmos_crypto_hd_v1_hd_proto_depIdxs,
		MessageInfos:      file_cosmos_crypto_hd_v1_hd_proto_msgTypes,
	}.Build()
	File_cosmos_crypto_hd_v1_hd_proto = out.File
	file_cosmos_crypto_hd_v1_hd_proto_rawDesc = nil
	file_cosmos_crypto_hd_v1_hd_proto_goTypes = nil
	file_cosmos_crypto_hd_v1_hd_proto_depIdxs = nil
}
