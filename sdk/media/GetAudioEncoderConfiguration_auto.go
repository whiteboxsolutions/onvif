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

// Call_GetAudioEncoderConfiguration forwards the call to dev.CallMethod() then parses the payload of the reply as a GetAudioEncoderConfigurationResponse.
func Call_GetAudioEncoderConfiguration(ctx context.Context, dev *onvif.Device, request media.GetAudioEncoderConfiguration) (media.GetAudioEncoderConfigurationResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetAudioEncoderConfigurationResponse media.GetAudioEncoderConfigurationResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.GetAudioEncoderConfigurationResponse, errors.Annotate(err, "call")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "GetAudioEncoderConfiguration")
		return reply.Body.GetAudioEncoderConfigurationResponse, errors.Annotate(err, "reply")
	}
}
