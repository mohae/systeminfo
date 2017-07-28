package systeminfo

import (
	"encoding/json"

	"github.com/golang/protobuf/proto"
	"github.com/mohae/joefriday/mem/membasic"
	"github.com/mohae/joefriday/net/netdev"
	"github.com/mohae/joefriday/processors"
	"github.com/mohae/joefriday/system/os"
	"github.com/mohae/joefriday/system/version"
)

type Error struct {
	Op  string
	Err error
}

func (e Error) Error() string {
	return e.Op + ": " + e.Err.Error()
}

func (s *System) Get() error {
	err := s.version()
	if err != nil {
		return err
	}
	err = s.os()
	if err != nil {
		return err
	}
	err = s.memory()
	if err != nil {
		return err
	}

	err = s.netdev()
	if err != nil {
		return err
	}

	err = s.processors()
	if err != nil {
		return err
	}

	return nil
}

func (s *System) version() error {
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
	return nil
}

func (s *System) os() error {
	// Get release info
	o, err := os.Get()
	if err != nil {
		return Error{Op: "os release info", Err: err}
	}
	s.OSName = o.Name
	s.OSID = o.ID
	s.OSIDLike = o.IDLike
	s.OSVersion = o.Version
	return nil
}

func (s *System) memory() error {
	m, err := membasic.Get()
	if err != nil {
		return Error{Op: "mem info", Err: err}
	}
	s.MemTotal = m.MemTotal
	s.SwapTotal = m.SwapTotal
	return nil
}

func (s *System) netdev() error {
	// Get network devices
	inf, err := netdev.Get()
	if err != nil {
		return Error{Op: "network devices info", Err: err}
	}
	s.NetDev = make([]string, len(inf.Device))
	for i := 0; i < len(inf.Device); i++ {
		s.NetDev[i] = inf.Device[i].Name
	}
	return nil
}

func (s *System) processors() error {
	// Get processors
	p, err := processors.Get()
	if err != nil {
		return Error{Op: "processor info", Err: err}
	}
	s.Socket = make([]*Processor, len(p.Socket))
	for i := 0; i < len(p.Socket); i++ {
		var proc Processor
		proc.PhysicalID = int32(p.Socket[i].PhysicalID)
		proc.VendorID = p.Socket[i].VendorID
		proc.CPUFamily = p.Socket[i].CPUFamily
		proc.Model = p.Socket[i].Model
		proc.ModelName = p.Socket[i].ModelName
		proc.Stepping = p.Socket[i].Stepping
		proc.Microcode = p.Socket[i].Microcode
		proc.CPUMHz = p.Socket[i].CPUMHz
		proc.BogoMIPS = p.Socket[i].BogoMIPS
		proc.CacheSize = p.Socket[i].CacheSize
		proc.CPUCores = int32(p.Socket[i].CPUCores)
		proc.Flags = make([]string, len(p.Socket[i].Flags))
		copy(proc.Flags, p.Socket[i].Flags)
		s.Socket[i] = &proc
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
