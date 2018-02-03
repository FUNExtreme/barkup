package barkup

import (
	"log"

	"github.com/minio/minio-go"
)

// DigitalOcean is a `Storer` interface that puts an ExportResult to the specified DigitalOcean space.
type DigitalOcean struct {
	// Available regions:
	// * nyc3
	// * ams3
	// * sgp1
	Region string
	// Name of the space
	Space string
	// DigitalOcean Spaces access key
	AccessKey string
	// DigitalOcean Spaces access secret
	ClientSecret string
}

// Store puts an `ExportResult` struct to an DigitalOcean space within the specified directory
func (x *DigitalOcean) Store(result *ExportResult, directory string) *Error {

	if result.Error != nil {
		return result.Error
	}

	ssl := true
	client, err := minio.New(x.Region+".digitaloceanspaces.com", x.AccessKey, x.ClientSecret, ssl)
	if err != nil {
		log.Fatal(err)
	}

	_, err = client.FPutObject(x.Space, result.Filename(), result.Path, minio.PutObjectOptions{ContentType: result.MIME})
	if err != nil {
		log.Fatalln(err)
	}

	return makeErr(err, "")
}
