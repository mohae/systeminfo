# systeminfo
Information about a system: CPU, RAM, network interfaces, OS, Kernel.  This only works on Linux, for now.

Only basic information about a system is gathered.  For more detailed informatino about different aspects of a system, use the relevant [JoeFriday](https://github.com/mohae/joefriday) packages.

Convenience functions are provided for Marshaling and Unmarshaling the data using either `JSON` or `ProtoBuf` serialization.

## Supported output
* Go struct
* JSON
* Protocol Buffers
* Flatbuffers

## TODO
Add disk information.  
Add support for other platforms.  

