# Room

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RoomId** | Pointer to **string** |  | [optional] 
**VideoId** | Pointer to **string** |  | [optional] 
**AudioCodec** | Pointer to **string** |  | [optional] 
**VideoCodec** | Pointer to **string** |  | [optional] 
**Record** | Pointer to **bool** |  | [optional] 
**RoomServerFqdn** | Pointer to **string** |  | [optional] 
**ParticipantsMax** | Pointer to **int32** |  | [optional] 
**Participants** | Pointer to [**[]Participant**](Participant.md) |  | [optional] 

## Methods

### NewRoom

`func NewRoom() *Room`

NewRoom instantiates a new Room object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoomWithDefaults

`func NewRoomWithDefaults() *Room`

NewRoomWithDefaults instantiates a new Room object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRoomId

`func (o *Room) GetRoomId() string`

GetRoomId returns the RoomId field if non-nil, zero value otherwise.

### GetRoomIdOk

`func (o *Room) GetRoomIdOk() (*string, bool)`

GetRoomIdOk returns a tuple with the RoomId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoomId

`func (o *Room) SetRoomId(v string)`

SetRoomId sets RoomId field to given value.

### HasRoomId

`func (o *Room) HasRoomId() bool`

HasRoomId returns a boolean if a field has been set.

### GetVideoId

`func (o *Room) GetVideoId() string`

GetVideoId returns the VideoId field if non-nil, zero value otherwise.

### GetVideoIdOk

`func (o *Room) GetVideoIdOk() (*string, bool)`

GetVideoIdOk returns a tuple with the VideoId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideoId

`func (o *Room) SetVideoId(v string)`

SetVideoId sets VideoId field to given value.

### HasVideoId

`func (o *Room) HasVideoId() bool`

HasVideoId returns a boolean if a field has been set.

### GetAudioCodec

`func (o *Room) GetAudioCodec() string`

GetAudioCodec returns the AudioCodec field if non-nil, zero value otherwise.

### GetAudioCodecOk

`func (o *Room) GetAudioCodecOk() (*string, bool)`

GetAudioCodecOk returns a tuple with the AudioCodec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudioCodec

`func (o *Room) SetAudioCodec(v string)`

SetAudioCodec sets AudioCodec field to given value.

### HasAudioCodec

`func (o *Room) HasAudioCodec() bool`

HasAudioCodec returns a boolean if a field has been set.

### GetVideoCodec

`func (o *Room) GetVideoCodec() string`

GetVideoCodec returns the VideoCodec field if non-nil, zero value otherwise.

### GetVideoCodecOk

`func (o *Room) GetVideoCodecOk() (*string, bool)`

GetVideoCodecOk returns a tuple with the VideoCodec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideoCodec

`func (o *Room) SetVideoCodec(v string)`

SetVideoCodec sets VideoCodec field to given value.

### HasVideoCodec

`func (o *Room) HasVideoCodec() bool`

HasVideoCodec returns a boolean if a field has been set.

### GetRecord

`func (o *Room) GetRecord() bool`

GetRecord returns the Record field if non-nil, zero value otherwise.

### GetRecordOk

`func (o *Room) GetRecordOk() (*bool, bool)`

GetRecordOk returns a tuple with the Record field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecord

`func (o *Room) SetRecord(v bool)`

SetRecord sets Record field to given value.

### HasRecord

`func (o *Room) HasRecord() bool`

HasRecord returns a boolean if a field has been set.

### GetRoomServerFqdn

`func (o *Room) GetRoomServerFqdn() string`

GetRoomServerFqdn returns the RoomServerFqdn field if non-nil, zero value otherwise.

### GetRoomServerFqdnOk

`func (o *Room) GetRoomServerFqdnOk() (*string, bool)`

GetRoomServerFqdnOk returns a tuple with the RoomServerFqdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoomServerFqdn

`func (o *Room) SetRoomServerFqdn(v string)`

SetRoomServerFqdn sets RoomServerFqdn field to given value.

### HasRoomServerFqdn

`func (o *Room) HasRoomServerFqdn() bool`

HasRoomServerFqdn returns a boolean if a field has been set.

### GetParticipantsMax

`func (o *Room) GetParticipantsMax() int32`

GetParticipantsMax returns the ParticipantsMax field if non-nil, zero value otherwise.

### GetParticipantsMaxOk

`func (o *Room) GetParticipantsMaxOk() (*int32, bool)`

GetParticipantsMaxOk returns a tuple with the ParticipantsMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParticipantsMax

`func (o *Room) SetParticipantsMax(v int32)`

SetParticipantsMax sets ParticipantsMax field to given value.

### HasParticipantsMax

`func (o *Room) HasParticipantsMax() bool`

HasParticipantsMax returns a boolean if a field has been set.

### GetParticipants

`func (o *Room) GetParticipants() []Participant`

GetParticipants returns the Participants field if non-nil, zero value otherwise.

### GetParticipantsOk

`func (o *Room) GetParticipantsOk() (*[]Participant, bool)`

GetParticipantsOk returns a tuple with the Participants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParticipants

`func (o *Room) SetParticipants(v []Participant)`

SetParticipants sets Participants field to given value.

### HasParticipants

`func (o *Room) HasParticipants() bool`

HasParticipants returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


