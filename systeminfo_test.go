package systeminfo

import "testing"

// IDLike isn't checked becuase it's not available on all platforms
// PhysicalId isn't checked becuase 0 is a valid value.
func TestSystemInfoGet(t *testing.T) {
	var s System
	err := s.Get()
	if err != nil {
		t.Error(err)
		return
	}
	if s.KernelOS == "" {
		t.Error("KernelOS: was empty, expected something")
	}
	if s.KernelVersion == "" {
		t.Error("KernelVersion: was empty, expected something")
	}
	if s.KernelArch == "" {
		t.Error("KernelArch: was empty, expected something")
	}
	if s.KernelType == "" {
		t.Error("KernelType: was empty, expected something")
	}
	if s.KernelCompileDate == "" {
		t.Error("KernelCompileDate: was empty, expected something")
	}
	if s.OSName == "" {
		t.Error("OSName: was empty, expected something")
	}
	if s.OSID == "" {
		t.Error("OSID: was empty, expected something")
	}
	if s.OSVersion == "" {
		t.Error("OSVersion: was empty, expected something")
	}
	if s.MemTotal == 0 {
		t.Error("MemTotal: was 0, expected a non-zero value")
	}
	if s.SwapTotal == 0 {
		t.Error("SwapTotal: was 0, expected a non-zero value")
	}
	if len(s.NetInfs) == 0 {
		t.Error("Len NetInfs was 0, expected something")
	}
	if len(s.Chips) == 0 {
		t.Error("Len Chips was 0, expected something")
	} else {
		for i, chip := range s.Chips {
			if chip.VendorID == "" {
				t.Errorf("%d: VendorID was empty, expected something", i)
			}
			if chip.CPUFamily == "" {
				t.Errorf("%d: CPUFamily was empty, expected something", i)
			}
			if chip.Model == "" {
				t.Errorf("%d: Model was empty, expected something", i)
			}
			if chip.ModelName == "" {
				t.Errorf("%d: ModelName was empty, expected something", i)
			}
			if chip.Stepping == "" {
				t.Errorf("%d: Stepping was empty, expected something", i)
			}
			if chip.Microcode == "" {
				t.Errorf("%d: Microcode was empty, expected something", i)
			}
			if chip.CPUMHz == 0 {
				t.Errorf("%d: CPUMHz was 0, expected non-zero value", i)
			}
			if chip.CacheSize == "" {
				t.Errorf("%d: CacheSize was empty, expected something", i)
			}
			if chip.CPUCores == 0 {
				t.Errorf("%d: CPUCores was 0, expected non-zero value", i)
			}
			if len(chip.Flags) == 0 {
				t.Errorf("%d: Flags was 0, expected non-zero value", i)
			}
		}
	}
}

func TestSystemInfoJSON(t *testing.T) {
	var s System
	err := s.Get()
	if err != nil {
		t.Error(err)
		return
	}

	p, err := s.JSONMarshal()
	if err != nil {
		t.Error(err)
		return
	}
	u, err := JSONUnmarshal(p)
	if err != nil {
		t.Error(err)
		return
	}

	Compare("TestSystemInfoJSON", s, *u, t)
}

func TestSystemInfoProto(t *testing.T) {
	var s System
	err := s.Get()
	if err != nil {
		t.Error(err)
		return
	}

	p, err := s.ProtoMarshal()
	if err != nil {
		t.Error(err)
		return
	}
	u, err := ProtoUnmarshal(p)
	if err != nil {
		t.Error(err)
		return
	}

	Compare("TestSystemInfoProto", s, *u, t)
}

func Compare(name string, s, u System, t *testing.T) {
	if s.KernelOS != u.KernelOS {
		t.Errorf("%s: KernelOS: got %s; want %s", name, u.KernelOS, s.KernelOS)
	}
	if s.KernelVersion != u.KernelVersion {
		t.Errorf("%s: KernelVersion: got %s; want %s", name, u.KernelVersion, s.KernelVersion)
	}
	if s.KernelArch != u.KernelArch {
		t.Errorf("%s: KernelArch: got %s; want %s", name, u.KernelArch, s.KernelArch)
	}
	if s.KernelType != u.KernelType {
		t.Errorf("%s: KernelType: got %s; want %s", name, u.KernelType, s.KernelType)
	}
	if s.KernelCompileDate != u.KernelCompileDate {
		t.Errorf("%s: KernelCompileDate: got %s; want %s", name, u.KernelCompileDate, s.KernelCompileDate)
	}
	if s.OSName != u.OSName {
		t.Errorf("%s: OSName: got %s; want %s", name, u.OSName, s.OSName)
	}
	if s.OSID != u.OSID {
		t.Errorf("%s: OSID: got %s; want %s", name, u.OSID, s.OSID)
	}
	if s.OSVersion != u.OSVersion {
		t.Errorf("%s: OSVersion: got %s; want %s", name, u.OSVersion, s.OSVersion)
	}
	if s.MemTotal != u.MemTotal {
		t.Errorf("%s: MemTotal: got %d; want %d", name, u.MemTotal, s.MemTotal)
	}
	if s.SwapTotal != u.SwapTotal {
		t.Errorf("%s: SwapTotal: got %d; want %d", name, u.SwapTotal, s.SwapTotal)
	}
	if len(s.NetInfs) != len(u.NetInfs) {
		t.Errorf("%s: Len NetInfs: got %d; want %d", name, len(u.NetInfs), len(s.NetInfs))
	} else {
		for i, v := range s.NetInfs {
			if v != u.NetInfs[i] {
				t.Errorf("%s: Netinfs[%d]: got %s, want %s", name, i, u.NetInfs[i], v)
			}
		}
	}
	if len(s.Chips) != len(u.Chips) {
		t.Errorf("%s: Len Chips: got %d, want %d", name, len(u.Chips), len(s.Chips))
	} else {
		for i, v := range s.Chips {
			if v.VendorID != u.Chips[i].VendorID {
				t.Errorf("%s: Chips[%d].VendorID: got %s, want %s", name, i, v.VendorID, u.Chips[i].VendorID)
			}
			if v.CPUFamily != u.Chips[i].CPUFamily {
				t.Errorf("%s: Chips[%d].CPUFamily: got %s, want %s", name, i, v.CPUFamily, u.Chips[i].CPUFamily)
			}
			if v.Model != u.Chips[i].Model {
				t.Errorf("%s: Chips[%d].Model: got %s, want %s", name, i, v.Model, u.Chips[i].Model)
			}
			if v.ModelName != u.Chips[i].ModelName {
				t.Errorf("%s: Chips[%d].ModelName: got %s, want %s", name, i, v.ModelName, u.Chips[i].ModelName)
			}
			if v.Stepping != u.Chips[i].Stepping {
				t.Errorf("%s: Chips[%d].Stepping: got %s, want %s", name, i, v.Stepping, u.Chips[i].Stepping)
			}
			if v.Microcode != u.Chips[i].Microcode {
				t.Errorf("%s: Chips[%d].Microcode: got %s, want %s", name, i, v.Microcode, u.Chips[i].Microcode)
			}
			if v.CPUMHz != u.Chips[i].CPUMHz {
				t.Errorf("%s: Chips[%d].CPUMHz: got %g, want %g", name, i, v.CPUMHz, u.Chips[i].CPUMHz)
			}
			if v.CacheSize != u.Chips[i].CacheSize {
				t.Errorf("%s: Chips[%d].CacheSize: got %s, want %s", name, i, v.CacheSize, u.Chips[i].CacheSize)
			}
			if v.CPUCores != u.Chips[i].CPUCores {
				t.Errorf("%s: Chips[%d].CPUCores: got %d, want %d", name, i, v.CPUCores, u.Chips[i].CPUCores)
			}
			if len(v.Flags) != len(u.Chips[i].Flags) {
				t.Errorf("%s: len Chips[%d].Flags: got %d, want %d", name, i, len(v.Flags), len(u.Chips[i].Flags))
			}

		}
	}

}
