package features

import (
	"context"
	"fmt"
	"io/fs"
	"path/filepath"
	"strings"

	"github.com/altierawr/vstreamer/ent"
	"github.com/altierawr/vstreamer/ent/library"
)

var validExtensions []string = []string{".mkv", ".mp4"}

func ScanLibrary(ctx context.Context, client *ent.Client, id int) {
	library, err := client.Library.Query().Where(library.ID(id)).Only(ctx)
	if err != nil {
		fmt.Printf("error querying library: %v\n", err)
		return
	}

	err = filepath.WalkDir(library.Path, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		info, err := d.Info()
		if err != nil {
			fmt.Printf("couldn't get info for file %s: %v", path, err)
		}

		if d.IsDir() {
			return nil
		}

		isValidExtension := false
		for _, ext := range validExtensions {
			if strings.HasSuffix(info.Name(), ext) {
				isValidExtension = true
				break
			}
		}

		if isValidExtension {
			video, err := client.Video.Create().
				SetLibraryID(id).
				SetPath(path).
				Save(ctx)

			if err != nil {
				fmt.Printf("error creating video %s: %v", path, err)
			} else {
				fmt.Printf("created video %s\n", video.Path)
			}
		}

		return nil
	})
	if err != nil {
		fmt.Printf("error walking directory: %v", err)
	}
}
