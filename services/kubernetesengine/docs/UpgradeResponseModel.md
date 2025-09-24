# UpgradeResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Upgradable** | Pointer to [**[]KubernetesEngineV1ApiListClusterUpgradableVersionsModelUpgradableVersionResponseModel**](KubernetesEngineV1ApiListClusterUpgradableVersionsModelUpgradableVersionResponseModel.md) |  | [optional] 
**Current** | [**CurrentOMTResponseModel**](CurrentOMTResponseModel.md) | 현재 클러스터 버전 | 

## Methods

### NewUpgradeResponseModel

`func NewUpgradeResponseModel(current CurrentOMTResponseModel, ) *UpgradeResponseModel`

NewUpgradeResponseModel instantiates a new UpgradeResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpgradeResponseModelWithDefaults

`func NewUpgradeResponseModelWithDefaults() *UpgradeResponseModel`

NewUpgradeResponseModelWithDefaults instantiates a new UpgradeResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUpgradable

`func (o *UpgradeResponseModel) GetUpgradable() []KubernetesEngineV1ApiListClusterUpgradableVersionsModelUpgradableVersionResponseModel`

GetUpgradable returns the Upgradable field if non-nil, zero value otherwise.

### GetUpgradableOk

`func (o *UpgradeResponseModel) GetUpgradableOk() (*[]KubernetesEngineV1ApiListClusterUpgradableVersionsModelUpgradableVersionResponseModel, bool)`

GetUpgradableOk returns a tuple with the Upgradable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpgradable

`func (o *UpgradeResponseModel) SetUpgradable(v []KubernetesEngineV1ApiListClusterUpgradableVersionsModelUpgradableVersionResponseModel)`

SetUpgradable sets Upgradable field to given value.

### HasUpgradable

`func (o *UpgradeResponseModel) HasUpgradable() bool`

HasUpgradable returns a boolean if a field has been set.

### SetUpgradableNil

`func (o *UpgradeResponseModel) SetUpgradableNil(b bool)`

 SetUpgradableNil sets the value for Upgradable to be an explicit nil

### UnsetUpgradable
`func (o *UpgradeResponseModel) UnsetUpgradable()`

UnsetUpgradable ensures that no value is present for Upgradable, not even an explicit nil
### GetCurrent

`func (o *UpgradeResponseModel) GetCurrent() CurrentOMTResponseModel`

GetCurrent returns the Current field if non-nil, zero value otherwise.

### GetCurrentOk

`func (o *UpgradeResponseModel) GetCurrentOk() (*CurrentOMTResponseModel, bool)`

GetCurrentOk returns a tuple with the Current field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrent

`func (o *UpgradeResponseModel) SetCurrent(v CurrentOMTResponseModel)`

SetCurrent sets Current field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


