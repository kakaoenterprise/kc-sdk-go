# ListTlsCertificatesSecret

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | **time.Time** | 리소스가 생성된 시간 - ISO 8601 형식 - UTC 기준 | 
**UpdatedAt** | **time.Time** | 리소스가 마지막으로 수정된 시간 - ISO 8601 형식 - UTC 기준 | 
**Status** | **string** | 인증서 상태 | 
**Name** | **string** | 인증서 이름 | 
**SecretType** | **string** | 인증서 유형 | 
**Expiration** | **time.Time** | 인증서 만료 일시 | 
**CreatorId** | **string** | 인증서 생성자 ID | 
**ContentTypes** | [**ContentTypesDto**](ContentTypesDto.md) | 인증서 콘텐츠 유형 목록 | 
**SecretRef** | **string** | 인증서 참조 ID | 

## Methods

### NewListTlsCertificatesSecret

`func NewListTlsCertificatesSecret(createdAt time.Time, updatedAt time.Time, status string, name string, secretType string, expiration time.Time, creatorId string, contentTypes ContentTypesDto, secretRef string, ) *ListTlsCertificatesSecret`

NewListTlsCertificatesSecret instantiates a new ListTlsCertificatesSecret object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListTlsCertificatesSecretWithDefaults

`func NewListTlsCertificatesSecretWithDefaults() *ListTlsCertificatesSecret`

NewListTlsCertificatesSecretWithDefaults instantiates a new ListTlsCertificatesSecret object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *ListTlsCertificatesSecret) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ListTlsCertificatesSecret) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ListTlsCertificatesSecret) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *ListTlsCertificatesSecret) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ListTlsCertificatesSecret) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ListTlsCertificatesSecret) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetStatus

`func (o *ListTlsCertificatesSecret) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ListTlsCertificatesSecret) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ListTlsCertificatesSecret) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetName

`func (o *ListTlsCertificatesSecret) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ListTlsCertificatesSecret) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ListTlsCertificatesSecret) SetName(v string)`

SetName sets Name field to given value.


### GetSecretType

`func (o *ListTlsCertificatesSecret) GetSecretType() string`

GetSecretType returns the SecretType field if non-nil, zero value otherwise.

### GetSecretTypeOk

`func (o *ListTlsCertificatesSecret) GetSecretTypeOk() (*string, bool)`

GetSecretTypeOk returns a tuple with the SecretType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretType

`func (o *ListTlsCertificatesSecret) SetSecretType(v string)`

SetSecretType sets SecretType field to given value.


### GetExpiration

`func (o *ListTlsCertificatesSecret) GetExpiration() time.Time`

GetExpiration returns the Expiration field if non-nil, zero value otherwise.

### GetExpirationOk

`func (o *ListTlsCertificatesSecret) GetExpirationOk() (*time.Time, bool)`

GetExpirationOk returns a tuple with the Expiration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiration

`func (o *ListTlsCertificatesSecret) SetExpiration(v time.Time)`

SetExpiration sets Expiration field to given value.


### GetCreatorId

`func (o *ListTlsCertificatesSecret) GetCreatorId() string`

GetCreatorId returns the CreatorId field if non-nil, zero value otherwise.

### GetCreatorIdOk

`func (o *ListTlsCertificatesSecret) GetCreatorIdOk() (*string, bool)`

GetCreatorIdOk returns a tuple with the CreatorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorId

`func (o *ListTlsCertificatesSecret) SetCreatorId(v string)`

SetCreatorId sets CreatorId field to given value.


### GetContentTypes

`func (o *ListTlsCertificatesSecret) GetContentTypes() ContentTypesDto`

GetContentTypes returns the ContentTypes field if non-nil, zero value otherwise.

### GetContentTypesOk

`func (o *ListTlsCertificatesSecret) GetContentTypesOk() (*ContentTypesDto, bool)`

GetContentTypesOk returns a tuple with the ContentTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentTypes

`func (o *ListTlsCertificatesSecret) SetContentTypes(v ContentTypesDto)`

SetContentTypes sets ContentTypes field to given value.


### GetSecretRef

`func (o *ListTlsCertificatesSecret) GetSecretRef() string`

GetSecretRef returns the SecretRef field if non-nil, zero value otherwise.

### GetSecretRefOk

`func (o *ListTlsCertificatesSecret) GetSecretRefOk() (*string, bool)`

GetSecretRefOk returns a tuple with the SecretRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretRef

`func (o *ListTlsCertificatesSecret) SetSecretRef(v string)`

SetSecretRef sets SecretRef field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


