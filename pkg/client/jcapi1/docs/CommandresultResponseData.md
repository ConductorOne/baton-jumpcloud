# CommandresultResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExitCode** | Pointer to **int32** | The stderr output from the command that ran. | [optional] 
**Output** | Pointer to **string** | The output of the command that was executed. | [optional] 

## Methods

### NewCommandresultResponseData

`func NewCommandresultResponseData() *CommandresultResponseData`

NewCommandresultResponseData instantiates a new CommandresultResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommandresultResponseDataWithDefaults

`func NewCommandresultResponseDataWithDefaults() *CommandresultResponseData`

NewCommandresultResponseDataWithDefaults instantiates a new CommandresultResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExitCode

`func (o *CommandresultResponseData) GetExitCode() int32`

GetExitCode returns the ExitCode field if non-nil, zero value otherwise.

### GetExitCodeOk

`func (o *CommandresultResponseData) GetExitCodeOk() (*int32, bool)`

GetExitCodeOk returns a tuple with the ExitCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExitCode

`func (o *CommandresultResponseData) SetExitCode(v int32)`

SetExitCode sets ExitCode field to given value.

### HasExitCode

`func (o *CommandresultResponseData) HasExitCode() bool`

HasExitCode returns a boolean if a field has been set.

### GetOutput

`func (o *CommandresultResponseData) GetOutput() string`

GetOutput returns the Output field if non-nil, zero value otherwise.

### GetOutputOk

`func (o *CommandresultResponseData) GetOutputOk() (*string, bool)`

GetOutputOk returns a tuple with the Output field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutput

`func (o *CommandresultResponseData) SetOutput(v string)`

SetOutput sets Output field to given value.

### HasOutput

`func (o *CommandresultResponseData) HasOutput() bool`

HasOutput returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


