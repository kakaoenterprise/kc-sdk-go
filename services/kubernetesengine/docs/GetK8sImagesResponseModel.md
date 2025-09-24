# GetK8sImagesResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Images** | [**[]ImageResponseModel**](ImageResponseModel.md) | 이미지 정보 | 

## Methods

### NewGetK8sImagesResponseModel

`func NewGetK8sImagesResponseModel(images []ImageResponseModel, ) *GetK8sImagesResponseModel`

NewGetK8sImagesResponseModel instantiates a new GetK8sImagesResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetK8sImagesResponseModelWithDefaults

`func NewGetK8sImagesResponseModelWithDefaults() *GetK8sImagesResponseModel`

NewGetK8sImagesResponseModelWithDefaults instantiates a new GetK8sImagesResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetImages

`func (o *GetK8sImagesResponseModel) GetImages() []ImageResponseModel`

GetImages returns the Images field if non-nil, zero value otherwise.

### GetImagesOk

`func (o *GetK8sImagesResponseModel) GetImagesOk() (*[]ImageResponseModel, bool)`

GetImagesOk returns a tuple with the Images field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImages

`func (o *GetK8sImagesResponseModel) SetImages(v []ImageResponseModel)`

SetImages sets Images field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


