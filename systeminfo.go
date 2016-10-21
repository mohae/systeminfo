package systeminfo

import (
	"encoding/json"

	"github.com/golang/protobuf/proto"
	mem "github.com/mohae/joefriday/mem/basic"
	"github.com/mohae/joefriday/net/info"
	"github.com/mohae/joefriday/platform/kernel"
	"github.com/mohae/joefriday/platform/release"
	"github.com/mohae/joefriday/processors"
)

type Error struct {
	Op  string
	Err error
}

func (e Error) Error() string {
	return e.Op + ": " + e.Err.Error()
}

func (s *System) Get() error {
	//Get Kernel info
	k, err := kernel.Get()
	if err != nil {
		return Error{Op: "kernel info", Err: err}
	}
	s.KernelOS = k.OS
	s.KernelVersion = k.Version
	s.KernelArch = k.Arch
	s.KernelType = k.Type
	s.KernelCompileDate = k.CompileDate
	// Get release info
	o, err := release.Get()
	if err != nil {
		return Error{Op: "release info", Err: err}
	}
	s.OSName = o.Name
	s.OSID = o.ID
	s.OSIDLike = o.IDLike
	s.OSVersion = o.Version
	// Get Memory info
	m, err := mem.Get()
	if err != nil {
		return Error{Op: "mem info", Err: err}
	}
	s.MemTotal = m.MemTotal
	s.SwapTotal = m.SwapTotal
	// Get network interfaces
	inf, err := info.Get()
	if err != nil {
		return Error{Op: "netinf info", Err: err}
	}
	s.NetInfs = make([]string, len(inf.Interfaces))
	for i := 0; i < len(inf.Interfaces); i++ {
		s.NetInfs[i] = inf.Interfaces[i].Name
	}
	// Get processors
	p, err := processors.Get()
	if err != nil {
		return Error{Op: "processor info", Err: err}
	}
	s.Chips = make([]*Chip, len(p.Chips))
	for i := 0; i < len(p.Chips); i++ {
		var chip Chip
		chip.PhysicalID = int32(p.Chips[i].PhysicalID)
		chip.VendorID = p.Chips[i].VendorID
		chip.CPUFamily = p.Chips[i].CPUFamily
		chip.Model = p.Chips[i].Model
		chip.ModelName = p.Chips[i].ModelName
		chip.Stepping = p.Chips[i].Stepping
		chip.Microcode = p.Chips[i].Microcode
		chip.CPUMHz = p.Chips[i].CPUMHz
		chip.CacheSize = p.Chips[i].CacheSize
		chip.CPUCores = int32(p.Chips[i].CPUCores)
		chip.Flags = make([]string, len(p.Chips[i].Flags))
		copy(chip.Flags, p.Chips[i].Flags)
		s.Chips[i] = &chip
	}
	return nil
}

// JSONMarshal marshals the System info as JSON and returns the result.
func (s *System) JSONMarshal() (p []byte, err error) {
	return json.Marshal(s)
}

// JSONUnmarshal unmarshals the received bytes into System.
func JSONUnmarshal(p []byte) (s *System, err error) {
	s = &System{}
	err = json.Unmarshal(p, &s)
	return s, err
}

// ProtoMarshal marshals the System info as Protobuf and returns the result.
func (s *System) ProtoMarshal() (p []byte, err error) {
	return proto.Marshal(s)
}

// ProtoUnmarshal unmarshals the received bytes into System.
func ProtoUnmarshal(p []byte) (s *System, err error) {
	s = &System{}
	err = proto.Unmarshal(p, s)
	return s, err
}
