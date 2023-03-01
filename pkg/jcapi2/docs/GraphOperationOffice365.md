# GraphOperationOffice365

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The ObjectID of graph object being added or removed as an association. | 
**Op** | **string** | How to modify the graph connection. | 
**Attributes** | Pointer to **map[string]interface{}** | The graph attributes. | [optional] 
**Type** | **string** | Targets which a \&quot;office_365\&quot; can be associated to. | 

## Methods

### NewGraphOperationOffice365

`func NewGraphOperationOffice365(id string, op string, type_ string, ) *GraphOperationOffice365`

NewGraphOperationOffice365 instantiates a new GraphOperationOffice365 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGraphOperationOffice365WithDefaults

`func NewGraphOperationOffice365WithDefaults() *GraphOperationOffice365`

NewGraphOperationOffice365WithDefaults instantiates a new GraphOperationOffice365 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GraphOperationOffice365) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GraphOperationOffice365) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GraphOperationOffice365) SetId(v string)`

SetId sets Id field to given value.


### GetOp

`func (o *GraphOperationOffice365) GetOp() string`

GetOp returns the Op field if non-nil, zero value otherwise.

### GetOpOk

`func (o *GraphOperationOffice365) GetOpOk() (*string, bool)`

GetOpOk returns a tuple with the Op field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOp

`func (o *GraphOperationOffice365) SetOp(v string)`

SetOp sets Op field to given value.


### GetAttributes

`func (o *GraphOperationOffice365) GetAttributes() map[string]interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *GraphOperationOffice365) GetAttributesOk() (*map[string]interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *GraphOperationOffice365) SetAttributes(v map[string]interface{})`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *GraphOperationOffice365) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetType

`func (o *GraphOperationOffice365) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GraphOperationOffice365) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GraphOperationOffice365) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


