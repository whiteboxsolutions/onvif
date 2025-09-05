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

// Call_GetServices forwards the call to dev.CallMethod() then parses the payload of the reply as a GetServicesResponse.
func Call_GetServices(ctx context.Context, dev *onvif.Device, request device.GetServices) (device.GetServicesResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetServicesResponse device.GetServicesResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.GetServicesResponse, errors.Annotate(err, "call")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "GetServices")
		return reply.Body.GetServicesResponse, errors.Annotate(err, "reply")
	}
}
