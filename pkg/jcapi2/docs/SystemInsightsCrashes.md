# SystemInsightsCrashes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CollectionTime** | Pointer to **string** |  | [optional] 
**CrashPath** | Pointer to **string** |  | [optional] 
**CrashedThread** | Pointer to **string** |  | [optional] 
**Datetime** | Pointer to **string** |  | [optional] 
**ExceptionCodes** | Pointer to **string** |  | [optional] 
**ExceptionNotes** | Pointer to **string** |  | [optional] 
**ExceptionType** | Pointer to **string** |  | [optional] 
**Identifier** | Pointer to **string** |  | [optional] 
**Parent** | Pointer to **string** |  | [optional] 
**Path** | Pointer to **string** |  | [optional] 
**Pid** | Pointer to **string** |  | [optional] 
**Registers** | Pointer to **string** |  | [optional] 
**Responsible** | Pointer to **string** |  | [optional] 
**StackTrace** | Pointer to **string** |  | [optional] 
**SystemId** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Uid** | Pointer to **int32** |  | [optional] 
**Version** | Pointer to **string** |  | [optional] 

## Methods

### NewSystemInsightsCrashes

`func NewSystemInsightsCrashes() *SystemInsightsCrashes`

NewSystemInsightsCrashes instantiates a new SystemInsightsCrashes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSystemInsightsCrashesWithDefaults

`func NewSystemInsightsCrashesWithDefaults() *SystemInsightsCrashes`

NewSystemInsightsCrashesWithDefaults instantiates a new SystemInsightsCrashes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCollectionTime

`func (o *SystemInsightsCrashes) GetCollectionTime() string`

GetCollectionTime returns the CollectionTime field if non-nil, zero value otherwise.

### GetCollectionTimeOk

`func (o *SystemInsightsCrashes) GetCollectionTimeOk() (*string, bool)`

GetCollectionTimeOk returns a tuple with the CollectionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionTime

`func (o *SystemInsightsCrashes) SetCollectionTime(v string)`

SetCollectionTime sets CollectionTime field to given value.

### HasCollectionTime

`func (o *SystemInsightsCrashes) HasCollectionTime() bool`

HasCollectionTime returns a boolean if a field has been set.

### GetCrashPath

`func (o *SystemInsightsCrashes) GetCrashPath() string`

GetCrashPath returns the CrashPath field if non-nil, zero value otherwise.

### GetCrashPathOk

`func (o *SystemInsightsCrashes) GetCrashPathOk() (*string, bool)`

GetCrashPathOk returns a tuple with the CrashPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrashPath

`func (o *SystemInsightsCrashes) SetCrashPath(v string)`

SetCrashPath sets CrashPath field to given value.

### HasCrashPath

`func (o *SystemInsightsCrashes) HasCrashPath() bool`

HasCrashPath returns a boolean if a field has been set.

### GetCrashedThread

`func (o *SystemInsightsCrashes) GetCrashedThread() string`

GetCrashedThread returns the CrashedThread field if non-nil, zero value otherwise.

### GetCrashedThreadOk

`func (o *SystemInsightsCrashes) GetCrashedThreadOk() (*string, bool)`

GetCrashedThreadOk returns a tuple with the CrashedThread field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrashedThread

`func (o *SystemInsightsCrashes) SetCrashedThread(v string)`

SetCrashedThread sets CrashedThread field to given value.

### HasCrashedThread

`func (o *SystemInsightsCrashes) HasCrashedThread() bool`

HasCrashedThread returns a boolean if a field has been set.

### GetDatetime

`func (o *SystemInsightsCrashes) GetDatetime() string`

GetDatetime returns the Datetime field if non-nil, zero value otherwise.

### GetDatetimeOk

`func (o *SystemInsightsCrashes) GetDatetimeOk() (*string, bool)`

GetDatetimeOk returns a tuple with the Datetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatetime

`func (o *SystemInsightsCrashes) SetDatetime(v string)`

SetDatetime sets Datetime field to given value.

### HasDatetime

`func (o *SystemInsightsCrashes) HasDatetime() bool`

HasDatetime returns a boolean if a field has been set.

### GetExceptionCodes

`func (o *SystemInsightsCrashes) GetExceptionCodes() string`

GetExceptionCodes returns the ExceptionCodes field if non-nil, zero value otherwise.

### GetExceptionCodesOk

`func (o *SystemInsightsCrashes) GetExceptionCodesOk() (*string, bool)`

GetExceptionCodesOk returns a tuple with the ExceptionCodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExceptionCodes

`func (o *SystemInsightsCrashes) SetExceptionCodes(v string)`

SetExceptionCodes sets ExceptionCodes field to given value.

### HasExceptionCodes

`func (o *SystemInsightsCrashes) HasExceptionCodes() bool`

HasExceptionCodes returns a boolean if a field has been set.

### GetExceptionNotes

`func (o *SystemInsightsCrashes) GetExceptionNotes() string`

GetExceptionNotes returns the ExceptionNotes field if non-nil, zero value otherwise.

### GetExceptionNotesOk

`func (o *SystemInsightsCrashes) GetExceptionNotesOk() (*string, bool)`

GetExceptionNotesOk returns a tuple with the ExceptionNotes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExceptionNotes

`func (o *SystemInsightsCrashes) SetExceptionNotes(v string)`

SetExceptionNotes sets ExceptionNotes field to given value.

### HasExceptionNotes

`func (o *SystemInsightsCrashes) HasExceptionNotes() bool`

HasExceptionNotes returns a boolean if a field has been set.

### GetExceptionType

`func (o *SystemInsightsCrashes) GetExceptionType() string`

GetExceptionType returns the ExceptionType field if non-nil, zero value otherwise.

### GetExceptionTypeOk

`func (o *SystemInsightsCrashes) GetExceptionTypeOk() (*string, bool)`

GetExceptionTypeOk returns a tuple with the ExceptionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExceptionType

`func (o *SystemInsightsCrashes) SetExceptionType(v string)`

SetExceptionType sets ExceptionType field to given value.

### HasExceptionType

`func (o *SystemInsightsCrashes) HasExceptionType() bool`

HasExceptionType returns a boolean if a field has been set.

### GetIdentifier

`func (o *SystemInsightsCrashes) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *SystemInsightsCrashes) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *SystemInsightsCrashes) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.

### HasIdentifier

`func (o *SystemInsightsCrashes) HasIdentifier() bool`

HasIdentifier returns a boolean if a field has been set.

### GetParent

`func (o *SystemInsightsCrashes) GetParent() string`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *SystemInsightsCrashes) GetParentOk() (*string, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *SystemInsightsCrashes) SetParent(v string)`

SetParent sets Parent field to given value.

### HasParent

`func (o *SystemInsightsCrashes) HasParent() bool`

HasParent returns a boolean if a field has been set.

### GetPath

`func (o *SystemInsightsCrashes) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *SystemInsightsCrashes) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *SystemInsightsCrashes) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *SystemInsightsCrashes) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetPid

`func (o *SystemInsightsCrashes) GetPid() string`

GetPid returns the Pid field if non-nil, zero value otherwise.

### GetPidOk

`func (o *SystemInsightsCrashes) GetPidOk() (*string, bool)`

GetPidOk returns a tuple with the Pid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPid

`func (o *SystemInsightsCrashes) SetPid(v string)`

SetPid sets Pid field to given value.

### HasPid

`func (o *SystemInsightsCrashes) HasPid() bool`

HasPid returns a boolean if a field has been set.

### GetRegisters

`func (o *SystemInsightsCrashes) GetRegisters() string`

GetRegisters returns the Registers field if non-nil, zero value otherwise.

### GetRegistersOk

`func (o *SystemInsightsCrashes) GetRegistersOk() (*string, bool)`

GetRegistersOk returns a tuple with the Registers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisters

`func (o *SystemInsightsCrashes) SetRegisters(v string)`

SetRegisters sets Registers field to given value.

### HasRegisters

`func (o *SystemInsightsCrashes) HasRegisters() bool`

HasRegisters returns a boolean if a field has been set.

### GetResponsible

`func (o *SystemInsightsCrashes) GetResponsible() string`

GetResponsible returns the Responsible field if non-nil, zero value otherwise.

### GetResponsibleOk

`func (o *SystemInsightsCrashes) GetResponsibleOk() (*string, bool)`

GetResponsibleOk returns a tuple with the Responsible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponsible

`func (o *SystemInsightsCrashes) SetResponsible(v string)`

SetResponsible sets Responsible field to given value.

### HasResponsible

`func (o *SystemInsightsCrashes) HasResponsible() bool`

HasResponsible returns a boolean if a field has been set.

### GetStackTrace

`func (o *SystemInsightsCrashes) GetStackTrace() string`

GetStackTrace returns the StackTrace field if non-nil, zero value otherwise.

### GetStackTraceOk

`func (o *SystemInsightsCrashes) GetStackTraceOk() (*string, bool)`

GetStackTraceOk returns a tuple with the StackTrace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackTrace

`func (o *SystemInsightsCrashes) SetStackTrace(v string)`

SetStackTrace sets StackTrace field to given value.

### HasStackTrace

`func (o *SystemInsightsCrashes) HasStackTrace() bool`

HasStackTrace returns a boolean if a field has been set.

### GetSystemId

`func (o *SystemInsightsCrashes) GetSystemId() string`

GetSystemId returns the SystemId field if non-nil, zero value otherwise.

### GetSystemIdOk

`func (o *SystemInsightsCrashes) GetSystemIdOk() (*string, bool)`

GetSystemIdOk returns a tuple with the SystemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemId

`func (o *SystemInsightsCrashes) SetSystemId(v string)`

SetSystemId sets SystemId field to given value.

### HasSystemId

`func (o *SystemInsightsCrashes) HasSystemId() bool`

HasSystemId returns a boolean if a field has been set.

### GetType

`func (o *SystemInsightsCrashes) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SystemInsightsCrashes) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SystemInsightsCrashes) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *SystemInsightsCrashes) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUid

`func (o *SystemInsightsCrashes) GetUid() int32`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *SystemInsightsCrashes) GetUidOk() (*int32, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *SystemInsightsCrashes) SetUid(v int32)`

SetUid sets Uid field to given value.

### HasUid

`func (o *SystemInsightsCrashes) HasUid() bool`

HasUid returns a boolean if a field has been set.

### GetVersion

`func (o *SystemInsightsCrashes) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *SystemInsightsCrashes) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *SystemInsightsCrashes) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *SystemInsightsCrashes) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


