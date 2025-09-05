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

// Call_GetSystemDateAndTime forwards the call to dev.CallMethod() then parses the payload of the reply as a GetSystemDateAndTimeResponse.
func Call_GetSystemDateAndTime(ctx context.Context, dev *onvif.Device, request device.GetSystemDateAndTime) (device.GetSystemDateAndTimeResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetSystemDateAndTimeResponse device.GetSystemDateAndTimeResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.GetSystemDateAndTimeResponse, errors.Annotate(err, "call")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "GetSystemDateAndTime")
		return reply.Body.GetSystemDateAndTimeResponse, errors.Annotate(err, "reply")
	}
}
