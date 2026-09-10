# UpdateListener

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultTlsContainerRef** | Pointer to **NullableString** |  | [optional] 
**SniContainerRefs** | Pointer to **[]string** |  | [optional] 
**TargetGroupId** | Pointer to **NullableString** |  | [optional] 
**TlsMinVersion** | Pointer to [**NullableTLSVersion**](TLSVersion.md) |  | [optional] 
**TimeoutClientData** | Pointer to **NullableInt32** |  | [optional] 
**InsertHeaders** | Pointer to [**NullableInsertHeaderRequest**](InsertHeaderRequest.md) |  | [optional] 

## Methods

### NewUpdateListener

`func NewUpdateListener() *UpdateListener`

NewUpdateListener instantiates a new UpdateListener object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateListenerWithDefaults

`func NewUpdateListenerWithDefaults() *UpdateListener`

NewUpdateListenerWithDefaults instantiates a new UpdateListener object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultTlsContainerRef

`func (o *UpdateListener) GetDefaultTlsContainerRef() string`

GetDefaultTlsContainerRef returns the DefaultTlsContainerRef field if non-nil, zero value otherwise.

### GetDefaultTlsContainerRefOk

`func (o *UpdateListener) GetDefaultTlsContainerRefOk() (*string, bool)`

GetDefaultTlsContainerRefOk returns a tuple with the DefaultTlsContainerRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultTlsContainerRef

`func (o *UpdateListener) SetDefaultTlsContainerRef(v string)`

SetDefaultTlsContainerRef sets DefaultTlsContainerRef field to given value.

### HasDefaultTlsContainerRef

`func (o *UpdateListener) HasDefaultTlsContainerRef() bool`

HasDefaultTlsContainerRef returns a boolean if a field has been set.

### SetDefaultTlsContainerRefNil

`func (o *UpdateListener) SetDefaultTlsContainerRefNil(b bool)`

 SetDefaultTlsContainerRefNil sets the value for DefaultTlsContainerRef to be an explicit nil

### UnsetDefaultTlsContainerRef
`func (o *UpdateListener) UnsetDefaultTlsContainerRef()`

UnsetDefaultTlsContainerRef ensures that no value is present for DefaultTlsContainerRef, not even an explicit nil
### GetSniContainerRefs

`func (o *UpdateListener) GetSniContainerRefs() []string`

GetSniContainerRefs returns the SniContainerRefs field if non-nil, zero value otherwise.

### GetSniContainerRefsOk

`func (o *UpdateListener) GetSniContainerRefsOk() (*[]string, bool)`

GetSniContainerRefsOk returns a tuple with the SniContainerRefs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSniContainerRefs

`func (o *UpdateListener) SetSniContainerRefs(v []string)`

SetSniContainerRefs sets SniContainerRefs field to given value.

### HasSniContainerRefs

`func (o *UpdateListener) HasSniContainerRefs() bool`

HasSniContainerRefs returns a boolean if a field has been set.

### SetSniContainerRefsNil

`func (o *UpdateListener) SetSniContainerRefsNil(b bool)`

 SetSniContainerRefsNil sets the value for SniContainerRefs to be an explicit nil

### UnsetSniContainerRefs
`func (o *UpdateListener) UnsetSniContainerRefs()`

UnsetSniContainerRefs ensures that no value is present for SniContainerRefs, not even an explicit nil
### GetTargetGroupId

`func (o *UpdateListener) GetTargetGroupId() string`

GetTargetGroupId returns the TargetGroupId field if non-nil, zero value otherwise.

### GetTargetGroupIdOk

`func (o *UpdateListener) GetTargetGroupIdOk() (*string, bool)`

GetTargetGroupIdOk returns a tuple with the TargetGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetGroupId

`func (o *UpdateListener) SetTargetGroupId(v string)`

SetTargetGroupId sets TargetGroupId field to given value.

### HasTargetGroupId

`func (o *UpdateListener) HasTargetGroupId() bool`

HasTargetGroupId returns a boolean if a field has been set.

### SetTargetGroupIdNil

`func (o *UpdateListener) SetTargetGroupIdNil(b bool)`

 SetTargetGroupIdNil sets the value for TargetGroupId to be an explicit nil

### UnsetTargetGroupId
`func (o *UpdateListener) UnsetTargetGroupId()`

UnsetTargetGroupId ensures that no value is present for TargetGroupId, not even an explicit nil
### GetTlsMinVersion

`func (o *UpdateListener) GetTlsMinVersion() TLSVersion`

GetTlsMinVersion returns the TlsMinVersion field if non-nil, zero value otherwise.

### GetTlsMinVersionOk

`func (o *UpdateListener) GetTlsMinVersionOk() (*TLSVersion, bool)`

GetTlsMinVersionOk returns a tuple with the TlsMinVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsMinVersion

`func (o *UpdateListener) SetTlsMinVersion(v TLSVersion)`

SetTlsMinVersion sets TlsMinVersion field to given value.

### HasTlsMinVersion

`func (o *UpdateListener) HasTlsMinVersion() bool`

HasTlsMinVersion returns a boolean if a field has been set.

### SetTlsMinVersionNil

`func (o *UpdateListener) SetTlsMinVersionNil(b bool)`

 SetTlsMinVersionNil sets the value for TlsMinVersion to be an explicit nil

### UnsetTlsMinVersion
`func (o *UpdateListener) UnsetTlsMinVersion()`

UnsetTlsMinVersion ensures that no value is present for TlsMinVersion, not even an explicit nil
### GetTimeoutClientData

`func (o *UpdateListener) GetTimeoutClientData() int32`

GetTimeoutClientData returns the TimeoutClientData field if non-nil, zero value otherwise.

### GetTimeoutClientDataOk

`func (o *UpdateListener) GetTimeoutClientDataOk() (*int32, bool)`

GetTimeoutClientDataOk returns a tuple with the TimeoutClientData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeoutClientData

`func (o *UpdateListener) SetTimeoutClientData(v int32)`

SetTimeoutClientData sets TimeoutClientData field to given value.

### HasTimeoutClientData

`func (o *UpdateListener) HasTimeoutClientData() bool`

HasTimeoutClientData returns a boolean if a field has been set.

### SetTimeoutClientDataNil

`func (o *UpdateListener) SetTimeoutClientDataNil(b bool)`

 SetTimeoutClientDataNil sets the value for TimeoutClientData to be an explicit nil

### UnsetTimeoutClientData
`func (o *UpdateListener) UnsetTimeoutClientData()`

UnsetTimeoutClientData ensures that no value is present for TimeoutClientData, not even an explicit nil
### GetInsertHeaders

`func (o *UpdateListener) GetInsertHeaders() InsertHeaderRequest`

GetInsertHeaders returns the InsertHeaders field if non-nil, zero value otherwise.

### GetInsertHeadersOk

`func (o *UpdateListener) GetInsertHeadersOk() (*InsertHeaderRequest, bool)`

GetInsertHeadersOk returns a tuple with the InsertHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsertHeaders

`func (o *UpdateListener) SetInsertHeaders(v InsertHeaderRequest)`

SetInsertHeaders sets InsertHeaders field to given value.

### HasInsertHeaders

`func (o *UpdateListener) HasInsertHeaders() bool`

HasInsertHeaders returns a boolean if a field has been set.

### SetInsertHeadersNil

`func (o *UpdateListener) SetInsertHeadersNil(b bool)`

 SetInsertHeadersNil sets the value for InsertHeaders to be an explicit nil

### UnsetInsertHeaders
`func (o *UpdateListener) UnsetInsertHeaders()`

UnsetInsertHeaders ensures that no value is present for InsertHeaders, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


