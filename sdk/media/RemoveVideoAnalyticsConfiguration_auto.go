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

// Call_RemoveVideoAnalyticsConfiguration forwards the call to dev.CallMethod() then parses the payload of the reply as a RemoveVideoAnalyticsConfigurationResponse.
func Call_RemoveVideoAnalyticsConfiguration(ctx context.Context, dev *onvif.Device, request media.RemoveVideoAnalyticsConfiguration) (media.RemoveVideoAnalyticsConfigurationResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			RemoveVideoAnalyticsConfigurationResponse media.RemoveVideoAnalyticsConfigurationResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.RemoveVideoAnalyticsConfigurationResponse, errors.Annotate(err, "call")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "RemoveVideoAnalyticsConfiguration")
		return reply.Body.RemoveVideoAnalyticsConfigurationResponse, errors.Annotate(err, "reply")
	}
}
