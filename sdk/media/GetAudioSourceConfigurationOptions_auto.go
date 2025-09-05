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

// Call_GetAudioSourceConfigurationOptions forwards the call to dev.CallMethod() then parses the payload of the reply as a GetAudioSourceConfigurationOptionsResponse.
func Call_GetAudioSourceConfigurationOptions(ctx context.Context, dev *onvif.Device, request media.GetAudioSourceConfigurationOptions) (media.GetAudioSourceConfigurationOptionsResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetAudioSourceConfigurationOptionsResponse media.GetAudioSourceConfigurationOptionsResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.GetAudioSourceConfigurationOptionsResponse, errors.Annotate(err, "call")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "GetAudioSourceConfigurationOptions")
		return reply.Body.GetAudioSourceConfigurationOptionsResponse, errors.Annotate(err, "reply")
	}
}
