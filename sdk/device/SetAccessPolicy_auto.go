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

// Call_SetAccessPolicy forwards the call to dev.CallMethod() then parses the payload of the reply as a SetAccessPolicyResponse.
func Call_SetAccessPolicy(ctx context.Context, dev *onvif.Device, request device.SetAccessPolicy) (device.SetAccessPolicyResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			SetAccessPolicyResponse device.SetAccessPolicyResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.SetAccessPolicyResponse, errors.Annotate(err, "call")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "SetAccessPolicy")
		return reply.Body.SetAccessPolicyResponse, errors.Annotate(err, "reply")
	}
}
