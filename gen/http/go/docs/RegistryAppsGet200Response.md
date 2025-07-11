# RegistryAppsGet200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**HttpListAppResponse**](HttpListAppResponse.md) |  | [optional] 
**ErrorCode** | Pointer to **string** |  | [optional] 
**Ok** | Pointer to **bool** |  | [optional] 

## Methods

### NewRegistryAppsGet200Response

`func NewRegistryAppsGet200Response() *RegistryAppsGet200Response`

NewRegistryAppsGet200Response instantiates a new RegistryAppsGet200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegistryAppsGet200ResponseWithDefaults

`func NewRegistryAppsGet200ResponseWithDefaults() *RegistryAppsGet200Response`

NewRegistryAppsGet200ResponseWithDefaults instantiates a new RegistryAppsGet200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *RegistryAppsGet200Response) GetData() HttpListAppResponse`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *RegistryAppsGet200Response) GetDataOk() (*HttpListAppResponse, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *RegistryAppsGet200Response) SetData(v HttpListAppResponse)`

SetData sets Data field to given value.

### HasData

`func (o *RegistryAppsGet200Response) HasData() bool`

HasData returns a boolean if a field has been set.

### GetErrorCode

`func (o *RegistryAppsGet200Response) GetErrorCode() string`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *RegistryAppsGet200Response) GetErrorCodeOk() (*string, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *RegistryAppsGet200Response) SetErrorCode(v string)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *RegistryAppsGet200Response) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### GetOk

`func (o *RegistryAppsGet200Response) GetOk() bool`

GetOk returns the Ok field if non-nil, zero value otherwise.

### GetOkOk

`func (o *RegistryAppsGet200Response) GetOkOk() (*bool, bool)`

GetOkOk returns a tuple with the Ok field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOk

`func (o *RegistryAppsGet200Response) SetOk(v bool)`

SetOk sets Ok field to given value.

### HasOk

`func (o *RegistryAppsGet200Response) HasOk() bool`

HasOk returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


