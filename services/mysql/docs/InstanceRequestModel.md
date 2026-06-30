# InstanceRequestModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bucket** | **string** | 로그를 저장할 Object Storage 버킷 이름 | 
**LogInfos** | [**[]LogInfoRequestModel**](LogInfoRequestModel.md) | 내보낼 로그 범위 목록 | 
**Path** | **string** | 로그가 저장될 Object Storage 내 경로 | 
**UserCredentialId** | **string** | Object Storage 접근을 위한 사용자 인증 ID | 
**UserCredentialSecret** | **string** | 해당 ID에 대한 Secret 키 | 

## Methods

### NewInstanceRequestModel

`func NewInstanceRequestModel(bucket string, logInfos []LogInfoRequestModel, path string, userCredentialId string, userCredentialSecret string, ) *InstanceRequestModel`

NewInstanceRequestModel instantiates a new InstanceRequestModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstanceRequestModelWithDefaults

`func NewInstanceRequestModelWithDefaults() *InstanceRequestModel`

NewInstanceRequestModelWithDefaults instantiates a new InstanceRequestModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBucket

`func (o *InstanceRequestModel) GetBucket() string`

GetBucket returns the Bucket field if non-nil, zero value otherwise.

### GetBucketOk

`func (o *InstanceRequestModel) GetBucketOk() (*string, bool)`

GetBucketOk returns a tuple with the Bucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucket

`func (o *InstanceRequestModel) SetBucket(v string)`

SetBucket sets Bucket field to given value.


### GetLogInfos

`func (o *InstanceRequestModel) GetLogInfos() []LogInfoRequestModel`

GetLogInfos returns the LogInfos field if non-nil, zero value otherwise.

### GetLogInfosOk

`func (o *InstanceRequestModel) GetLogInfosOk() (*[]LogInfoRequestModel, bool)`

GetLogInfosOk returns a tuple with the LogInfos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogInfos

`func (o *InstanceRequestModel) SetLogInfos(v []LogInfoRequestModel)`

SetLogInfos sets LogInfos field to given value.


### GetPath

`func (o *InstanceRequestModel) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *InstanceRequestModel) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *InstanceRequestModel) SetPath(v string)`

SetPath sets Path field to given value.


### GetUserCredentialId

`func (o *InstanceRequestModel) GetUserCredentialId() string`

GetUserCredentialId returns the UserCredentialId field if non-nil, zero value otherwise.

### GetUserCredentialIdOk

`func (o *InstanceRequestModel) GetUserCredentialIdOk() (*string, bool)`

GetUserCredentialIdOk returns a tuple with the UserCredentialId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserCredentialId

`func (o *InstanceRequestModel) SetUserCredentialId(v string)`

SetUserCredentialId sets UserCredentialId field to given value.


### GetUserCredentialSecret

`func (o *InstanceRequestModel) GetUserCredentialSecret() string`

GetUserCredentialSecret returns the UserCredentialSecret field if non-nil, zero value otherwise.

### GetUserCredentialSecretOk

`func (o *InstanceRequestModel) GetUserCredentialSecretOk() (*string, bool)`

GetUserCredentialSecretOk returns a tuple with the UserCredentialSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserCredentialSecret

`func (o *InstanceRequestModel) SetUserCredentialSecret(v string)`

SetUserCredentialSecret sets UserCredentialSecret field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


