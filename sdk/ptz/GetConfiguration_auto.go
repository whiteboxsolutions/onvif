// Code generated : DO NOT EDIT.
// Copyright (c) 2022 Jean-Francois SMIGIELSKI
// Distributed under the MIT License

package ptz

import (
	"context"

	"github.com/juju/errors"
	"github.com/whiteboxsolutions/onvif"
	"github.com/whiteboxsolutions/onvif/ptz"
	"github.com/whiteboxsolutions/onvif/sdk"
)

// Call_GetConfiguration forwards the call to dev.CallMethod() then parses the payload of the reply as a GetConfigurationResponse.
func Call_GetConfiguration(ctx context.Context, dev *onvif.Device, request ptz.GetConfiguration) (ptz.GetConfigurationResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetConfigurationResponse ptz.GetConfigurationResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.GetConfigurationResponse, errors.Annotate(err, "call")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "GetConfiguration")
		return reply.Body.GetConfigurationResponse, errors.Annotate(err, "reply")
	}
}
