// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/TheThingsNetwork/ttn/api/user.proto

package ttnpb

import proto "github.com/gogo/protobuf/proto"
import golang_proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"
import _ "github.com/gogo/protobuf/types"
import _ "github.com/gogo/protobuf/types"

import time "time"

import types "github.com/gogo/protobuf/types"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = golang_proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// User is the message that defines an user on the network.
type User struct {
	// User identifiers.
	UserIdentifiers `protobuf:"bytes,1,opt,name=ids,embedded=ids" json:"ids"`
	// password is the user's password.
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	// name is the user's full name.
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	// validated_at denotes the time the email was validated at.
	// This is a read-only field.
	ValidatedAt *time.Time `protobuf:"bytes,4,opt,name=validated_at,json=validatedAt,stdtime" json:"validated_at,omitempty"`
	// admin denotes whether or not the user has administrative rights within the tenancy.
	// This field can only be modified by admins.
	Admin bool `protobuf:"varint,5,opt,name=admin,proto3" json:"admin,omitempty"`
	// state denotes the reviewing state of the user by admin.
	// This field can only be modified by admins.
	State ReviewingState `protobuf:"varint,6,opt,name=state,proto3,enum=ttn.v3.ReviewingState" json:"state,omitempty"`
	// created_at denotes when the user was created.
	// This is a read-only field.
	CreatedAt time.Time `protobuf:"bytes,7,opt,name=created_at,json=createdAt,stdtime" json:"created_at"`
	// updated_at is the last time the user was updated.
	// This is a read-only field.
	UpdatedAt time.Time `protobuf:"bytes,8,opt,name=updated_at,json=updatedAt,stdtime" json:"updated_at"`
	// The last time the password was updated (read-only field).
	PasswordUpdatedAt time.Time `protobuf:"bytes,9,opt,name=password_updated_at,json=passwordUpdatedAt,stdtime" json:"password_updated_at"`
	// Require the user to update its password (modifiable only by admins).
	RequirePasswordUpdate bool `protobuf:"varint,10,opt,name=require_password_update,json=requirePasswordUpdate,proto3" json:"require_password_update,omitempty"`
}

func (m *User) Reset()                    { *m = User{} }
func (m *User) String() string            { return proto.CompactTextString(m) }
func (*User) ProtoMessage()               {}
func (*User) Descriptor() ([]byte, []int) { return fileDescriptorUser, []int{0} }

func (m *User) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *User) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *User) GetValidatedAt() *time.Time {
	if m != nil {
		return m.ValidatedAt
	}
	return nil
}

func (m *User) GetAdmin() bool {
	if m != nil {
		return m.Admin
	}
	return false
}

func (m *User) GetState() ReviewingState {
	if m != nil {
		return m.State
	}
	return STATE_PENDING
}

func (m *User) GetCreatedAt() time.Time {
	if m != nil {
		return m.CreatedAt
	}
	return time.Time{}
}

func (m *User) GetUpdatedAt() time.Time {
	if m != nil {
		return m.UpdatedAt
	}
	return time.Time{}
}

func (m *User) GetPasswordUpdatedAt() time.Time {
	if m != nil {
		return m.PasswordUpdatedAt
	}
	return time.Time{}
}

func (m *User) GetRequirePasswordUpdate() bool {
	if m != nil {
		return m.RequirePasswordUpdate
	}
	return false
}

func init() {
	proto.RegisterType((*User)(nil), "ttn.v3.User")
	golang_proto.RegisterType((*User)(nil), "ttn.v3.User")
}
func (this *User) VerboseEqual(that interface{}) error {
	if that == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that == nil && this != nil")
	}

	that1, ok := that.(*User)
	if !ok {
		that2, ok := that.(User)
		if ok {
			that1 = &that2
		} else {
			return fmt.Errorf("that is not of type *User")
		}
	}
	if that1 == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that is type *User but is nil && this != nil")
	} else if this == nil {
		return fmt.Errorf("that is type *User but is not nil && this == nil")
	}
	if !this.UserIdentifiers.Equal(&that1.UserIdentifiers) {
		return fmt.Errorf("UserIdentifiers this(%v) Not Equal that(%v)", this.UserIdentifiers, that1.UserIdentifiers)
	}
	if this.Password != that1.Password {
		return fmt.Errorf("Password this(%v) Not Equal that(%v)", this.Password, that1.Password)
	}
	if this.Name != that1.Name {
		return fmt.Errorf("Name this(%v) Not Equal that(%v)", this.Name, that1.Name)
	}
	if that1.ValidatedAt == nil {
		if this.ValidatedAt != nil {
			return fmt.Errorf("this.ValidatedAt != nil && that1.ValidatedAt == nil")
		}
	} else if !this.ValidatedAt.Equal(*that1.ValidatedAt) {
		return fmt.Errorf("ValidatedAt this(%v) Not Equal that(%v)", this.ValidatedAt, that1.ValidatedAt)
	}
	if this.Admin != that1.Admin {
		return fmt.Errorf("Admin this(%v) Not Equal that(%v)", this.Admin, that1.Admin)
	}
	if this.State != that1.State {
		return fmt.Errorf("State this(%v) Not Equal that(%v)", this.State, that1.State)
	}
	if !this.CreatedAt.Equal(that1.CreatedAt) {
		return fmt.Errorf("CreatedAt this(%v) Not Equal that(%v)", this.CreatedAt, that1.CreatedAt)
	}
	if !this.UpdatedAt.Equal(that1.UpdatedAt) {
		return fmt.Errorf("UpdatedAt this(%v) Not Equal that(%v)", this.UpdatedAt, that1.UpdatedAt)
	}
	if !this.PasswordUpdatedAt.Equal(that1.PasswordUpdatedAt) {
		return fmt.Errorf("PasswordUpdatedAt this(%v) Not Equal that(%v)", this.PasswordUpdatedAt, that1.PasswordUpdatedAt)
	}
	if this.RequirePasswordUpdate != that1.RequirePasswordUpdate {
		return fmt.Errorf("RequirePasswordUpdate this(%v) Not Equal that(%v)", this.RequirePasswordUpdate, that1.RequirePasswordUpdate)
	}
	return nil
}
func (this *User) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*User)
	if !ok {
		that2, ok := that.(User)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.UserIdentifiers.Equal(&that1.UserIdentifiers) {
		return false
	}
	if this.Password != that1.Password {
		return false
	}
	if this.Name != that1.Name {
		return false
	}
	if that1.ValidatedAt == nil {
		if this.ValidatedAt != nil {
			return false
		}
	} else if !this.ValidatedAt.Equal(*that1.ValidatedAt) {
		return false
	}
	if this.Admin != that1.Admin {
		return false
	}
	if this.State != that1.State {
		return false
	}
	if !this.CreatedAt.Equal(that1.CreatedAt) {
		return false
	}
	if !this.UpdatedAt.Equal(that1.UpdatedAt) {
		return false
	}
	if !this.PasswordUpdatedAt.Equal(that1.PasswordUpdatedAt) {
		return false
	}
	if this.RequirePasswordUpdate != that1.RequirePasswordUpdate {
		return false
	}
	return true
}
func (m *User) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *User) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0xa
	i++
	i = encodeVarintUser(dAtA, i, uint64(m.UserIdentifiers.Size()))
	n1, err := m.UserIdentifiers.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n1
	if len(m.Password) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintUser(dAtA, i, uint64(len(m.Password)))
		i += copy(dAtA[i:], m.Password)
	}
	if len(m.Name) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintUser(dAtA, i, uint64(len(m.Name)))
		i += copy(dAtA[i:], m.Name)
	}
	if m.ValidatedAt != nil {
		dAtA[i] = 0x22
		i++
		i = encodeVarintUser(dAtA, i, uint64(types.SizeOfStdTime(*m.ValidatedAt)))
		n2, err := types.StdTimeMarshalTo(*m.ValidatedAt, dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	if m.Admin {
		dAtA[i] = 0x28
		i++
		if m.Admin {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if m.State != 0 {
		dAtA[i] = 0x30
		i++
		i = encodeVarintUser(dAtA, i, uint64(m.State))
	}
	dAtA[i] = 0x3a
	i++
	i = encodeVarintUser(dAtA, i, uint64(types.SizeOfStdTime(m.CreatedAt)))
	n3, err := types.StdTimeMarshalTo(m.CreatedAt, dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n3
	dAtA[i] = 0x42
	i++
	i = encodeVarintUser(dAtA, i, uint64(types.SizeOfStdTime(m.UpdatedAt)))
	n4, err := types.StdTimeMarshalTo(m.UpdatedAt, dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n4
	dAtA[i] = 0x4a
	i++
	i = encodeVarintUser(dAtA, i, uint64(types.SizeOfStdTime(m.PasswordUpdatedAt)))
	n5, err := types.StdTimeMarshalTo(m.PasswordUpdatedAt, dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n5
	if m.RequirePasswordUpdate {
		dAtA[i] = 0x50
		i++
		if m.RequirePasswordUpdate {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	return i, nil
}

func encodeVarintUser(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func NewPopulatedUser(r randyUser, easy bool) *User {
	this := &User{}
	v1 := NewPopulatedUserIdentifiers(r, easy)
	this.UserIdentifiers = *v1
	this.Password = randStringUser(r)
	this.Name = randStringUser(r)
	if r.Intn(10) != 0 {
		this.ValidatedAt = types.NewPopulatedStdTime(r, easy)
	}
	this.Admin = bool(r.Intn(2) == 0)
	this.State = ReviewingState([]int32{0, 1, 2}[r.Intn(3)])
	v2 := types.NewPopulatedStdTime(r, easy)
	this.CreatedAt = *v2
	v3 := types.NewPopulatedStdTime(r, easy)
	this.UpdatedAt = *v3
	v4 := types.NewPopulatedStdTime(r, easy)
	this.PasswordUpdatedAt = *v4
	this.RequirePasswordUpdate = bool(r.Intn(2) == 0)
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

type randyUser interface {
	Float32() float32
	Float64() float64
	Int63() int64
	Int31() int32
	Uint32() uint32
	Intn(n int) int
}

func randUTF8RuneUser(r randyUser) rune {
	ru := r.Intn(62)
	if ru < 10 {
		return rune(ru + 48)
	} else if ru < 36 {
		return rune(ru + 55)
	}
	return rune(ru + 61)
}
func randStringUser(r randyUser) string {
	v5 := r.Intn(100)
	tmps := make([]rune, v5)
	for i := 0; i < v5; i++ {
		tmps[i] = randUTF8RuneUser(r)
	}
	return string(tmps)
}
func randUnrecognizedUser(r randyUser, maxFieldNumber int) (dAtA []byte) {
	l := r.Intn(5)
	for i := 0; i < l; i++ {
		wire := r.Intn(4)
		if wire == 3 {
			wire = 5
		}
		fieldNumber := maxFieldNumber + r.Intn(100)
		dAtA = randFieldUser(dAtA, r, fieldNumber, wire)
	}
	return dAtA
}
func randFieldUser(dAtA []byte, r randyUser, fieldNumber int, wire int) []byte {
	key := uint32(fieldNumber)<<3 | uint32(wire)
	switch wire {
	case 0:
		dAtA = encodeVarintPopulateUser(dAtA, uint64(key))
		v6 := r.Int63()
		if r.Intn(2) == 0 {
			v6 *= -1
		}
		dAtA = encodeVarintPopulateUser(dAtA, uint64(v6))
	case 1:
		dAtA = encodeVarintPopulateUser(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	case 2:
		dAtA = encodeVarintPopulateUser(dAtA, uint64(key))
		ll := r.Intn(100)
		dAtA = encodeVarintPopulateUser(dAtA, uint64(ll))
		for j := 0; j < ll; j++ {
			dAtA = append(dAtA, byte(r.Intn(256)))
		}
	default:
		dAtA = encodeVarintPopulateUser(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	}
	return dAtA
}
func encodeVarintPopulateUser(dAtA []byte, v uint64) []byte {
	for v >= 1<<7 {
		dAtA = append(dAtA, uint8(v&0x7f|0x80))
		v >>= 7
	}
	dAtA = append(dAtA, uint8(v))
	return dAtA
}
func (m *User) Size() (n int) {
	var l int
	_ = l
	l = m.UserIdentifiers.Size()
	n += 1 + l + sovUser(uint64(l))
	l = len(m.Password)
	if l > 0 {
		n += 1 + l + sovUser(uint64(l))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovUser(uint64(l))
	}
	if m.ValidatedAt != nil {
		l = types.SizeOfStdTime(*m.ValidatedAt)
		n += 1 + l + sovUser(uint64(l))
	}
	if m.Admin {
		n += 2
	}
	if m.State != 0 {
		n += 1 + sovUser(uint64(m.State))
	}
	l = types.SizeOfStdTime(m.CreatedAt)
	n += 1 + l + sovUser(uint64(l))
	l = types.SizeOfStdTime(m.UpdatedAt)
	n += 1 + l + sovUser(uint64(l))
	l = types.SizeOfStdTime(m.PasswordUpdatedAt)
	n += 1 + l + sovUser(uint64(l))
	if m.RequirePasswordUpdate {
		n += 2
	}
	return n
}

func sovUser(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozUser(x uint64) (n int) {
	return sovUser((x << 1) ^ uint64((int64(x) >> 63)))
}
func (m *User) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowUser
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: User: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: User: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UserIdentifiers", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUser
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthUser
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.UserIdentifiers.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Password", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUser
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthUser
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Password = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUser
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthUser
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ValidatedAt", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUser
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthUser
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.ValidatedAt == nil {
				m.ValidatedAt = new(time.Time)
			}
			if err := types.StdTimeUnmarshal(m.ValidatedAt, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Admin", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUser
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Admin = bool(v != 0)
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field State", wireType)
			}
			m.State = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUser
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.State |= (ReviewingState(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CreatedAt", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUser
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthUser
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := types.StdTimeUnmarshal(&m.CreatedAt, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UpdatedAt", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUser
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthUser
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := types.StdTimeUnmarshal(&m.UpdatedAt, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PasswordUpdatedAt", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUser
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthUser
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := types.StdTimeUnmarshal(&m.PasswordUpdatedAt, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 10:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RequirePasswordUpdate", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUser
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.RequirePasswordUpdate = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipUser(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthUser
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipUser(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowUser
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowUser
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowUser
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthUser
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowUser
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipUser(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthUser = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowUser   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("github.com/TheThingsNetwork/ttn/api/user.proto", fileDescriptorUser) }
func init() {
	golang_proto.RegisterFile("github.com/TheThingsNetwork/ttn/api/user.proto", fileDescriptorUser)
}

var fileDescriptorUser = []byte{
	// 530 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x93, 0x3d, 0x6c, 0xd3, 0x40,
	0x14, 0xc7, 0xef, 0xd1, 0x24, 0x24, 0x57, 0x84, 0x84, 0xf9, 0xa8, 0x15, 0xa4, 0x97, 0x88, 0x29,
	0x48, 0xe0, 0xa0, 0x46, 0xb0, 0x27, 0x9d, 0x58, 0x10, 0x32, 0xe9, 0xc2, 0x12, 0x39, 0xf1, 0xd5,
	0x39, 0x35, 0xfe, 0xc0, 0xbe, 0x24, 0x62, 0xeb, 0xd8, 0x81, 0xa1, 0x23, 0x1b, 0x8c, 0x1d, 0x3b,
	0x76, 0xec, 0x98, 0x31, 0x63, 0xa7, 0x52, 0x9f, 0x97, 0x8e, 0x1d, 0x3b, 0xa2, 0x9c, 0xed, 0xf4,
	0x63, 0x69, 0x3a, 0xf9, 0xde, 0xbd, 0xff, 0xef, 0x7f, 0xff, 0xf7, 0x24, 0x53, 0xc3, 0xe1, 0x62,
	0x38, 0xee, 0x1b, 0x03, 0xdf, 0x6d, 0x76, 0x87, 0xac, 0x3b, 0xe4, 0x9e, 0x13, 0x7d, 0x61, 0x62,
	0xea, 0x87, 0xbb, 0x4d, 0x21, 0xbc, 0xa6, 0x15, 0xf0, 0xe6, 0x38, 0x62, 0xa1, 0x11, 0x84, 0xbe,
	0xf0, 0xb5, 0x92, 0x10, 0x9e, 0x31, 0x69, 0x55, 0x3f, 0xac, 0xc2, 0x0d, 0x46, 0x9c, 0x79, 0x22,
	0x25, 0xab, 0x1f, 0x57, 0x21, 0xb8, 0xcd, 0x3c, 0xc1, 0x77, 0x38, 0x0b, 0xa3, 0x0c, 0x5b, 0xe9,
	0xa1, 0x90, 0x3b, 0x43, 0x91, 0x13, 0xef, 0x6f, 0x10, 0x8e, 0xef, 0xf8, 0x4d, 0x75, 0xdd, 0x1f,
	0xef, 0xa8, 0x4a, 0x15, 0xea, 0x94, 0xc9, 0x5f, 0x3b, 0xbe, 0xef, 0x8c, 0xd8, 0xb5, 0x8a, 0xb9,
	0x81, 0xf8, 0x99, 0x35, 0x6b, 0x77, 0x9b, 0x82, 0xbb, 0x2c, 0x12, 0x96, 0x1b, 0xa4, 0x82, 0x37,
	0xbf, 0x0a, 0xb4, 0xb0, 0x1d, 0xb1, 0x50, 0x6b, 0xd1, 0x35, 0x6e, 0x47, 0x3a, 0xd4, 0xa1, 0xb1,
	0xbe, 0xb9, 0x61, 0xa4, 0x6b, 0x32, 0x16, 0xad, 0xcf, 0xd7, 0x33, 0x75, 0xca, 0xb3, 0xb3, 0x1a,
	0x99, 0x9f, 0xd5, 0xc0, 0x5c, 0xa8, 0xb5, 0x2a, 0x2d, 0x07, 0x56, 0x14, 0x4d, 0xfd, 0xd0, 0xd6,
	0x1f, 0xd5, 0xa1, 0x51, 0x31, 0x97, 0xb5, 0xa6, 0xd1, 0x82, 0x67, 0xb9, 0x4c, 0x5f, 0x53, 0xf7,
	0xea, 0xac, 0x6d, 0xd1, 0x27, 0x13, 0x6b, 0xc4, 0x6d, 0x4b, 0x30, 0xbb, 0x67, 0x09, 0xbd, 0xa0,
	0x5e, 0xab, 0x1a, 0x69, 0x4a, 0x23, 0x4f, 0x69, 0x74, 0xf3, 0x94, 0x9d, 0xc2, 0xc1, 0xbf, 0x1a,
	0x98, 0xeb, 0x4b, 0xaa, 0x2d, 0xb4, 0x17, 0xb4, 0x68, 0xd9, 0x2e, 0xf7, 0xf4, 0x62, 0x1d, 0x1a,
	0x65, 0x33, 0x2d, 0xb4, 0x77, 0xb4, 0x18, 0x09, 0x4b, 0x30, 0xbd, 0x54, 0x87, 0xc6, 0xd3, 0xcd,
	0x57, 0xf9, 0x04, 0x26, 0x9b, 0x70, 0x36, 0xe5, 0x9e, 0xf3, 0x6d, 0xd1, 0x35, 0x53, 0x91, 0xb6,
	0x45, 0xe9, 0x20, 0x64, 0x79, 0x8c, 0xc7, 0xf7, 0xc6, 0x50, 0x73, 0xab, 0x28, 0x95, 0x8c, 0x6b,
	0x8b, 0x85, 0xc9, 0x38, 0x58, 0xce, 0x52, 0x7e, 0x88, 0x49, 0xc6, 0xb5, 0x85, 0xd6, 0xa5, 0xcf,
	0xf3, 0x95, 0xf5, 0x6e, 0xb8, 0x55, 0x1e, 0xe0, 0xf6, 0x2c, 0x37, 0xd8, 0x5e, 0xba, 0x7e, 0xa2,
	0x1b, 0x21, 0xfb, 0x31, 0xe6, 0x21, 0xeb, 0xdd, 0x71, 0xd7, 0xa9, 0xda, 0xda, 0xcb, 0xac, 0xfd,
	0xf5, 0x16, 0xda, 0xf9, 0x03, 0xb3, 0x18, 0x61, 0x1e, 0x23, 0x9c, 0xc6, 0x08, 0xe7, 0x31, 0xc2,
	0x45, 0x8c, 0xe4, 0x32, 0x46, 0x72, 0x15, 0x23, 0xec, 0x49, 0x24, 0xfb, 0x12, 0xc9, 0xa1, 0x44,
	0x38, 0x92, 0x48, 0x8e, 0x25, 0xc2, 0x89, 0x44, 0x98, 0x49, 0x84, 0xb9, 0x44, 0x38, 0x95, 0x48,
	0xce, 0x25, 0xc2, 0x85, 0x44, 0x72, 0x29, 0x11, 0xae, 0x24, 0x92, 0xbd, 0x04, 0xc9, 0x7e, 0x82,
	0x70, 0x90, 0x20, 0xf9, 0x9d, 0x20, 0xfc, 0x4d, 0x90, 0x1c, 0x26, 0x48, 0x8e, 0x12, 0x84, 0xe3,
	0x04, 0xe1, 0x24, 0x41, 0xf8, 0xfe, 0xf6, 0xbe, 0x9f, 0x24, 0xd8, 0x75, 0x16, 0xdf, 0xa0, 0xdf,
	0x2f, 0xa9, 0x55, 0xb4, 0xfe, 0x07, 0x00, 0x00, 0xff, 0xff, 0xf7, 0x0a, 0xe9, 0xb5, 0xf9, 0x03,
	0x00, 0x00,
}
