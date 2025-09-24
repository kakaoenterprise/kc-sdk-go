# BnsLoadBalancerV1ApiListTlsCertificatesModelSecretModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | **time.Time** | 리소스가 생성된 시간 &lt;br/&gt; - ISO_8601 형식  &lt;br/&gt; - UTC 기준 | 
**UpdatedAt** | **time.Time** | 리소스가 마지막으로 수정된 시간 &lt;br/&gt; - ISO_8601 형식  &lt;br/&gt; - UTC 기준 | 
**Status** | **string** | 인증서 상태 | 
**Name** | **string** | 인증서 이름 | 
**SecretType** | **string** | 인증서 유형 | 
**Expiration** | **string** | 인증서 만료 일시 | 
**CreatorId** | **string** | 인증서 생성자 ID | 
**ContentTypes** | [**ContentType**](ContentType.md) | 인증서 콘텐츠 유형 목록 | 
**SecretRef** | **string** | 인증서 참조 ID | 

## Methods

### NewBnsLoadBalancerV1ApiListTlsCertificatesModelSecretModel

`func NewBnsLoadBalancerV1ApiListTlsCertificatesModelSecretModel(createdAt time.Time, updatedAt time.Time, status string, name string, secretType string, expiration string, creatorId string, contentTypes ContentType, secretRef string, ) *BnsLoadBalancerV1ApiListTlsCertificatesModelSecretModel`

NewBnsLoadBalancerV1ApiListTlsCertificatesModelSecretModel instantiates a new BnsLoadBalancerV1ApiListTlsCertificatesModelSecretModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBnsLoadBalancerV1ApiListTlsCertificatesModelSecretModelWithDefaults

`func NewBnsLoadBalancerV1ApiListTlsCertificatesModelSecretModelWithDefaults() *BnsLoadBalancerV1ApiListTlsCertificatesModelSecretModel`

NewBnsLoadBalancerV1ApiListTlsCertificatesModelSecretModelWithDefaults instantiates a new BnsLoadBalancerV1ApiListTlsCertificatesModelSecretModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *BnsLoadBalancerV1ApiListTlsCertificatesModelSecretModel) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *BnsLoadBalancerV1ApiListTlsCertificatesModelSecretModel) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *BnsLoadBalancerV1ApiListTlsCertificatesModelSecretModel) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *BnsLoadBalancerV1ApiListTlsCertificatesModelSecretModel) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *BnsLoadBalancerV1ApiListTlsCertificatesModelSecretModel) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *BnsLoadBalancerV1ApiListTlsCertificatesModelSecretModel) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetStatus

`func (o *BnsLoadBalancerV1ApiListTlsCertificatesModelSecretModel) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BnsLoadBalancerV1ApiListTlsCertificatesModelSecretModel) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BnsLoadBalancerV1ApiListTlsCertificatesModelSecretModel) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetName

`func (o *BnsLoadBalancerV1ApiListTlsCertificatesModelSecretModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BnsLoadBalancerV1ApiListTlsCertificatesModelSecretModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BnsLoadBalancerV1ApiListTlsCertificatesModelSecretModel) SetName(v string)`

SetName sets Name field to given value.


### GetSecretType

`func (o *BnsLoadBalancerV1ApiListTlsCertificatesModelSecretModel) GetSecretType() string`

GetSecretType returns the SecretType field if non-nil, zero value otherwise.

### GetSecretTypeOk

`func (o *BnsLoadBalancerV1ApiListTlsCertificatesModelSecretModel) GetSecretTypeOk() (*string, bool)`

GetSecretTypeOk returns a tuple with the SecretType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretType

`func (o *BnsLoadBalancerV1ApiListTlsCertificatesModelSecretModel) SetSecretType(v string)`

SetSecretType sets SecretType field to given value.


### GetExpiration

`func (o *BnsLoadBalancerV1ApiListTlsCertificatesModelSecretModel) GetExpiration() string`

GetExpiration returns the Expiration field if non-nil, zero value otherwise.

### GetExpirationOk

`func (o *BnsLoadBalancerV1ApiListTlsCertificatesModelSecretModel) GetExpirationOk() (*string, bool)`

GetExpirationOk returns a tuple with the Expiration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiration

`func (o *BnsLoadBalancerV1ApiListTlsCertificatesModelSecretModel) SetExpiration(v string)`

SetExpiration sets Expiration field to given value.


### GetCreatorId

`func (o *BnsLoadBalancerV1ApiListTlsCertificatesModelSecretModel) GetCreatorId() string`

GetCreatorId returns the CreatorId field if non-nil, zero value otherwise.

### GetCreatorIdOk

`func (o *BnsLoadBalancerV1ApiListTlsCertificatesModelSecretModel) GetCreatorIdOk() (*string, bool)`

GetCreatorIdOk returns a tuple with the CreatorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorId

`func (o *BnsLoadBalancerV1ApiListTlsCertificatesModelSecretModel) SetCreatorId(v string)`

SetCreatorId sets CreatorId field to given value.


### GetContentTypes

`func (o *BnsLoadBalancerV1ApiListTlsCertificatesModelSecretModel) GetContentTypes() ContentType`

GetContentTypes returns the ContentTypes field if non-nil, zero value otherwise.

### GetContentTypesOk

`func (o *BnsLoadBalancerV1ApiListTlsCertificatesModelSecretModel) GetContentTypesOk() (*ContentType, bool)`

GetContentTypesOk returns a tuple with the ContentTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentTypes

`func (o *BnsLoadBalancerV1ApiListTlsCertificatesModelSecretModel) SetContentTypes(v ContentType)`

SetContentTypes sets ContentTypes field to given value.


### GetSecretRef

`func (o *BnsLoadBalancerV1ApiListTlsCertificatesModelSecretModel) GetSecretRef() string`

GetSecretRef returns the SecretRef field if non-nil, zero value otherwise.

### GetSecretRefOk

`func (o *BnsLoadBalancerV1ApiListTlsCertificatesModelSecretModel) GetSecretRefOk() (*string, bool)`

GetSecretRefOk returns a tuple with the SecretRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretRef

`func (o *BnsLoadBalancerV1ApiListTlsCertificatesModelSecretModel) SetSecretRef(v string)`

SetSecretRef sets SecretRef field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


