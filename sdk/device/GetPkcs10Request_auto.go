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

// Call_GetPkcs10Request forwards the call to dev.CallMethod() then parses the payload of the reply as a GetPkcs10RequestResponse.
func Call_GetPkcs10Request(ctx context.Context, dev *onvif.Device, request device.GetPkcs10Request) (device.GetPkcs10RequestResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetPkcs10RequestResponse device.GetPkcs10RequestResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.GetPkcs10RequestResponse, errors.Annotate(err, "call")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "GetPkcs10Request")
		return reply.Body.GetPkcs10RequestResponse, errors.Annotate(err, "reply")
	}
}
