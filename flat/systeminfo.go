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
	SystemStartNetInfsVector(bldr, len(inf.Device))
	for i := len(uof) - 1; i >= 0; i-- {
		bldr.PrependUOffsetT(uof[i])
	}
	netinfs := bldr.EndVector(len(uof))
	// Get processors
	p, err := processors.Get()
	if err != nil {
		return nil, sysinfo.Error{Op: "flat processor info", Err: err}
	}

	uoff := make([]fb.UOffsetT, len(p.Socket))
	for i, processor := range p.Socket {
		uoff[i] = serializeProcessor(bldr, &processor)
	}
	SystemStartChipsVector(bldr, len(uoff))
	for i := len(uoff) - 1; i >= 0; i-- {
		bldr.PrependUOffsetT(uoff[i])
	}
	chips := bldr.EndVector(len(uoff))
	SystemStart(bldr)
	SystemAddKernelOS(bldr, osS)
	SystemAddKernelVersion(bldr, version)
	SystemAddKernelArch(bldr, arch)
	SystemAddKernelCompileDate(bldr, compileDate)
	SystemAddOSName(bldr, oName)
	SystemAddOSID(bldr, oID)
	SystemAddOSIDLike(bldr, oIDLike)
	SystemAddOSVersion(bldr, oVersion)
	SystemAddMemTotal(bldr, m.MemTotal)
	SystemAddSwapTotal(bldr, m.SwapTotal)
	SystemAddNetInfs(bldr, netinfs)
	SystemAddChips(bldr, chips)
	bldr.Finish(SystemEnd(bldr))
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
	ChipStartFlagsVector(bldr, len(uoffs))
	for i := len(uoffs) - 1; i >= 0; i-- {
		bldr.PrependUOffsetT(uoffs[i])
	}
	flags := bldr.EndVector(len(uoffs))
	ChipStart(bldr)
	ChipAddPhysicalID(bldr, int32(c.PhysicalID))
	ChipAddVendorID(bldr, vendorID)
	ChipAddCPUFamily(bldr, cpuFamily)
	ChipAddModel(bldr, model)
	ChipAddModelName(bldr, modelName)
	ChipAddStepping(bldr, stepping)
	ChipAddMicrocode(bldr, microcode)
	ChipAddCPUMHz(bldr, c.CPUMHz)
	ChipAddCacheSize(bldr, cacheSize)
	ChipAddCPUCores(bldr, int32(c.CPUCores))
	ChipAddFlags(bldr, flags)
	return ChipEnd(bldr)
}

// Deserialize deserializes bytes representing a flatbuffer serialized System
// as *sysinformation.System.  If you want to work with System directly (the
// flatbuffers version), use GetRootAsSystem instead.
func Deserialize(p []byte) *sysinfo.System {
	var sys sysinfo.System
	s := GetRootAsSystem(p, 0)
	c := &Chip{}
	chip := &sysinfo.Chip{}
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
	sys.NetInfs = make([]string, s.NetInfsLength())
	for i := 0; i < len(sys.NetInfs); i++ {
		sys.NetInfs[i] = string(s.NetInfs(i))
	}
	sys.Chips = make([]*sysinfo.Chip, s.ChipsLength())
	for i := 0; i < len(sys.Chips); i++ {
		if !s.Chips(c, i) {
			continue
		}
		chip.PhysicalID = c.PhysicalID()
		chip.VendorID = string(c.VendorID())
		chip.CPUFamily = string(c.CPUFamily())
		chip.Model = string(c.Model())
		chip.ModelName = string(c.ModelName())
		chip.Stepping = string(c.Stepping())
		chip.Microcode = string(c.Microcode())
		chip.CPUMHz = c.CPUMHz()
		chip.CacheSize = string(c.CacheSize())
		chip.CPUCores = c.CPUCores()
		chip.Flags = make([]string, c.FlagsLength())
		for j := 0; j < len(chip.Flags); j++ {
			chip.Flags[j] = string(c.Flags(j))
		}
		sys.Chips[i] = chip
	}
	return &sys
}
