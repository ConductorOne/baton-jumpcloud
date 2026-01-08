# GraphOperationUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The ObjectID of graph object being added or removed as an association. | 
**Op** | **string** | How to modify the graph connection. | 
**Attributes** | Pointer to [**GraphOperationSystemAllOfAttributes**](GraphOperationSystemAllOfAttributes.md) |  | [optional] 
**Type** | **string** | Targets which a \&quot;user\&quot; can be associated to. | 

## Methods

### NewGraphOperationUser

`func NewGraphOperationUser(id string, op string, type_ string, ) *GraphOperationUser`

NewGraphOperationUser instantiates a new GraphOperationUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGraphOperationUserWithDefaults

`func NewGraphOperationUserWithDefaults() *GraphOperationUser`

NewGraphOperationUserWithDefaults instantiates a new GraphOperationUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GraphOperationUser) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GraphOperationUser) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GraphOperationUser) SetId(v string)`

SetId sets Id field to given value.


### GetOp

`func (o *GraphOperationUser) GetOp() string`

GetOp returns the Op field if non-nil, zero value otherwise.

### GetOpOk

`func (o *GraphOperationUser) GetOpOk() (*string, bool)`

GetOpOk returns a tuple with the Op field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOp

`func (o *GraphOperationUser) SetOp(v string)`

SetOp sets Op field to given value.


### GetAttributes

`func (o *GraphOperationUser) GetAttributes() GraphOperationSystemAllOfAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *GraphOperationUser) GetAttributesOk() (*GraphOperationSystemAllOfAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *GraphOperationUser) SetAttributes(v GraphOperationSystemAllOfAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *GraphOperationUser) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetType

`func (o *GraphOperationUser) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GraphOperationUser) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GraphOperationUser) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


