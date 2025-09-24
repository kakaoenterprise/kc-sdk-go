# AccessLogsModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bucket** | **string** | 로그를 저장하는 대상 버킷 이름 | 
**AccessKey** | **string** | 버킷 접근에 사용하는 액세스 키 | 
**SecretKey** | **string** | 버킷 접근에 사용하는 인증서 키 | 

## Methods

### NewAccessLogsModel

`func NewAccessLogsModel(bucket string, accessKey string, secretKey string, ) *AccessLogsModel`

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


### GetAccessKey

`func (o *AccessLogsModel) GetAccessKey() string`

GetAccessKey returns the AccessKey field if non-nil, zero value otherwise.

### GetAccessKeyOk

`func (o *AccessLogsModel) GetAccessKeyOk() (*string, bool)`

GetAccessKeyOk returns a tuple with the AccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessKey

`func (o *AccessLogsModel) SetAccessKey(v string)`

SetAccessKey sets AccessKey field to given value.


### GetSecretKey

`func (o *AccessLogsModel) GetSecretKey() string`

GetSecretKey returns the SecretKey field if non-nil, zero value otherwise.

### GetSecretKeyOk

`func (o *AccessLogsModel) GetSecretKeyOk() (*string, bool)`

GetSecretKeyOk returns a tuple with the SecretKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretKey

`func (o *AccessLogsModel) SetSecretKey(v string)`

SetSecretKey sets SecretKey field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


