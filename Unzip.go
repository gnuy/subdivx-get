package main

import (
	"compress/flate"
	"log"

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
	rar = archiver.Rar{
		MkdirAll:               true,
		ContinueOnError:        false,
		OverwriteExisting:      false,
		ImplicitTopLevelFolder: false,
	}
)

func unzip(file string, dest string) {

	error := zip.Unarchive(file, dest)
	if error != nil {
		error := rar.Unarchive(file, dest)
		if error != nil {
			log.Fatal(error)
		}
	}
}
