package exporters

import (
	"os/exec"

	"github.com/FUNExtreme/barkup"
)

// FileSystem is an `Exporter` interface that backs up a folder on disk
type FileSystem struct {
	// Path (e.g. /var/backup)
	Path string
}

// Export produces a `mysqldump` of the specified database, and creates a gzip compressed tarball archive.
func (x FileSystem) Export(filename string) *barkup.ExportResult {
	result := &barkup.ExportResult{MIME: "application/x-tar"}

	result.Path = filename + ".tar.gz"
	out, err := exec.Command(barkup.TarCmd, "-czf", result.Path, x.Path).Output()
	if err != nil {
		result.Error = barkup.MakeErr(err, string(out))
		return result
	}

	return result
}
