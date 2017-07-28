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
	if len(s.Socket) != len(u.Socket) {
		t.Errorf("Len Socket: got %d, want %d", len(u.Socket), len(s.Socket))
	} else {
		for i, v := range s.Socket {
			if v.VendorID != u.Socket[i].VendorID {
				t.Errorf("Socket[%d].VendorID: got %s, want %s", i, v.VendorID, u.Socket[i].VendorID)
			}
			if v.CPUFamily != u.Socket[i].CPUFamily {
				t.Errorf("Socket[%d].CPUFamily: got %s, want %s", i, v.CPUFamily, u.Socket[i].CPUFamily)
			}
			if v.Model != u.Socket[i].Model {
				t.Errorf("Socket[%d].Model: got %s, want %s", i, v.Model, u.Socket[i].Model)
			}
			if v.ModelName != u.Socket[i].ModelName {
				t.Errorf("Socket[%d].ModelName: got %s, want %s", i, v.ModelName, u.Socket[i].ModelName)
			}
			if v.Stepping != u.Socket[i].Stepping {
				t.Errorf("Socket[%d].Stepping: got %s, want %s", i, v.Stepping, u.Socket[i].Stepping)
			}
			if v.Microcode != u.Socket[i].Microcode {
				t.Errorf("Socket[%d].Microcode: got %s, want %s", i, v.Microcode, u.Socket[i].Microcode)
			}
			if v.CPUMHz != u.Socket[i].CPUMHz {
				t.Errorf("Socket[%d].CPUMHz: got %g, want %g", i, v.CPUMHz, u.Socket[i].CPUMHz)
			}
			if v.BogoMIPS != u.Socket[i].BogoMIPS {
				t.Errorf("Socket[%d].BogoMIPS: got %g, want %g", i, v.BogoMIPS, u.Socket[i].BogoMIPS)
			}
			if v.CacheSize != u.Socket[i].CacheSize {
				t.Errorf("Socket[%d].CacheSize: got %s, want %s", i, v.CacheSize, u.Socket[i].CacheSize)
			}
			if v.CPUCores != u.Socket[i].CPUCores {
				t.Errorf("Socket[%d].CPUCores: got %d, want %d", i, v.CPUCores, u.Socket[i].CPUCores)
			}
			if len(v.Flags) != len(u.Socket[i].Flags) {
				t.Errorf("len Socket[%d].Flags: got %d, want %d", i, len(v.Flags), len(u.Socket[i].Flags))
			}

		}
	}
}
