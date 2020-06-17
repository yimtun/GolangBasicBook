```
package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	err := filepath.Walk("/opt/bozz/",
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			//fmt.Println(path, info.Size())
			//fmt.Println(path)
			//fmt.Println(info.Name())
			name := info.Name()
			//fmt.Println(strings.HasSuffix(name, "war"))
			if strings.HasSuffix(name, "war") {
				fmt.Println(path)
			}
			if strings.HasSuffix(name, "Dockerfile") {
				fmt.Println(path)
			}

			return nil
		})
	if err != nil {
		log.Println(err)
	}

}
```
