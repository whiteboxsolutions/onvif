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

// Call_GetUsers forwards the call to dev.CallMethod() then parses the payload of the reply as a GetUsersResponse.
func Call_GetUsers(ctx context.Context, dev *onvif.Device, request device.GetUsers) (device.GetUsersResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetUsersResponse device.GetUsersResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.GetUsersResponse, errors.Annotate(err, "call")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "GetUsers")
		return reply.Body.GetUsersResponse, errors.Annotate(err, "reply")
	}
}
