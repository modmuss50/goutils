package goutils

import (
	"fmt"
	"github.com/mholt/archiver"
)

func ExtractZip(zip string, dest string) {
	fmt.Println("Extracting " + zip)
	archiver.DefaultZip.Unarchive(zip, dest)
}
