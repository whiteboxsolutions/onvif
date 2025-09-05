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

// Call_GetHostname forwards the call to dev.CallMethod() then parses the payload of the reply as a GetHostnameResponse.
func Call_GetHostname(ctx context.Context, dev *onvif.Device, request device.GetHostname) (device.GetHostnameResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetHostnameResponse device.GetHostnameResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.GetHostnameResponse, errors.Annotate(err, "call")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "GetHostname")
		return reply.Body.GetHostnameResponse, errors.Annotate(err, "reply")
	}
}
