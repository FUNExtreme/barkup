package exporters

import (
	"fmt"
	"os/exec"
	"time"

	"github.com/FUNExtreme/barkup"
)

// FileSystem is an `Exporter` interface that backs up a folder on disk
type FileSystem struct {
	// Path (e.g. /var/backup)
	Path string
}

// Export produces a `mysqldump` of the specified database, and creates a gzip compressed tarball archive.
func (x FileSystem) Export() *barkup.ExportResult {
	result := &barkup.ExportResult{MIME: "application/x-tar"}

	dumpPath := fmt.Sprintf(`bu_%v_%v.sql`, x.Path, time.Now().Unix())

	result.Path = dumpPath + ".tar.gz"
	out, err := exec.Command(barkup.TarCmd, "-czf", result.Path, x.Path).Output()
	if err != nil {
		result.Error = barkup.MakeErr(err, string(out))
		return result
	}

	return result
}
