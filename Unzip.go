package main

import (
	"compress/flate"
	"log"
	"strings"

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
		if strings.Contains(error.Error(), "not a valid zip file") {
			error := rar.Unarchive(file, dest)
			if error != nil {
				log.Fatal(error)
			}

		} else {
			log.Fatal(error)
		}
	}
}
