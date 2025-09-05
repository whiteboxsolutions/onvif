// Code generated : DO NOT EDIT.
// Copyright (c) 2022 Jean-Francois SMIGIELSKI
// Distributed under the MIT License

package device

import (
	"context"

	"github.com/juju/errors"
	"github.com/whiteboxsolutions/onvif"
	"github.com/whiteboxsolutions/onvif/device"
	"github.com/whiteboxsolutions/onvif/sdk"
)

// Call_CreateCertificate forwards the call to dev.CallMethod() then parses the payload of the reply as a CreateCertificateResponse.
func Call_CreateCertificate(ctx context.Context, dev *onvif.Device, request device.CreateCertificate) (device.CreateCertificateResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			CreateCertificateResponse device.CreateCertificateResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.CreateCertificateResponse, errors.Annotate(err, "call")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "CreateCertificate")
		return reply.Body.CreateCertificateResponse, errors.Annotate(err, "reply")
	}
}
