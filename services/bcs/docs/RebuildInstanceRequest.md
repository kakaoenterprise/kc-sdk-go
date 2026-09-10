# RebuildInstanceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Rebuild** | [**RebuildInstance**](RebuildInstance.md) | 인스턴스 재구성 요청 정보 | 

## Methods

### NewRebuildInstanceRequest

`func NewRebuildInstanceRequest(rebuild RebuildInstance, ) *RebuildInstanceRequest`

NewRebuildInstanceRequest instantiates a new RebuildInstanceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRebuildInstanceRequestWithDefaults

`func NewRebuildInstanceRequestWithDefaults() *RebuildInstanceRequest`

NewRebuildInstanceRequestWithDefaults instantiates a new RebuildInstanceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRebuild

`func (o *RebuildInstanceRequest) GetRebuild() RebuildInstance`

GetRebuild returns the Rebuild field if non-nil, zero value otherwise.

### GetRebuildOk

`func (o *RebuildInstanceRequest) GetRebuildOk() (*RebuildInstance, bool)`

GetRebuildOk returns a tuple with the Rebuild field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRebuild

`func (o *RebuildInstanceRequest) SetRebuild(v RebuildInstance)`

SetRebuild sets Rebuild field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


