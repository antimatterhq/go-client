// Code generated by ogen, DO NOT EDIT.

package spec

import (
	"context"
)

// Handler handles operations described by OpenAPI v3 specification.
type Handler interface {
	// CapsuleGetById implements capsuleGetById operation.
	//
	// This endpoint allows you to get a Capsule without knowing the Domain ID. It will redirect the user
	// to the full `/domains/{domainID}/capsules/{capsuleID}` path.
	//
	// GET /capsules/{capsuleID}
	CapsuleGetById(ctx context.Context, params CapsuleGetByIdParams) (CapsuleGetByIdRes, error)
	// DomainAddAccessLogEntry implements domainAddAccessLogEntry operation.
	//
	// Adds a data-plane audit log entry for this capsule. Contains information about the originating
	// principal and about read tag rollups. Contains an open capsule token (read from the file) to
	// ensure that you legitimately read the capsule. Note that not all audit log entry types may be
	// added with this method. Some (like open records) are generated server side.
	//
	// POST /domains/{domainID}/capsules/{capsuleID}/log
	DomainAddAccessLogEntry(ctx context.Context, req *AddCapsuleLogEntryRequest, params DomainAddAccessLogEntryParams) (DomainAddAccessLogEntryRes, error)
	// DomainAddExternalRootEncryptionKey implements domainAddExternalRootEncryptionKey operation.
	//
	// Add a new external root encryption key with its supporting access configuration.
	//
	// POST /domains/{domainID}/control/keys
	DomainAddExternalRootEncryptionKey(ctx context.Context, req *KeyInfos, params DomainAddExternalRootEncryptionKeyParams) (DomainAddExternalRootEncryptionKeyRes, error)
	// DomainAddNew implements domainAddNew operation.
	//
	// Add a new domain with no default peer relationships. You will need to confirm the email address
	// before the domain is able to be interacted with.
	//
	// POST /domains
	DomainAddNew(ctx context.Context, req *NewDomain) (DomainAddNewRes, error)
	// DomainAddReadContextRule implements domainAddReadContextRule operation.
	//
	// Read context configuration is rule based, much like domain policy. This adds a new rule to the
	// read context. Rules are processed in priority order, stopping with the first matching rule.
	//
	// POST /domains/{domainID}/control/read-context/{contextName}/config
	DomainAddReadContextRule(ctx context.Context, req *NewReadContextConfigRule, params DomainAddReadContextRuleParams) (DomainAddReadContextRuleRes, error)
	// DomainAuthenticate implements domainAuthenticate operation.
	//
	// Use an authentication method to obtain a domain ID token which is used as the bearer for all other
	// endpoints. You can use the `/public-info` route to obtain a list of identity providers supported
	// by this endpoint.
	//
	// POST /domains/{domainID}/authenticate
	DomainAuthenticate(ctx context.Context, req *DomainAuthenticate, params DomainAuthenticateParams) (DomainAuthenticateRes, error)
	// DomainContactIssueVerify implements domainContactIssueVerify operation.
	//
	// Issue a new verification request to a pending contact email associated with the domain. There is a
	// rate limiter on this endpoint, you may need to wait between invocations.
	//
	// POST /domains/{domainID}/account/verify
	DomainContactIssueVerify(ctx context.Context, req *DomainContactIssueVerifyReq, params DomainContactIssueVerifyParams) (DomainContactIssueVerifyRes, error)
	// DomainContactVerify implements domainContactVerify operation.
	//
	// Verify an admin contact email recently associated with a domain. The token will have been emailed
	// (in the form of a link) to the email address when `/account/verify` is called, the domain was
	// initially created, or the email was added via the settings endpoint.
	//
	// GET /domains/{domainID}/account/verify
	DomainContactVerify(ctx context.Context, params DomainContactVerifyParams) (DomainContactVerifyRes, error)
	// DomainCreateCapsule implements domainCreateCapsule operation.
	//
	// Create a new capsule. The ID will be returned. Capsule will be "unsealed" first, meaning it's
	// still in a creating state. Returns a capsule create token that can be used to feed in additional
	// data about the capsule while it's still unsealed. Also returns a DEK and an encrypted DEK.
	//
	// POST /domains/{domainID}/capsules
	DomainCreateCapsule(ctx context.Context, req *DomainCreateCapsuleReq, params DomainCreateCapsuleParams) (DomainCreateCapsuleRes, error)
	// DomainCreatePeerDomain implements domainCreatePeerDomain operation.
	//
	// Create a domain with a default "subordinate" peering relationship with the current domain.
	// Namely, the current "parent" domain will be configured to allow the new "child" domain to use the
	// parent's billing and admin contact settings, and the child domain will be configured to import
	// those settings.
	// Optionally, similar linking can be performed for identity providers, read/write contexts and facts
	// by setting the appropriate linkX parameter to true. In most cases, what you want is to set
	// `linkAll=true`.
	// Note, that a "subdomain" is just shorthand for a domain with the above-described peering config.
	// This peering can be changed at any time, and there is no permanent difference between a domain
	// created in this way, and a domain created with POST /domains.
	//
	// POST /domains/{domainID}/peer-domain
	DomainCreatePeerDomain(ctx context.Context, req *CreatePeerDomain, params DomainCreatePeerDomainParams) (DomainCreatePeerDomainRes, error)
	// DomainCreatePolicyRule implements domainCreatePolicyRule operation.
	//
	// Create a domain policy rule.
	//
	// POST /domains/{domainID}/control/policy
	DomainCreatePolicyRule(ctx context.Context, req *NewDomainPolicyRule, params DomainCreatePolicyRuleParams) (DomainCreatePolicyRuleRes, error)
	// DomainDataTaggingHookInvoke implements domainDataTaggingHookInvoke operation.
	//
	// Invoke a hook that operates on data and returns tags.
	//
	// POST /domains/{domainID}/hooks/data-tagging/{hookName}/invoke
	DomainDataTaggingHookInvoke(ctx context.Context, req *DataTaggingHookInput, params DomainDataTaggingHookInvokeParams) (DomainDataTaggingHookInvokeRes, error)
	// DomainDeleteCapability implements domainDeleteCapability operation.
	//
	// Delete a capability. All domain policy rules that reference the capability must have already been
	// deleted, or you will receive a 409 error.
	//
	// DELETE /domains/{domainID}/control/capabilities/{capability}
	DomainDeleteCapability(ctx context.Context, params DomainDeleteCapabilityParams) (DomainDeleteCapabilityRes, error)
	// DomainDeleteCapsuleTags implements domainDeleteCapsuleTags operation.
	//
	// Delete capsule-level tags.
	//
	// POST /domains/{domainID}/capsules/{capsuleID}/capsule-tags/delete
	DomainDeleteCapsuleTags(ctx context.Context, req *DeleteTags, params DomainDeleteCapsuleTagsParams) (DomainDeleteCapsuleTagsRes, error)
	// DomainDeleteExternalRootEncryptionKey implements domainDeleteExternalRootEncryptionKey operation.
	//
	// Delete an external root encryption key using its ID. This operation is only successful if the
	// external root encryption key is not in use by any key encryption keys. Call the /keys/rotate
	// endpoint to ensure that all KEKs have been migrated to the active REK.
	//
	// DELETE /domains/{domainID}/control/keys/{rootEncryptionKeyID}
	DomainDeleteExternalRootEncryptionKey(ctx context.Context, params DomainDeleteExternalRootEncryptionKeyParams) (DomainDeleteExternalRootEncryptionKeyRes, error)
	// DomainDeleteFactByID implements domainDeleteFactByID operation.
	//
	// Delete a fact by ID.
	//
	// DELETE /domains/{domainID}/control/facts/{factType}/{factID}
	DomainDeleteFactByID(ctx context.Context, params DomainDeleteFactByIDParams) (DomainDeleteFactByIDRes, error)
	// DomainDeleteFactType implements domainDeleteFactType operation.
	//
	// Deletes a fact type and all facts inside it.
	//
	// DELETE /domains/{domainID}/control/facts/{factType}
	DomainDeleteFactType(ctx context.Context, params DomainDeleteFactTypeParams) (DomainDeleteFactTypeRes, error)
	// DomainDeleteIdentityProvider implements domainDeleteIdentityProvider operation.
	//
	// Delete an identity provider. All domain tokens created using this identity provider will be
	// invalidated. Take care not to remove the identity provider that is providing you admin access to
	// your domain, as you may "lock yourself out".
	//
	// DELETE /domains/{domainID}/control/identities/{identityProviderName}
	DomainDeleteIdentityProvider(ctx context.Context, params DomainDeleteIdentityProviderParams) (DomainDeleteIdentityProviderRes, error)
	// DomainDeleteIdentityProviderPrincipal implements domainDeleteIdentityProviderPrincipal operation.
	//
	// Delete an identity provider principal.
	//
	// DELETE /domains/{domainID}/control/identities/{identityProviderName}/principals/{principalID}
	DomainDeleteIdentityProviderPrincipal(ctx context.Context, params DomainDeleteIdentityProviderPrincipalParams) (DomainDeleteIdentityProviderPrincipalRes, error)
	// DomainDeletePeer implements domainDeletePeer operation.
	//
	// Removes the peering relationship with the given domain.
	//
	// DELETE /domains/{domainID}/control/peers/{peerDomainID}
	DomainDeletePeer(ctx context.Context, params DomainDeletePeerParams) (DomainDeletePeerRes, error)
	// DomainDeletePolicyRule implements domainDeletePolicyRule operation.
	//
	// Delete a domain policy rule by ID.
	//
	// DELETE /domains/{domainID}/control/policy/{ruleID}
	DomainDeletePolicyRule(ctx context.Context, params DomainDeletePolicyRuleParams) (DomainDeletePolicyRuleRes, error)
	// DomainDeleteReadContext implements domainDeleteReadContext operation.
	//
	// Delete a read context. All configuration associated with this read context will also be deleted.
	// Domain policy rules referencing this read context will be left as-is.
	//
	// DELETE /domains/{domainID}/control/read-context/{contextName}
	DomainDeleteReadContext(ctx context.Context, params DomainDeleteReadContextParams) (DomainDeleteReadContextRes, error)
	// DomainDeleteReadContextRule implements domainDeleteReadContextRule operation.
	//
	// Deletes a read context configuration rule by ID.
	//
	// DELETE /domains/{domainID}/control/read-context/{contextName}/config/{ruleID}
	DomainDeleteReadContextRule(ctx context.Context, params DomainDeleteReadContextRuleParams) (DomainDeleteReadContextRuleRes, error)
	// DomainDeleteWriteContext implements domainDeleteWriteContext operation.
	//
	// Delete a write context. All configuration associated with this write context will also be deleted.
	// Domain policy rules referencing this write context will be left as-is.
	//
	// DELETE /domains/{domainID}/control/write-context/{contextName}
	DomainDeleteWriteContext(ctx context.Context, params DomainDeleteWriteContextParams) (DomainDeleteWriteContextRes, error)
	// DomainDeleteWriteContextRegexRule implements domainDeleteWriteContextRegexRule operation.
	//
	// Delete a regex classifier rule for the context.
	//
	// DELETE /domains/{domainID}/control/write-context/{contextName}/regex-rule/{ruleID}
	DomainDeleteWriteContextRegexRule(ctx context.Context, params DomainDeleteWriteContextRegexRuleParams) (DomainDeleteWriteContextRegexRuleRes, error)
	// DomainDescribeWriteContext implements domainDescribeWriteContext operation.
	//
	// Returns a detailed description of a write context.
	//
	// GET /domains/{domainID}/control/write-context/{contextName}
	DomainDescribeWriteContext(ctx context.Context, params DomainDescribeWriteContextParams) (DomainDescribeWriteContextRes, error)
	// DomainExternalRootEncryptionKeyTest implements domainExternalRootEncryptionKeyTest operation.
	//
	// Attempts to use a root encryption key to encrypt and decrypt, validating its availability.
	//
	// POST /domains/{domainID}/control/keys/{rootEncryptionKeyID}/test
	DomainExternalRootEncryptionKeyTest(ctx context.Context, req *DomainExternalRootEncryptionKeyTestReq, params DomainExternalRootEncryptionKeyTestParams) (DomainExternalRootEncryptionKeyTestRes, error)
	// DomainFlushEncryptionKeys implements domainFlushEncryptionKeys operation.
	//
	// Flush all keys in memory. The keys will be immediately reloaded from persistent storage, forcing a
	// check that the domain's root encryption key is still available.
	//
	// POST /domains/{domainID}/encryption/flush
	DomainFlushEncryptionKeys(ctx context.Context, req *DomainFlushEncryptionKeysReq, params DomainFlushEncryptionKeysParams) (DomainFlushEncryptionKeysRes, error)
	// DomainGetActiveExternalRootEncryptionKey implements domainGetActiveExternalRootEncryptionKey operation.
	//
	// Return the details about the current active root encryption key used by the domain.
	//
	// GET /domains/{domainID}/control/keys/active
	DomainGetActiveExternalRootEncryptionKey(ctx context.Context, params DomainGetActiveExternalRootEncryptionKeyParams) (DomainGetActiveExternalRootEncryptionKeyRes, error)
	// DomainGetCapabilities implements domainGetCapabilities operation.
	//
	// Get the capabilities configured within the domain. A capability is a key/value pair that can be
	// attached to a principal by an identity provider. The capabilities can be referenced by the domain
	// policy rules.
	//
	// GET /domains/{domainID}/control/capabilities
	DomainGetCapabilities(ctx context.Context, params DomainGetCapabilitiesParams) (DomainGetCapabilitiesRes, error)
	// DomainGetCapability implements domainGetCapability operation.
	//
	// Get a capability. A capability is a key/value pair that can be  attached to a principal by an
	// identity provider. The capabilities can be referenced by the domain policy rules.
	//
	// GET /domains/{domainID}/control/capabilities/{capability}
	DomainGetCapability(ctx context.Context, params DomainGetCapabilityParams) (DomainGetCapabilityRes, error)
	// DomainGetCapsuleInfo implements domainGetCapsuleInfo operation.
	//
	// Get the summary information about this capsule.
	//
	// GET /domains/{domainID}/capsules/{capsuleID}
	DomainGetCapsuleInfo(ctx context.Context, params DomainGetCapsuleInfoParams) (DomainGetCapsuleInfoRes, error)
	// DomainGetExternalRootEncryptionKeyProviders implements domainGetExternalRootEncryptionKeyProviders operation.
	//
	// Returns a list of available root encryption key providers, along with their description and, if
	// relevant, any additional information required to use them (e.g. for the delegated key provider
	// `aws_am` the AWS account number to delegate to is returned).
	//
	// GET /domains/{domainID}/control/keys/providers
	DomainGetExternalRootEncryptionKeyProviders(ctx context.Context, params DomainGetExternalRootEncryptionKeyProvidersParams) (DomainGetExternalRootEncryptionKeyProvidersRes, error)
	// DomainGetFactByID implements domainGetFactByID operation.
	//
	// Returns the fact with the given ID.
	//
	// GET /domains/{domainID}/control/facts/{factType}/{factID}
	DomainGetFactByID(ctx context.Context, params DomainGetFactByIDParams) (DomainGetFactByIDRes, error)
	// DomainGetFactType implements domainGetFactType operation.
	//
	// Get the definition of the given fact type.
	//
	// GET /domains/{domainID}/control/facts/{factType}
	DomainGetFactType(ctx context.Context, params DomainGetFactTypeParams) (DomainGetFactTypeRes, error)
	// DomainGetIdentityProvider implements domainGetIdentityProvider operation.
	//
	// Retrieve detailed information and configuration of an identity provider.
	//
	// GET /domains/{domainID}/control/identities/{identityProviderName}
	DomainGetIdentityProvider(ctx context.Context, params DomainGetIdentityProviderParams) (DomainGetIdentityProviderRes, error)
	// DomainGetIdentityProviderPrincipal implements domainGetIdentityProviderPrincipal operation.
	//
	// Retrieve detailed information about an identity provider principal.
	//
	// GET /domains/{domainID}/control/identities/{identityProviderName}/principals/{principalID}
	DomainGetIdentityProviderPrincipal(ctx context.Context, params DomainGetIdentityProviderPrincipalParams) (DomainGetIdentityProviderPrincipalRes, error)
	// DomainGetIdentityProviderPrincipals implements domainGetIdentityProviderPrincipals operation.
	//
	// Retrieve a list of principals for an identity provider.
	//
	// GET /domains/{domainID}/control/identities/{identityProviderName}/principals
	DomainGetIdentityProviderPrincipals(ctx context.Context, params DomainGetIdentityProviderPrincipalsParams) (DomainGetIdentityProviderPrincipalsRes, error)
	// DomainGetPeer implements domainGetPeer operation.
	//
	// Retrieve the details of a domain that is configured as a peer of this domain, by using its alias
	// or one of its nicknames.
	//
	// GET /domains/{domainID}/peer-domain
	DomainGetPeer(ctx context.Context, params DomainGetPeerParams) (DomainGetPeerRes, error)
	// DomainGetPeerConfig implements domainGetPeerConfig operation.
	//
	// Get the configuration for this peer.
	//
	// GET /domains/{domainID}/control/peers/{peerDomainID}
	DomainGetPeerConfig(ctx context.Context, params DomainGetPeerConfigParams) (DomainGetPeerConfigRes, error)
	// DomainGetPrivateInfo implements domainGetPrivateInfo operation.
	//
	// Returns a Domain's summary information. This may include more information than the `public-info`
	// endpoint but requires authentication.
	//
	// GET /domains/{domainID}/info
	DomainGetPrivateInfo(ctx context.Context, params DomainGetPrivateInfoParams) (DomainGetPrivateInfoRes, error)
	// DomainGetPublicInfo implements domainGetPublicInfo operation.
	//
	// Returns a Domain's summary information. This endpoint does not require authorization. This
	// endpoint can be used to determine which identity providers the `/authenticate` endpoint supports.
	//
	// GET /domains/{domainID}/public-info
	DomainGetPublicInfo(ctx context.Context, params DomainGetPublicInfoParams) (DomainGetPublicInfoRes, error)
	// DomainGetReadContext implements domainGetReadContext operation.
	//
	// Returns information about a read context.
	//
	// GET /domains/{domainID}/control/read-context/{contextName}
	DomainGetReadContext(ctx context.Context, params DomainGetReadContextParams) (DomainGetReadContextRes, error)
	// DomainGetSettings implements domainGetSettings operation.
	//
	// Get the domain settings. This contains configuration for the contact email addresses as well as
	// the display name for the domain.
	//
	// GET /domains/{domainID}/control/settings
	DomainGetSettings(ctx context.Context, params DomainGetSettingsParams) (DomainGetSettingsRes, error)
	// DomainGetStatus implements domainGetStatus operation.
	//
	// The domain status object contains important notifications for administrators of the domain.
	//
	// GET /domains/{domainID}/control/status
	DomainGetStatus(ctx context.Context, params DomainGetStatusParams) (DomainGetStatusRes, error)
	// DomainGetTagInfo implements domainGetTagInfo operation.
	//
	// Get an ordered list of the top 100 tags. The ordering is: - Tags emitted by hooks - Tags
	// referenced in read context rules - Capsule and span tags that appear in the capsule manifest
	// ordered by number of appearances This list will be truncated (and `has_more` will be true) if the
	// above yields more than 100 tags. There is currently no endpoint to receive a complete list of tags.
	//
	// GET /domains/{domainID}/tag-info
	DomainGetTagInfo(ctx context.Context, params DomainGetTagInfoParams) (DomainGetTagInfoRes, error)
	// DomainGetWriteContextRegexRules implements domainGetWriteContextRegexRules operation.
	//
	// Get a full listing of all regex rules for the context.
	//
	// GET /domains/{domainID}/control/write-context/{contextName}/regex-rule
	DomainGetWriteContextRegexRules(ctx context.Context, params DomainGetWriteContextRegexRulesParams) (DomainGetWriteContextRegexRulesRes, error)
	// DomainInsertIdentityProviderPrincipal implements domainInsertIdentityProviderPrincipal operation.
	//
	// Create a new principal for the provider. Note that the identityProviderName must refer to an
	// existing identity provider or the response will be a 400.
	//
	// POST /domains/{domainID}/control/identities/{identityProviderName}/principals
	DomainInsertIdentityProviderPrincipal(ctx context.Context, req *DomainIdentityProviderPrincipalParams, params DomainInsertIdentityProviderPrincipalParams) (DomainInsertIdentityProviderPrincipalRes, error)
	// DomainInsertWriteContextRegexRule implements domainInsertWriteContextRegexRule operation.
	//
	// Create a new regex rule for a write context.
	//
	// POST /domains/{domainID}/control/write-context/{contextName}/regex-rule
	DomainInsertWriteContextRegexRule(ctx context.Context, req *WriteContextRegexRule, params DomainInsertWriteContextRegexRuleParams) (DomainInsertWriteContextRegexRuleRes, error)
	// DomainListCapsules implements domainListCapsules operation.
	//
	// Get information about capsules.
	//
	// GET /domains/{domainID}/capsules
	DomainListCapsules(ctx context.Context, params DomainListCapsulesParams) (DomainListCapsulesRes, error)
	// DomainListExternalRootEncryptionKey implements domainListExternalRootEncryptionKey operation.
	//
	// List all external root encryption keys for the domain.
	//
	// GET /domains/{domainID}/control/keys
	DomainListExternalRootEncryptionKey(ctx context.Context, params DomainListExternalRootEncryptionKeyParams) (DomainListExternalRootEncryptionKeyRes, error)
	// DomainListFactTypes implements domainListFactTypes operation.
	//
	// Get a list of the fact types in this domain. Facts are used by domain policy rules and read
	// context rules.
	//
	// GET /domains/{domainID}/control/facts
	DomainListFactTypes(ctx context.Context, params DomainListFactTypesParams) (DomainListFactTypesRes, error)
	// DomainListFacts implements domainListFacts operation.
	//
	// Get the facts within a fact type.
	//
	// GET /domains/{domainID}/control/facts/{factType}/list
	DomainListFacts(ctx context.Context, params DomainListFactsParams) (DomainListFactsRes, error)
	// DomainListHooks implements domainListHooks operation.
	//
	// Get a list of available hooks in this domain. Hooks can be added to write contexts to classify
	// data.
	//
	// GET /domains/{domainID}/hooks
	DomainListHooks(ctx context.Context, params DomainListHooksParams) (DomainListHooksRes, error)
	// DomainListIdentityProviders implements domainListIdentityProviders operation.
	//
	// Retrieve the domain's identity providers and a brief overview of their configuration. This
	// endpoint requires authentication, but you can obtain an abridged list of the domain identity
	// providers prior to authentication by using the `/public-info` endpoint.
	//
	// GET /domains/{domainID}/control/identities
	DomainListIdentityProviders(ctx context.Context, params DomainListIdentityProvidersParams) (DomainListIdentityProvidersRes, error)
	// DomainListPeers implements domainListPeers operation.
	//
	// Returns a list of this domains peers.
	//
	// GET /domains/{domainID}/control/peers
	DomainListPeers(ctx context.Context, params DomainListPeersParams) (DomainListPeersRes, error)
	// DomainListPolicyRules implements domainListPolicyRules operation.
	//
	// Get the domain policy rules. These govern which resources in the domain can be interacted with.
	// Note that the peers "bypass" these rules, in that a peer domain can retrieve policy and
	// configuration that has been allowed by peering configuration without needing an allowing domain
	// policy rule, but they cannot access data within this domain.
	//
	// GET /domains/{domainID}/control/policy
	DomainListPolicyRules(ctx context.Context, params DomainListPolicyRulesParams) (DomainListPolicyRulesRes, error)
	// DomainListReadContexts implements domainListReadContexts operation.
	//
	// List the domain read contexts. If a user has view permissions on this resource, they may list all
	// read contexts, even if they do not have view, edit or use permissions on some of the read contexts
	// in the list.
	//
	// GET /domains/{domainID}/control/read-context
	DomainListReadContexts(ctx context.Context, params DomainListReadContextsParams) (DomainListReadContextsRes, error)
	// DomainListResources implements domainListResources operation.
	//
	// Gets a list of resource strings that can be used in policy rules, and the set of permissions that
	// you can assign to them. The return value from this endpoint is useful as a reference when
	// authoring custom domain policy for new capabilities.
	//
	// GET /domains/{domainID}/control/resources
	DomainListResources(ctx context.Context, params DomainListResourcesParams) (DomainListResourcesRes, error)
	// DomainListWriteContexts implements domainListWriteContexts operation.
	//
	// List the domain write contexts. If a user has view permissions on this resource, they may list all
	// write contexts, even if they do not have view, edit or use permissions on some of the write
	// contexts in the list.
	//
	// GET /domains/{domainID}/control/write-context
	DomainListWriteContexts(ctx context.Context, params DomainListWriteContextsParams) (DomainListWriteContextsRes, error)
	// DomainOpenCapsule implements domainOpenCapsule operation.
	//
	// Given the encrypted DEK for this capsule, get back the decrypted DEK. contains the read context.
	//
	// POST /domains/{domainID}/capsules/{capsuleID}/open
	DomainOpenCapsule(ctx context.Context, req *CapsuleOpenRequest, params DomainOpenCapsuleParams) (DomainOpenCapsuleRes, error)
	// DomainPatchSettings implements domainPatchSettings operation.
	//
	// Applies the given patch to the domain settings.
	//
	// PATCH /domains/{domainID}/control/settings
	DomainPatchSettings(ctx context.Context, req *DomainSettingsPatch, params DomainPatchSettingsParams) (DomainPatchSettingsRes, error)
	// DomainPolicyFlush implements domainPolicyFlush operation.
	//
	// Flush the policy cache so that changes to permissions take effect.
	//
	// POST /domains/{domainID}/control/policy/flush
	DomainPolicyFlush(ctx context.Context, params DomainPolicyFlushParams) (DomainPolicyFlushRes, error)
	// DomainPutCapability implements domainPutCapability operation.
	//
	// Create or update a capability. If you want to return an error if the capability already existed,
	// set `createonly` to true.
	//
	// PUT /domains/{domainID}/control/capabilities/{capability}
	DomainPutCapability(ctx context.Context, req *NewCapabilityDefinition, params DomainPutCapabilityParams) (DomainPutCapabilityRes, error)
	// DomainPutFactType implements domainPutFactType operation.
	//
	// Facts are used to store ancillary information that helps express domain policy rules and read
	// context configuration rules. This endpoint allows you to register a new fact type. To create a
	// fact within an existing type, use `/control/facts/{factType}/new`.
	//
	// PUT /domains/{domainID}/control/facts/{factType}
	DomainPutFactType(ctx context.Context, req *NewFactTypeDefinition, params DomainPutFactTypeParams) (DomainPutFactTypeRes, error)
	// DomainQueryAccessLog implements domainQueryAccessLog operation.
	//
	// Query the data access log for this domain. This contains all operations interacting with capsules
	// within this domain. Results are returned in reverse chronological order.
	//
	// GET /domains/{domainID}/log
	DomainQueryAccessLog(ctx context.Context, params DomainQueryAccessLogParams) (DomainQueryAccessLogRes, error)
	// DomainQueryAccessLogSingleCapsule implements domainQueryAccessLogSingleCapsule operation.
	//
	// Query the data-plane access log for this capsule. Results are returned in reverse chronological
	// order.
	//
	// GET /domains/{domainID}/capsules/{capsuleID}/log
	DomainQueryAccessLogSingleCapsule(ctx context.Context, params DomainQueryAccessLogSingleCapsuleParams) (DomainQueryAccessLogSingleCapsuleRes, error)
	// DomainQueryControlLog implements domainQueryControlLog operation.
	//
	// Query the domain control-plane audit log. Results are returned in reverse chronological order.
	//
	// GET /domains/{domainID}/control/log
	DomainQueryControlLog(ctx context.Context, params DomainQueryControlLogParams) (DomainQueryControlLogRes, error)
	// DomainReadContextFlush implements domainReadContextFlush operation.
	//
	// Flush the read context cache so that changes to permissions take effect.
	//
	// POST /domains/{domainID}/control/read-context/{contextName}/flush
	DomainReadContextFlush(ctx context.Context, params DomainReadContextFlushParams) (DomainReadContextFlushRes, error)
	// DomainRenumberPolicyRules implements domainRenumberPolicyRules operation.
	//
	// Re-assign rule priority numbers to integer multiples of 10.
	//
	// POST /domains/{domainID}/control/policy/renumber
	DomainRenumberPolicyRules(ctx context.Context, params DomainRenumberPolicyRulesParams) (DomainRenumberPolicyRulesRes, error)
	// DomainRotateRootEncryptionKeys implements domainRotateRootEncryptionKeys operation.
	//
	// Collects key encryption keys not encrypted with the current active root encryption key, decrypts
	// them with their original root encryption key, and then encrypts them with the active root
	// encryption key. This is a batched operation and only 100 keys will be processed in a single call.
	// In the response, "has_more" will be true if there are more KEKs that can be rotated. Usually the
	// caller will call this endpoint in a loop until has_more is false.
	//
	// POST /domains/{domainID}/control/keys/rotate
	DomainRotateRootEncryptionKeys(ctx context.Context, req *DomainRotateRootEncryptionKeysReq, params DomainRotateRootEncryptionKeysParams) (DomainRotateRootEncryptionKeysRes, error)
	// DomainSealCapsule implements domainSealCapsule operation.
	//
	// Seal this capsule, if it's unsealed. Requires capsule create token.
	//
	// POST /domains/{domainID}/capsules/{capsuleID}/seal
	DomainSealCapsule(ctx context.Context, req *CapsuleSealRequest, params DomainSealCapsuleParams) (DomainSealCapsuleRes, error)
	// DomainSetActiveExternalRootEncryptionKey implements domainSetActiveExternalRootEncryptionKey operation.
	//
	// This will set which root encryption is active: i.e. is used for new capsules, or is used to
	// encrypt KEKs when `/keys/rotate` is called or when new capsules are created.
	//
	// POST /domains/{domainID}/control/keys/active
	DomainSetActiveExternalRootEncryptionKey(ctx context.Context, req *ActiveRootEncryptionKeyID, params DomainSetActiveExternalRootEncryptionKeyParams) (DomainSetActiveExternalRootEncryptionKeyRes, error)
	// DomainUpdateIdentityProviderPrincipal implements domainUpdateIdentityProviderPrincipal operation.
	//
	// Update the set of capabilities assigned to an identity provider principal. The capabilities must
	// exist.
	//
	// PUT /domains/{domainID}/control/identities/{identityProviderName}/principals/{principalID}
	DomainUpdateIdentityProviderPrincipal(ctx context.Context, req *CapabilityList, params DomainUpdateIdentityProviderPrincipalParams) (DomainUpdateIdentityProviderPrincipalRes, error)
	// DomainUpdatePeer implements domainUpdatePeer operation.
	//
	// Create or update the configuration for this peer. Please note, if the configuration already exists,
	//  it is updated to reflect the values in the request. This will include setting the fields to their
	// default value if not supplied.
	//
	// PUT /domains/{domainID}/control/peers/{peerDomainID}
	DomainUpdatePeer(ctx context.Context, req *DomainPeerConfig, params DomainUpdatePeerParams) (DomainUpdatePeerRes, error)
	// DomainUpdatePolicyRule implements domainUpdatePolicyRule operation.
	//
	// Update a domain policy rule.
	//
	// PUT /domains/{domainID}/control/policy/{ruleID}
	DomainUpdatePolicyRule(ctx context.Context, req *NewDomainPolicyRule, params DomainUpdatePolicyRuleParams) (DomainUpdatePolicyRuleRes, error)
	// DomainUpdateReadContextRule implements domainUpdateReadContextRule operation.
	//
	// Update a read context configuration rule. The rule must already exist.
	//
	// PUT /domains/{domainID}/control/read-context/{contextName}/config/{ruleID}
	DomainUpdateReadContextRule(ctx context.Context, req *NewReadContextConfigRule, params DomainUpdateReadContextRuleParams) (DomainUpdateReadContextRuleRes, error)
	// DomainUpsertCapsuleTags implements domainUpsertCapsuleTags operation.
	//
	// Upsert capsule-level tags. This is permitted even after a capsule is sealed.
	//
	// POST /domains/{domainID}/capsules/{capsuleID}/capsule-tags
	DomainUpsertCapsuleTags(ctx context.Context, req []Tag, params DomainUpsertCapsuleTagsParams) (DomainUpsertCapsuleTagsRes, error)
	// DomainUpsertFact implements domainUpsertFact operation.
	//
	// Create a new fact. The fact type must have been previously registered using
	// `/control/facts/{factType}`. If an identical fact exists (having the same value for all fields),
	// this call is a no-op and returns the same ID.
	//
	// POST /domains/{domainID}/control/facts/{factType}/new
	DomainUpsertFact(ctx context.Context, req *NewFact, params DomainUpsertFactParams) (DomainUpsertFactRes, error)
	// DomainUpsertIdentityProvider implements domainUpsertIdentityProvider operation.
	//
	// Create or configure an identity provider.
	//
	// PUT /domains/{domainID}/control/identities/{identityProviderName}
	DomainUpsertIdentityProvider(ctx context.Context, req DomainIdentityProviderDetails, params DomainUpsertIdentityProviderParams) (DomainUpsertIdentityProviderRes, error)
	// DomainUpsertReadContext implements domainUpsertReadContext operation.
	//
	// Update or create a read context.
	//
	// PUT /domains/{domainID}/control/read-context/{contextName}
	DomainUpsertReadContext(ctx context.Context, req *AddReadContext, params DomainUpsertReadContextParams) (DomainUpsertReadContextRes, error)
	// DomainUpsertSpanTags implements domainUpsertSpanTags operation.
	//
	// Upsert span tag rollups. This is only permitted when a capsule is not sealed. It requires a
	// special "capsule owner" token that is returned by create capsule. Note that the rollup
	// calculations must be done on the client side. This method only permits storing the entire rollup,
	// not aggregating serverside. This is idempotent.
	//
	// PUT /domains/{domainID}/capsules/{capsuleID}/span-tags
	DomainUpsertSpanTags(ctx context.Context, req *UpsertSpanTagsRequest, params DomainUpsertSpanTagsParams) (DomainUpsertSpanTagsRes, error)
	// DomainUpsertWriteContext implements domainUpsertWriteContext operation.
	//
	// Create or update an existing write context. If the config is omitted, it will be left as-is
	// (existing write contexts) or created as blank (new write contexts).
	//
	// PUT /domains/{domainID}/control/write-context/{contextName}
	DomainUpsertWriteContext(ctx context.Context, req *AddWriteContext, params DomainUpsertWriteContextParams) (DomainUpsertWriteContextRes, error)
	// DomainUpsertWriteContextConfiguration implements domainUpsertWriteContextConfiguration operation.
	//
	// Update a write context configuration. The write context must already exist.
	//
	// PUT /domains/{domainID}/control/write-context/{contextName}/config
	DomainUpsertWriteContextConfiguration(ctx context.Context, req *WriteContextConfigInfo, params DomainUpsertWriteContextConfigurationParams) (DomainUpsertWriteContextConfigurationRes, error)
	// StarredDomainAdd implements starredDomainAdd operation.
	//
	// Adds the domain to the list of starred domains for the user.
	//
	// PUT /global/starred-domains/{domainID}
	StarredDomainAdd(ctx context.Context, req *StarredDomainAddReq, params StarredDomainAddParams) (StarredDomainAddRes, error)
	// StarredDomainList implements starredDomainList operation.
	//
	// Returns a list of domains that the user has starred. This is a list of domain IDs, not domain
	// names. The user must be authenticated to call this method.
	//
	// GET /global/starred-domains
	StarredDomainList(ctx context.Context) (StarredDomainListRes, error)
	// StarredDomainRemove implements starredDomainRemove operation.
	//
	// Removes the domain from the list of starred domains for the user.
	//
	// DELETE /global/starred-domains/{domainID}
	StarredDomainRemove(ctx context.Context, params StarredDomainRemoveParams) (StarredDomainRemoveRes, error)
	// NewError creates *ErrorStatusCode from error returned by handler.
	//
	// Used for common default response.
	NewError(ctx context.Context, err error) *ErrorStatusCode
}

// Server implements http server based on OpenAPI v3 specification and
// calls Handler to handle requests.
type Server struct {
	h   Handler
	sec SecurityHandler
	baseServer
}

// NewServer creates new Server.
func NewServer(h Handler, sec SecurityHandler, opts ...ServerOption) (*Server, error) {
	s, err := newServerConfig(opts...).baseServer()
	if err != nil {
		return nil, err
	}
	return &Server{
		h:          h,
		sec:        sec,
		baseServer: s,
	}, nil
}
