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

// Call_RemoveVideoSourceConfiguration forwards the call to dev.CallMethod() then parses the payload of the reply as a RemoveVideoSourceConfigurationResponse.
func Call_RemoveVideoSourceConfiguration(ctx context.Context, dev *onvif.Device, request media.RemoveVideoSourceConfiguration) (media.RemoveVideoSourceConfigurationResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			RemoveVideoSourceConfigurationResponse media.RemoveVideoSourceConfigurationResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.RemoveVideoSourceConfigurationResponse, errors.Annotate(err, "call")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "RemoveVideoSourceConfiguration")
		return reply.Body.RemoveVideoSourceConfigurationResponse, errors.Annotate(err, "reply")
	}
}
