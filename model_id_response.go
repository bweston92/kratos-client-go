/*
 * Ory Kratos API
 *
 * Documentation for all public and administrative Ory Kratos APIs. Public and administrative APIs are exposed on different ports. Public APIs can face the public internet without any protection while administrative APIs should never be exposed without prior authorization. To protect the administative API port you should use something like Nginx, Ory Oathkeeper, or any other technology capable of authorizing incoming requests. 
 *
 * API version: v0.6.0-alpha.16
 * Contact: hi@ory.sh
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// IdResponse IDResponse Response to an API call that returns just an Id
type IdResponse struct {
	// The id of the newly created object.
	Id string `json:"Id"`
}

// NewIdResponse instantiates a new IdResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIdResponse(id string) *IdResponse {
	this := IdResponse{}
	this.Id = id
	return &this
}

// NewIdResponseWithDefaults instantiates a new IdResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIdResponseWithDefaults() *IdResponse {
	this := IdResponse{}
	return &this
}

// GetId returns the Id field value
func (o *IdResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *IdResponse) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *IdResponse) SetId(v string) {
	o.Id = v
}

func (o IdResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["Id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableIdResponse struct {
	value *IdResponse
	isSet bool
}

func (v NullableIdResponse) Get() *IdResponse {
	return v.value
}

func (v *NullableIdResponse) Set(val *IdResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableIdResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableIdResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdResponse(val *IdResponse) *NullableIdResponse {
	return &NullableIdResponse{value: val, isSet: true}
}

func (v NullableIdResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


