# GetK8sClusterUpgradableVersionsResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Upgrade** | [**UpgradeResponseModel**](UpgradeResponseModel.md) | 클러스터 업그레이드 정보 | 

## Methods

### NewGetK8sClusterUpgradableVersionsResponseModel

`func NewGetK8sClusterUpgradableVersionsResponseModel(upgrade UpgradeResponseModel, ) *GetK8sClusterUpgradableVersionsResponseModel`

NewGetK8sClusterUpgradableVersionsResponseModel instantiates a new GetK8sClusterUpgradableVersionsResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetK8sClusterUpgradableVersionsResponseModelWithDefaults

`func NewGetK8sClusterUpgradableVersionsResponseModelWithDefaults() *GetK8sClusterUpgradableVersionsResponseModel`

NewGetK8sClusterUpgradableVersionsResponseModelWithDefaults instantiates a new GetK8sClusterUpgradableVersionsResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUpgrade

`func (o *GetK8sClusterUpgradableVersionsResponseModel) GetUpgrade() UpgradeResponseModel`

GetUpgrade returns the Upgrade field if non-nil, zero value otherwise.

### GetUpgradeOk

`func (o *GetK8sClusterUpgradableVersionsResponseModel) GetUpgradeOk() (*UpgradeResponseModel, bool)`

GetUpgradeOk returns a tuple with the Upgrade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpgrade

`func (o *GetK8sClusterUpgradableVersionsResponseModel) SetUpgrade(v UpgradeResponseModel)`

SetUpgrade sets Upgrade field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


