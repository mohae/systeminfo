package systeminfo

import (
	"os"
	"sync"

	fb "github.com/google/flatbuffers/go"
	"github.com/mohae/joefriday/cpu/cpuinfo"
	mem "github.com/mohae/joefriday/mem/membasic"
	"github.com/mohae/joefriday/net/netdev"
	sysos "github.com/mohae/joefriday/system/os"
	"github.com/mohae/joefriday/system/version"
	sysinfo "github.com/mohae/systeminfo"
	"github.com/mohae/systeminfo/flat/structs"
)

var (
	b   = fb.NewBuilder(0)
	bMu sync.Mutex
)

// Get get's the system info using the package global builder.  This is safe
// for concurrent use.
func Get() ([]byte, error) {
	bMu.Lock()
	defer bMu.Unlock()
	b.Reset()
	return GetWBuilder(b)
}

// GetWBuilder get's the system info using the passed builder.  It is assumed
// that the builder is ready to use.
func GetWBuilder(bldr *fb.Builder) ([]byte, error) {
	hostname, err := os.Hostname()
	if err != nil {
		return nil, sysinfo.Error{Op: "flat hostname", Err: err}
	}
	hostnameS := bldr.CreateString(hostname)
	//Get Kernel info
	k, err := version.Get()
	if err != nil {
		return nil, sysinfo.Error{Op: "flat kernel version info", Err: err}
	}
	osS := bldr.CreateString(k.OS)
	version := bldr.CreateString(k.Version)
	arch := bldr.CreateString(k.Arch)
	compileDate := bldr.CreateString(k.CompileDate)
	// Get release info
	o, err := sysos.Get()
	if err != nil {
		return nil, sysinfo.Error{Op: "flat OS release info", Err: err}
	}
	oName := bldr.CreateString(o.Name)
	oID := bldr.CreateString(o.ID)
	oIDLike := bldr.CreateString(o.IDLike)
	oVersion := bldr.CreateString(o.Version)
	// Get Memory info
	m, err := mem.Get()
	if err != nil {
		return nil, sysinfo.Error{Op: "flat mem info", Err: err}
	}

	// Get network interfaces
	inf, err := netdev.Get()
	if err != nil {
		return nil, sysinfo.Error{Op: "flat netdev info", Err: err}
	}
	uof := make([]fb.UOffsetT, len(inf.Device))
	for i, nic := range inf.Device {
		uof[i] = bldr.CreateString(nic.Name)
	}
	structs.SystemStartNetDevVector(bldr, len(inf.Device))
	for i := len(uof) - 1; i >= 0; i-- {
		bldr.PrependUOffsetT(uof[i])
	}
	netdevs := bldr.EndVector(len(uof))
	// Get processors
	p, err := cpuinfo.Get()
	if err != nil {
		return nil, sysinfo.Error{Op: "flat processor info", Err: err}
	}

	uoff := make([]fb.UOffsetT, len(p.CPU))
	for i, cpu := range p.CPU {
		uoff[i] = serializeCPU(bldr, &cpu)
	}
	structs.SystemStartCPUVector(bldr, len(uoff))
	for i := len(uoff) - 1; i >= 0; i-- {
		bldr.PrependUOffsetT(uoff[i])
	}
	cpus := bldr.EndVector(len(uoff))
	structs.SystemStart(bldr)
	structs.SystemAddHostname(bldr, hostnameS)
	structs.SystemAddKernelOS(bldr, osS)
	structs.SystemAddKernelVersion(bldr, version)
	structs.SystemAddKernelArch(bldr, arch)
	structs.SystemAddKernelCompileDate(bldr, compileDate)
	structs.SystemAddOSName(bldr, oName)
	structs.SystemAddOSID(bldr, oID)
	structs.SystemAddOSIDLike(bldr, oIDLike)
	structs.SystemAddOSVersion(bldr, oVersion)
	structs.SystemAddMemTotal(bldr, m.MemTotal)
	structs.SystemAddSwapTotal(bldr, m.SwapTotal)
	structs.SystemAddNetDev(bldr, netdevs)
	structs.SystemAddSockets(bldr, int32(p.Sockets))
	structs.SystemAddCPU(bldr, cpus)
	bldr.Finish(structs.SystemEnd(bldr))
	bs := bldr.Bytes[bldr.Head():]
	tmp := make([]byte, len(bs))
	copy(tmp, bs)
	return tmp, nil
}

func serializeCPU(bldr *fb.Builder, c *cpuinfo.CPU) fb.UOffsetT {
	vendorID := bldr.CreateString(c.VendorID)
	cpuFamily := bldr.CreateString(c.CPUFamily)
	model := bldr.CreateString(c.Model)
	modelName := bldr.CreateString(c.ModelName)
	stepping := bldr.CreateString(c.Stepping)
	microcode := bldr.CreateString(c.Microcode)
	cacheSize := bldr.CreateString(c.CacheSize)
	// flags
	uoffs := make([]fb.UOffsetT, len(c.Flags))
	for i, flag := range c.Flags {
		uoffs[i] = bldr.CreateString(flag)
	}
	structs.CPUStartFlagsVector(bldr, len(uoffs))
	for i := len(uoffs) - 1; i >= 0; i-- {
		bldr.PrependUOffsetT(uoffs[i])
	}
	flags := bldr.EndVector(len(uoffs))
	// bugs
	uoffs = make([]fb.UOffsetT, len(c.Bugs))
	for i, bug := range c.Bugs {
		uoffs[i] = bldr.CreateString(bug)
	}
	structs.CPUStartBugsVector(bldr, len(uoffs))
	for i := len(uoffs) - 1; i >= 0; i-- {
		bldr.PrependUOffsetT(uoffs[i])
	}
	bugs := bldr.EndVector(len(uoffs))
	structs.CPUStart(bldr)
	structs.CPUAddPhysicalID(bldr, int32(c.PhysicalID))
	structs.CPUAddCoreID(bldr, int32(c.CoreID))
	structs.CPUAddSiblings(bldr, int32(c.Siblings))
	structs.CPUAddVendorID(bldr, vendorID)
	structs.CPUAddCPUFamily(bldr, cpuFamily)
	structs.CPUAddModel(bldr, model)
	structs.CPUAddModelName(bldr, modelName)
	structs.CPUAddStepping(bldr, stepping)
	structs.CPUAddMicrocode(bldr, microcode)
	structs.CPUAddCPUMHz(bldr, c.CPUMHz)
	structs.CPUAddBogoMIPS(bldr, c.BogoMIPS)
	structs.CPUAddCacheSize(bldr, cacheSize)
	structs.CPUAddCPUCores(bldr, int32(c.CPUCores))
	structs.CPUAddFlags(bldr, flags)
	structs.CPUAddBugs(bldr, bugs)
	return structs.CPUEnd(bldr)
}

// Deserialize deserializes bytes representing a flatbuffer serialized System
// as *sysinformation.System.  If you want to work with System directly (the
// flatbuffers version), use GetRootAsSystem instead.
func Deserialize(p []byte) *sysinfo.System {
	var sys sysinfo.System
	s := structs.GetRootAsSystem(p, 0)
	cpuF := &structs.CPU{}
	cpu := &sysinfo.CPU{}
	sys.Hostname = string(s.Hostname())
	sys.KernelOS = string(s.KernelOS())
	sys.KernelVersion = string(s.KernelVersion())
	sys.KernelArch = string(s.KernelArch())
	sys.KernelCompileDate = string(s.KernelCompileDate())
	sys.OSName = string(s.OSName())
	sys.OSID = string(s.OSID())
	sys.OSIDLike = string(s.OSIDLike())
	sys.OSVersion = string(s.OSVersion())
	sys.MemTotal = s.MemTotal()
	sys.SwapTotal = s.SwapTotal()
	sys.NetDev = make([]string, s.NetDevLength())
	for i := 0; i < len(sys.NetDev); i++ {
		sys.NetDev[i] = string(s.NetDev(i))
	}
	sys.Sockets = s.Sockets()
	sys.CPU = make([]*sysinfo.CPU, s.CPULength())
	for i := 0; i < len(sys.CPU); i++ {
		if !s.CPU(cpuF, i) {
			continue
		}
		cpu.PhysicalID = cpuF.PhysicalID()
		cpu.CoreID = cpuF.CoreID()
		cpu.Siblings = cpuF.Siblings()
		cpu.VendorID = string(cpuF.VendorID())
		cpu.CPUFamily = string(cpuF.CPUFamily())
		cpu.Model = string(cpuF.Model())
		cpu.ModelName = string(cpuF.ModelName())
		cpu.Stepping = string(cpuF.Stepping())
		cpu.Microcode = string(cpuF.Microcode())
		cpu.CPUMHz = cpuF.CPUMHz()
		cpu.BogoMIPS = cpuF.BogoMIPS()
		cpu.CacheSize = string(cpuF.CacheSize())
		cpu.CPUCores = cpuF.CPUCores()
		cpu.Flags = make([]string, cpuF.FlagsLength())
		for j := 0; j < len(cpu.Flags); j++ {
			cpu.Flags[j] = string(cpuF.Flags(j))
		}
		cpu.Bugs = make([]string, cpuF.BugsLength())
		for j := 0; j < len(cpu.Bugs); j++ {
			cpu.Bugs[j] = string(cpuF.Bugs(j))
		}
		sys.CPU[i] = cpu
	}
	return &sys
}
