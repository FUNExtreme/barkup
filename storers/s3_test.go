package storers

import (
	"os"
	"testing"

	"github.com/FUNExtreme/barkup"
	"launchpad.net/goamz/aws"
)

func Test_S3_Store_Success(t *testing.T) {
	server := testServer(200, "", "application/json")
	aws.Regions["fake-region"] = aws.Region{
		Name:       "fake-region",
		S3Endpoint: server.URL,
	}

	s := &S3{
		Region:       "fake-region",
		Bucket:       "cheese",
		AccessKey:    "afjlsdfjkaf",
		ClientSecret: "adsfljdsahfl",
	}

	err := s.Store(&ExportResult{"test/test.txt", "text/plain", nil}, "test/")
	expect(t, err, (*barkup.Error)(nil))
}

func Test_S3_Store_Fail(t *testing.T) {
	server := testServer(500, "", "application/json")
	aws.Regions["fake-region"] = aws.Region{
		Name:       "fake-region",
		S3Endpoint: server.URL,
	}

	s := &S3{
		Region:       "fake-region",
		Bucket:       "cheese",
		AccessKey:    "afjlsdfjkaf",
		ClientSecret: "adsfljdsahfl",
	}

	_, _ = os.Create("test/test.txt")
	err := s.Store(&ExportResult{"test/test.txt", "text/plain", nil}, "test/")
	refute(t, err, (*barkup.Error)(nil))
}

func Test_S3_Store_ExportError(t *testing.T) {
	s := &S3{
		Region:       "fake-region",
		Bucket:       "cheese",
		AccessKey:    "afjlsdfjkaf",
		ClientSecret: "adsfljdsahfl",
	}

	_, _ = os.Create("test/test.txt")
	err := s.Store(&barkup.ExportResult{"", "text/plain", &barkup.Error{}}, "test/")
	refute(t, err, nil)
}
