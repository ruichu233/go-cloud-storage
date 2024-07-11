// Code generated by Fastpb v0.0.2. DO NOT EDIT.

package v1

import (
	fmt "fmt"
	fastpb "github.com/cloudwego/fastpb"
)

var (
	_ = fmt.Errorf
	_ = fastpb.Skip
)

func (x *SignupRequest) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 2:
		offset, err = x.fastReadField2(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_SignupRequest[number], err)
}

func (x *SignupRequest) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.UserName, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *SignupRequest) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.Password, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *SignupResponse) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 2:
		offset, err = x.fastReadField2(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 3:
		offset, err = x.fastReadField3(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 4:
		offset, err = x.fastReadField4(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_SignupResponse[number], err)
}

func (x *SignupResponse) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.Code, offset, err = fastpb.ReadInt32(buf, _type)
	return offset, err
}

func (x *SignupResponse) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.Message, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *SignupResponse) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	x.UserId, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *SignupResponse) fastReadField4(buf []byte, _type int8) (offset int, err error) {
	x.Token, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *LoginRequest) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 2:
		offset, err = x.fastReadField2(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_LoginRequest[number], err)
}

func (x *LoginRequest) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.UserName, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *LoginRequest) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.Password, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *LoginResponse) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 2:
		offset, err = x.fastReadField2(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 3:
		offset, err = x.fastReadField3(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 4:
		offset, err = x.fastReadField4(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_LoginResponse[number], err)
}

func (x *LoginResponse) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.Code, offset, err = fastpb.ReadInt32(buf, _type)
	return offset, err
}

func (x *LoginResponse) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.Message, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *LoginResponse) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	x.UserId, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *LoginResponse) fastReadField4(buf []byte, _type int8) (offset int, err error) {
	x.Token, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *SignupRequest) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	return offset
}

func (x *SignupRequest) fastWriteField1(buf []byte) (offset int) {
	if x.UserName == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 1, x.GetUserName())
	return offset
}

func (x *SignupRequest) fastWriteField2(buf []byte) (offset int) {
	if x.Password == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 2, x.GetPassword())
	return offset
}

func (x *SignupResponse) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	offset += x.fastWriteField3(buf[offset:])
	offset += x.fastWriteField4(buf[offset:])
	return offset
}

func (x *SignupResponse) fastWriteField1(buf []byte) (offset int) {
	if x.Code == 0 {
		return offset
	}
	offset += fastpb.WriteInt32(buf[offset:], 1, x.GetCode())
	return offset
}

func (x *SignupResponse) fastWriteField2(buf []byte) (offset int) {
	if x.Message == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 2, x.GetMessage())
	return offset
}

func (x *SignupResponse) fastWriteField3(buf []byte) (offset int) {
	if x.UserId == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 3, x.GetUserId())
	return offset
}

func (x *SignupResponse) fastWriteField4(buf []byte) (offset int) {
	if x.Token == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 4, x.GetToken())
	return offset
}

func (x *LoginRequest) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	return offset
}

func (x *LoginRequest) fastWriteField1(buf []byte) (offset int) {
	if x.UserName == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 1, x.GetUserName())
	return offset
}

func (x *LoginRequest) fastWriteField2(buf []byte) (offset int) {
	if x.Password == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 2, x.GetPassword())
	return offset
}

func (x *LoginResponse) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	offset += x.fastWriteField3(buf[offset:])
	offset += x.fastWriteField4(buf[offset:])
	return offset
}

func (x *LoginResponse) fastWriteField1(buf []byte) (offset int) {
	if x.Code == 0 {
		return offset
	}
	offset += fastpb.WriteInt32(buf[offset:], 1, x.GetCode())
	return offset
}

func (x *LoginResponse) fastWriteField2(buf []byte) (offset int) {
	if x.Message == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 2, x.GetMessage())
	return offset
}

func (x *LoginResponse) fastWriteField3(buf []byte) (offset int) {
	if x.UserId == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 3, x.GetUserId())
	return offset
}

func (x *LoginResponse) fastWriteField4(buf []byte) (offset int) {
	if x.Token == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 4, x.GetToken())
	return offset
}

func (x *SignupRequest) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	return n
}

func (x *SignupRequest) sizeField1() (n int) {
	if x.UserName == "" {
		return n
	}
	n += fastpb.SizeString(1, x.GetUserName())
	return n
}

func (x *SignupRequest) sizeField2() (n int) {
	if x.Password == "" {
		return n
	}
	n += fastpb.SizeString(2, x.GetPassword())
	return n
}

func (x *SignupResponse) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	n += x.sizeField3()
	n += x.sizeField4()
	return n
}

func (x *SignupResponse) sizeField1() (n int) {
	if x.Code == 0 {
		return n
	}
	n += fastpb.SizeInt32(1, x.GetCode())
	return n
}

func (x *SignupResponse) sizeField2() (n int) {
	if x.Message == "" {
		return n
	}
	n += fastpb.SizeString(2, x.GetMessage())
	return n
}

func (x *SignupResponse) sizeField3() (n int) {
	if x.UserId == 0 {
		return n
	}
	n += fastpb.SizeInt64(3, x.GetUserId())
	return n
}

func (x *SignupResponse) sizeField4() (n int) {
	if x.Token == "" {
		return n
	}
	n += fastpb.SizeString(4, x.GetToken())
	return n
}

func (x *LoginRequest) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	return n
}

func (x *LoginRequest) sizeField1() (n int) {
	if x.UserName == "" {
		return n
	}
	n += fastpb.SizeString(1, x.GetUserName())
	return n
}

func (x *LoginRequest) sizeField2() (n int) {
	if x.Password == "" {
		return n
	}
	n += fastpb.SizeString(2, x.GetPassword())
	return n
}

func (x *LoginResponse) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	n += x.sizeField3()
	n += x.sizeField4()
	return n
}

func (x *LoginResponse) sizeField1() (n int) {
	if x.Code == 0 {
		return n
	}
	n += fastpb.SizeInt32(1, x.GetCode())
	return n
}

func (x *LoginResponse) sizeField2() (n int) {
	if x.Message == "" {
		return n
	}
	n += fastpb.SizeString(2, x.GetMessage())
	return n
}

func (x *LoginResponse) sizeField3() (n int) {
	if x.UserId == 0 {
		return n
	}
	n += fastpb.SizeInt64(3, x.GetUserId())
	return n
}

func (x *LoginResponse) sizeField4() (n int) {
	if x.Token == "" {
		return n
	}
	n += fastpb.SizeString(4, x.GetToken())
	return n
}

var fieldIDToName_SignupRequest = map[int32]string{
	1: "UserName",
	2: "Password",
}

var fieldIDToName_SignupResponse = map[int32]string{
	1: "Code",
	2: "Message",
	3: "UserId",
	4: "Token",
}

var fieldIDToName_LoginRequest = map[int32]string{
	1: "UserName",
	2: "Password",
}

var fieldIDToName_LoginResponse = map[int32]string{
	1: "Code",
	2: "Message",
	3: "UserId",
	4: "Token",
}
