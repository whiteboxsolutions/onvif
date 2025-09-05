// Code generated : DO NOT EDIT.
// Copyright (c) 2022 Jean-Francois SMIGIELSKI
// Distributed under the MIT License

package media

import (
	"context"

	"github.com/juju/errors"
	"github.com/whiteboxsolutions/onvif"
	"github.com/whiteboxsolutions/onvif/media"
	"github.com/whiteboxsolutions/onvif/sdk"
)

// Call_GetProfile forwards the call to dev.CallMethod() then parses the payload of the reply as a GetProfileResponse.
func Call_GetProfile(ctx context.Context, dev *onvif.Device, request media.GetProfile) (media.GetProfileResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetProfileResponse media.GetProfileResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.GetProfileResponse, errors.Annotate(err, "call")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "GetProfile")
		return reply.Body.GetProfileResponse, errors.Annotate(err, "reply")
	}
}
