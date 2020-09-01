package main

import (
	"google.golang.org/protobuf/compiler/protogen"
)

func main() {
	protogen.Options{}.Run(func(p *protogen.Plugin) error {
		for _, f := range p.Files {
			if f.Generate {
				GenerateFile(p, f)
			}
		}

		return nil
	})
}
