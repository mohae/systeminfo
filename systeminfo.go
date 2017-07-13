package systeminfo

import (
	"encoding/json"

	"github.com/golang/protobuf/proto"
	"github.com/mohae/joefriday/mem/membasic"
	"github.com/mohae/joefriday/net/netdev"
	"github.com/mohae/joefriday/system/version"
	"github.com/mohae/joefriday/system/os"
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
	k, err := version.Get()
	if err != nil {
		return Error{Op: "kernel info", Err: err}
	}
	s.KernelOS = k.OS
	s.KernelVersion = k.Version
	s.KernelArch = k.Arch
	s.KernelType = k.Type
	s.KernelCompileDate = k.CompileDate
	// Get release info
	o, err := os.Get()
	if err != nil {
		return Error{Op: "os release info", Err: err}
	}
	s.OSName = o.Name
	s.OSID = o.ID
	s.OSIDLike = o.IDLike
	s.OSVersion = o.Version
	// Get Memory info
	m, err := membasic.Get()
	if err != nil {
		return Error{Op: "mem info", Err: err}
	}
	s.MemTotal = m.MemTotal
	s.SwapTotal = m.SwapTotal
	// Get network devices
	inf, err := netdev.Get()
	if err != nil {
		return Error{Op: "network devices info", Err: err}
	}
	s.NetInfs = make([]string, len(inf.Device))
	for i := 0; i < len(inf.Device); i++ {
		s.NetInfs[i] = inf.Device[i].Name
	}
	// Get processors
	p, err := processors.Get()
	if err != nil {
		return Error{Op: "processor info", Err: err}
	}
	s.Chips = make([]*Chip, len(p.Socket))
	for i := 0; i < len(p.Socket); i++ {
		var chip Chip
		chip.PhysicalID = int32(p.Socket[i].PhysicalID)
		chip.VendorID = p.Socket[i].VendorID
		chip.CPUFamily = p.Socket[i].CPUFamily
		chip.Model = p.Socket[i].Model
		chip.ModelName = p.Socket[i].ModelName
		chip.Stepping = p.Socket[i].Stepping
		chip.Microcode = p.Socket[i].Microcode
		chip.CPUMHz = p.Socket[i].CPUMHz
		chip.CacheSize = p.Socket[i].CacheSize
		chip.CPUCores = int32(p.Socket[i].CPUCores)
		chip.Flags = make([]string, len(p.Socket[i].Flags))
		copy(chip.Flags, p.Socket[i].Flags)
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
