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

// Call_RemoveScopes forwards the call to dev.CallMethod() then parses the payload of the reply as a RemoveScopesResponse.
func Call_RemoveScopes(ctx context.Context, dev *onvif.Device, request device.RemoveScopes) (device.RemoveScopesResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			RemoveScopesResponse device.RemoveScopesResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.RemoveScopesResponse, errors.Annotate(err, "call")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "RemoveScopes")
		return reply.Body.RemoveScopesResponse, errors.Annotate(err, "reply")
	}
}
