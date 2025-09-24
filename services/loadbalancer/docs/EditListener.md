# EditListener

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultTlsContainerRef** | Pointer to **NullableString** |  | [optional] 
**SniContainerRefs** | Pointer to **[]string** |  | [optional] 
**ConnectionLimit** | Pointer to **NullableInt32** |  | [optional] 
**TargetGroupId** | Pointer to **NullableString** |  | [optional] 
**TlsMinVersion** | Pointer to [**NullableTLSVersion**](TLSVersion.md) |  | [optional] 
**TimeoutClientData** | Pointer to **NullableInt32** |  | [optional] 
**InsertHeaders** | Pointer to [**NullableInsertHeaderModel**](InsertHeaderModel.md) |  | [optional] 

## Methods

### NewEditListener

`func NewEditListener() *EditListener`

NewEditListener instantiates a new EditListener object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEditListenerWithDefaults

`func NewEditListenerWithDefaults() *EditListener`

NewEditListenerWithDefaults instantiates a new EditListener object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultTlsContainerRef

`func (o *EditListener) GetDefaultTlsContainerRef() string`

GetDefaultTlsContainerRef returns the DefaultTlsContainerRef field if non-nil, zero value otherwise.

### GetDefaultTlsContainerRefOk

`func (o *EditListener) GetDefaultTlsContainerRefOk() (*string, bool)`

GetDefaultTlsContainerRefOk returns a tuple with the DefaultTlsContainerRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultTlsContainerRef

`func (o *EditListener) SetDefaultTlsContainerRef(v string)`

SetDefaultTlsContainerRef sets DefaultTlsContainerRef field to given value.

### HasDefaultTlsContainerRef

`func (o *EditListener) HasDefaultTlsContainerRef() bool`

HasDefaultTlsContainerRef returns a boolean if a field has been set.

### SetDefaultTlsContainerRefNil

`func (o *EditListener) SetDefaultTlsContainerRefNil(b bool)`

 SetDefaultTlsContainerRefNil sets the value for DefaultTlsContainerRef to be an explicit nil

### UnsetDefaultTlsContainerRef
`func (o *EditListener) UnsetDefaultTlsContainerRef()`

UnsetDefaultTlsContainerRef ensures that no value is present for DefaultTlsContainerRef, not even an explicit nil
### GetSniContainerRefs

`func (o *EditListener) GetSniContainerRefs() []string`

GetSniContainerRefs returns the SniContainerRefs field if non-nil, zero value otherwise.

### GetSniContainerRefsOk

`func (o *EditListener) GetSniContainerRefsOk() (*[]string, bool)`

GetSniContainerRefsOk returns a tuple with the SniContainerRefs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSniContainerRefs

`func (o *EditListener) SetSniContainerRefs(v []string)`

SetSniContainerRefs sets SniContainerRefs field to given value.

### HasSniContainerRefs

`func (o *EditListener) HasSniContainerRefs() bool`

HasSniContainerRefs returns a boolean if a field has been set.

### SetSniContainerRefsNil

`func (o *EditListener) SetSniContainerRefsNil(b bool)`

 SetSniContainerRefsNil sets the value for SniContainerRefs to be an explicit nil

### UnsetSniContainerRefs
`func (o *EditListener) UnsetSniContainerRefs()`

UnsetSniContainerRefs ensures that no value is present for SniContainerRefs, not even an explicit nil
### GetConnectionLimit

`func (o *EditListener) GetConnectionLimit() int32`

GetConnectionLimit returns the ConnectionLimit field if non-nil, zero value otherwise.

### GetConnectionLimitOk

`func (o *EditListener) GetConnectionLimitOk() (*int32, bool)`

GetConnectionLimitOk returns a tuple with the ConnectionLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionLimit

`func (o *EditListener) SetConnectionLimit(v int32)`

SetConnectionLimit sets ConnectionLimit field to given value.

### HasConnectionLimit

`func (o *EditListener) HasConnectionLimit() bool`

HasConnectionLimit returns a boolean if a field has been set.

### SetConnectionLimitNil

`func (o *EditListener) SetConnectionLimitNil(b bool)`

 SetConnectionLimitNil sets the value for ConnectionLimit to be an explicit nil

### UnsetConnectionLimit
`func (o *EditListener) UnsetConnectionLimit()`

UnsetConnectionLimit ensures that no value is present for ConnectionLimit, not even an explicit nil
### GetTargetGroupId

`func (o *EditListener) GetTargetGroupId() string`

GetTargetGroupId returns the TargetGroupId field if non-nil, zero value otherwise.

### GetTargetGroupIdOk

`func (o *EditListener) GetTargetGroupIdOk() (*string, bool)`

GetTargetGroupIdOk returns a tuple with the TargetGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetGroupId

`func (o *EditListener) SetTargetGroupId(v string)`

SetTargetGroupId sets TargetGroupId field to given value.

### HasTargetGroupId

`func (o *EditListener) HasTargetGroupId() bool`

HasTargetGroupId returns a boolean if a field has been set.

### SetTargetGroupIdNil

`func (o *EditListener) SetTargetGroupIdNil(b bool)`

 SetTargetGroupIdNil sets the value for TargetGroupId to be an explicit nil

### UnsetTargetGroupId
`func (o *EditListener) UnsetTargetGroupId()`

UnsetTargetGroupId ensures that no value is present for TargetGroupId, not even an explicit nil
### GetTlsMinVersion

`func (o *EditListener) GetTlsMinVersion() TLSVersion`

GetTlsMinVersion returns the TlsMinVersion field if non-nil, zero value otherwise.

### GetTlsMinVersionOk

`func (o *EditListener) GetTlsMinVersionOk() (*TLSVersion, bool)`

GetTlsMinVersionOk returns a tuple with the TlsMinVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsMinVersion

`func (o *EditListener) SetTlsMinVersion(v TLSVersion)`

SetTlsMinVersion sets TlsMinVersion field to given value.

### HasTlsMinVersion

`func (o *EditListener) HasTlsMinVersion() bool`

HasTlsMinVersion returns a boolean if a field has been set.

### SetTlsMinVersionNil

`func (o *EditListener) SetTlsMinVersionNil(b bool)`

 SetTlsMinVersionNil sets the value for TlsMinVersion to be an explicit nil

### UnsetTlsMinVersion
`func (o *EditListener) UnsetTlsMinVersion()`

UnsetTlsMinVersion ensures that no value is present for TlsMinVersion, not even an explicit nil
### GetTimeoutClientData

`func (o *EditListener) GetTimeoutClientData() int32`

GetTimeoutClientData returns the TimeoutClientData field if non-nil, zero value otherwise.

### GetTimeoutClientDataOk

`func (o *EditListener) GetTimeoutClientDataOk() (*int32, bool)`

GetTimeoutClientDataOk returns a tuple with the TimeoutClientData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeoutClientData

`func (o *EditListener) SetTimeoutClientData(v int32)`

SetTimeoutClientData sets TimeoutClientData field to given value.

### HasTimeoutClientData

`func (o *EditListener) HasTimeoutClientData() bool`

HasTimeoutClientData returns a boolean if a field has been set.

### SetTimeoutClientDataNil

`func (o *EditListener) SetTimeoutClientDataNil(b bool)`

 SetTimeoutClientDataNil sets the value for TimeoutClientData to be an explicit nil

### UnsetTimeoutClientData
`func (o *EditListener) UnsetTimeoutClientData()`

UnsetTimeoutClientData ensures that no value is present for TimeoutClientData, not even an explicit nil
### GetInsertHeaders

`func (o *EditListener) GetInsertHeaders() InsertHeaderModel`

GetInsertHeaders returns the InsertHeaders field if non-nil, zero value otherwise.

### GetInsertHeadersOk

`func (o *EditListener) GetInsertHeadersOk() (*InsertHeaderModel, bool)`

GetInsertHeadersOk returns a tuple with the InsertHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsertHeaders

`func (o *EditListener) SetInsertHeaders(v InsertHeaderModel)`

SetInsertHeaders sets InsertHeaders field to given value.

### HasInsertHeaders

`func (o *EditListener) HasInsertHeaders() bool`

HasInsertHeaders returns a boolean if a field has been set.

### SetInsertHeadersNil

`func (o *EditListener) SetInsertHeadersNil(b bool)`

 SetInsertHeadersNil sets the value for InsertHeaders to be an explicit nil

### UnsetInsertHeaders
`func (o *EditListener) UnsetInsertHeaders()`

UnsetInsertHeaders ensures that no value is present for InsertHeaders, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


