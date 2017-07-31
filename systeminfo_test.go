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
	if s.Hostname == "" {
		t.Errorf("Hostname: was empty; expected something")
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
	if len(s.NetDev) == 0 {
		t.Error("Len NetInfs was 0, expected something")
	}
	if s.Sockets <= 0 {
		t.Errorf("expected at least 1 socket; got %d", s.Sockets)
	}
	if len(s.CPU) == 0 {
		t.Error("Len Socket was 0, expected something")
	} else {
		for i, cpu := range s.CPU {
			if cpu.VendorID == "" {
				t.Errorf("%d: VendorID was empty, expected something", i)
			}
			if cpu.CPUFamily == "" {
				t.Errorf("%d: CPUFamily was empty, expected something", i)
			}
			if cpu.Model == "" {
				t.Errorf("%d: Model was empty, expected something", i)
			}
			if cpu.ModelName == "" {
				t.Errorf("%d: ModelName was empty, expected something", i)
			}
			if cpu.Stepping == "" {
				t.Errorf("%d: Stepping was empty, expected something", i)
			}
			if cpu.Microcode == "" {
				t.Errorf("%d: Microcode was empty, expected something", i)
			}
			if cpu.CPUMHz == 0 {
				t.Errorf("%d: CPUMHz was 0, expected non-zero value", i)
			}
			if cpu.CacheSize == "" {
				t.Errorf("%d: CacheSize was empty, expected something", i)
			}
			if cpu.CPUCores == 0 {
				t.Errorf("%d: CPUCores was 0, expected non-zero value", i)
			}
			if len(cpu.Flags) == 0 {
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
	if len(s.NetDev) != len(u.NetDev) {
		t.Errorf("%s: Len NetInfs: got %d; want %d", name, len(u.NetDev), len(s.NetDev))
	} else {
		for i, v := range s.NetDev {
			if v != u.NetDev[i] {
				t.Errorf("%s: Netinfs[%d]: got %s, want %s", name, i, u.NetDev[i], v)
			}
		}
	}
	if len(s.CPU) != len(u.CPU) {
		t.Errorf("%s: Len CPU: got %d, want %d", name, len(u.CPU), len(s.CPU))
	} else {
		for i, v := range s.CPU {
			if v.VendorID != u.CPU[i].VendorID {
				t.Errorf("%s: CPU[%d].VendorID: got %s, want %s", name, i, v.VendorID, u.CPU[i].VendorID)
			}
			if v.CPUFamily != u.CPU[i].CPUFamily {
				t.Errorf("%s: CPU[%d].CPUFamily: got %s, want %s", name, i, v.CPUFamily, u.CPU[i].CPUFamily)
			}
			if v.Model != u.CPU[i].Model {
				t.Errorf("%s: CPU[%d].Model: got %s, want %s", name, i, v.Model, u.CPU[i].Model)
			}
			if v.ModelName != u.CPU[i].ModelName {
				t.Errorf("%s: CPU[%d].ModelName: got %s, want %s", name, i, v.ModelName, u.CPU[i].ModelName)
			}
			if v.Stepping != u.CPU[i].Stepping {
				t.Errorf("%s: CPU[%d].Stepping: got %s, want %s", name, i, v.Stepping, u.CPU[i].Stepping)
			}
			if v.Microcode != u.CPU[i].Microcode {
				t.Errorf("%s: CPU[%d].Microcode: got %s, want %s", name, i, v.Microcode, u.CPU[i].Microcode)
			}
			if v.CPUMHz != u.CPU[i].CPUMHz {
				t.Errorf("%s: CPU[%d].CPUMHz: got %g, want %g", name, i, v.CPUMHz, u.CPU[i].CPUMHz)
			}
			if v.BogoMIPS != u.CPU[i].BogoMIPS {
				t.Errorf("%s: CPU[%d].BogoMIPS: got %g, want %g", name, i, v.BogoMIPS, u.CPU[i].BogoMIPS)
			}
			if v.CacheSize != u.CPU[i].CacheSize {
				t.Errorf("%s: CPU[%d].CacheSize: got %s, want %s", name, i, v.CacheSize, u.CPU[i].CacheSize)
			}
			if v.CPUCores != u.CPU[i].CPUCores {
				t.Errorf("%s: CPU[%d].CPUCores: got %d, want %d", name, i, v.CPUCores, u.CPU[i].CPUCores)
			}
			if len(v.Flags) != len(u.CPU[i].Flags) {
				t.Errorf("%s: len CPU[%d].Flags: got %d, want %d", name, i, len(v.Flags), len(u.CPU[i].Flags))
			}

		}
	}

}
