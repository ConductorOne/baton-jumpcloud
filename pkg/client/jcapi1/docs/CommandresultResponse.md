# CommandresultResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**CommandresultResponseData**](CommandresultResponseData.md) |  | [optional] 
**Error** | Pointer to **string** | The stderr output from the command that ran. | [optional] 
**Id** | Pointer to **string** | ID of the response. | [optional] 

## Methods

### NewCommandresultResponse

`func NewCommandresultResponse() *CommandresultResponse`

NewCommandresultResponse instantiates a new CommandresultResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommandresultResponseWithDefaults

`func NewCommandresultResponseWithDefaults() *CommandresultResponse`

NewCommandresultResponseWithDefaults instantiates a new CommandresultResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *CommandresultResponse) GetData() CommandresultResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CommandresultResponse) GetDataOk() (*CommandresultResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CommandresultResponse) SetData(v CommandresultResponseData)`

SetData sets Data field to given value.

### HasData

`func (o *CommandresultResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetError

`func (o *CommandresultResponse) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *CommandresultResponse) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *CommandresultResponse) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *CommandresultResponse) HasError() bool`

HasError returns a boolean if a field has been set.

### GetId

`func (o *CommandresultResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CommandresultResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CommandresultResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CommandresultResponse) HasId() bool`

HasId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


