# GetK8sUpgradeVersionsResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Versions** | [**[]KubernetesEngineV1ApiListAvailableKubernetesVersionsModelUpgradableVersionResponseModel**](KubernetesEngineV1ApiListAvailableKubernetesVersionsModelUpgradableVersionResponseModel.md) | Kubernetes 버전 정보 목록 | 

## Methods

### NewGetK8sUpgradeVersionsResponseModel

`func NewGetK8sUpgradeVersionsResponseModel(versions []KubernetesEngineV1ApiListAvailableKubernetesVersionsModelUpgradableVersionResponseModel, ) *GetK8sUpgradeVersionsResponseModel`

NewGetK8sUpgradeVersionsResponseModel instantiates a new GetK8sUpgradeVersionsResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetK8sUpgradeVersionsResponseModelWithDefaults

`func NewGetK8sUpgradeVersionsResponseModelWithDefaults() *GetK8sUpgradeVersionsResponseModel`

NewGetK8sUpgradeVersionsResponseModelWithDefaults instantiates a new GetK8sUpgradeVersionsResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVersions

`func (o *GetK8sUpgradeVersionsResponseModel) GetVersions() []KubernetesEngineV1ApiListAvailableKubernetesVersionsModelUpgradableVersionResponseModel`

GetVersions returns the Versions field if non-nil, zero value otherwise.

### GetVersionsOk

`func (o *GetK8sUpgradeVersionsResponseModel) GetVersionsOk() (*[]KubernetesEngineV1ApiListAvailableKubernetesVersionsModelUpgradableVersionResponseModel, bool)`

GetVersionsOk returns a tuple with the Versions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersions

`func (o *GetK8sUpgradeVersionsResponseModel) SetVersions(v []KubernetesEngineV1ApiListAvailableKubernetesVersionsModelUpgradableVersionResponseModel)`

SetVersions sets Versions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


