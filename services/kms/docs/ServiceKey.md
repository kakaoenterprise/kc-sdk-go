# ServiceKey

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** |  | [optional] 
**DefaultVersion** | Pointer to **NullableInt32** |  | [optional] 
**Service** | Pointer to **NullableString** |  | [optional] 
**Status** | Pointer to **NullableString** |  | [optional] 
**RotationPeriod** | Pointer to **NullableInt32** |  | [optional] 
**LastRotationAt** | Pointer to **NullableString** |  | [optional] 
**NextRotationAt** | Pointer to **NullableString** |  | [optional] 
**CreatedAt** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewServiceKey

`func NewServiceKey() *ServiceKey`

NewServiceKey instantiates a new ServiceKey object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceKeyWithDefaults

`func NewServiceKeyWithDefaults() *ServiceKey`

NewServiceKeyWithDefaults instantiates a new ServiceKey object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ServiceKey) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ServiceKey) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ServiceKey) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ServiceKey) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *ServiceKey) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ServiceKey) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDefaultVersion

`func (o *ServiceKey) GetDefaultVersion() int32`

GetDefaultVersion returns the DefaultVersion field if non-nil, zero value otherwise.

### GetDefaultVersionOk

`func (o *ServiceKey) GetDefaultVersionOk() (*int32, bool)`

GetDefaultVersionOk returns a tuple with the DefaultVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultVersion

`func (o *ServiceKey) SetDefaultVersion(v int32)`

SetDefaultVersion sets DefaultVersion field to given value.

### HasDefaultVersion

`func (o *ServiceKey) HasDefaultVersion() bool`

HasDefaultVersion returns a boolean if a field has been set.

### SetDefaultVersionNil

`func (o *ServiceKey) SetDefaultVersionNil(b bool)`

 SetDefaultVersionNil sets the value for DefaultVersion to be an explicit nil

### UnsetDefaultVersion
`func (o *ServiceKey) UnsetDefaultVersion()`

UnsetDefaultVersion ensures that no value is present for DefaultVersion, not even an explicit nil
### GetService

`func (o *ServiceKey) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *ServiceKey) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *ServiceKey) SetService(v string)`

SetService sets Service field to given value.

### HasService

`func (o *ServiceKey) HasService() bool`

HasService returns a boolean if a field has been set.

### SetServiceNil

`func (o *ServiceKey) SetServiceNil(b bool)`

 SetServiceNil sets the value for Service to be an explicit nil

### UnsetService
`func (o *ServiceKey) UnsetService()`

UnsetService ensures that no value is present for Service, not even an explicit nil
### GetStatus

`func (o *ServiceKey) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ServiceKey) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ServiceKey) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ServiceKey) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *ServiceKey) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *ServiceKey) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetRotationPeriod

`func (o *ServiceKey) GetRotationPeriod() int32`

GetRotationPeriod returns the RotationPeriod field if non-nil, zero value otherwise.

### GetRotationPeriodOk

`func (o *ServiceKey) GetRotationPeriodOk() (*int32, bool)`

GetRotationPeriodOk returns a tuple with the RotationPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRotationPeriod

`func (o *ServiceKey) SetRotationPeriod(v int32)`

SetRotationPeriod sets RotationPeriod field to given value.

### HasRotationPeriod

`func (o *ServiceKey) HasRotationPeriod() bool`

HasRotationPeriod returns a boolean if a field has been set.

### SetRotationPeriodNil

`func (o *ServiceKey) SetRotationPeriodNil(b bool)`

 SetRotationPeriodNil sets the value for RotationPeriod to be an explicit nil

### UnsetRotationPeriod
`func (o *ServiceKey) UnsetRotationPeriod()`

UnsetRotationPeriod ensures that no value is present for RotationPeriod, not even an explicit nil
### GetLastRotationAt

`func (o *ServiceKey) GetLastRotationAt() string`

GetLastRotationAt returns the LastRotationAt field if non-nil, zero value otherwise.

### GetLastRotationAtOk

`func (o *ServiceKey) GetLastRotationAtOk() (*string, bool)`

GetLastRotationAtOk returns a tuple with the LastRotationAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastRotationAt

`func (o *ServiceKey) SetLastRotationAt(v string)`

SetLastRotationAt sets LastRotationAt field to given value.

### HasLastRotationAt

`func (o *ServiceKey) HasLastRotationAt() bool`

HasLastRotationAt returns a boolean if a field has been set.

### SetLastRotationAtNil

`func (o *ServiceKey) SetLastRotationAtNil(b bool)`

 SetLastRotationAtNil sets the value for LastRotationAt to be an explicit nil

### UnsetLastRotationAt
`func (o *ServiceKey) UnsetLastRotationAt()`

UnsetLastRotationAt ensures that no value is present for LastRotationAt, not even an explicit nil
### GetNextRotationAt

`func (o *ServiceKey) GetNextRotationAt() string`

GetNextRotationAt returns the NextRotationAt field if non-nil, zero value otherwise.

### GetNextRotationAtOk

`func (o *ServiceKey) GetNextRotationAtOk() (*string, bool)`

GetNextRotationAtOk returns a tuple with the NextRotationAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextRotationAt

`func (o *ServiceKey) SetNextRotationAt(v string)`

SetNextRotationAt sets NextRotationAt field to given value.

### HasNextRotationAt

`func (o *ServiceKey) HasNextRotationAt() bool`

HasNextRotationAt returns a boolean if a field has been set.

### SetNextRotationAtNil

`func (o *ServiceKey) SetNextRotationAtNil(b bool)`

 SetNextRotationAtNil sets the value for NextRotationAt to be an explicit nil

### UnsetNextRotationAt
`func (o *ServiceKey) UnsetNextRotationAt()`

UnsetNextRotationAt ensures that no value is present for NextRotationAt, not even an explicit nil
### GetCreatedAt

`func (o *ServiceKey) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ServiceKey) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ServiceKey) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ServiceKey) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *ServiceKey) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *ServiceKey) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


