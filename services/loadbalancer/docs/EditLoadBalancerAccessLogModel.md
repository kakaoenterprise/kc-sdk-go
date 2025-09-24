# EditLoadBalancerAccessLogModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bucket** | **string** | 로그를 저장하는 대상 버킷 이름 | 
**AccessKey** | **string** | 버킷 접근에 사용하는 액세스 키 | 
**SecretKey** | **string** | 버킷 접근에 사용하는 보안 액세스 키 | 

## Methods

### NewEditLoadBalancerAccessLogModel

`func NewEditLoadBalancerAccessLogModel(bucket string, accessKey string, secretKey string, ) *EditLoadBalancerAccessLogModel`

NewEditLoadBalancerAccessLogModel instantiates a new EditLoadBalancerAccessLogModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEditLoadBalancerAccessLogModelWithDefaults

`func NewEditLoadBalancerAccessLogModelWithDefaults() *EditLoadBalancerAccessLogModel`

NewEditLoadBalancerAccessLogModelWithDefaults instantiates a new EditLoadBalancerAccessLogModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBucket

`func (o *EditLoadBalancerAccessLogModel) GetBucket() string`

GetBucket returns the Bucket field if non-nil, zero value otherwise.

### GetBucketOk

`func (o *EditLoadBalancerAccessLogModel) GetBucketOk() (*string, bool)`

GetBucketOk returns a tuple with the Bucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucket

`func (o *EditLoadBalancerAccessLogModel) SetBucket(v string)`

SetBucket sets Bucket field to given value.


### GetAccessKey

`func (o *EditLoadBalancerAccessLogModel) GetAccessKey() string`

GetAccessKey returns the AccessKey field if non-nil, zero value otherwise.

### GetAccessKeyOk

`func (o *EditLoadBalancerAccessLogModel) GetAccessKeyOk() (*string, bool)`

GetAccessKeyOk returns a tuple with the AccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessKey

`func (o *EditLoadBalancerAccessLogModel) SetAccessKey(v string)`

SetAccessKey sets AccessKey field to given value.


### GetSecretKey

`func (o *EditLoadBalancerAccessLogModel) GetSecretKey() string`

GetSecretKey returns the SecretKey field if non-nil, zero value otherwise.

### GetSecretKeyOk

`func (o *EditLoadBalancerAccessLogModel) GetSecretKeyOk() (*string, bool)`

GetSecretKeyOk returns a tuple with the SecretKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretKey

`func (o *EditLoadBalancerAccessLogModel) SetSecretKey(v string)`

SetSecretKey sets SecretKey field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


