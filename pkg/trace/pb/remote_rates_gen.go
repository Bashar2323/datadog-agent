package pb

// Code generated by github.com/tinylib/msgp DO NOT EDIT.

import (
	"github.com/tinylib/msgp/msgp"
)

// DecodeMsg implements msgp.Decodable
func (z *APMSampling) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, err = dc.ReadMapHeader()
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		switch msgp.UnsafeString(field) {
		case "0":
			var zb0002 uint32
			zb0002, err = dc.ReadArrayHeader()
			if err != nil {
				err = msgp.WrapError(err, "TargetTPS")
				return
			}
			if cap(z.TargetTPS) >= int(zb0002) {
				z.TargetTPS = (z.TargetTPS)[:zb0002]
			} else {
				z.TargetTPS = make([]TargetTPS, zb0002)
			}
			for za0001 := range z.TargetTPS {
				err = z.TargetTPS[za0001].DecodeMsg(dc)
				if err != nil {
					err = msgp.WrapError(err, "TargetTPS", za0001)
					return
				}
			}
		default:
			err = dc.Skip()
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *APMSampling) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 1
	// write "0"
	err = en.Append(0x81, 0xa1, 0x30)
	if err != nil {
		return
	}
	err = en.WriteArrayHeader(uint32(len(z.TargetTPS)))
	if err != nil {
		err = msgp.WrapError(err, "TargetTPS")
		return
	}
	for za0001 := range z.TargetTPS {
		err = z.TargetTPS[za0001].EncodeMsg(en)
		if err != nil {
			err = msgp.WrapError(err, "TargetTPS", za0001)
			return
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *APMSampling) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 1
	// string "0"
	o = append(o, 0x81, 0xa1, 0x30)
	o = msgp.AppendArrayHeader(o, uint32(len(z.TargetTPS)))
	for za0001 := range z.TargetTPS {
		o, err = z.TargetTPS[za0001].MarshalMsg(o)
		if err != nil {
			err = msgp.WrapError(err, "TargetTPS", za0001)
			return
		}
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *APMSampling) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		switch msgp.UnsafeString(field) {
		case "0":
			var zb0002 uint32
			zb0002, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "TargetTPS")
				return
			}
			if cap(z.TargetTPS) >= int(zb0002) {
				z.TargetTPS = (z.TargetTPS)[:zb0002]
			} else {
				z.TargetTPS = make([]TargetTPS, zb0002)
			}
			for za0001 := range z.TargetTPS {
				bts, err = z.TargetTPS[za0001].UnmarshalMsg(bts)
				if err != nil {
					err = msgp.WrapError(err, "TargetTPS", za0001)
					return
				}
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *APMSampling) Msgsize() (s int) {
	s = 1 + 2 + msgp.ArrayHeaderSize
	for za0001 := range z.TargetTPS {
		s += z.TargetTPS[za0001].Msgsize()
	}
	return
}

// DecodeMsg implements msgp.Decodable
func (z *SamplingMechanism) DecodeMsg(dc *msgp.Reader) (err error) {
	{
		var zb0001 uint32
		zb0001, err = dc.ReadUint32()
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		(*z) = SamplingMechanism(zb0001)
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z SamplingMechanism) EncodeMsg(en *msgp.Writer) (err error) {
	err = en.WriteUint32(uint32(z))
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z SamplingMechanism) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	o = msgp.AppendUint32(o, uint32(z))
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *SamplingMechanism) UnmarshalMsg(bts []byte) (o []byte, err error) {
	{
		var zb0001 uint32
		zb0001, bts, err = msgp.ReadUint32Bytes(bts)
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		(*z) = SamplingMechanism(zb0001)
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z SamplingMechanism) Msgsize() (s int) {
	s = msgp.Uint32Size
	return
}

// DecodeMsg implements msgp.Decodable
func (z *TargetTPS) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, err = dc.ReadMapHeader()
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		switch msgp.UnsafeString(field) {
		case "0":
			z.Service, err = dc.ReadString()
			if err != nil {
				err = msgp.WrapError(err, "Service")
				return
			}
		case "1":
			z.Env, err = dc.ReadString()
			if err != nil {
				err = msgp.WrapError(err, "Env")
				return
			}
		case "2":
			z.Value, err = dc.ReadFloat64()
			if err != nil {
				err = msgp.WrapError(err, "Value")
				return
			}
		case "3":
			z.Rank, err = dc.ReadUint32()
			if err != nil {
				err = msgp.WrapError(err, "Rank")
				return
			}
		case "4":
			{
				var zb0002 uint32
				zb0002, err = dc.ReadUint32()
				if err != nil {
					err = msgp.WrapError(err, "Mechanism")
					return
				}
				z.Mechanism = SamplingMechanism(zb0002)
			}
		default:
			err = dc.Skip()
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *TargetTPS) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 5
	// write "0"
	err = en.Append(0x85, 0xa1, 0x30)
	if err != nil {
		return
	}
	err = en.WriteString(z.Service)
	if err != nil {
		err = msgp.WrapError(err, "Service")
		return
	}
	// write "1"
	err = en.Append(0xa1, 0x31)
	if err != nil {
		return
	}
	err = en.WriteString(z.Env)
	if err != nil {
		err = msgp.WrapError(err, "Env")
		return
	}
	// write "2"
	err = en.Append(0xa1, 0x32)
	if err != nil {
		return
	}
	err = en.WriteFloat64(z.Value)
	if err != nil {
		err = msgp.WrapError(err, "Value")
		return
	}
	// write "3"
	err = en.Append(0xa1, 0x33)
	if err != nil {
		return
	}
	err = en.WriteUint32(z.Rank)
	if err != nil {
		err = msgp.WrapError(err, "Rank")
		return
	}
	// write "4"
	err = en.Append(0xa1, 0x34)
	if err != nil {
		return
	}
	err = en.WriteUint32(uint32(z.Mechanism))
	if err != nil {
		err = msgp.WrapError(err, "Mechanism")
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *TargetTPS) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 5
	// string "0"
	o = append(o, 0x85, 0xa1, 0x30)
	o = msgp.AppendString(o, z.Service)
	// string "1"
	o = append(o, 0xa1, 0x31)
	o = msgp.AppendString(o, z.Env)
	// string "2"
	o = append(o, 0xa1, 0x32)
	o = msgp.AppendFloat64(o, z.Value)
	// string "3"
	o = append(o, 0xa1, 0x33)
	o = msgp.AppendUint32(o, z.Rank)
	// string "4"
	o = append(o, 0xa1, 0x34)
	o = msgp.AppendUint32(o, uint32(z.Mechanism))
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *TargetTPS) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		switch msgp.UnsafeString(field) {
		case "0":
			z.Service, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Service")
				return
			}
		case "1":
			z.Env, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Env")
				return
			}
		case "2":
			z.Value, bts, err = msgp.ReadFloat64Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Value")
				return
			}
		case "3":
			z.Rank, bts, err = msgp.ReadUint32Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Rank")
				return
			}
		case "4":
			{
				var zb0002 uint32
				zb0002, bts, err = msgp.ReadUint32Bytes(bts)
				if err != nil {
					err = msgp.WrapError(err, "Mechanism")
					return
				}
				z.Mechanism = SamplingMechanism(zb0002)
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *TargetTPS) Msgsize() (s int) {
	s = 1 + 2 + msgp.StringPrefixSize + len(z.Service) + 2 + msgp.StringPrefixSize + len(z.Env) + 2 + msgp.Float64Size + 2 + msgp.Uint32Size + 2 + msgp.Uint32Size
	return
}
