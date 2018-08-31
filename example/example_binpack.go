package example

import (
	"encoding/binary"
	"errors"
	"strconv"
)

func (v *Device) MarshalLength() int {
	length := 0
	if v.Type != 0 {
		length += 2
	}
	if v.Vendor != nil {
		length += 3 + v.Vendor.MarshalLength()
	}
	if v.Model != nil {
		length += 3 + v.Model.MarshalLength()
	}
	if v.OS != nil {
		length += 3 + v.OS.MarshalLength()
	}
	if v.Browser != nil {
		length += 3 + v.Browser.MarshalLength()
	}
	if v.Language != nil {
		length += 3 + v.Language.MarshalLength()
	}
	return length
}

func (v *Device) Marshal() []byte {
	if v == nil {
		return nil
	}
	length := v.MarshalLength()
	if length == 0 {
		return nil
	}
	data := make([]byte, length)
	v.MarshalTo(data)
	return data
}

func (v *Device) MarshalTo(data []byte) {
	i := 0
	if v.Type != 0 {
		data[i] = 1
		i++
		data[i] = byte(v.Type)
		i++
	}
	if v.Vendor != nil {
		data[i] = 2
		i++
		length := v.Vendor.MarshalLength()
		binary.LittleEndian.PutUint16(data[i:], uint16(length))
		i += 2
		v.Vendor.MarshalTo(data[i:])
		i += length
	}
	if v.Model != nil {
		data[i] = 3
		i++
		length := v.Model.MarshalLength()
		binary.LittleEndian.PutUint16(data[i:], uint16(length))
		i += 2
		v.Model.MarshalTo(data[i:])
		i += length
	}
	if v.OS != nil {
		data[i] = 4
		i++
		length := v.OS.MarshalLength()
		binary.LittleEndian.PutUint16(data[i:], uint16(length))
		i += 2
		v.OS.MarshalTo(data[i:])
		i += length
	}
	if v.Browser != nil {
		data[i] = 5
		i++
		length := v.Browser.MarshalLength()
		binary.LittleEndian.PutUint16(data[i:], uint16(length))
		i += 2
		v.Browser.MarshalTo(data[i:])
		i += length
	}
	if v.Language != nil {
		data[i] = 6
		i++
		length := v.Language.MarshalLength()
		binary.LittleEndian.PutUint16(data[i:], uint16(length))
		i += 2
		v.Language.MarshalTo(data[i:])
		i += length
	}
}

func (v *Device) Unmarshal(data []byte) error {
	if len(data) == 0 {
		return nil
	}
	i := 0
	for {
		switch data[i] {
		case 1:
			i++
			if len(data) < i+1 {
				return errors.New("Invalid Device.Type data length size")
			}
			v.Type = DeviceType(data[i])
			i++
		case 2:
			i++
			if len(data) < i+2 {
				return errors.New("Invalid Device.Vendor data size")
			}
			length := int(binary.LittleEndian.Uint16(data[i : i+2]))
			i += 2
			if len(data) < i+length {
				return errors.New("Invalid Device.Vendor data length size")
			}
			if v.Vendor == nil {
				v.Vendor = &KeyValue{}
			}
			err := v.Vendor.Unmarshal(data[i : i+length])
			if err != nil {
				return err
			}
			i += length
		case 3:
			i++
			if len(data) < i+2 {
				return errors.New("Invalid Device.Model data size")
			}
			length := int(binary.LittleEndian.Uint16(data[i : i+2]))
			i += 2
			if len(data) < i+length {
				return errors.New("Invalid Device.Model data length size")
			}
			if v.Model == nil {
				v.Model = &KeyValue{}
			}
			err := v.Model.Unmarshal(data[i : i+length])
			if err != nil {
				return err
			}
			i += length
		case 4:
			i++
			if len(data) < i+2 {
				return errors.New("Invalid Device.OS data size")
			}
			length := int(binary.LittleEndian.Uint16(data[i : i+2]))
			i += 2
			if len(data) < i+length {
				return errors.New("Invalid Device.OS data length size")
			}
			if v.OS == nil {
				v.OS = &KeyValueVersion{}
			}
			err := v.OS.Unmarshal(data[i : i+length])
			if err != nil {
				return err
			}
			i += length
		case 5:
			i++
			if len(data) < i+2 {
				return errors.New("Invalid Device.Browser data size")
			}
			length := int(binary.LittleEndian.Uint16(data[i : i+2]))
			i += 2
			if len(data) < i+length {
				return errors.New("Invalid Device.Browser data length size")
			}
			if v.Browser == nil {
				v.Browser = &KeyValueVersion{}
			}
			err := v.Browser.Unmarshal(data[i : i+length])
			if err != nil {
				return err
			}
			i += length
		case 6:
			i++
			if len(data) < i+2 {
				return errors.New("Invalid Device.Language data size")
			}
			length := int(binary.LittleEndian.Uint16(data[i : i+2]))
			i += 2
			if len(data) < i+length {
				return errors.New("Invalid Device.Language data length size")
			}
			if v.Language == nil {
				v.Language = &KeyValue{}
			}
			err := v.Language.Unmarshal(data[i : i+length])
			if err != nil {
				return err
			}
			i += length
		default:
			return errors.New("Invalid Device index: " + strconv.Itoa(int(data[i])))
		}
		if len(data) == i {
			break
		}
	}
	return nil
}

func (v *KeyValue) MarshalLength() int {
	length := 0
	if v.ID != 0 {
		length += 9
	}
	if len(v.Name) != 0 {
		length += 2 + len(v.Name)
	}
	return length
}

func (v *KeyValue) Marshal() []byte {
	if v == nil {
		return nil
	}
	length := v.MarshalLength()
	if length == 0 {
		return nil
	}
	data := make([]byte, length)
	v.MarshalTo(data)
	return data
}

func (v *KeyValue) MarshalTo(data []byte) {
	i := 0
	if v.ID != 0 {
		data[i] = 1
		i++
		binary.LittleEndian.PutUint64(data[i:], v.ID)
		i += 8
	}
	if len(v.Name) != 0 {
		data[i] = 2
		i++
		data[i] = byte(len(v.Name))
		i++
		copy(data[i:], []byte(v.Name))
	}
}

func (v *KeyValue) Unmarshal(data []byte) error {
	if len(data) == 0 {
		return nil
	}
	i := 0
	for {
		switch data[i] {
		case 1:
			i++
			if len(data) < i+8 {
				return errors.New("Invalid KeyValue.ID data size")
			}
			v.ID = binary.LittleEndian.Uint64(data[i:])
			i += 8
		case 2:
			i++
			if len(data) < i {
				return errors.New("Invalid KeyValue.Name data size")
			}
			length := int(data[i])
			i++
			if len(data) < i+length {
				return errors.New("Invalid KeyValue.Name data length size")
			}
			v.Name = string(data[i : i+length])
			i += length
		default:
			return errors.New("Invalid KeyValue index: " + strconv.Itoa(int(data[i])))
		}
		if len(data) == i {
			break
		}
	}
	return nil
}

func (v *KeyValueVersion) MarshalLength() int {
	length := 0
	if v.ID != 0 {
		length += 9 // index + size
	}
	if len(v.Name) != 0 {
		length += 2 + len(v.Name) // index + size + len
	}
	if v.Version != nil {
		length += 2 + v.Version.MarshalLength() // index + size + len
	}
	return length
}

func (v *KeyValueVersion) Marshal() []byte {
	if v == nil {
		return nil
	}
	length := v.MarshalLength()
	if length == 0 {
		return nil
	}
	data := make([]byte, length)
	v.MarshalTo(data)
	return data
}

func (v *KeyValueVersion) MarshalTo(data []byte) {
	i := 0
	if v.ID != 0 {
		data[i] = 1
		i++
		binary.LittleEndian.PutUint64(data[i:], v.ID)
		i += 8
	}
	if len(v.Name) != 0 {
		data[i] = 2
		i++
		data[i] = byte(len(v.Name))
		i++
		copy(data[i:], []byte(v.Name))
		i += len(v.Name)
	}
	if v.Version != nil {
		data[i] = 3
		i++
		length := v.Version.MarshalLength()
		data[i] = uint8(length)
		i++
		v.Version.MarshalTo(data[i:])
		i += length
	}
}

func (v *KeyValueVersion) Unmarshal(data []byte) error {
	if len(data) == 0 {
		return nil
	}
	i := 0
	for {
		switch data[i] {
		case 1:
			i++
			if len(data) < i+8 {
				return errors.New("Invalid KeyValueVersion.ID data size")
			}
			v.ID = binary.LittleEndian.Uint64(data[i:])
			i += 8
		case 2:
			i++
			if len(data) < i {
				return errors.New("Invalid KeyValueVersion.Name data size")
			}
			length := int(data[i])
			i++
			if len(data) < i+length {
				return errors.New("Invalid KeyValueVersion.Name data length size")
			}
			v.Name = string(data[i : i+length])
			i += length
		case 3:
			i++
			if len(data) < i {
				return errors.New("Invalid KeyValueVersion.Version data size")
			}
			length := int(data[i])
			i++
			if len(data) < i+length {
				return errors.New("Invalid KeyValueVersion.Version data length size")
			}
			if v.Version == nil {
				v.Version = &Version{}
			}
			err := v.Version.Unmarshal(data[i : i+length])
			if err != nil {
				return err
			}
			i += length
		default:
			return errors.New("Invalid KeyValueVersion index: " + strconv.Itoa(int(data[i])))
		}
		if len(data) == i {
			break
		}
	}
	return nil
}

func (v *Version) MarshalLength() int {
	length := 0
	if len(v.Major) != 0 {
		length += 2 + len(v.Major)
	}
	if len(v.Minor) != 0 {
		length += 2 + len(v.Minor)
	}
	return length
}

func (v *Version) Marshal() []byte {
	if v == nil {
		return nil
	}
	length := v.MarshalLength()
	if length == 0 {
		return nil
	}
	data := make([]byte, length)
	v.MarshalTo(data)
	return data
}

func (v *Version) MarshalTo(data []byte) {
	i := 0
	if len(v.Major) != 0 {
		data[i] = 1
		i++
		data[i] = byte(len(v.Major))
		i++
		copy(data[i:], []byte(v.Major))
		i += len(v.Major)
	}
	if len(v.Minor) != 0 {
		data[i] = 2
		i++
		data[i] = byte(len(v.Minor))
		i++
		copy(data[i:], []byte(v.Minor))
		i += len(v.Minor)
	}
}

func (v *Version) Unmarshal(data []byte) error {
	if len(data) == 0 {
		return nil
	}
	i := 0
	for {
		switch data[i] {
		case 1:
			i++
			if len(data) < i {
				return errors.New("Invalid Version.Major data size")
			}
			length := int(data[i])
			i++
			if len(data) < i+length {
				return errors.New("Invalid Version.Major data length size")
			}
			v.Major = string(data[i : i+length])
			i += length
		case 2:
			i++
			if len(data) < i {
				return errors.New("Invalid Version.Minor data size")
			}
			length := int(data[i])
			i++
			if len(data) < i+length {
				return errors.New("Invalid Version.Minor data length size")
			}
			v.Minor = string(data[i : i+length])
			i += length
		default:
			return errors.New("Invalid Version index: " + strconv.Itoa(int(data[i])))
		}
		if len(data) == i {
			break
		}
	}
	return nil
}
