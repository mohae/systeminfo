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
	if len(s.NetInfs) != len(u.NetInfs) {
		t.Errorf("Len NetInfs: got %d; want %d", len(u.NetInfs), len(s.NetInfs))
	} else {
		for i, v := range s.NetInfs {
			if v != u.NetInfs[i] {
				t.Errorf("Netinfs[%d]: got %s, want %s", i, u.NetInfs[i], v)
			}
		}
	}
	if len(s.Chips) != len(u.Chips) {
		t.Errorf("Len Chips: got %d, want %d", len(u.Chips), len(s.Chips))
	} else {
		for i, v := range s.Chips {
			if v.VendorID != u.Chips[i].VendorID {
				t.Errorf("Chips[%d].VendorID: got %s, want %s", i, v.VendorID, u.Chips[i].VendorID)
			}
			if v.CPUFamily != u.Chips[i].CPUFamily {
				t.Errorf("Chips[%d].CPUFamily: got %s, want %s", i, v.CPUFamily, u.Chips[i].CPUFamily)
			}
			if v.Model != u.Chips[i].Model {
				t.Errorf("Chips[%d].Model: got %s, want %s", i, v.Model, u.Chips[i].Model)
			}
			if v.ModelName != u.Chips[i].ModelName {
				t.Errorf("Chips[%d].ModelName: got %s, want %s", i, v.ModelName, u.Chips[i].ModelName)
			}
			if v.Stepping != u.Chips[i].Stepping {
				t.Errorf("Chips[%d].Stepping: got %s, want %s", i, v.Stepping, u.Chips[i].Stepping)
			}
			if v.Microcode != u.Chips[i].Microcode {
				t.Errorf("Chips[%d].Microcode: got %s, want %s", i, v.Microcode, u.Chips[i].Microcode)
			}
			if v.CPUMHz != u.Chips[i].CPUMHz {
				t.Errorf("Chips[%d].CPUMHz: got %g, want %g", i, v.CPUMHz, u.Chips[i].CPUMHz)
			}
			if v.CacheSize != u.Chips[i].CacheSize {
				t.Errorf("Chips[%d].CacheSize: got %s, want %s", i, v.CacheSize, u.Chips[i].CacheSize)
			}
			if v.CPUCores != u.Chips[i].CPUCores {
				t.Errorf("Chips[%d].CPUCores: got %d, want %d", i, v.CPUCores, u.Chips[i].CPUCores)
			}
			if len(v.Flags) != len(u.Chips[i].Flags) {
				t.Errorf("len Chips[%d].Flags: got %d, want %d", i, len(v.Flags), len(u.Chips[i].Flags))
			}

		}
	}
}
