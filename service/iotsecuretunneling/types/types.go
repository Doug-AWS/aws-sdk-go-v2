// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	"time"
)

// The state of a connection.
type ConnectionState struct {

	// The last time the connection status was updated.
	LastUpdatedAt *time.Time

	// The connection status of the tunnel. Valid values are CONNECTED and
	// DISCONNECTED.
	Status ConnectionStatus
}

// The destination configuration.
type DestinationConfig struct {

	// A list of service names that identity the target application. Currently, you can
	// only specify a single name. The AWS IoT client running on the destination device
	// reads this value and uses it to look up a port or an IP address and a port. The
	// AWS IoT client instantiates the local proxy which uses this information to
	// connect to the destination application.
	//
	// This member is required.
	Services []*string

	// The name of the IoT thing to which you want to connect.
	//
	// This member is required.
	ThingName *string
}

// An arbitary key/value pair used to add searchable metadata to secure tunnel
// resources.
type Tag struct {

	// The key of the tag.
	//
	// This member is required.
	Key *string

	// The value of the tag.
	//
	// This member is required.
	Value *string
}

// Tunnel timeout configuration.
type TimeoutConfig struct {

	// The maximum amount of time (in minutes) a tunnel can remain open. If not
	// specified, maxLifetimeTimeoutMinutes defaults to 720 minutes. Valid values are
	// from 1 minute to 12 hours (720 minutes)
	MaxLifetimeTimeoutMinutes *int32
}

// A connection between a source computer and a destination device.
type Tunnel struct {

	// The time when the tunnel was created.
	CreatedAt *time.Time

	// A description of the tunnel.
	Description *string

	// The destination configuration that specifies the thing name of the destination
	// device and a service name that the local proxy uses to connect to the
	// destination application.
	DestinationConfig *DestinationConfig

	// The connection state of the destination application.
	DestinationConnectionState *ConnectionState

	// The last time the tunnel was updated.
	LastUpdatedAt *time.Time

	// The connection state of the source application.
	SourceConnectionState *ConnectionState

	// The status of a tunnel. Valid values are: Open and Closed.
	Status TunnelStatus

	// A list of tag metadata associated with the secure tunnel.
	Tags []*Tag

	// Timeout configuration for the tunnel.
	TimeoutConfig *TimeoutConfig

	// The Amazon Resource Name (ARN) of a tunnel. The tunnel ARN format is
	// arn:aws:tunnel:::tunnel/
	TunnelArn *string

	// A unique alpha-numeric ID that identifies a tunnel.
	TunnelId *string
}

// Information about the tunnel.
type TunnelSummary struct {

	// The time the tunnel was created.
	CreatedAt *time.Time

	// A description of the tunnel.
	Description *string

	// The time the tunnel was last updated.
	LastUpdatedAt *time.Time

	// The status of a tunnel. Valid values are: Open and Closed.
	Status TunnelStatus

	// The Amazon Resource Name of the tunnel. The tunnel ARN format is
	// arn:aws:tunnel:::tunnel/
	TunnelArn *string

	// The unique alpha-numeric identifier for the tunnel.
	TunnelId *string
}
