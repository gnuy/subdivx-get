package main

import (
	"compress/flate"
	"fmt"
	"log"
	"strings"

	"github.com/mholt/archiver/v3"
)

var (
	z = archiver.Zip{
		CompressionLevel:       flate.DefaultCompression,
		MkdirAll:               true,
		SelectiveCompression:   true,
		ContinueOnError:        false,
		OverwriteExisting:      false,
		ImplicitTopLevelFolder: false,
	}
	r = archiver.Rar{
		MkdirAll:               true,
		ContinueOnError:        false,
		OverwriteExisting:      false,
		ImplicitTopLevelFolder: false,
	}
	u archiver.Unarchiver
	w archiver.Walker
)

func scan(file string) {

	err := z.Walk(file, func(f archiver.File) error {
		fmt.Println(f.Name(), f.Size())

		return nil
	})

	if err != nil {
		if strings.Contains(err.Error(), "not a valid zip file") {
			err := r.Walk(file, func(f archiver.File) error {
				fmt.Println(f.Name(), f.Size())

				return nil
			})
			if err != nil {
				log.Fatal(err)
			}

		} else {
			log.Fatal(err)
		}
	}

}

func unzip(file string, dest string) {
	error := z.Unarchive(file, dest)
	if error != nil {
		if strings.Contains(error.Error(), "not a valid zip file") {
			error := r.Unarchive(file, dest)
			if error != nil {
				log.Fatal(error)
			}

		} else {
			log.Fatal(error)
		}
	}
}
