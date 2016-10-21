# systeminfo
[![GoDoc](https://godoc.org/github.com/mohae/systeminfo.svg)](https://godoc.org/github.com/mohae/systeminfo)[![Build Status](https://travis-ci.org/mohae/systeminfo.png)](https://travis-ci.org/mohae/systeminfo)

Information about a system: CPU, RAM, network interfaces, OS, Kernel.  This only works on Linux, for now.

Only basic information about a system is gathered.  For more detailed informatino about different aspects of a system, use the relevant [JoeFriday](https://github.com/mohae/joefriday) packages.

Convenience functions are provided for Marshaling and Unmarshaling the data using either `JSON` or `ProtoBuf` serialization.

## Supported output
* Go struct
* JSON
* Protocol Buffers
* Flatbuffers

## Notes
If using the flatbuffers implementation, the package name is still `systeminfo` even though the import path is `github.com/mohae/systeminfo/flat`.  If one is also using the `github.com/mohae/systeminfo.System` struct one of the packages will need to be aliased.  If this ends up being problematic, it may be changed in the future.  (In typing this I'm becoming less sure of this decision, but it will remain for now.)

## TODO
Add disk information.  
Add support for other platforms.  

