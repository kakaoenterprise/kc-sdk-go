# KubernetesEngineV1ApiListClusterUpgradableVersionsModelUpgradableVersionResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsDeprecated** | **bool** | 해당 버전의 사용 중단 여부 | 
**Eol** | **string** | 지원 종료 시점(ISO 8601 형식, UTC) | 
**MinorVersion** | **string** | Kubernetes 마이너 버전 | 
**PatchVersion** | **string** | Kubernetes 패치 버전 | 
**NextVersion** | **string** | 업그레이드 가능한 다음 버전 | 

## Methods

### NewKubernetesEngineV1ApiListClusterUpgradableVersionsModelUpgradableVersionResponseModel

`func NewKubernetesEngineV1ApiListClusterUpgradableVersionsModelUpgradableVersionResponseModel(isDeprecated bool, eol string, minorVersion string, patchVersion string, nextVersion string, ) *KubernetesEngineV1ApiListClusterUpgradableVersionsModelUpgradableVersionResponseModel`

NewKubernetesEngineV1ApiListClusterUpgradableVersionsModelUpgradableVersionResponseModel instantiates a new KubernetesEngineV1ApiListClusterUpgradableVersionsModelUpgradableVersionResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubernetesEngineV1ApiListClusterUpgradableVersionsModelUpgradableVersionResponseModelWithDefaults

`func NewKubernetesEngineV1ApiListClusterUpgradableVersionsModelUpgradableVersionResponseModelWithDefaults() *KubernetesEngineV1ApiListClusterUpgradableVersionsModelUpgradableVersionResponseModel`

NewKubernetesEngineV1ApiListClusterUpgradableVersionsModelUpgradableVersionResponseModelWithDefaults instantiates a new KubernetesEngineV1ApiListClusterUpgradableVersionsModelUpgradableVersionResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsDeprecated

`func (o *KubernetesEngineV1ApiListClusterUpgradableVersionsModelUpgradableVersionResponseModel) GetIsDeprecated() bool`

GetIsDeprecated returns the IsDeprecated field if non-nil, zero value otherwise.

### GetIsDeprecatedOk

`func (o *KubernetesEngineV1ApiListClusterUpgradableVersionsModelUpgradableVersionResponseModel) GetIsDeprecatedOk() (*bool, bool)`

GetIsDeprecatedOk returns a tuple with the IsDeprecated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeprecated

`func (o *KubernetesEngineV1ApiListClusterUpgradableVersionsModelUpgradableVersionResponseModel) SetIsDeprecated(v bool)`

SetIsDeprecated sets IsDeprecated field to given value.


### GetEol

`func (o *KubernetesEngineV1ApiListClusterUpgradableVersionsModelUpgradableVersionResponseModel) GetEol() string`

GetEol returns the Eol field if non-nil, zero value otherwise.

### GetEolOk

`func (o *KubernetesEngineV1ApiListClusterUpgradableVersionsModelUpgradableVersionResponseModel) GetEolOk() (*string, bool)`

GetEolOk returns a tuple with the Eol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEol

`func (o *KubernetesEngineV1ApiListClusterUpgradableVersionsModelUpgradableVersionResponseModel) SetEol(v string)`

SetEol sets Eol field to given value.


### GetMinorVersion

`func (o *KubernetesEngineV1ApiListClusterUpgradableVersionsModelUpgradableVersionResponseModel) GetMinorVersion() string`

GetMinorVersion returns the MinorVersion field if non-nil, zero value otherwise.

### GetMinorVersionOk

`func (o *KubernetesEngineV1ApiListClusterUpgradableVersionsModelUpgradableVersionResponseModel) GetMinorVersionOk() (*string, bool)`

GetMinorVersionOk returns a tuple with the MinorVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinorVersion

`func (o *KubernetesEngineV1ApiListClusterUpgradableVersionsModelUpgradableVersionResponseModel) SetMinorVersion(v string)`

SetMinorVersion sets MinorVersion field to given value.


### GetPatchVersion

`func (o *KubernetesEngineV1ApiListClusterUpgradableVersionsModelUpgradableVersionResponseModel) GetPatchVersion() string`

GetPatchVersion returns the PatchVersion field if non-nil, zero value otherwise.

### GetPatchVersionOk

`func (o *KubernetesEngineV1ApiListClusterUpgradableVersionsModelUpgradableVersionResponseModel) GetPatchVersionOk() (*string, bool)`

GetPatchVersionOk returns a tuple with the PatchVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPatchVersion

`func (o *KubernetesEngineV1ApiListClusterUpgradableVersionsModelUpgradableVersionResponseModel) SetPatchVersion(v string)`

SetPatchVersion sets PatchVersion field to given value.


### GetNextVersion

`func (o *KubernetesEngineV1ApiListClusterUpgradableVersionsModelUpgradableVersionResponseModel) GetNextVersion() string`

GetNextVersion returns the NextVersion field if non-nil, zero value otherwise.

### GetNextVersionOk

`func (o *KubernetesEngineV1ApiListClusterUpgradableVersionsModelUpgradableVersionResponseModel) GetNextVersionOk() (*string, bool)`

GetNextVersionOk returns a tuple with the NextVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextVersion

`func (o *KubernetesEngineV1ApiListClusterUpgradableVersionsModelUpgradableVersionResponseModel) SetNextVersion(v string)`

SetNextVersion sets NextVersion field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


