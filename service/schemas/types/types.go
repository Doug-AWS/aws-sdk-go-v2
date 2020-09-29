// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	"time"
)

type DiscovererSummary struct {
	// The state of the discoverer.
	State DiscovererState
	// The ARN of the event bus.
	SourceArn *string
	// The ID of the discoverer.
	DiscovererId *string
	// Tags associated with the resource.
	Tags map[string]*string
	// The ARN of the discoverer.
	DiscovererArn *string
}

type RegistrySummary struct {
	// The ARN of the registry.
	RegistryArn *string
	// The name of the registry.
	RegistryName *string
	// Tags associated with the registry.
	Tags map[string]*string
}

// A summary of schema details.
type SchemaSummary struct {
	// The name of the schema.
	SchemaName *string
	// The ARN of the schema.
	SchemaArn *string
	// The number of versions available for the schema.
	VersionCount *int64
	// The date and time that schema was modified.
	LastModified *time.Time
	// Tags associated with the schema.
	Tags map[string]*string
}

type SchemaVersionSummary struct {
	// The version number of the schema.
	SchemaVersion *string
	// The name of the schema.
	SchemaName *string
	// The ARN of the schema version.
	SchemaArn *string
}

type SearchSchemaSummary struct {
	// An array of schema version summaries.
	SchemaVersions []*SearchSchemaVersionSummary
	// The name of the schema.
	SchemaName *string
	// The ARN of the schema.
	SchemaArn *string
	// The name of the registry.
	RegistryName *string
}

type SearchSchemaVersionSummary struct {
	// The version number of the schema
	SchemaVersion *string
	// The date the schema version was created.
	CreatedDate *time.Time
}
