// Autogenerated by Thrift Compiler (0.9.3)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package request

import (
	"bytes"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
	"github.com/caijinlin/learning-pratice/thrift/lib/types"
)

// (needed to ensure safety because of naive import list construction.)
var _ = thrift.ZERO
var _ = fmt.Printf
var _ = bytes.Equal

var _ = types.GoUnusedProtection__
var GoUnusedProtection__ int

// Attributes:
//  - Logid
//  - UID
type GetUserRequest struct {
	Logid string `thrift:"logid,1,required" json:"logid"`
	UID   int32  `thrift:"uid,2,required" json:"uid"`
}

func NewGetUserRequest() *GetUserRequest {
	return &GetUserRequest{}
}

func (p *GetUserRequest) GetLogid() string {
	return p.Logid
}

func (p *GetUserRequest) GetUID() int32 {
	return p.UID
}
func (p *GetUserRequest) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	var issetLogid bool = false
	var issetUID bool = false

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if err := p.readField1(iprot); err != nil {
				return err
			}
			issetLogid = true
		case 2:
			if err := p.readField2(iprot); err != nil {
				return err
			}
			issetUID = true
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	if !issetLogid {
		return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("Required field Logid is not set"))
	}
	if !issetUID {
		return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("Required field UID is not set"))
	}
	return nil
}

func (p *GetUserRequest) readField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.Logid = v
	}
	return nil
}

func (p *GetUserRequest) readField2(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return thrift.PrependError("error reading field 2: ", err)
	} else {
		p.UID = v
	}
	return nil
}

func (p *GetUserRequest) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("GetUserRequest"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := p.writeField2(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *GetUserRequest) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("logid", thrift.STRING, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:logid: ", p), err)
	}
	if err := oprot.WriteString(string(p.Logid)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.logid (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:logid: ", p), err)
	}
	return err
}

func (p *GetUserRequest) writeField2(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("uid", thrift.I32, 2); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:uid: ", p), err)
	}
	if err := oprot.WriteI32(int32(p.UID)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.uid (2) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 2:uid: ", p), err)
	}
	return err
}

func (p *GetUserRequest) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("GetUserRequest(%+v)", *p)
}

// Attributes:
//  - Logid
//  - Userlist
type SayHelloRequest struct {
	Logid    string          `thrift:"logid,1,required" json:"logid"`
	Userlist *types.UserList `thrift:"userlist,2,required" json:"userlist"`
}

func NewSayHelloRequest() *SayHelloRequest {
	return &SayHelloRequest{}
}

func (p *SayHelloRequest) GetLogid() string {
	return p.Logid
}

var SayHelloRequest_Userlist_DEFAULT *types.UserList

func (p *SayHelloRequest) GetUserlist() *types.UserList {
	if !p.IsSetUserlist() {
		return SayHelloRequest_Userlist_DEFAULT
	}
	return p.Userlist
}
func (p *SayHelloRequest) IsSetUserlist() bool {
	return p.Userlist != nil
}

func (p *SayHelloRequest) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	var issetLogid bool = false
	var issetUserlist bool = false

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if err := p.readField1(iprot); err != nil {
				return err
			}
			issetLogid = true
		case 2:
			if err := p.readField2(iprot); err != nil {
				return err
			}
			issetUserlist = true
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	if !issetLogid {
		return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("Required field Logid is not set"))
	}
	if !issetUserlist {
		return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("Required field Userlist is not set"))
	}
	return nil
}

func (p *SayHelloRequest) readField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.Logid = v
	}
	return nil
}

func (p *SayHelloRequest) readField2(iprot thrift.TProtocol) error {
	p.Userlist = &types.UserList{}
	if err := p.Userlist.Read(iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Userlist), err)
	}
	return nil
}

func (p *SayHelloRequest) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("SayHelloRequest"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := p.writeField2(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *SayHelloRequest) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("logid", thrift.STRING, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:logid: ", p), err)
	}
	if err := oprot.WriteString(string(p.Logid)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.logid (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:logid: ", p), err)
	}
	return err
}

func (p *SayHelloRequest) writeField2(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("userlist", thrift.STRUCT, 2); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:userlist: ", p), err)
	}
	if err := p.Userlist.Write(oprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Userlist), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 2:userlist: ", p), err)
	}
	return err
}

func (p *SayHelloRequest) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("SayHelloRequest(%+v)", *p)
}