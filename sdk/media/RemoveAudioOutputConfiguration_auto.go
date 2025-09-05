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

// Call_RemoveAudioOutputConfiguration forwards the call to dev.CallMethod() then parses the payload of the reply as a RemoveAudioOutputConfigurationResponse.
func Call_RemoveAudioOutputConfiguration(ctx context.Context, dev *onvif.Device, request media.RemoveAudioOutputConfiguration) (media.RemoveAudioOutputConfigurationResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			RemoveAudioOutputConfigurationResponse media.RemoveAudioOutputConfigurationResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.RemoveAudioOutputConfigurationResponse, errors.Annotate(err, "call")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "RemoveAudioOutputConfiguration")
		return reply.Body.RemoveAudioOutputConfigurationResponse, errors.Annotate(err, "reply")
	}
}
