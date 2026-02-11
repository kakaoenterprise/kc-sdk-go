# AccessLogsModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bucket** | **string** | 로그를 저장하는 대상 버킷 이름 | 

## Methods

### NewAccessLogsModel

`func NewAccessLogsModel(bucket string, ) *AccessLogsModel`

NewAccessLogsModel instantiates a new AccessLogsModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccessLogsModelWithDefaults

`func NewAccessLogsModelWithDefaults() *AccessLogsModel`

NewAccessLogsModelWithDefaults instantiates a new AccessLogsModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBucket

`func (o *AccessLogsModel) GetBucket() string`

GetBucket returns the Bucket field if non-nil, zero value otherwise.

### GetBucketOk

`func (o *AccessLogsModel) GetBucketOk() (*string, bool)`

GetBucketOk returns a tuple with the Bucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucket

`func (o *AccessLogsModel) SetBucket(v string)`

SetBucket sets Bucket field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


