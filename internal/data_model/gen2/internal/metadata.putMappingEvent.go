// Copyright 2023 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
package internal

import (
	"github.com/vkcom/statshouse/internal/vkgo/basictl"
)

var _ = basictl.NatWrite

type MetadataPutMappingEvent struct {
	FieldsMask uint32
	Keys       []string
	Value      []int32
}

func (MetadataPutMappingEvent) TLName() string { return "metadata.putMappingEvent" }
func (MetadataPutMappingEvent) TLTag() uint32  { return 0x12345676 }

func (item *MetadataPutMappingEvent) Reset() {
	item.FieldsMask = 0
	item.Keys = item.Keys[:0]
	item.Value = item.Value[:0]
}

func (item *MetadataPutMappingEvent) Read(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatRead(w, &item.FieldsMask); err != nil {
		return w, err
	}
	if w, err = VectorString0Read(w, &item.Keys); err != nil {
		return w, err
	}
	return VectorInt0Read(w, &item.Value)
}

func (item *MetadataPutMappingEvent) Write(w []byte) (_ []byte, err error) {
	w = basictl.NatWrite(w, item.FieldsMask)
	if w, err = VectorString0Write(w, item.Keys); err != nil {
		return w, err
	}
	return VectorInt0Write(w, item.Value)
}

func (item *MetadataPutMappingEvent) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0x12345676); err != nil {
		return w, err
	}
	return item.Read(w)
}

func (item *MetadataPutMappingEvent) WriteBoxed(w []byte) ([]byte, error) {
	w = basictl.NatWrite(w, 0x12345676)
	return item.Write(w)
}

func (item MetadataPutMappingEvent) String() string {
	w, err := item.WriteJSON(nil)
	if err != nil {
		return err.Error()
	}
	return string(w)
}

func MetadataPutMappingEvent__ReadJSON(item *MetadataPutMappingEvent, j interface{}) error {
	return item.readJSON(j)
}
func (item *MetadataPutMappingEvent) readJSON(j interface{}) error {
	_jm, _ok := j.(map[string]interface{})
	if j != nil && !_ok {
		return ErrorInvalidJSON("metadata.putMappingEvent", "expected json object")
	}
	_jFieldsMask := _jm["fields_mask"]
	delete(_jm, "fields_mask")
	if err := JsonReadUint32(_jFieldsMask, &item.FieldsMask); err != nil {
		return err
	}
	_jKeys := _jm["keys"]
	delete(_jm, "keys")
	_jValue := _jm["value"]
	delete(_jm, "value")
	for k := range _jm {
		return ErrorInvalidJSONExcessElement("metadata.putMappingEvent", k)
	}
	if err := VectorString0ReadJSON(_jKeys, &item.Keys); err != nil {
		return err
	}
	if err := VectorInt0ReadJSON(_jValue, &item.Value); err != nil {
		return err
	}
	return nil
}

func (item *MetadataPutMappingEvent) WriteJSON(w []byte) (_ []byte, err error) {
	w = append(w, '{')
	if item.FieldsMask != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"fields_mask":`...)
		w = basictl.JSONWriteUint32(w, item.FieldsMask)
	}
	if len(item.Keys) != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"keys":`...)
		if w, err = VectorString0WriteJSON(w, item.Keys); err != nil {
			return w, err
		}
	}
	if len(item.Value) != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"value":`...)
		if w, err = VectorInt0WriteJSON(w, item.Value); err != nil {
			return w, err
		}
	}
	return append(w, '}'), nil
}

func (item *MetadataPutMappingEvent) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil)
}

func (item *MetadataPutMappingEvent) UnmarshalJSON(b []byte) error {
	j, err := JsonBytesToInterface(b)
	if err != nil {
		return ErrorInvalidJSON("metadata.putMappingEvent", err.Error())
	}
	if err = item.readJSON(j); err != nil {
		return ErrorInvalidJSON("metadata.putMappingEvent", err.Error())
	}
	return nil
}
