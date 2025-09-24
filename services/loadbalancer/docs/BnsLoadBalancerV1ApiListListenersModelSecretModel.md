# BnsLoadBalancerV1ApiListListenersModelSecretModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | 인증서 ID | 
**Name** | Pointer to **NullableString** |  | [optional] 
**Expiration** | Pointer to **NullableString** |  | [optional] 
**Status** | Pointer to **NullableString** |  | [optional] 
**SecretType** | Pointer to **NullableString** |  | [optional] 
**IsDefault** | Pointer to **NullableBool** |  | [optional] 
**CreatorId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewBnsLoadBalancerV1ApiListListenersModelSecretModel

`func NewBnsLoadBalancerV1ApiListListenersModelSecretModel(id string, ) *BnsLoadBalancerV1ApiListListenersModelSecretModel`

NewBnsLoadBalancerV1ApiListListenersModelSecretModel instantiates a new BnsLoadBalancerV1ApiListListenersModelSecretModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBnsLoadBalancerV1ApiListListenersModelSecretModelWithDefaults

`func NewBnsLoadBalancerV1ApiListListenersModelSecretModelWithDefaults() *BnsLoadBalancerV1ApiListListenersModelSecretModel`

NewBnsLoadBalancerV1ApiListListenersModelSecretModelWithDefaults instantiates a new BnsLoadBalancerV1ApiListListenersModelSecretModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BnsLoadBalancerV1ApiListListenersModelSecretModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BnsLoadBalancerV1ApiListListenersModelSecretModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BnsLoadBalancerV1ApiListListenersModelSecretModel) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *BnsLoadBalancerV1ApiListListenersModelSecretModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BnsLoadBalancerV1ApiListListenersModelSecretModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BnsLoadBalancerV1ApiListListenersModelSecretModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BnsLoadBalancerV1ApiListListenersModelSecretModel) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *BnsLoadBalancerV1ApiListListenersModelSecretModel) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *BnsLoadBalancerV1ApiListListenersModelSecretModel) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetExpiration

`func (o *BnsLoadBalancerV1ApiListListenersModelSecretModel) GetExpiration() string`

GetExpiration returns the Expiration field if non-nil, zero value otherwise.

### GetExpirationOk

`func (o *BnsLoadBalancerV1ApiListListenersModelSecretModel) GetExpirationOk() (*string, bool)`

GetExpirationOk returns a tuple with the Expiration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiration

`func (o *BnsLoadBalancerV1ApiListListenersModelSecretModel) SetExpiration(v string)`

SetExpiration sets Expiration field to given value.

### HasExpiration

`func (o *BnsLoadBalancerV1ApiListListenersModelSecretModel) HasExpiration() bool`

HasExpiration returns a boolean if a field has been set.

### SetExpirationNil

`func (o *BnsLoadBalancerV1ApiListListenersModelSecretModel) SetExpirationNil(b bool)`

 SetExpirationNil sets the value for Expiration to be an explicit nil

### UnsetExpiration
`func (o *BnsLoadBalancerV1ApiListListenersModelSecretModel) UnsetExpiration()`

UnsetExpiration ensures that no value is present for Expiration, not even an explicit nil
### GetStatus

`func (o *BnsLoadBalancerV1ApiListListenersModelSecretModel) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BnsLoadBalancerV1ApiListListenersModelSecretModel) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BnsLoadBalancerV1ApiListListenersModelSecretModel) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *BnsLoadBalancerV1ApiListListenersModelSecretModel) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *BnsLoadBalancerV1ApiListListenersModelSecretModel) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *BnsLoadBalancerV1ApiListListenersModelSecretModel) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetSecretType

`func (o *BnsLoadBalancerV1ApiListListenersModelSecretModel) GetSecretType() string`

GetSecretType returns the SecretType field if non-nil, zero value otherwise.

### GetSecretTypeOk

`func (o *BnsLoadBalancerV1ApiListListenersModelSecretModel) GetSecretTypeOk() (*string, bool)`

GetSecretTypeOk returns a tuple with the SecretType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretType

`func (o *BnsLoadBalancerV1ApiListListenersModelSecretModel) SetSecretType(v string)`

SetSecretType sets SecretType field to given value.

### HasSecretType

`func (o *BnsLoadBalancerV1ApiListListenersModelSecretModel) HasSecretType() bool`

HasSecretType returns a boolean if a field has been set.

### SetSecretTypeNil

`func (o *BnsLoadBalancerV1ApiListListenersModelSecretModel) SetSecretTypeNil(b bool)`

 SetSecretTypeNil sets the value for SecretType to be an explicit nil

### UnsetSecretType
`func (o *BnsLoadBalancerV1ApiListListenersModelSecretModel) UnsetSecretType()`

UnsetSecretType ensures that no value is present for SecretType, not even an explicit nil
### GetIsDefault

`func (o *BnsLoadBalancerV1ApiListListenersModelSecretModel) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *BnsLoadBalancerV1ApiListListenersModelSecretModel) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *BnsLoadBalancerV1ApiListListenersModelSecretModel) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.

### HasIsDefault

`func (o *BnsLoadBalancerV1ApiListListenersModelSecretModel) HasIsDefault() bool`

HasIsDefault returns a boolean if a field has been set.

### SetIsDefaultNil

`func (o *BnsLoadBalancerV1ApiListListenersModelSecretModel) SetIsDefaultNil(b bool)`

 SetIsDefaultNil sets the value for IsDefault to be an explicit nil

### UnsetIsDefault
`func (o *BnsLoadBalancerV1ApiListListenersModelSecretModel) UnsetIsDefault()`

UnsetIsDefault ensures that no value is present for IsDefault, not even an explicit nil
### GetCreatorId

`func (o *BnsLoadBalancerV1ApiListListenersModelSecretModel) GetCreatorId() string`

GetCreatorId returns the CreatorId field if non-nil, zero value otherwise.

### GetCreatorIdOk

`func (o *BnsLoadBalancerV1ApiListListenersModelSecretModel) GetCreatorIdOk() (*string, bool)`

GetCreatorIdOk returns a tuple with the CreatorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorId

`func (o *BnsLoadBalancerV1ApiListListenersModelSecretModel) SetCreatorId(v string)`

SetCreatorId sets CreatorId field to given value.

### HasCreatorId

`func (o *BnsLoadBalancerV1ApiListListenersModelSecretModel) HasCreatorId() bool`

HasCreatorId returns a boolean if a field has been set.

### SetCreatorIdNil

`func (o *BnsLoadBalancerV1ApiListListenersModelSecretModel) SetCreatorIdNil(b bool)`

 SetCreatorIdNil sets the value for CreatorId to be an explicit nil

### UnsetCreatorId
`func (o *BnsLoadBalancerV1ApiListListenersModelSecretModel) UnsetCreatorId()`

UnsetCreatorId ensures that no value is present for CreatorId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


