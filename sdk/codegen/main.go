// Copyright (c) 2022 Jean-Francois SMIGIELSKI
// Distributed under the MIT License

package main

import (
	"flag"
	"log"
	"os"
	"text/template"
)

var mainTemplate = `// Code generated : DO NOT EDIT.
// Copyright (c) 2022 Jean-Francois SMIGIELSKI
// Distributed under the MIT License

package {{.Package}}

import (
	"context"
	"github.com/juju/errors"
	"github.com/whiteboxsolutions/onvif"
	"github.com/whiteboxsolutions/onvif/sdk"
	"github.com/whiteboxsolutions/onvif/{{.StructPackage}}"
)

// Call_{{.TypeRequest}} forwards the call to dev.CallMethod() then parses the payload of the reply as a {{.TypeReply}}.
func Call_{{.TypeRequest}}(ctx context.Context, dev *onvif.Device, request {{.StructPackage}}.{{.TypeRequest}}) ({{.StructPackage}}.{{.TypeReply}}, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			{{.TypeReply}} {{.StructPackage}}.{{.TypeReply}}
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.{{.TypeReply}}, errors.Annotate(err, "call")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "{{.TypeRequest}}")
		return reply.Body.{{.TypeReply}}, errors.Annotate(err, "reply")
	}
}
`

type parserEnv struct {
	Package       string
	StructPackage string
	TypeReply     string
	TypeRequest   string
}

func main() {
	flag.Parse()
	env := parserEnv{
		Package:       flag.Arg(0),
		StructPackage: flag.Arg(1),
		TypeRequest:   flag.Arg(2),
		TypeReply:     flag.Arg(2) + "Response",
	}

	log.Println(env)

	body, err := template.New("body").Parse(mainTemplate)
	if err != nil {
		log.Fatalln(err)
	}

	fout, err := os.Create(env.TypeRequest + "_auto.go")
	if err != nil {
		log.Fatalln(err)
	}
	defer fout.Close()

	err = body.Execute(fout, &env)
	if err != nil {
		log.Fatalln(err)
	}
}
