# LoadBalancerListener

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | 리스너 ID | 
**Name** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**Protocol** | Pointer to [**NullableProtocol**](Protocol.md) |  | [optional] 
**IsEnabled** | Pointer to **NullableBool** |  | [optional] 
**Secrets** | Pointer to [**[]Secret**](Secret.md) |  | [optional] 
**L7Policies** | Pointer to [**[]L7Policy**](L7Policy.md) |  | [optional] 
**TlsCiphers** | Pointer to **NullableString** |  | [optional] 
**TlsVersions** | Pointer to [**[]TLSVersion**](TLSVersion.md) |  | [optional] 
**ProjectId** | Pointer to **NullableString** |  | [optional] 
**ProtocolPort** | Pointer to **NullableInt32** |  | [optional] 
**LoadBalancerId** | Pointer to **NullableString** |  | [optional] 
**TlsCertificateId** | Pointer to **NullableString** |  | [optional] 
**ProvisioningStatus** | Pointer to [**NullableProvisioningStatus**](ProvisioningStatus.md) |  | [optional] 
**OperatingStatus** | Pointer to [**NullableLoadBalancerOperatingStatus**](LoadBalancerOperatingStatus.md) |  | [optional] 
**InsertHeaders** | Pointer to [**map[string]LoadBalancerListenerInsertHeadersValue**](LoadBalancerListenerInsertHeadersValue.md) |  | [optional] 
**CreatedAt** | Pointer to **NullableTime** |  | [optional] 
**UpdatedAt** | Pointer to **NullableTime** |  | [optional] 
**TimeoutClientData** | Pointer to **NullableInt32** |  | [optional] 
**DefaultTargetGroupName** | Pointer to **NullableString** |  | [optional] 
**DefaultTargetGroupId** | Pointer to **NullableString** |  | [optional] 
**LoadBalancerType** | Pointer to [**NullableLoadBalancerType**](LoadBalancerType.md) |  | [optional] 

## Methods

### NewLoadBalancerListener

`func NewLoadBalancerListener(id string, ) *LoadBalancerListener`

NewLoadBalancerListener instantiates a new LoadBalancerListener object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoadBalancerListenerWithDefaults

`func NewLoadBalancerListenerWithDefaults() *LoadBalancerListener`

NewLoadBalancerListenerWithDefaults instantiates a new LoadBalancerListener object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *LoadBalancerListener) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LoadBalancerListener) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LoadBalancerListener) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *LoadBalancerListener) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LoadBalancerListener) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LoadBalancerListener) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *LoadBalancerListener) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *LoadBalancerListener) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *LoadBalancerListener) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDescription

`func (o *LoadBalancerListener) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *LoadBalancerListener) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *LoadBalancerListener) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *LoadBalancerListener) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *LoadBalancerListener) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *LoadBalancerListener) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetProtocol

`func (o *LoadBalancerListener) GetProtocol() Protocol`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *LoadBalancerListener) GetProtocolOk() (*Protocol, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *LoadBalancerListener) SetProtocol(v Protocol)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *LoadBalancerListener) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### SetProtocolNil

`func (o *LoadBalancerListener) SetProtocolNil(b bool)`

 SetProtocolNil sets the value for Protocol to be an explicit nil

### UnsetProtocol
`func (o *LoadBalancerListener) UnsetProtocol()`

UnsetProtocol ensures that no value is present for Protocol, not even an explicit nil
### GetIsEnabled

`func (o *LoadBalancerListener) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *LoadBalancerListener) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *LoadBalancerListener) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.

### HasIsEnabled

`func (o *LoadBalancerListener) HasIsEnabled() bool`

HasIsEnabled returns a boolean if a field has been set.

### SetIsEnabledNil

`func (o *LoadBalancerListener) SetIsEnabledNil(b bool)`

 SetIsEnabledNil sets the value for IsEnabled to be an explicit nil

### UnsetIsEnabled
`func (o *LoadBalancerListener) UnsetIsEnabled()`

UnsetIsEnabled ensures that no value is present for IsEnabled, not even an explicit nil
### GetSecrets

`func (o *LoadBalancerListener) GetSecrets() []Secret`

GetSecrets returns the Secrets field if non-nil, zero value otherwise.

### GetSecretsOk

`func (o *LoadBalancerListener) GetSecretsOk() (*[]Secret, bool)`

GetSecretsOk returns a tuple with the Secrets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecrets

`func (o *LoadBalancerListener) SetSecrets(v []Secret)`

SetSecrets sets Secrets field to given value.

### HasSecrets

`func (o *LoadBalancerListener) HasSecrets() bool`

HasSecrets returns a boolean if a field has been set.

### SetSecretsNil

`func (o *LoadBalancerListener) SetSecretsNil(b bool)`

 SetSecretsNil sets the value for Secrets to be an explicit nil

### UnsetSecrets
`func (o *LoadBalancerListener) UnsetSecrets()`

UnsetSecrets ensures that no value is present for Secrets, not even an explicit nil
### GetL7Policies

`func (o *LoadBalancerListener) GetL7Policies() []L7Policy`

GetL7Policies returns the L7Policies field if non-nil, zero value otherwise.

### GetL7PoliciesOk

`func (o *LoadBalancerListener) GetL7PoliciesOk() (*[]L7Policy, bool)`

GetL7PoliciesOk returns a tuple with the L7Policies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetL7Policies

`func (o *LoadBalancerListener) SetL7Policies(v []L7Policy)`

SetL7Policies sets L7Policies field to given value.

### HasL7Policies

`func (o *LoadBalancerListener) HasL7Policies() bool`

HasL7Policies returns a boolean if a field has been set.

### SetL7PoliciesNil

`func (o *LoadBalancerListener) SetL7PoliciesNil(b bool)`

 SetL7PoliciesNil sets the value for L7Policies to be an explicit nil

### UnsetL7Policies
`func (o *LoadBalancerListener) UnsetL7Policies()`

UnsetL7Policies ensures that no value is present for L7Policies, not even an explicit nil
### GetTlsCiphers

`func (o *LoadBalancerListener) GetTlsCiphers() string`

GetTlsCiphers returns the TlsCiphers field if non-nil, zero value otherwise.

### GetTlsCiphersOk

`func (o *LoadBalancerListener) GetTlsCiphersOk() (*string, bool)`

GetTlsCiphersOk returns a tuple with the TlsCiphers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsCiphers

`func (o *LoadBalancerListener) SetTlsCiphers(v string)`

SetTlsCiphers sets TlsCiphers field to given value.

### HasTlsCiphers

`func (o *LoadBalancerListener) HasTlsCiphers() bool`

HasTlsCiphers returns a boolean if a field has been set.

### SetTlsCiphersNil

`func (o *LoadBalancerListener) SetTlsCiphersNil(b bool)`

 SetTlsCiphersNil sets the value for TlsCiphers to be an explicit nil

### UnsetTlsCiphers
`func (o *LoadBalancerListener) UnsetTlsCiphers()`

UnsetTlsCiphers ensures that no value is present for TlsCiphers, not even an explicit nil
### GetTlsVersions

`func (o *LoadBalancerListener) GetTlsVersions() []*TLSVersion`

GetTlsVersions returns the TlsVersions field if non-nil, zero value otherwise.

### GetTlsVersionsOk

`func (o *LoadBalancerListener) GetTlsVersionsOk() (*[]*TLSVersion, bool)`

GetTlsVersionsOk returns a tuple with the TlsVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsVersions

`func (o *LoadBalancerListener) SetTlsVersions(v []*TLSVersion)`

SetTlsVersions sets TlsVersions field to given value.

### HasTlsVersions

`func (o *LoadBalancerListener) HasTlsVersions() bool`

HasTlsVersions returns a boolean if a field has been set.

### SetTlsVersionsNil

`func (o *LoadBalancerListener) SetTlsVersionsNil(b bool)`

 SetTlsVersionsNil sets the value for TlsVersions to be an explicit nil

### UnsetTlsVersions
`func (o *LoadBalancerListener) UnsetTlsVersions()`

UnsetTlsVersions ensures that no value is present for TlsVersions, not even an explicit nil
### GetProjectId

`func (o *LoadBalancerListener) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *LoadBalancerListener) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *LoadBalancerListener) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *LoadBalancerListener) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### SetProjectIdNil

`func (o *LoadBalancerListener) SetProjectIdNil(b bool)`

 SetProjectIdNil sets the value for ProjectId to be an explicit nil

### UnsetProjectId
`func (o *LoadBalancerListener) UnsetProjectId()`

UnsetProjectId ensures that no value is present for ProjectId, not even an explicit nil
### GetProtocolPort

`func (o *LoadBalancerListener) GetProtocolPort() int32`

GetProtocolPort returns the ProtocolPort field if non-nil, zero value otherwise.

### GetProtocolPortOk

`func (o *LoadBalancerListener) GetProtocolPortOk() (*int32, bool)`

GetProtocolPortOk returns a tuple with the ProtocolPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocolPort

`func (o *LoadBalancerListener) SetProtocolPort(v int32)`

SetProtocolPort sets ProtocolPort field to given value.

### HasProtocolPort

`func (o *LoadBalancerListener) HasProtocolPort() bool`

HasProtocolPort returns a boolean if a field has been set.

### SetProtocolPortNil

`func (o *LoadBalancerListener) SetProtocolPortNil(b bool)`

 SetProtocolPortNil sets the value for ProtocolPort to be an explicit nil

### UnsetProtocolPort
`func (o *LoadBalancerListener) UnsetProtocolPort()`

UnsetProtocolPort ensures that no value is present for ProtocolPort, not even an explicit nil
### GetLoadBalancerId

`func (o *LoadBalancerListener) GetLoadBalancerId() string`

GetLoadBalancerId returns the LoadBalancerId field if non-nil, zero value otherwise.

### GetLoadBalancerIdOk

`func (o *LoadBalancerListener) GetLoadBalancerIdOk() (*string, bool)`

GetLoadBalancerIdOk returns a tuple with the LoadBalancerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadBalancerId

`func (o *LoadBalancerListener) SetLoadBalancerId(v string)`

SetLoadBalancerId sets LoadBalancerId field to given value.

### HasLoadBalancerId

`func (o *LoadBalancerListener) HasLoadBalancerId() bool`

HasLoadBalancerId returns a boolean if a field has been set.

### SetLoadBalancerIdNil

`func (o *LoadBalancerListener) SetLoadBalancerIdNil(b bool)`

 SetLoadBalancerIdNil sets the value for LoadBalancerId to be an explicit nil

### UnsetLoadBalancerId
`func (o *LoadBalancerListener) UnsetLoadBalancerId()`

UnsetLoadBalancerId ensures that no value is present for LoadBalancerId, not even an explicit nil
### GetTlsCertificateId

`func (o *LoadBalancerListener) GetTlsCertificateId() string`

GetTlsCertificateId returns the TlsCertificateId field if non-nil, zero value otherwise.

### GetTlsCertificateIdOk

`func (o *LoadBalancerListener) GetTlsCertificateIdOk() (*string, bool)`

GetTlsCertificateIdOk returns a tuple with the TlsCertificateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsCertificateId

`func (o *LoadBalancerListener) SetTlsCertificateId(v string)`

SetTlsCertificateId sets TlsCertificateId field to given value.

### HasTlsCertificateId

`func (o *LoadBalancerListener) HasTlsCertificateId() bool`

HasTlsCertificateId returns a boolean if a field has been set.

### SetTlsCertificateIdNil

`func (o *LoadBalancerListener) SetTlsCertificateIdNil(b bool)`

 SetTlsCertificateIdNil sets the value for TlsCertificateId to be an explicit nil

### UnsetTlsCertificateId
`func (o *LoadBalancerListener) UnsetTlsCertificateId()`

UnsetTlsCertificateId ensures that no value is present for TlsCertificateId, not even an explicit nil
### GetProvisioningStatus

`func (o *LoadBalancerListener) GetProvisioningStatus() ProvisioningStatus`

GetProvisioningStatus returns the ProvisioningStatus field if non-nil, zero value otherwise.

### GetProvisioningStatusOk

`func (o *LoadBalancerListener) GetProvisioningStatusOk() (*ProvisioningStatus, bool)`

GetProvisioningStatusOk returns a tuple with the ProvisioningStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioningStatus

`func (o *LoadBalancerListener) SetProvisioningStatus(v ProvisioningStatus)`

SetProvisioningStatus sets ProvisioningStatus field to given value.

### HasProvisioningStatus

`func (o *LoadBalancerListener) HasProvisioningStatus() bool`

HasProvisioningStatus returns a boolean if a field has been set.

### SetProvisioningStatusNil

`func (o *LoadBalancerListener) SetProvisioningStatusNil(b bool)`

 SetProvisioningStatusNil sets the value for ProvisioningStatus to be an explicit nil

### UnsetProvisioningStatus
`func (o *LoadBalancerListener) UnsetProvisioningStatus()`

UnsetProvisioningStatus ensures that no value is present for ProvisioningStatus, not even an explicit nil
### GetOperatingStatus

`func (o *LoadBalancerListener) GetOperatingStatus() LoadBalancerOperatingStatus`

GetOperatingStatus returns the OperatingStatus field if non-nil, zero value otherwise.

### GetOperatingStatusOk

`func (o *LoadBalancerListener) GetOperatingStatusOk() (*LoadBalancerOperatingStatus, bool)`

GetOperatingStatusOk returns a tuple with the OperatingStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatingStatus

`func (o *LoadBalancerListener) SetOperatingStatus(v LoadBalancerOperatingStatus)`

SetOperatingStatus sets OperatingStatus field to given value.

### HasOperatingStatus

`func (o *LoadBalancerListener) HasOperatingStatus() bool`

HasOperatingStatus returns a boolean if a field has been set.

### SetOperatingStatusNil

`func (o *LoadBalancerListener) SetOperatingStatusNil(b bool)`

 SetOperatingStatusNil sets the value for OperatingStatus to be an explicit nil

### UnsetOperatingStatus
`func (o *LoadBalancerListener) UnsetOperatingStatus()`

UnsetOperatingStatus ensures that no value is present for OperatingStatus, not even an explicit nil
### GetInsertHeaders

`func (o *LoadBalancerListener) GetInsertHeaders() map[string]LoadBalancerListenerInsertHeadersValue`

GetInsertHeaders returns the InsertHeaders field if non-nil, zero value otherwise.

### GetInsertHeadersOk

`func (o *LoadBalancerListener) GetInsertHeadersOk() (*map[string]LoadBalancerListenerInsertHeadersValue, bool)`

GetInsertHeadersOk returns a tuple with the InsertHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsertHeaders

`func (o *LoadBalancerListener) SetInsertHeaders(v map[string]LoadBalancerListenerInsertHeadersValue)`

SetInsertHeaders sets InsertHeaders field to given value.

### HasInsertHeaders

`func (o *LoadBalancerListener) HasInsertHeaders() bool`

HasInsertHeaders returns a boolean if a field has been set.

### SetInsertHeadersNil

`func (o *LoadBalancerListener) SetInsertHeadersNil(b bool)`

 SetInsertHeadersNil sets the value for InsertHeaders to be an explicit nil

### UnsetInsertHeaders
`func (o *LoadBalancerListener) UnsetInsertHeaders()`

UnsetInsertHeaders ensures that no value is present for InsertHeaders, not even an explicit nil
### GetCreatedAt

`func (o *LoadBalancerListener) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *LoadBalancerListener) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *LoadBalancerListener) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *LoadBalancerListener) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *LoadBalancerListener) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *LoadBalancerListener) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetUpdatedAt

`func (o *LoadBalancerListener) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *LoadBalancerListener) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *LoadBalancerListener) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *LoadBalancerListener) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *LoadBalancerListener) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *LoadBalancerListener) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetTimeoutClientData

`func (o *LoadBalancerListener) GetTimeoutClientData() int32`

GetTimeoutClientData returns the TimeoutClientData field if non-nil, zero value otherwise.

### GetTimeoutClientDataOk

`func (o *LoadBalancerListener) GetTimeoutClientDataOk() (*int32, bool)`

GetTimeoutClientDataOk returns a tuple with the TimeoutClientData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeoutClientData

`func (o *LoadBalancerListener) SetTimeoutClientData(v int32)`

SetTimeoutClientData sets TimeoutClientData field to given value.

### HasTimeoutClientData

`func (o *LoadBalancerListener) HasTimeoutClientData() bool`

HasTimeoutClientData returns a boolean if a field has been set.

### SetTimeoutClientDataNil

`func (o *LoadBalancerListener) SetTimeoutClientDataNil(b bool)`

 SetTimeoutClientDataNil sets the value for TimeoutClientData to be an explicit nil

### UnsetTimeoutClientData
`func (o *LoadBalancerListener) UnsetTimeoutClientData()`

UnsetTimeoutClientData ensures that no value is present for TimeoutClientData, not even an explicit nil
### GetDefaultTargetGroupName

`func (o *LoadBalancerListener) GetDefaultTargetGroupName() string`

GetDefaultTargetGroupName returns the DefaultTargetGroupName field if non-nil, zero value otherwise.

### GetDefaultTargetGroupNameOk

`func (o *LoadBalancerListener) GetDefaultTargetGroupNameOk() (*string, bool)`

GetDefaultTargetGroupNameOk returns a tuple with the DefaultTargetGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultTargetGroupName

`func (o *LoadBalancerListener) SetDefaultTargetGroupName(v string)`

SetDefaultTargetGroupName sets DefaultTargetGroupName field to given value.

### HasDefaultTargetGroupName

`func (o *LoadBalancerListener) HasDefaultTargetGroupName() bool`

HasDefaultTargetGroupName returns a boolean if a field has been set.

### SetDefaultTargetGroupNameNil

`func (o *LoadBalancerListener) SetDefaultTargetGroupNameNil(b bool)`

 SetDefaultTargetGroupNameNil sets the value for DefaultTargetGroupName to be an explicit nil

### UnsetDefaultTargetGroupName
`func (o *LoadBalancerListener) UnsetDefaultTargetGroupName()`

UnsetDefaultTargetGroupName ensures that no value is present for DefaultTargetGroupName, not even an explicit nil
### GetDefaultTargetGroupId

`func (o *LoadBalancerListener) GetDefaultTargetGroupId() string`

GetDefaultTargetGroupId returns the DefaultTargetGroupId field if non-nil, zero value otherwise.

### GetDefaultTargetGroupIdOk

`func (o *LoadBalancerListener) GetDefaultTargetGroupIdOk() (*string, bool)`

GetDefaultTargetGroupIdOk returns a tuple with the DefaultTargetGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultTargetGroupId

`func (o *LoadBalancerListener) SetDefaultTargetGroupId(v string)`

SetDefaultTargetGroupId sets DefaultTargetGroupId field to given value.

### HasDefaultTargetGroupId

`func (o *LoadBalancerListener) HasDefaultTargetGroupId() bool`

HasDefaultTargetGroupId returns a boolean if a field has been set.

### SetDefaultTargetGroupIdNil

`func (o *LoadBalancerListener) SetDefaultTargetGroupIdNil(b bool)`

 SetDefaultTargetGroupIdNil sets the value for DefaultTargetGroupId to be an explicit nil

### UnsetDefaultTargetGroupId
`func (o *LoadBalancerListener) UnsetDefaultTargetGroupId()`

UnsetDefaultTargetGroupId ensures that no value is present for DefaultTargetGroupId, not even an explicit nil
### GetLoadBalancerType

`func (o *LoadBalancerListener) GetLoadBalancerType() LoadBalancerType`

GetLoadBalancerType returns the LoadBalancerType field if non-nil, zero value otherwise.

### GetLoadBalancerTypeOk

`func (o *LoadBalancerListener) GetLoadBalancerTypeOk() (*LoadBalancerType, bool)`

GetLoadBalancerTypeOk returns a tuple with the LoadBalancerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadBalancerType

`func (o *LoadBalancerListener) SetLoadBalancerType(v LoadBalancerType)`

SetLoadBalancerType sets LoadBalancerType field to given value.

### HasLoadBalancerType

`func (o *LoadBalancerListener) HasLoadBalancerType() bool`

HasLoadBalancerType returns a boolean if a field has been set.

### SetLoadBalancerTypeNil

`func (o *LoadBalancerListener) SetLoadBalancerTypeNil(b bool)`

 SetLoadBalancerTypeNil sets the value for LoadBalancerType to be an explicit nil

### UnsetLoadBalancerType
`func (o *LoadBalancerListener) UnsetLoadBalancerType()`

UnsetLoadBalancerType ensures that no value is present for LoadBalancerType, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


