package exes

import (
	"github.com/CptOfEvilMinions/osquery-memory-forensics/assets"
	"github.com/CptOfEvilMinions/osquery-memory-forensics/pkg/hash"
	"github.com/getlantern/byteexec"
)

// DumpExecutablesToDisk input: None
// Extracts executables from go-bindata module and writes them to
// disk located at `%APPDATA%\\byteexec\\*.exe`
// DumpExecutablesToDisk output: Returns result and error (if any)
func DumpExecutablesToDisk(winPmemExecutable *byteexec.Exec, procDumpExecutable *byteexec.Exec, verification int, winAppDataDirPath string) (bool, error) {
	// Extract ProdDump executable byte array
	procDumpByteData, err := assets.Asset("bins/procdump.exe")
	if err != nil {
		return false, err
	}

	// Write executable to disk
	procDumpExecutable, err = byteexec.New(procDumpByteData, "procdump")
	if err != nil {
		return false, err
	}

	// Extract WinPmem executable byte array
	WinPmemData, err := assets.Asset("bins/winpmem.exe")
	if err != nil {
		return false, err
	}

	// Write executable to disk
	winPmemExecutable, err = byteexec.New(WinPmemData, "winpmem")
	if err != nil {
		return false, err
	}

	// Verify Binary
	if verification == 1 {
		status, err := hash.VerifyBinaries(winAppDataDirPath)
		if err != nil {
			return status, err
		}
	}

	// Return pointers to executablse
	return true, nil

}