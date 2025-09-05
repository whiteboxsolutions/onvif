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

// Call_SetOSD forwards the call to dev.CallMethod() then parses the payload of the reply as a SetOSDResponse.
func Call_SetOSD(ctx context.Context, dev *onvif.Device, request media.SetOSD) (media.SetOSDResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			SetOSDResponse media.SetOSDResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.SetOSDResponse, errors.Annotate(err, "call")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "SetOSD")
		return reply.Body.SetOSDResponse, errors.Annotate(err, "reply")
	}
}
