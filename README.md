# Osquery-memory-forensics

## Download bins
1. Download [ProcDump](https://docs.microsoft.com/en-us/sysinternals/downloads/procdump)
1. Download [WinPmem](https://github.com/Velocidex/c-aff4/releases/download/v3.3.rc3/winpmem_v3.3.rc3.exe)
1. Copy binaries do `bins/`
1. `go get -u github.com/go-bindata/go-bindata/...`
1. `go install github.com/go-bindata/go-bindata/...`
1. `~/go/bin/go-bindata -o assets/bindata.go -pkg assets bins/...`
1. `ls -lh assets/bindata.go`

## Setup dev env
1. `go mod init github.com/CptOfEvilMinions/osquery-memory-forensics`
1. `go get`


## References
* [Combine absolute path and relative path to get a new absolute path](https://stackoverflow.com/questions/13078314/combine-absolute-path-and-relative-path-to-get-a-new-absolute-path)
* [Go by Example: Epoch](https://gobyexample.com/epoch)
* [PsExec gets stuck on licence prompt when running non-interactively](https://stackoverflow.com/questions/5151034/psexec-gets-stuck-on-licence-prompt-when-running-non-interactively)
* [Package strconv](https://golang.org/pkg/strconv/)
* [ProcDump v9.0](https://docs.microsoft.com/en-us/sysinternals/downloads/procdump)
* [Go and file perms on Windows](https://medium.com/@MichalPristas/go-and-file-perms-on-windows-3c944d55dd44)
* [Hash checksums: MD5, SHA-1, SHA-256](https://yourbasic.org/golang/hash-md5-sha256-string-file/)
* [Github - getlantern/byteexec](https://github.com/getlantern/byteexec)
* [package byteexec](https://pkg.go.dev/github.com/getlantern/byteexec?tab=doc#Exec)
* [Embedding data in Go executables](https://scene-si.org/2017/08/22/embedding-data-in-go-executables/)
* [Github - go-bindata/go-bindata](https://github.com/go-bindata/go-bindata)
* []()
* []()
* []()
* []()
* []()