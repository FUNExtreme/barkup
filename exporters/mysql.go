package exporters

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/FUNExtreme/barkup"
)

// MySQL is an `Exporter` interface that backs up a MySQL database via the `mysqldump` command
type MySQL struct {
	// DB Host (e.g. 127.0.0.1)
	Host string
	// DB Port (e.g. 3306)
	Port string
	// DB Name
	DB string
	// DB User
	User string
	// DB Password
	Password string
	// Extra mysqldump options
	// e.g []string{"--extended-insert"}
	Options []string
}

// Export produces a `mysqldump` of the specified database, and creates a gzip compressed tarball archive.
func (x MySQL) Export(filename string) *barkup.ExportResult {
	result := &barkup.ExportResult{MIME: "application/gzip"}

	dumpPath := filename + ".sql"

	options := append(x.dumpOptions(), fmt.Sprintf(`-r%v`, dumpPath))
	out, err := exec.Command(barkup.MysqlDumpCmd, options...).Output()
	if err != nil {
		result.Error = barkup.MakeErr(err, string(out))
		return result
	}

	result.Path = dumpPath + ".tar.gz"
	_, err = exec.Command(barkup.TarCmd, "-czf", result.Path, dumpPath).Output()
	if err != nil {
		result.Error = barkup.MakeErr(err, string(out))
		return result
	}
	os.Remove(dumpPath)

	return result
}

func (x MySQL) dumpOptions() []string {
	options := x.Options
	options = append(options, fmt.Sprintf(`-h%v`, x.Host))
	options = append(options, fmt.Sprintf(`-P%v`, x.Port))
	options = append(options, fmt.Sprintf(`-u%v`, x.User))
	if x.Password != "" {
		options = append(options, fmt.Sprintf(`-p%v`, x.Password))
	}
	options = append(options, x.DB)
	return options
}
