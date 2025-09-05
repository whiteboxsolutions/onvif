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

// Call_GetNetworkProtocols forwards the call to dev.CallMethod() then parses the payload of the reply as a GetNetworkProtocolsResponse.
func Call_GetNetworkProtocols(ctx context.Context, dev *onvif.Device, request device.GetNetworkProtocols) (device.GetNetworkProtocolsResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetNetworkProtocolsResponse device.GetNetworkProtocolsResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.GetNetworkProtocolsResponse, errors.Annotate(err, "call")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "GetNetworkProtocols")
		return reply.Body.GetNetworkProtocolsResponse, errors.Annotate(err, "reply")
	}
}
