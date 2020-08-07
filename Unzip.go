package main

import (
	"compress/flate"

	"github.com/mholt/archiver/v3"
)

var (
	zip = archiver.Zip{
		CompressionLevel:       flate.DefaultCompression,
		MkdirAll:               true,
		SelectiveCompression:   true,
		ContinueOnError:        false,
		OverwriteExisting:      false,
		ImplicitTopLevelFolder: false,
	}
)

func unzip(file string, dest string) {
	// z.Extract(file, "*.srt", dest)
	zip.Unarchive(file, dest)
}
