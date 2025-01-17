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

type MetadataPutMapping struct {
	FieldMask uint32
	Keys      []string
	Value     []int32
}

func (MetadataPutMapping) TLName() string { return "metadata.putMapping" }
func (MetadataPutMapping) TLTag() uint32  { return 0x9faf5281 }

func (item *MetadataPutMapping) Reset() {
	item.FieldMask = 0
	item.Keys = item.Keys[:0]
	item.Value = item.Value[:0]
}

func (item *MetadataPutMapping) Read(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatRead(w, &item.FieldMask); err != nil {
		return w, err
	}
	if w, err = VectorString0Read(w, &item.Keys); err != nil {
		return w, err
	}
	return VectorInt0Read(w, &item.Value)
}

func (item *MetadataPutMapping) Write(w []byte) (_ []byte, err error) {
	w = basictl.NatWrite(w, item.FieldMask)
	if w, err = VectorString0Write(w, item.Keys); err != nil {
		return w, err
	}
	return VectorInt0Write(w, item.Value)
}

func (item *MetadataPutMapping) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0x9faf5281); err != nil {
		return w, err
	}
	return item.Read(w)
}

func (item *MetadataPutMapping) WriteBoxed(w []byte) ([]byte, error) {
	w = basictl.NatWrite(w, 0x9faf5281)
	return item.Write(w)
}

func (item *MetadataPutMapping) ReadResult(w []byte, ret *MetadataPutMappingResponse) (_ []byte, err error) {
	return ret.ReadBoxed(w)
}

func (item *MetadataPutMapping) WriteResult(w []byte, ret MetadataPutMappingResponse) (_ []byte, err error) {
	return ret.WriteBoxed(w)
}

func (item *MetadataPutMapping) ReadResultJSON(j interface{}, ret *MetadataPutMappingResponse) error {
	if err := MetadataPutMappingResponse__ReadJSON(ret, j); err != nil {
		return err
	}
	return nil
}

func (item *MetadataPutMapping) WriteResultJSON(w []byte, ret MetadataPutMappingResponse) (_ []byte, err error) {
	if w, err = ret.WriteJSON(w); err != nil {
		return w, err
	}
	return w, nil
}

func (item *MetadataPutMapping) ReadResultWriteResultJSON(r []byte, w []byte) (_ []byte, _ []byte, err error) {
	var ret MetadataPutMappingResponse
	if r, err = item.ReadResult(r, &ret); err != nil {
		return r, w, err
	}
	w, err = item.WriteResultJSON(w, ret)
	return r, w, err
}

func (item *MetadataPutMapping) ReadResultJSONWriteResult(r []byte, w []byte) ([]byte, []byte, error) {
	j, err := JsonBytesToInterface(r)
	if err != nil {
		return r, w, ErrorInvalidJSON("metadata.putMapping", err.Error())
	}
	var ret MetadataPutMappingResponse
	if err = item.ReadResultJSON(j, &ret); err != nil {
		return r, w, err
	}
	w, err = item.WriteResult(w, ret)
	return r, w, err
}

func (item MetadataPutMapping) String() string {
	w, err := item.WriteJSON(nil)
	if err != nil {
		return err.Error()
	}
	return string(w)
}

func MetadataPutMapping__ReadJSON(item *MetadataPutMapping, j interface{}) error {
	return item.readJSON(j)
}
func (item *MetadataPutMapping) readJSON(j interface{}) error {
	_jm, _ok := j.(map[string]interface{})
	if j != nil && !_ok {
		return ErrorInvalidJSON("metadata.putMapping", "expected json object")
	}
	_jFieldMask := _jm["field_mask"]
	delete(_jm, "field_mask")
	if err := JsonReadUint32(_jFieldMask, &item.FieldMask); err != nil {
		return err
	}
	_jKeys := _jm["keys"]
	delete(_jm, "keys")
	_jValue := _jm["value"]
	delete(_jm, "value")
	for k := range _jm {
		return ErrorInvalidJSONExcessElement("metadata.putMapping", k)
	}
	if err := VectorString0ReadJSON(_jKeys, &item.Keys); err != nil {
		return err
	}
	if err := VectorInt0ReadJSON(_jValue, &item.Value); err != nil {
		return err
	}
	return nil
}

func (item *MetadataPutMapping) WriteJSON(w []byte) (_ []byte, err error) {
	w = append(w, '{')
	if item.FieldMask != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"field_mask":`...)
		w = basictl.JSONWriteUint32(w, item.FieldMask)
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

func (item *MetadataPutMapping) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil)
}

func (item *MetadataPutMapping) UnmarshalJSON(b []byte) error {
	j, err := JsonBytesToInterface(b)
	if err != nil {
		return ErrorInvalidJSON("metadata.putMapping", err.Error())
	}
	if err = item.readJSON(j); err != nil {
		return ErrorInvalidJSON("metadata.putMapping", err.Error())
	}
	return nil
}
