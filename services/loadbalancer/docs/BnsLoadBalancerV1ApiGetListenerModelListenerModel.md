# BnsLoadBalancerV1ApiGetListenerModelListenerModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | 리스너 ID | 
**Name** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**Protocol** | Pointer to [**NullableProtocol**](Protocol.md) |  | [optional] 
**IsEnabled** | Pointer to **NullableBool** |  | [optional] 
**Secrets** | Pointer to [**[]BnsLoadBalancerV1ApiGetListenerModelSecretModel**](BnsLoadBalancerV1ApiGetListenerModelSecretModel.md) |  | [optional] 
**L7Policies** | Pointer to [**[]BnsLoadBalancerV1ApiGetListenerModelL7PolicyModel**](BnsLoadBalancerV1ApiGetListenerModelL7PolicyModel.md) |  | [optional] 
**TlsCiphers** | Pointer to **NullableString** |  | [optional] 
**TlsVersions** | Pointer to [**[]TLSVersion**](TLSVersion.md) |  | [optional] 
**AlpnProtocols** | Pointer to **[]string** |  | [optional] 
**ProjectId** | Pointer to **NullableString** |  | [optional] 
**ProtocolPort** | Pointer to **NullableInt32** |  | [optional] 
**ConnectionLimit** | Pointer to **NullableInt32** |  | [optional] 
**LoadBalancerId** | Pointer to **NullableString** |  | [optional] 
**TlsCertificateId** | Pointer to **NullableString** |  | [optional] 
**ProvisioningStatus** | Pointer to [**NullableProvisioningStatus**](ProvisioningStatus.md) |  | [optional] 
**OperatingStatus** | Pointer to [**NullableLoadBalancerOperatingStatus**](LoadBalancerOperatingStatus.md) |  | [optional] 
**InsertHeaders** | Pointer to [**map[string]BnsLoadBalancerV1ApiGetListenerModelListenerModelInsertHeadersValue**](BnsLoadBalancerV1ApiGetListenerModelListenerModelInsertHeadersValue.md) |  | [optional] 
**CreatedAt** | Pointer to **NullableTime** |  | [optional] 
**UpdatedAt** | Pointer to **NullableTime** |  | [optional] 
**TimeoutClientData** | Pointer to **NullableInt32** |  | [optional] 
**DefaultTargetGroupName** | Pointer to **NullableString** |  | [optional] 
**DefaultTargetGroupId** | Pointer to **NullableString** |  | [optional] 
**LoadBalancerType** | Pointer to [**NullableLoadBalancerType**](LoadBalancerType.md) |  | [optional] 

## Methods

### NewBnsLoadBalancerV1ApiGetListenerModelListenerModel

`func NewBnsLoadBalancerV1ApiGetListenerModelListenerModel(id string, ) *BnsLoadBalancerV1ApiGetListenerModelListenerModel`

NewBnsLoadBalancerV1ApiGetListenerModelListenerModel instantiates a new BnsLoadBalancerV1ApiGetListenerModelListenerModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBnsLoadBalancerV1ApiGetListenerModelListenerModelWithDefaults

`func NewBnsLoadBalancerV1ApiGetListenerModelListenerModelWithDefaults() *BnsLoadBalancerV1ApiGetListenerModelListenerModel`

NewBnsLoadBalancerV1ApiGetListenerModelListenerModelWithDefaults instantiates a new BnsLoadBalancerV1ApiGetListenerModelListenerModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BnsLoadBalancerV1ApiGetListenerModelListenerModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BnsLoadBalancerV1ApiGetListenerModelListenerModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BnsLoadBalancerV1ApiGetListenerModelListenerModel) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *BnsLoadBalancerV1ApiGetListenerModelListenerModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BnsLoadBalancerV1ApiGetListenerModelListenerModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BnsLoadBalancerV1ApiGetListenerModelListenerModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BnsLoadBalancerV1ApiGetListenerModelListenerModel) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *BnsLoadBalancerV1ApiGetListenerModelListenerModel) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *BnsLoadBalancerV1ApiGetListenerModelListenerModel) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDescription

`func (o *BnsLoadBalancerV1ApiGetListenerModelListenerModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BnsLoadBalancerV1ApiGetListenerModelListenerModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BnsLoadBalancerV1ApiGetListenerModelListenerModel) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BnsLoadBalancerV1ApiGetListenerModelListenerModel) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *BnsLoadBalancerV1ApiGetListenerModelListenerModel) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *BnsLoadBalancerV1ApiGetListenerModelListenerModel) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetProtocol

`func (o *BnsLoadBalancerV1ApiGetListenerModelListenerModel) GetProtocol() Protocol`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *BnsLoadBalancerV1ApiGetListenerModelListenerModel) GetProtocolOk() (*Protocol, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *BnsLoadBalancerV1ApiGetListenerModelListenerModel) SetProtocol(v Protocol)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *BnsLoadBalancerV1ApiGetListenerModelListenerModel) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### SetProtocolNil

`func (o *BnsLoadBalancerV1ApiGetListenerModelListenerModel) SetProtocolNil(b bool)`

 SetProtocolNil sets the value for Protocol to be an explicit nil

### UnsetProtocol
`func (o *BnsLoadBalancerV1ApiGetListenerModelListenerModel) UnsetProtocol()`

UnsetProtocol ensures that no value is present for Protocol, not even an explicit nil
### GetIsEnabled

`func (o *BnsLoadBalancerV1ApiGetListenerModelListenerModel) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *BnsLoadBalancerV1ApiGetListenerModelListenerModel) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *BnsLoadBalancerV1ApiGetListenerModelListenerModel) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.

### HasIsEnabled

`func (o *BnsLoadBalancerV1ApiGetListenerModelListenerModel) HasIsEnabled() bool`

HasIsEnabled returns a boolean if a field has been set.

### SetIsEnabledNil

`func (o *BnsLoadBalancerV1ApiGetListenerModelListenerModel) SetIsEnabledNil(b bool)`

 SetIsEnabledNil sets the value for IsEnabled to be an explicit nil

### UnsetIsEnabled
`func (o *BnsLoadBalancerV1ApiGetListenerModelListenerModel) UnsetIsEnabled()`

UnsetIsEnabled ensures that no value is present for IsEnabled, not even an explicit nil
### GetSecrets

`func (o *BnsLoadBalancerV1ApiGetListenerModelListenerModel) GetSecrets() []BnsLoadBalancerV1ApiGetListenerModelSecretModel`

GetSecrets returns the Secrets field if non-nil, zero value otherwise.

### GetSecretsOk

`func (o *BnsLoadBalancerV1ApiGetListenerModelListenerModel) GetSecretsOk() (*[]BnsLoadBalancerV1ApiGetListenerModelSecretModel, bool)`

GetSecretsOk returns a tuple with the Secrets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecrets

`func (o *BnsLoadBalancerV1ApiGetListenerModelListenerModel) SetSecrets(v []BnsLoadBalancerV1ApiGetListenerModelSecretModel)`

SetSecrets sets Secrets field to given value.

### HasSecrets

`func (o *BnsLoadBalancerV1ApiGetListenerModelListenerModel) HasSecrets() bool`

HasSecrets returns a boolean if a field has been set.

### SetSecretsNil

`func (o *BnsLoadBalancerV1ApiGetListenerModelListenerModel) SetSecretsNil(b bool)`

 SetSecretsNil sets the value for Secrets to be an explicit nil

### UnsetSecrets
`func (o *BnsLoadBalancerV1ApiGetListenerModelListenerModel) UnsetSecrets()`

UnsetSecrets ensures that no value is present for Secrets, not even an explicit nil
### GetL7Policies

`func (o *BnsLoadBalancerV1ApiGetListenerModelListenerModel) GetL7Policies() []BnsLoadBalancerV1ApiGetListenerModelL7PolicyModel`

GetL7Policies returns the L7Policies field if non-nil, zero value otherwise.

### GetL7PoliciesOk

`func (o *BnsLoadBalancerV1ApiGetListenerModelListenerModel) GetL7PoliciesOk() (*[]BnsLoadBalancerV1ApiGetListenerModelL7PolicyModel, bool)`

GetL7PoliciesOk returns a tuple with the L7Policies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetL7Policies

`func (o *BnsLoadBalancerV1ApiGetListenerModelListenerModel) SetL7Policies(v []BnsLoadBalancerV1ApiGetListenerModelL7PolicyModel)`

SetL7Policies sets L7Policies field to given value.

### HasL7Policies

`func (o *BnsLoadBalancerV1ApiGetListenerModelListenerModel) HasL7Policies() bool`

HasL7Policies returns a boolean if a field has been set.

### SetL7PoliciesNil

`func (o *BnsLoadBalancerV1ApiGetListenerModelListenerModel) SetL7PoliciesNil(b bool)`

 SetL7PoliciesNil sets the value for L7Policies to be an explicit nil

### UnsetL7Policies
`func (o *BnsLoadBalancerV1ApiGetListenerModelListenerModel) UnsetL7Policies()`

UnsetL7Policies ensures that no value is present for L7Policies, not even an explicit nil
### GetTlsCiphers

`func (o *BnsLoadBalancerV1ApiGetListenerModelListenerModel) GetTlsCiphers() string`

GetTlsCiphers returns the TlsCiphers field if non-nil, zero value otherwise.

### GetTlsCiphersOk

`func (o *BnsLoadBalancerV1ApiGetListenerModelListenerModel) GetTlsCiphersOk() (*string, bool)`

GetTlsCiphersOk returns a tuple with the TlsCiphers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsCiphers

`func (o *BnsLoadBalancerV1ApiGetListenerModelListenerModel) SetTlsCiphers(v string)`

SetTlsCiphers sets TlsCiphers field to given value.

### HasTlsCiphers

`func (o *BnsLoadBalancerV1ApiGetListenerModelListenerModel) HasTlsCiphers() bool`

HasTlsCiphers returns a boolean if a field has been set.

### SetTlsCiphersNil

`func (o *BnsLoadBalancerV1ApiGetListenerModelListenerModel) SetTlsCiphersNil(b bool)`

 SetTlsCiphersNil sets the value for TlsCiphers to be an explicit nil

### UnsetTlsCiphers
`func (o *BnsLoadBalancerV1ApiGetListenerModelListenerModel) UnsetTlsCiphers()`

UnsetTlsCiphers ensures that no value is present for TlsCiphers, not even an explicit nil
### GetTlsVersions

`func (o *BnsLoadBalancerV1ApiGetListenerModelListenerModel) GetTlsVersions() []TLSVersion`

GetTlsVersions returns the TlsVersions field if non-nil, zero value otherwise.

### GetTlsVersionsOk

`func (o *BnsLoadBalancerV1ApiGetListenerModelListenerModel) GetTlsVersionsOk() (*[]TLSVersion, bool)`

GetTlsVersionsOk returns a tuple with the TlsVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsVersions

`func (o *BnsLoadBalancerV1ApiGetListenerModelListenerModel) SetTlsVersions(v []TLSVersion)`

SetTlsVersions sets TlsVersions field to given value.

### HasTlsVersions

`func (o *BnsLoadBalancerV1ApiGetListenerModelListenerModel) HasTlsVersions() bool`

HasTlsVersions returns a boolean if a field has been set.

### SetTlsVersionsNil

`func (o *BnsLoadBalancerV1ApiGetListenerModelListenerModel) SetTlsVersionsNil(b bool)`

 SetTlsVersionsNil sets the value for TlsVersions to be an explicit nil

### UnsetTlsVersions
`func (o *BnsLoadBalancerV1ApiGetListenerModelListenerModel) UnsetTlsVersions()`

UnsetTlsVersions ensures that no value is present for TlsVersions, not even an explicit nil
### GetAlpnProtocols

`func (o *BnsLoadBalancerV1ApiGetListenerModelListenerModel) GetAlpnProtocols() []string`

GetAlpnProtocols returns the AlpnProtocols field if non-nil, zero value otherwise.

### GetAlpnProtocolsOk

`func (o *BnsLoadBalancerV1ApiGetListenerModelListenerModel) GetAlpnProtocolsOk() (*[]string, bool)`

GetAlpnProtocolsOk returns a tuple with the AlpnProtocols field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlpnProtocols

`func (o *BnsLoadBalancerV1ApiGetListenerModelListenerModel) SetAlpnProtocols(v []string)`

SetAlpnProtocols sets AlpnProtocols field to given value.

### HasAlpnProtocols

`func (o *BnsLoadBalancerV1ApiGetListenerModelListenerModel) HasAlpnProtocols() bool`

HasAlpnProtocols returns a boolean if a field has been set.

### SetAlpnProtocolsNil

`func (o *BnsLoadBalancerV1ApiGetListenerModelListenerModel) SetAlpnProtocolsNil(b bool)`

 SetAlpnProtocolsNil sets the value for AlpnProtocols to be an explicit nil

### UnsetAlpnProtocols
`func (o *BnsLoadBalancerV1ApiGetListenerModelListenerModel) UnsetAlpnProtocols()`

UnsetAlpnProtocols ensures that no value is present for AlpnProtocols, not even an explicit nil
### GetProjectId

`func (o *BnsLoadBalancerV1ApiGetListenerModelListenerModel) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *BnsLoadBalancerV1ApiGetListenerModelListenerModel) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *BnsLoadBalancerV1ApiGetListenerModelListenerModel) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *BnsLoadBalancerV1ApiGetListenerModelListenerModel) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### SetProjectIdNil

`func (o *BnsLoadBalancerV1ApiGetListenerModelListenerModel) SetProjectIdNil(b bool)`

 SetProjectIdNil sets the value for ProjectId to be an explicit nil

### UnsetProjectId
`func (o *BnsLoadBalancerV1ApiGetListenerModelListenerModel) UnsetProjectId()`

UnsetProjectId ensures that no value is present for ProjectId, not even an explicit nil
### GetProtocolPort

`func (o *BnsLoadBalancerV1ApiGetListenerModelListenerModel) GetProtocolPort() int32`

GetProtocolPort returns the ProtocolPort field if non-nil, zero value otherwise.

### GetProtocolPortOk

`func (o *BnsLoadBalancerV1ApiGetListenerModelListenerModel) GetProtocolPortOk() (*int32, bool)`

GetProtocolPortOk returns a tuple with the ProtocolPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocolPort

`func (o *BnsLoadBalancerV1ApiGetListenerModelListenerModel) SetProtocolPort(v int32)`

SetProtocolPort sets ProtocolPort field to given value.

### HasProtocolPort

`func (o *BnsLoadBalancerV1ApiGetListenerModelListenerModel) HasProtocolPort() bool`

HasProtocolPort returns a boolean if a field has been set.

### SetProtocolPortNil

`func (o *BnsLoadBalancerV1ApiGetListenerModelListenerModel) SetProtocolPortNil(b bool)`

 SetProtocolPortNil sets the value for ProtocolPort to be an explicit nil

### UnsetProtocolPort
`func (o *BnsLoadBalancerV1ApiGetListenerModelListenerModel) UnsetProtocolPort()`

UnsetProtocolPort ensures that no value is present for ProtocolPort, not even an explicit nil
### GetConnectionLimit

`func (o *BnsLoadBalancerV1ApiGetListenerModelListenerModel) GetConnectionLimit() int32`

GetConnectionLimit returns the ConnectionLimit field if non-nil, zero value otherwise.

### GetConnectionLimitOk

`func (o *BnsLoadBalancerV1ApiGetListenerModelListenerModel) GetConnectionLimitOk() (*int32, bool)`

GetConnectionLimitOk returns a tuple with the ConnectionLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionLimit

`func (o *BnsLoadBalancerV1ApiGetListenerModelListenerModel) SetConnectionLimit(v int32)`

SetConnectionLimit sets ConnectionLimit field to given value.

### HasConnectionLimit

`func (o *BnsLoadBalancerV1ApiGetListenerModelListenerModel) HasConnectionLimit() bool`

HasConnectionLimit returns a boolean if a field has been set.

### SetConnectionLimitNil

`func (o *BnsLoadBalancerV1ApiGetListenerModelListenerModel) SetConnectionLimitNil(b bool)`

 SetConnectionLimitNil sets the value for ConnectionLimit to be an explicit nil

### UnsetConnectionLimit
`func (o *BnsLoadBalancerV1ApiGetListenerModelListenerModel) UnsetConnectionLimit()`

UnsetConnectionLimit ensures that no value is present for ConnectionLimit, not even an explicit nil
### GetLoadBalancerId

`func (o *BnsLoadBalancerV1ApiGetListenerModelListenerModel) GetLoadBalancerId() string`

GetLoadBalancerId returns the LoadBalancerId field if non-nil, zero value otherwise.

### GetLoadBalancerIdOk

`func (o *BnsLoadBalancerV1ApiGetListenerModelListenerModel) GetLoadBalancerIdOk() (*string, bool)`

GetLoadBalancerIdOk returns a tuple with the LoadBalancerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadBalancerId

`func (o *BnsLoadBalancerV1ApiGetListenerModelListenerModel) SetLoadBalancerId(v string)`

SetLoadBalancerId sets LoadBalancerId field to given value.

### HasLoadBalancerId

`func (o *BnsLoadBalancerV1ApiGetListenerModelListenerModel) HasLoadBalancerId() bool`

HasLoadBalancerId returns a boolean if a field has been set.

### SetLoadBalancerIdNil

`func (o *BnsLoadBalancerV1ApiGetListenerModelListenerModel) SetLoadBalancerIdNil(b bool)`

 SetLoadBalancerIdNil sets the value for LoadBalancerId to be an explicit nil

### UnsetLoadBalancerId
`func (o *BnsLoadBalancerV1ApiGetListenerModelListenerModel) UnsetLoadBalancerId()`

UnsetLoadBalancerId ensures that no value is present for LoadBalancerId, not even an explicit nil
### GetTlsCertificateId

`func (o *BnsLoadBalancerV1ApiGetListenerModelListenerModel) GetTlsCertificateId() string`

GetTlsCertificateId returns the TlsCertificateId field if non-nil, zero value otherwise.

### GetTlsCertificateIdOk

`func (o *BnsLoadBalancerV1ApiGetListenerModelListenerModel) GetTlsCertificateIdOk() (*string, bool)`

GetTlsCertificateIdOk returns a tuple with the TlsCertificateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsCertificateId

`func (o *BnsLoadBalancerV1ApiGetListenerModelListenerModel) SetTlsCertificateId(v string)`

SetTlsCertificateId sets TlsCertificateId field to given value.

### HasTlsCertificateId

`func (o *BnsLoadBalancerV1ApiGetListenerModelListenerModel) HasTlsCertificateId() bool`

HasTlsCertificateId returns a boolean if a field has been set.

### SetTlsCertificateIdNil

`func (o *BnsLoadBalancerV1ApiGetListenerModelListenerModel) SetTlsCertificateIdNil(b bool)`

 SetTlsCertificateIdNil sets the value for TlsCertificateId to be an explicit nil

### UnsetTlsCertificateId
`func (o *BnsLoadBalancerV1ApiGetListenerModelListenerModel) UnsetTlsCertificateId()`

UnsetTlsCertificateId ensures that no value is present for TlsCertificateId, not even an explicit nil
### GetProvisioningStatus

`func (o *BnsLoadBalancerV1ApiGetListenerModelListenerModel) GetProvisioningStatus() ProvisioningStatus`

GetProvisioningStatus returns the ProvisioningStatus field if non-nil, zero value otherwise.

### GetProvisioningStatusOk

`func (o *BnsLoadBalancerV1ApiGetListenerModelListenerModel) GetProvisioningStatusOk() (*ProvisioningStatus, bool)`

GetProvisioningStatusOk returns a tuple with the ProvisioningStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioningStatus

`func (o *BnsLoadBalancerV1ApiGetListenerModelListenerModel) SetProvisioningStatus(v ProvisioningStatus)`

SetProvisioningStatus sets ProvisioningStatus field to given value.

### HasProvisioningStatus

`func (o *BnsLoadBalancerV1ApiGetListenerModelListenerModel) HasProvisioningStatus() bool`

HasProvisioningStatus returns a boolean if a field has been set.

### SetProvisioningStatusNil

`func (o *BnsLoadBalancerV1ApiGetListenerModelListenerModel) SetProvisioningStatusNil(b bool)`

 SetProvisioningStatusNil sets the value for ProvisioningStatus to be an explicit nil

### UnsetProvisioningStatus
`func (o *BnsLoadBalancerV1ApiGetListenerModelListenerModel) UnsetProvisioningStatus()`

UnsetProvisioningStatus ensures that no value is present for ProvisioningStatus, not even an explicit nil
### GetOperatingStatus

`func (o *BnsLoadBalancerV1ApiGetListenerModelListenerModel) GetOperatingStatus() LoadBalancerOperatingStatus`

GetOperatingStatus returns the OperatingStatus field if non-nil, zero value otherwise.

### GetOperatingStatusOk

`func (o *BnsLoadBalancerV1ApiGetListenerModelListenerModel) GetOperatingStatusOk() (*LoadBalancerOperatingStatus, bool)`

GetOperatingStatusOk returns a tuple with the OperatingStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatingStatus

`func (o *BnsLoadBalancerV1ApiGetListenerModelListenerModel) SetOperatingStatus(v LoadBalancerOperatingStatus)`

SetOperatingStatus sets OperatingStatus field to given value.

### HasOperatingStatus

`func (o *BnsLoadBalancerV1ApiGetListenerModelListenerModel) HasOperatingStatus() bool`

HasOperatingStatus returns a boolean if a field has been set.

### SetOperatingStatusNil

`func (o *BnsLoadBalancerV1ApiGetListenerModelListenerModel) SetOperatingStatusNil(b bool)`

 SetOperatingStatusNil sets the value for OperatingStatus to be an explicit nil

### UnsetOperatingStatus
`func (o *BnsLoadBalancerV1ApiGetListenerModelListenerModel) UnsetOperatingStatus()`

UnsetOperatingStatus ensures that no value is present for OperatingStatus, not even an explicit nil
### GetInsertHeaders

`func (o *BnsLoadBalancerV1ApiGetListenerModelListenerModel) GetInsertHeaders() map[string]BnsLoadBalancerV1ApiGetListenerModelListenerModelInsertHeadersValue`

GetInsertHeaders returns the InsertHeaders field if non-nil, zero value otherwise.

### GetInsertHeadersOk

`func (o *BnsLoadBalancerV1ApiGetListenerModelListenerModel) GetInsertHeadersOk() (*map[string]BnsLoadBalancerV1ApiGetListenerModelListenerModelInsertHeadersValue, bool)`

GetInsertHeadersOk returns a tuple with the InsertHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsertHeaders

`func (o *BnsLoadBalancerV1ApiGetListenerModelListenerModel) SetInsertHeaders(v map[string]BnsLoadBalancerV1ApiGetListenerModelListenerModelInsertHeadersValue)`

SetInsertHeaders sets InsertHeaders field to given value.

### HasInsertHeaders

`func (o *BnsLoadBalancerV1ApiGetListenerModelListenerModel) HasInsertHeaders() bool`

HasInsertHeaders returns a boolean if a field has been set.

### SetInsertHeadersNil

`func (o *BnsLoadBalancerV1ApiGetListenerModelListenerModel) SetInsertHeadersNil(b bool)`

 SetInsertHeadersNil sets the value for InsertHeaders to be an explicit nil

### UnsetInsertHeaders
`func (o *BnsLoadBalancerV1ApiGetListenerModelListenerModel) UnsetInsertHeaders()`

UnsetInsertHeaders ensures that no value is present for InsertHeaders, not even an explicit nil
### GetCreatedAt

`func (o *BnsLoadBalancerV1ApiGetListenerModelListenerModel) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *BnsLoadBalancerV1ApiGetListenerModelListenerModel) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *BnsLoadBalancerV1ApiGetListenerModelListenerModel) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *BnsLoadBalancerV1ApiGetListenerModelListenerModel) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *BnsLoadBalancerV1ApiGetListenerModelListenerModel) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *BnsLoadBalancerV1ApiGetListenerModelListenerModel) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetUpdatedAt

`func (o *BnsLoadBalancerV1ApiGetListenerModelListenerModel) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *BnsLoadBalancerV1ApiGetListenerModelListenerModel) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *BnsLoadBalancerV1ApiGetListenerModelListenerModel) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *BnsLoadBalancerV1ApiGetListenerModelListenerModel) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *BnsLoadBalancerV1ApiGetListenerModelListenerModel) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *BnsLoadBalancerV1ApiGetListenerModelListenerModel) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetTimeoutClientData

`func (o *BnsLoadBalancerV1ApiGetListenerModelListenerModel) GetTimeoutClientData() int32`

GetTimeoutClientData returns the TimeoutClientData field if non-nil, zero value otherwise.

### GetTimeoutClientDataOk

`func (o *BnsLoadBalancerV1ApiGetListenerModelListenerModel) GetTimeoutClientDataOk() (*int32, bool)`

GetTimeoutClientDataOk returns a tuple with the TimeoutClientData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeoutClientData

`func (o *BnsLoadBalancerV1ApiGetListenerModelListenerModel) SetTimeoutClientData(v int32)`

SetTimeoutClientData sets TimeoutClientData field to given value.

### HasTimeoutClientData

`func (o *BnsLoadBalancerV1ApiGetListenerModelListenerModel) HasTimeoutClientData() bool`

HasTimeoutClientData returns a boolean if a field has been set.

### SetTimeoutClientDataNil

`func (o *BnsLoadBalancerV1ApiGetListenerModelListenerModel) SetTimeoutClientDataNil(b bool)`

 SetTimeoutClientDataNil sets the value for TimeoutClientData to be an explicit nil

### UnsetTimeoutClientData
`func (o *BnsLoadBalancerV1ApiGetListenerModelListenerModel) UnsetTimeoutClientData()`

UnsetTimeoutClientData ensures that no value is present for TimeoutClientData, not even an explicit nil
### GetDefaultTargetGroupName

`func (o *BnsLoadBalancerV1ApiGetListenerModelListenerModel) GetDefaultTargetGroupName() string`

GetDefaultTargetGroupName returns the DefaultTargetGroupName field if non-nil, zero value otherwise.

### GetDefaultTargetGroupNameOk

`func (o *BnsLoadBalancerV1ApiGetListenerModelListenerModel) GetDefaultTargetGroupNameOk() (*string, bool)`

GetDefaultTargetGroupNameOk returns a tuple with the DefaultTargetGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultTargetGroupName

`func (o *BnsLoadBalancerV1ApiGetListenerModelListenerModel) SetDefaultTargetGroupName(v string)`

SetDefaultTargetGroupName sets DefaultTargetGroupName field to given value.

### HasDefaultTargetGroupName

`func (o *BnsLoadBalancerV1ApiGetListenerModelListenerModel) HasDefaultTargetGroupName() bool`

HasDefaultTargetGroupName returns a boolean if a field has been set.

### SetDefaultTargetGroupNameNil

`func (o *BnsLoadBalancerV1ApiGetListenerModelListenerModel) SetDefaultTargetGroupNameNil(b bool)`

 SetDefaultTargetGroupNameNil sets the value for DefaultTargetGroupName to be an explicit nil

### UnsetDefaultTargetGroupName
`func (o *BnsLoadBalancerV1ApiGetListenerModelListenerModel) UnsetDefaultTargetGroupName()`

UnsetDefaultTargetGroupName ensures that no value is present for DefaultTargetGroupName, not even an explicit nil
### GetDefaultTargetGroupId

`func (o *BnsLoadBalancerV1ApiGetListenerModelListenerModel) GetDefaultTargetGroupId() string`

GetDefaultTargetGroupId returns the DefaultTargetGroupId field if non-nil, zero value otherwise.

### GetDefaultTargetGroupIdOk

`func (o *BnsLoadBalancerV1ApiGetListenerModelListenerModel) GetDefaultTargetGroupIdOk() (*string, bool)`

GetDefaultTargetGroupIdOk returns a tuple with the DefaultTargetGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultTargetGroupId

`func (o *BnsLoadBalancerV1ApiGetListenerModelListenerModel) SetDefaultTargetGroupId(v string)`

SetDefaultTargetGroupId sets DefaultTargetGroupId field to given value.

### HasDefaultTargetGroupId

`func (o *BnsLoadBalancerV1ApiGetListenerModelListenerModel) HasDefaultTargetGroupId() bool`

HasDefaultTargetGroupId returns a boolean if a field has been set.

### SetDefaultTargetGroupIdNil

`func (o *BnsLoadBalancerV1ApiGetListenerModelListenerModel) SetDefaultTargetGroupIdNil(b bool)`

 SetDefaultTargetGroupIdNil sets the value for DefaultTargetGroupId to be an explicit nil

### UnsetDefaultTargetGroupId
`func (o *BnsLoadBalancerV1ApiGetListenerModelListenerModel) UnsetDefaultTargetGroupId()`

UnsetDefaultTargetGroupId ensures that no value is present for DefaultTargetGroupId, not even an explicit nil
### GetLoadBalancerType

`func (o *BnsLoadBalancerV1ApiGetListenerModelListenerModel) GetLoadBalancerType() LoadBalancerType`

GetLoadBalancerType returns the LoadBalancerType field if non-nil, zero value otherwise.

### GetLoadBalancerTypeOk

`func (o *BnsLoadBalancerV1ApiGetListenerModelListenerModel) GetLoadBalancerTypeOk() (*LoadBalancerType, bool)`

GetLoadBalancerTypeOk returns a tuple with the LoadBalancerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadBalancerType

`func (o *BnsLoadBalancerV1ApiGetListenerModelListenerModel) SetLoadBalancerType(v LoadBalancerType)`

SetLoadBalancerType sets LoadBalancerType field to given value.

### HasLoadBalancerType

`func (o *BnsLoadBalancerV1ApiGetListenerModelListenerModel) HasLoadBalancerType() bool`

HasLoadBalancerType returns a boolean if a field has been set.

### SetLoadBalancerTypeNil

`func (o *BnsLoadBalancerV1ApiGetListenerModelListenerModel) SetLoadBalancerTypeNil(b bool)`

 SetLoadBalancerTypeNil sets the value for LoadBalancerType to be an explicit nil

### UnsetLoadBalancerType
`func (o *BnsLoadBalancerV1ApiGetListenerModelListenerModel) UnsetLoadBalancerType()`

UnsetLoadBalancerType ensures that no value is present for LoadBalancerType, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


