package systeminfo

import (
	"testing"

	sysinfo "github.com/mohae/systeminfo"
)

func TestGet(t *testing.T) {
	var s sysinfo.System
	err := s.Get()
	if err != nil {
		t.Errorf("unexpected error: %s", err)
		return
	}
	b, err := Get()
	if err != nil {
		t.Errorf("unexpected error: %s", err)
		return
	}
	u := Deserialize(b)
	if s.KernelOS != u.KernelOS {
		t.Errorf("KernelOS: got %s; want %s", u.KernelOS, s.KernelOS)
	}
	if s.KernelVersion != u.KernelVersion {
		t.Errorf("KernelVersion: got %s; want %s", u.KernelVersion, s.KernelVersion)
	}
	if s.KernelArch != u.KernelArch {
		t.Errorf("KernelArch: got %s; want %s", u.KernelArch, s.KernelArch)
	}
	if s.KernelCompileDate != u.KernelCompileDate {
		t.Errorf("KernelCompileDate: got %s; want %s", u.KernelCompileDate, s.KernelCompileDate)
	}
	if s.OSName != u.OSName {
		t.Errorf("OSName: got %s; want %s", u.OSName, s.OSName)
	}
	if s.OSID != u.OSID {
		t.Errorf("OSID: got %s; want %s", u.OSID, s.OSID)
	}
	if s.OSVersion != u.OSVersion {
		t.Errorf("OSVersion: got %s; want %s", u.OSVersion, s.OSVersion)
	}
	if s.MemTotal != u.MemTotal {
		t.Errorf("MemTotal: got %d; want %d", u.MemTotal, s.MemTotal)
	}
	if s.SwapTotal != u.SwapTotal {
		t.Errorf("SwapTotal: got %d; want %d", u.SwapTotal, s.SwapTotal)
	}
	if len(s.NetDev) != len(u.NetDev) {
		t.Errorf("Len NetInfs: got %d; want %d", len(u.NetDev), len(s.NetDev))
	} else {
		for i, v := range s.NetDev {
			if v != u.NetDev[i] {
				t.Errorf("Netinfs[%d]: got %s, want %s", i, u.NetDev[i], v)
			}
		}
	}
	if s.Sockets != u.Sockets {
		t.Errorf("Sockets: got %d; want %d", u.Sockets, s.Sockets)
	}
	if len(s.CPU) != len(u.CPU) {
		t.Errorf("Len CPU: got %d, want %d", len(u.CPU), len(s.CPU))
	} else {
		for i, v := range s.CPU {
			if v.PhysicalID != u.CPU[i].PhysicalID {
				t.Errorf("CPU[%d].PhysicalID: got %s, want %s", i, v.PhysicalID, u.CPU[i].PhysicalID)
			}
			if v.CoreID != u.CPU[i].CoreID {
				t.Errorf("CPU[%d].CoreID: got %s, want %s", i, v.CoreID, u.CPU[i].CoreID)
			}
			if v.Siblings != u.CPU[i].Siblings {
				t.Errorf("CPU[%d].Siblings: got %s, want %s", i, v.Siblings, u.CPU[i].Siblings)
			}
			if v.VendorID != u.CPU[i].VendorID {
				t.Errorf("CPU[%d].VendorID: got %s, want %s", i, v.VendorID, u.CPU[i].VendorID)
			}
			if v.CPUFamily != u.CPU[i].CPUFamily {
				t.Errorf("CPU[%d].CPUFamily: got %s, want %s", i, v.CPUFamily, u.CPU[i].CPUFamily)
			}
			if v.Model != u.CPU[i].Model {
				t.Errorf("CPU[%d].Model: got %s, want %s", i, v.Model, u.CPU[i].Model)
			}
			if v.ModelName != u.CPU[i].ModelName {
				t.Errorf("CPU[%d].ModelName: got %s, want %s", i, v.ModelName, u.CPU[i].ModelName)
			}
			if v.Stepping != u.CPU[i].Stepping {
				t.Errorf("CPU[%d].Stepping: got %s, want %s", i, v.Stepping, u.CPU[i].Stepping)
			}
			if v.Microcode != u.CPU[i].Microcode {
				t.Errorf("CPU[%d].Microcode: got %s, want %s", i, v.Microcode, u.CPU[i].Microcode)
			}
			if v.CPUMHz != u.CPU[i].CPUMHz {
				t.Errorf("CPU[%d].CPUMHz: got %g, want %g", i, v.CPUMHz, u.CPU[i].CPUMHz)
			}
			if v.BogoMIPS != u.CPU[i].BogoMIPS {
				t.Errorf("CPU[%d].BogoMIPS: got %g, want %g", i, v.BogoMIPS, u.CPU[i].BogoMIPS)
			}
			if v.CacheSize != u.CPU[i].CacheSize {
				t.Errorf("CPU[%d].CacheSize: got %s, want %s", i, v.CacheSize, u.CPU[i].CacheSize)
			}
			if v.CPUCores != u.CPU[i].CPUCores {
				t.Errorf("CPU[%d].CPUCores: got %d, want %d", i, v.CPUCores, u.CPU[i].CPUCores)
			}
			if len(v.Flags) != len(u.CPU[i].Flags) {
				t.Errorf("len CPU[%d].Flags: got %d, want %d", i, len(v.Flags), len(u.CPU[i].Flags))
			}
			if len(v.Bugs) != len(u.CPU[i].Bugs) {
				t.Errorf("len CPU[%d].Bugs: got %d, want %d", i, len(v.Bugs), len(u.CPU[i].Bugs))
			}
		}
	}
}
