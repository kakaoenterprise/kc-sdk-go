# BnsLoadBalancerV1ApiListListenersModelListenerModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | 리스너 ID | 
**Name** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**Protocol** | Pointer to [**NullableProtocol**](Protocol.md) |  | [optional] 
**IsEnabled** | Pointer to **NullableBool** |  | [optional] 
**Secrets** | Pointer to [**[]BnsLoadBalancerV1ApiListListenersModelSecretModel**](BnsLoadBalancerV1ApiListListenersModelSecretModel.md) |  | [optional] 
**L7Policies** | Pointer to [**[]BnsLoadBalancerV1ApiListListenersModelL7PolicyModel**](BnsLoadBalancerV1ApiListListenersModelL7PolicyModel.md) |  | [optional] 
**TlsCiphers** | Pointer to **NullableString** |  | [optional] 
**TlsVersions** | Pointer to [**[]TLSVersion**](TLSVersion.md) |  | [optional] 
**ProjectId** | Pointer to **NullableString** |  | [optional] 
**ProtocolPort** | Pointer to **NullableInt32** |  | [optional] 
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

### NewBnsLoadBalancerV1ApiListListenersModelListenerModel

`func NewBnsLoadBalancerV1ApiListListenersModelListenerModel(id string, ) *BnsLoadBalancerV1ApiListListenersModelListenerModel`

NewBnsLoadBalancerV1ApiListListenersModelListenerModel instantiates a new BnsLoadBalancerV1ApiListListenersModelListenerModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBnsLoadBalancerV1ApiListListenersModelListenerModelWithDefaults

`func NewBnsLoadBalancerV1ApiListListenersModelListenerModelWithDefaults() *BnsLoadBalancerV1ApiListListenersModelListenerModel`

NewBnsLoadBalancerV1ApiListListenersModelListenerModelWithDefaults instantiates a new BnsLoadBalancerV1ApiListListenersModelListenerModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BnsLoadBalancerV1ApiListListenersModelListenerModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BnsLoadBalancerV1ApiListListenersModelListenerModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BnsLoadBalancerV1ApiListListenersModelListenerModel) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *BnsLoadBalancerV1ApiListListenersModelListenerModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BnsLoadBalancerV1ApiListListenersModelListenerModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BnsLoadBalancerV1ApiListListenersModelListenerModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BnsLoadBalancerV1ApiListListenersModelListenerModel) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *BnsLoadBalancerV1ApiListListenersModelListenerModel) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *BnsLoadBalancerV1ApiListListenersModelListenerModel) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDescription

`func (o *BnsLoadBalancerV1ApiListListenersModelListenerModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BnsLoadBalancerV1ApiListListenersModelListenerModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BnsLoadBalancerV1ApiListListenersModelListenerModel) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BnsLoadBalancerV1ApiListListenersModelListenerModel) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *BnsLoadBalancerV1ApiListListenersModelListenerModel) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *BnsLoadBalancerV1ApiListListenersModelListenerModel) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetProtocol

`func (o *BnsLoadBalancerV1ApiListListenersModelListenerModel) GetProtocol() Protocol`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *BnsLoadBalancerV1ApiListListenersModelListenerModel) GetProtocolOk() (*Protocol, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *BnsLoadBalancerV1ApiListListenersModelListenerModel) SetProtocol(v Protocol)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *BnsLoadBalancerV1ApiListListenersModelListenerModel) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### SetProtocolNil

`func (o *BnsLoadBalancerV1ApiListListenersModelListenerModel) SetProtocolNil(b bool)`

 SetProtocolNil sets the value for Protocol to be an explicit nil

### UnsetProtocol
`func (o *BnsLoadBalancerV1ApiListListenersModelListenerModel) UnsetProtocol()`

UnsetProtocol ensures that no value is present for Protocol, not even an explicit nil
### GetIsEnabled

`func (o *BnsLoadBalancerV1ApiListListenersModelListenerModel) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *BnsLoadBalancerV1ApiListListenersModelListenerModel) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *BnsLoadBalancerV1ApiListListenersModelListenerModel) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.

### HasIsEnabled

`func (o *BnsLoadBalancerV1ApiListListenersModelListenerModel) HasIsEnabled() bool`

HasIsEnabled returns a boolean if a field has been set.

### SetIsEnabledNil

`func (o *BnsLoadBalancerV1ApiListListenersModelListenerModel) SetIsEnabledNil(b bool)`

 SetIsEnabledNil sets the value for IsEnabled to be an explicit nil

### UnsetIsEnabled
`func (o *BnsLoadBalancerV1ApiListListenersModelListenerModel) UnsetIsEnabled()`

UnsetIsEnabled ensures that no value is present for IsEnabled, not even an explicit nil
### GetSecrets

`func (o *BnsLoadBalancerV1ApiListListenersModelListenerModel) GetSecrets() []BnsLoadBalancerV1ApiListListenersModelSecretModel`

GetSecrets returns the Secrets field if non-nil, zero value otherwise.

### GetSecretsOk

`func (o *BnsLoadBalancerV1ApiListListenersModelListenerModel) GetSecretsOk() (*[]BnsLoadBalancerV1ApiListListenersModelSecretModel, bool)`

GetSecretsOk returns a tuple with the Secrets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecrets

`func (o *BnsLoadBalancerV1ApiListListenersModelListenerModel) SetSecrets(v []BnsLoadBalancerV1ApiListListenersModelSecretModel)`

SetSecrets sets Secrets field to given value.

### HasSecrets

`func (o *BnsLoadBalancerV1ApiListListenersModelListenerModel) HasSecrets() bool`

HasSecrets returns a boolean if a field has been set.

### SetSecretsNil

`func (o *BnsLoadBalancerV1ApiListListenersModelListenerModel) SetSecretsNil(b bool)`

 SetSecretsNil sets the value for Secrets to be an explicit nil

### UnsetSecrets
`func (o *BnsLoadBalancerV1ApiListListenersModelListenerModel) UnsetSecrets()`

UnsetSecrets ensures that no value is present for Secrets, not even an explicit nil
### GetL7Policies

`func (o *BnsLoadBalancerV1ApiListListenersModelListenerModel) GetL7Policies() []BnsLoadBalancerV1ApiListListenersModelL7PolicyModel`

GetL7Policies returns the L7Policies field if non-nil, zero value otherwise.

### GetL7PoliciesOk

`func (o *BnsLoadBalancerV1ApiListListenersModelListenerModel) GetL7PoliciesOk() (*[]BnsLoadBalancerV1ApiListListenersModelL7PolicyModel, bool)`

GetL7PoliciesOk returns a tuple with the L7Policies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetL7Policies

`func (o *BnsLoadBalancerV1ApiListListenersModelListenerModel) SetL7Policies(v []BnsLoadBalancerV1ApiListListenersModelL7PolicyModel)`

SetL7Policies sets L7Policies field to given value.

### HasL7Policies

`func (o *BnsLoadBalancerV1ApiListListenersModelListenerModel) HasL7Policies() bool`

HasL7Policies returns a boolean if a field has been set.

### SetL7PoliciesNil

`func (o *BnsLoadBalancerV1ApiListListenersModelListenerModel) SetL7PoliciesNil(b bool)`

 SetL7PoliciesNil sets the value for L7Policies to be an explicit nil

### UnsetL7Policies
`func (o *BnsLoadBalancerV1ApiListListenersModelListenerModel) UnsetL7Policies()`

UnsetL7Policies ensures that no value is present for L7Policies, not even an explicit nil
### GetTlsCiphers

`func (o *BnsLoadBalancerV1ApiListListenersModelListenerModel) GetTlsCiphers() string`

GetTlsCiphers returns the TlsCiphers field if non-nil, zero value otherwise.

### GetTlsCiphersOk

`func (o *BnsLoadBalancerV1ApiListListenersModelListenerModel) GetTlsCiphersOk() (*string, bool)`

GetTlsCiphersOk returns a tuple with the TlsCiphers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsCiphers

`func (o *BnsLoadBalancerV1ApiListListenersModelListenerModel) SetTlsCiphers(v string)`

SetTlsCiphers sets TlsCiphers field to given value.

### HasTlsCiphers

`func (o *BnsLoadBalancerV1ApiListListenersModelListenerModel) HasTlsCiphers() bool`

HasTlsCiphers returns a boolean if a field has been set.

### SetTlsCiphersNil

`func (o *BnsLoadBalancerV1ApiListListenersModelListenerModel) SetTlsCiphersNil(b bool)`

 SetTlsCiphersNil sets the value for TlsCiphers to be an explicit nil

### UnsetTlsCiphers
`func (o *BnsLoadBalancerV1ApiListListenersModelListenerModel) UnsetTlsCiphers()`

UnsetTlsCiphers ensures that no value is present for TlsCiphers, not even an explicit nil
### GetTlsVersions

`func (o *BnsLoadBalancerV1ApiListListenersModelListenerModel) GetTlsVersions() []TLSVersion`

GetTlsVersions returns the TlsVersions field if non-nil, zero value otherwise.

### GetTlsVersionsOk

`func (o *BnsLoadBalancerV1ApiListListenersModelListenerModel) GetTlsVersionsOk() (*[]TLSVersion, bool)`

GetTlsVersionsOk returns a tuple with the TlsVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsVersions

`func (o *BnsLoadBalancerV1ApiListListenersModelListenerModel) SetTlsVersions(v []TLSVersion)`

SetTlsVersions sets TlsVersions field to given value.

### HasTlsVersions

`func (o *BnsLoadBalancerV1ApiListListenersModelListenerModel) HasTlsVersions() bool`

HasTlsVersions returns a boolean if a field has been set.

### SetTlsVersionsNil

`func (o *BnsLoadBalancerV1ApiListListenersModelListenerModel) SetTlsVersionsNil(b bool)`

 SetTlsVersionsNil sets the value for TlsVersions to be an explicit nil

### UnsetTlsVersions
`func (o *BnsLoadBalancerV1ApiListListenersModelListenerModel) UnsetTlsVersions()`

UnsetTlsVersions ensures that no value is present for TlsVersions, not even an explicit nil
### GetProjectId

`func (o *BnsLoadBalancerV1ApiListListenersModelListenerModel) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *BnsLoadBalancerV1ApiListListenersModelListenerModel) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *BnsLoadBalancerV1ApiListListenersModelListenerModel) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *BnsLoadBalancerV1ApiListListenersModelListenerModel) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### SetProjectIdNil

`func (o *BnsLoadBalancerV1ApiListListenersModelListenerModel) SetProjectIdNil(b bool)`

 SetProjectIdNil sets the value for ProjectId to be an explicit nil

### UnsetProjectId
`func (o *BnsLoadBalancerV1ApiListListenersModelListenerModel) UnsetProjectId()`

UnsetProjectId ensures that no value is present for ProjectId, not even an explicit nil
### GetProtocolPort

`func (o *BnsLoadBalancerV1ApiListListenersModelListenerModel) GetProtocolPort() int32`

GetProtocolPort returns the ProtocolPort field if non-nil, zero value otherwise.

### GetProtocolPortOk

`func (o *BnsLoadBalancerV1ApiListListenersModelListenerModel) GetProtocolPortOk() (*int32, bool)`

GetProtocolPortOk returns a tuple with the ProtocolPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocolPort

`func (o *BnsLoadBalancerV1ApiListListenersModelListenerModel) SetProtocolPort(v int32)`

SetProtocolPort sets ProtocolPort field to given value.

### HasProtocolPort

`func (o *BnsLoadBalancerV1ApiListListenersModelListenerModel) HasProtocolPort() bool`

HasProtocolPort returns a boolean if a field has been set.

### SetProtocolPortNil

`func (o *BnsLoadBalancerV1ApiListListenersModelListenerModel) SetProtocolPortNil(b bool)`

 SetProtocolPortNil sets the value for ProtocolPort to be an explicit nil

### UnsetProtocolPort
`func (o *BnsLoadBalancerV1ApiListListenersModelListenerModel) UnsetProtocolPort()`

UnsetProtocolPort ensures that no value is present for ProtocolPort, not even an explicit nil
### GetLoadBalancerId

`func (o *BnsLoadBalancerV1ApiListListenersModelListenerModel) GetLoadBalancerId() string`

GetLoadBalancerId returns the LoadBalancerId field if non-nil, zero value otherwise.

### GetLoadBalancerIdOk

`func (o *BnsLoadBalancerV1ApiListListenersModelListenerModel) GetLoadBalancerIdOk() (*string, bool)`

GetLoadBalancerIdOk returns a tuple with the LoadBalancerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadBalancerId

`func (o *BnsLoadBalancerV1ApiListListenersModelListenerModel) SetLoadBalancerId(v string)`

SetLoadBalancerId sets LoadBalancerId field to given value.

### HasLoadBalancerId

`func (o *BnsLoadBalancerV1ApiListListenersModelListenerModel) HasLoadBalancerId() bool`

HasLoadBalancerId returns a boolean if a field has been set.

### SetLoadBalancerIdNil

`func (o *BnsLoadBalancerV1ApiListListenersModelListenerModel) SetLoadBalancerIdNil(b bool)`

 SetLoadBalancerIdNil sets the value for LoadBalancerId to be an explicit nil

### UnsetLoadBalancerId
`func (o *BnsLoadBalancerV1ApiListListenersModelListenerModel) UnsetLoadBalancerId()`

UnsetLoadBalancerId ensures that no value is present for LoadBalancerId, not even an explicit nil
### GetTlsCertificateId

`func (o *BnsLoadBalancerV1ApiListListenersModelListenerModel) GetTlsCertificateId() string`

GetTlsCertificateId returns the TlsCertificateId field if non-nil, zero value otherwise.

### GetTlsCertificateIdOk

`func (o *BnsLoadBalancerV1ApiListListenersModelListenerModel) GetTlsCertificateIdOk() (*string, bool)`

GetTlsCertificateIdOk returns a tuple with the TlsCertificateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsCertificateId

`func (o *BnsLoadBalancerV1ApiListListenersModelListenerModel) SetTlsCertificateId(v string)`

SetTlsCertificateId sets TlsCertificateId field to given value.

### HasTlsCertificateId

`func (o *BnsLoadBalancerV1ApiListListenersModelListenerModel) HasTlsCertificateId() bool`

HasTlsCertificateId returns a boolean if a field has been set.

### SetTlsCertificateIdNil

`func (o *BnsLoadBalancerV1ApiListListenersModelListenerModel) SetTlsCertificateIdNil(b bool)`

 SetTlsCertificateIdNil sets the value for TlsCertificateId to be an explicit nil

### UnsetTlsCertificateId
`func (o *BnsLoadBalancerV1ApiListListenersModelListenerModel) UnsetTlsCertificateId()`

UnsetTlsCertificateId ensures that no value is present for TlsCertificateId, not even an explicit nil
### GetProvisioningStatus

`func (o *BnsLoadBalancerV1ApiListListenersModelListenerModel) GetProvisioningStatus() ProvisioningStatus`

GetProvisioningStatus returns the ProvisioningStatus field if non-nil, zero value otherwise.

### GetProvisioningStatusOk

`func (o *BnsLoadBalancerV1ApiListListenersModelListenerModel) GetProvisioningStatusOk() (*ProvisioningStatus, bool)`

GetProvisioningStatusOk returns a tuple with the ProvisioningStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioningStatus

`func (o *BnsLoadBalancerV1ApiListListenersModelListenerModel) SetProvisioningStatus(v ProvisioningStatus)`

SetProvisioningStatus sets ProvisioningStatus field to given value.

### HasProvisioningStatus

`func (o *BnsLoadBalancerV1ApiListListenersModelListenerModel) HasProvisioningStatus() bool`

HasProvisioningStatus returns a boolean if a field has been set.

### SetProvisioningStatusNil

`func (o *BnsLoadBalancerV1ApiListListenersModelListenerModel) SetProvisioningStatusNil(b bool)`

 SetProvisioningStatusNil sets the value for ProvisioningStatus to be an explicit nil

### UnsetProvisioningStatus
`func (o *BnsLoadBalancerV1ApiListListenersModelListenerModel) UnsetProvisioningStatus()`

UnsetProvisioningStatus ensures that no value is present for ProvisioningStatus, not even an explicit nil
### GetOperatingStatus

`func (o *BnsLoadBalancerV1ApiListListenersModelListenerModel) GetOperatingStatus() LoadBalancerOperatingStatus`

GetOperatingStatus returns the OperatingStatus field if non-nil, zero value otherwise.

### GetOperatingStatusOk

`func (o *BnsLoadBalancerV1ApiListListenersModelListenerModel) GetOperatingStatusOk() (*LoadBalancerOperatingStatus, bool)`

GetOperatingStatusOk returns a tuple with the OperatingStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatingStatus

`func (o *BnsLoadBalancerV1ApiListListenersModelListenerModel) SetOperatingStatus(v LoadBalancerOperatingStatus)`

SetOperatingStatus sets OperatingStatus field to given value.

### HasOperatingStatus

`func (o *BnsLoadBalancerV1ApiListListenersModelListenerModel) HasOperatingStatus() bool`

HasOperatingStatus returns a boolean if a field has been set.

### SetOperatingStatusNil

`func (o *BnsLoadBalancerV1ApiListListenersModelListenerModel) SetOperatingStatusNil(b bool)`

 SetOperatingStatusNil sets the value for OperatingStatus to be an explicit nil

### UnsetOperatingStatus
`func (o *BnsLoadBalancerV1ApiListListenersModelListenerModel) UnsetOperatingStatus()`

UnsetOperatingStatus ensures that no value is present for OperatingStatus, not even an explicit nil
### GetInsertHeaders

`func (o *BnsLoadBalancerV1ApiListListenersModelListenerModel) GetInsertHeaders() map[string]BnsLoadBalancerV1ApiGetListenerModelListenerModelInsertHeadersValue`

GetInsertHeaders returns the InsertHeaders field if non-nil, zero value otherwise.

### GetInsertHeadersOk

`func (o *BnsLoadBalancerV1ApiListListenersModelListenerModel) GetInsertHeadersOk() (*map[string]BnsLoadBalancerV1ApiGetListenerModelListenerModelInsertHeadersValue, bool)`

GetInsertHeadersOk returns a tuple with the InsertHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsertHeaders

`func (o *BnsLoadBalancerV1ApiListListenersModelListenerModel) SetInsertHeaders(v map[string]BnsLoadBalancerV1ApiGetListenerModelListenerModelInsertHeadersValue)`

SetInsertHeaders sets InsertHeaders field to given value.

### HasInsertHeaders

`func (o *BnsLoadBalancerV1ApiListListenersModelListenerModel) HasInsertHeaders() bool`

HasInsertHeaders returns a boolean if a field has been set.

### SetInsertHeadersNil

`func (o *BnsLoadBalancerV1ApiListListenersModelListenerModel) SetInsertHeadersNil(b bool)`

 SetInsertHeadersNil sets the value for InsertHeaders to be an explicit nil

### UnsetInsertHeaders
`func (o *BnsLoadBalancerV1ApiListListenersModelListenerModel) UnsetInsertHeaders()`

UnsetInsertHeaders ensures that no value is present for InsertHeaders, not even an explicit nil
### GetCreatedAt

`func (o *BnsLoadBalancerV1ApiListListenersModelListenerModel) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *BnsLoadBalancerV1ApiListListenersModelListenerModel) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *BnsLoadBalancerV1ApiListListenersModelListenerModel) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *BnsLoadBalancerV1ApiListListenersModelListenerModel) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *BnsLoadBalancerV1ApiListListenersModelListenerModel) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *BnsLoadBalancerV1ApiListListenersModelListenerModel) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetUpdatedAt

`func (o *BnsLoadBalancerV1ApiListListenersModelListenerModel) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *BnsLoadBalancerV1ApiListListenersModelListenerModel) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *BnsLoadBalancerV1ApiListListenersModelListenerModel) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *BnsLoadBalancerV1ApiListListenersModelListenerModel) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *BnsLoadBalancerV1ApiListListenersModelListenerModel) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *BnsLoadBalancerV1ApiListListenersModelListenerModel) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetTimeoutClientData

`func (o *BnsLoadBalancerV1ApiListListenersModelListenerModel) GetTimeoutClientData() int32`

GetTimeoutClientData returns the TimeoutClientData field if non-nil, zero value otherwise.

### GetTimeoutClientDataOk

`func (o *BnsLoadBalancerV1ApiListListenersModelListenerModel) GetTimeoutClientDataOk() (*int32, bool)`

GetTimeoutClientDataOk returns a tuple with the TimeoutClientData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeoutClientData

`func (o *BnsLoadBalancerV1ApiListListenersModelListenerModel) SetTimeoutClientData(v int32)`

SetTimeoutClientData sets TimeoutClientData field to given value.

### HasTimeoutClientData

`func (o *BnsLoadBalancerV1ApiListListenersModelListenerModel) HasTimeoutClientData() bool`

HasTimeoutClientData returns a boolean if a field has been set.

### SetTimeoutClientDataNil

`func (o *BnsLoadBalancerV1ApiListListenersModelListenerModel) SetTimeoutClientDataNil(b bool)`

 SetTimeoutClientDataNil sets the value for TimeoutClientData to be an explicit nil

### UnsetTimeoutClientData
`func (o *BnsLoadBalancerV1ApiListListenersModelListenerModel) UnsetTimeoutClientData()`

UnsetTimeoutClientData ensures that no value is present for TimeoutClientData, not even an explicit nil
### GetDefaultTargetGroupName

`func (o *BnsLoadBalancerV1ApiListListenersModelListenerModel) GetDefaultTargetGroupName() string`

GetDefaultTargetGroupName returns the DefaultTargetGroupName field if non-nil, zero value otherwise.

### GetDefaultTargetGroupNameOk

`func (o *BnsLoadBalancerV1ApiListListenersModelListenerModel) GetDefaultTargetGroupNameOk() (*string, bool)`

GetDefaultTargetGroupNameOk returns a tuple with the DefaultTargetGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultTargetGroupName

`func (o *BnsLoadBalancerV1ApiListListenersModelListenerModel) SetDefaultTargetGroupName(v string)`

SetDefaultTargetGroupName sets DefaultTargetGroupName field to given value.

### HasDefaultTargetGroupName

`func (o *BnsLoadBalancerV1ApiListListenersModelListenerModel) HasDefaultTargetGroupName() bool`

HasDefaultTargetGroupName returns a boolean if a field has been set.

### SetDefaultTargetGroupNameNil

`func (o *BnsLoadBalancerV1ApiListListenersModelListenerModel) SetDefaultTargetGroupNameNil(b bool)`

 SetDefaultTargetGroupNameNil sets the value for DefaultTargetGroupName to be an explicit nil

### UnsetDefaultTargetGroupName
`func (o *BnsLoadBalancerV1ApiListListenersModelListenerModel) UnsetDefaultTargetGroupName()`

UnsetDefaultTargetGroupName ensures that no value is present for DefaultTargetGroupName, not even an explicit nil
### GetDefaultTargetGroupId

`func (o *BnsLoadBalancerV1ApiListListenersModelListenerModel) GetDefaultTargetGroupId() string`

GetDefaultTargetGroupId returns the DefaultTargetGroupId field if non-nil, zero value otherwise.

### GetDefaultTargetGroupIdOk

`func (o *BnsLoadBalancerV1ApiListListenersModelListenerModel) GetDefaultTargetGroupIdOk() (*string, bool)`

GetDefaultTargetGroupIdOk returns a tuple with the DefaultTargetGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultTargetGroupId

`func (o *BnsLoadBalancerV1ApiListListenersModelListenerModel) SetDefaultTargetGroupId(v string)`

SetDefaultTargetGroupId sets DefaultTargetGroupId field to given value.

### HasDefaultTargetGroupId

`func (o *BnsLoadBalancerV1ApiListListenersModelListenerModel) HasDefaultTargetGroupId() bool`

HasDefaultTargetGroupId returns a boolean if a field has been set.

### SetDefaultTargetGroupIdNil

`func (o *BnsLoadBalancerV1ApiListListenersModelListenerModel) SetDefaultTargetGroupIdNil(b bool)`

 SetDefaultTargetGroupIdNil sets the value for DefaultTargetGroupId to be an explicit nil

### UnsetDefaultTargetGroupId
`func (o *BnsLoadBalancerV1ApiListListenersModelListenerModel) UnsetDefaultTargetGroupId()`

UnsetDefaultTargetGroupId ensures that no value is present for DefaultTargetGroupId, not even an explicit nil
### GetLoadBalancerType

`func (o *BnsLoadBalancerV1ApiListListenersModelListenerModel) GetLoadBalancerType() LoadBalancerType`

GetLoadBalancerType returns the LoadBalancerType field if non-nil, zero value otherwise.

### GetLoadBalancerTypeOk

`func (o *BnsLoadBalancerV1ApiListListenersModelListenerModel) GetLoadBalancerTypeOk() (*LoadBalancerType, bool)`

GetLoadBalancerTypeOk returns a tuple with the LoadBalancerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadBalancerType

`func (o *BnsLoadBalancerV1ApiListListenersModelListenerModel) SetLoadBalancerType(v LoadBalancerType)`

SetLoadBalancerType sets LoadBalancerType field to given value.

### HasLoadBalancerType

`func (o *BnsLoadBalancerV1ApiListListenersModelListenerModel) HasLoadBalancerType() bool`

HasLoadBalancerType returns a boolean if a field has been set.

### SetLoadBalancerTypeNil

`func (o *BnsLoadBalancerV1ApiListListenersModelListenerModel) SetLoadBalancerTypeNil(b bool)`

 SetLoadBalancerTypeNil sets the value for LoadBalancerType to be an explicit nil

### UnsetLoadBalancerType
`func (o *BnsLoadBalancerV1ApiListListenersModelListenerModel) UnsetLoadBalancerType()`

UnsetLoadBalancerType ensures that no value is present for LoadBalancerType, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


