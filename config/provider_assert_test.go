// Code generated by github.com/aws/aws-sdk-go-v2/config. DO NOT EDIT.

package config

// APIOptionsProvider implementor assertions
var (
	_ APIOptionsProvider = WithAPIOptions(nil)
)

// AssumeRoleCredentialOptionsProvider implementor assertions
var (
	_ AssumeRoleCredentialOptionsProvider = WithAssumeRoleCredentialOptions(nil)
)

// CredentialsProviderProvider implementor assertions
var (
	_ CredentialsProviderProvider = WithCredentialsProvider(nil)
)

// CustomCABundleProvider implementor assertions
var (
	_ CustomCABundleProvider = &EnvConfig{}
	_ CustomCABundleProvider = WithCustomCABundle(nil)
)

// DefaultRegionProvider implementor assertions
var (
	_ DefaultRegionProvider = WithDefaultRegion("")
)

// EC2RoleCredentialOptionsProvider implementor assertions
var (
	_ EC2RoleCredentialOptionsProvider = WithEC2RoleCredentialOptions(nil)
)

// EndpointCredentialOptionsProvider implementor assertions
var (
	_ EndpointCredentialOptionsProvider = WithEndpointCredentialOptions(nil)
)

// EndpointResolverProvider implementor assertions
var (
	_ EndpointResolverProvider = WithEndpointResolver(nil)
)

// HTTPClientProvider implementor assertions
var (
	_ HTTPClientProvider = WithHTTPClient(nil)
)

// RegionProvider implementor assertions
var (
	_ RegionProvider = &EnvConfig{}
	_ RegionProvider = &SharedConfig{}
	_ RegionProvider = WithRegion("")
	_ RegionProvider = WithEC2IMDSRegion{}
)

// SharedConfigFilesProvider implementor assertions
var (
	_ SharedConfigFilesProvider = &EnvConfig{}
	_ SharedConfigFilesProvider = WithSharedConfigFiles(nil)
)

// SharedConfigProfileProvider implementor assertions
var (
	_ SharedConfigProfileProvider = &EnvConfig{}
	_ SharedConfigProfileProvider = WithSharedConfigProfile("")
)

// WebIdentityRoleCredentialOptionsProvider implementor assertions
var (
	_ WebIdentityRoleCredentialOptionsProvider = WithWebIdentityRoleCredentialOptions(nil)
)
