package systeminfo

import (
	"sync"

	fb "github.com/google/flatbuffers/go"
	mem "github.com/mohae/joefriday/mem/membasic"
	"github.com/mohae/joefriday/net/netdev"
	"github.com/mohae/joefriday/system/os"
	"github.com/mohae/joefriday/system/version"
	"github.com/mohae/joefriday/processors"
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
	o, err := os.Get()
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
	p, err := processors.Get()
	if err != nil {
		return nil, sysinfo.Error{Op: "flat processor info", Err: err}
	}

	uoff := make([]fb.UOffsetT, len(p.Socket))
	for i, processor := range p.Socket {
		uoff[i] = serializeProcessor(bldr, &processor)
	}
	structs.SystemStartSocketVector(bldr, len(uoff))
	for i := len(uoff) - 1; i >= 0; i-- {
		bldr.PrependUOffsetT(uoff[i])
	}
	procs := bldr.EndVector(len(uoff))
	structs.SystemStart(bldr)
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
	structs.SystemAddSocket(bldr, procs)
	bldr.Finish(structs.SystemEnd(bldr))
	bs := bldr.Bytes[bldr.Head():]
	tmp := make([]byte, len(bs))
	copy(tmp, bs)
	return tmp, nil
}

func serializeProcessor(bldr *fb.Builder, c *processors.Processor) fb.UOffsetT {
	vendorID := bldr.CreateString(c.VendorID)
	cpuFamily := bldr.CreateString(c.CPUFamily)
	model := bldr.CreateString(c.Model)
	modelName := bldr.CreateString(c.ModelName)
	stepping := bldr.CreateString(c.Stepping)
	microcode := bldr.CreateString(c.Microcode)
	cacheSize := bldr.CreateString(c.CacheSize)
	uoffs := make([]fb.UOffsetT, len(c.Flags))
	for i, flag := range c.Flags {
		uoffs[i] = bldr.CreateString(flag)
	}
	structs.ProcessorStartFlagsVector(bldr, len(uoffs))
	for i := len(uoffs) - 1; i >= 0; i-- {
		bldr.PrependUOffsetT(uoffs[i])
	}
	flags := bldr.EndVector(len(uoffs))
	structs.ProcessorStart(bldr)
	structs.ProcessorAddPhysicalID(bldr, int32(c.PhysicalID))
	structs.ProcessorAddVendorID(bldr, vendorID)
	structs.ProcessorAddCPUFamily(bldr, cpuFamily)
	structs.ProcessorAddModel(bldr, model)
	structs.ProcessorAddModelName(bldr, modelName)
	structs.ProcessorAddStepping(bldr, stepping)
	structs.ProcessorAddMicrocode(bldr, microcode)
	structs.ProcessorAddCPUMHz(bldr, c.CPUMHz)
	structs.ProcessorAddCacheSize(bldr, cacheSize)
	structs.ProcessorAddCPUCores(bldr, int32(c.CPUCores))
	structs.ProcessorAddFlags(bldr, flags)
	return structs.ProcessorEnd(bldr)
}

// Deserialize deserializes bytes representing a flatbuffer serialized System
// as *sysinformation.System.  If you want to work with System directly (the
// flatbuffers version), use GetRootAsSystem instead.
func Deserialize(p []byte) *sysinfo.System {
	var sys sysinfo.System
	s := structs.GetRootAsSystem(p, 0)
	procF := &structs.Processor{}
	proc := &sysinfo.Processor{}
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
	sys.Socket = make([]*sysinfo.Processor, s.SocketLength())
	for i := 0; i < len(sys.Socket); i++ {
		if !s.Socket(procF, i) {
			continue
		}
		proc.PhysicalID = procF.PhysicalID()
		proc.VendorID = string(procF.VendorID())
		proc.CPUFamily = string(procF.CPUFamily())
		proc.Model = string(procF.Model())
		proc.ModelName = string(procF.ModelName())
		proc.Stepping = string(procF.Stepping())
		proc.Microcode = string(procF.Microcode())
		proc.CPUMHz = procF.CPUMHz()
		proc.CacheSize = string(procF.CacheSize())
		proc.CPUCores = procF.CPUCores()
		proc.Flags = make([]string, procF.FlagsLength())
		for j := 0; j < len(proc.Flags); j++ {
			proc.Flags[j] = string(procF.Flags(j))
		}
		sys.Socket[i] = proc
	}
	return &sys
}
