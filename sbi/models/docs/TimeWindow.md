# TimeWindow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StartTime** | **string** | string with format \&quot;date-time\&quot; as defined in OpenAPI. | 
**StopTime** | **string** | string with format \&quot;date-time\&quot; as defined in OpenAPI. | 

## Methods

### NewTimeWindow

`func NewTimeWindow(startTime string, stopTime string, ) *TimeWindow`

NewTimeWindow instantiates a new TimeWindow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTimeWindowWithDefaults

`func NewTimeWindowWithDefaults() *TimeWindow`

NewTimeWindowWithDefaults instantiates a new TimeWindow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStartTime

`func (o *TimeWindow) GetStartTime() string`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *TimeWindow) GetStartTimeOk() (*string, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *TimeWindow) SetStartTime(v string)`

SetStartTime sets StartTime field to given value.


### GetStopTime

`func (o *TimeWindow) GetStopTime() string`

GetStopTime returns the StopTime field if non-nil, zero value otherwise.

### GetStopTimeOk

`func (o *TimeWindow) GetStopTimeOk() (*string, bool)`

GetStopTimeOk returns a tuple with the StopTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStopTime

`func (o *TimeWindow) SetStopTime(v string)`

SetStopTime sets StopTime field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


