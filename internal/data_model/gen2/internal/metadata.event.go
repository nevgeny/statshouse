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

type MetadataEvent struct {
	FieldMask  uint32
	Id         int64
	Name       string
	EventType  int32
	Unused     uint32
	Version    int64
	UpdateTime uint32
	Data       string
}

func (MetadataEvent) TLName() string { return "metadata.event" }
func (MetadataEvent) TLTag() uint32  { return 0x9286affa }

func (item *MetadataEvent) Reset() {
	item.FieldMask = 0
	item.Id = 0
	item.Name = ""
	item.EventType = 0
	item.Unused = 0
	item.Version = 0
	item.UpdateTime = 0
	item.Data = ""
}

func (item *MetadataEvent) Read(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatRead(w, &item.FieldMask); err != nil {
		return w, err
	}
	if w, err = basictl.LongRead(w, &item.Id); err != nil {
		return w, err
	}
	if w, err = basictl.StringRead(w, &item.Name); err != nil {
		return w, err
	}
	if w, err = basictl.IntRead(w, &item.EventType); err != nil {
		return w, err
	}
	if w, err = basictl.NatRead(w, &item.Unused); err != nil {
		return w, err
	}
	if w, err = basictl.LongRead(w, &item.Version); err != nil {
		return w, err
	}
	if w, err = basictl.NatRead(w, &item.UpdateTime); err != nil {
		return w, err
	}
	return basictl.StringRead(w, &item.Data)
}

func (item *MetadataEvent) Write(w []byte) (_ []byte, err error) {
	w = basictl.NatWrite(w, item.FieldMask)
	w = basictl.LongWrite(w, item.Id)
	if w, err = basictl.StringWrite(w, item.Name); err != nil {
		return w, err
	}
	w = basictl.IntWrite(w, item.EventType)
	w = basictl.NatWrite(w, item.Unused)
	w = basictl.LongWrite(w, item.Version)
	w = basictl.NatWrite(w, item.UpdateTime)
	return basictl.StringWrite(w, item.Data)
}

func (item *MetadataEvent) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0x9286affa); err != nil {
		return w, err
	}
	return item.Read(w)
}

func (item *MetadataEvent) WriteBoxed(w []byte) ([]byte, error) {
	w = basictl.NatWrite(w, 0x9286affa)
	return item.Write(w)
}

func (item MetadataEvent) String() string {
	w, err := item.WriteJSON(nil)
	if err != nil {
		return err.Error()
	}
	return string(w)
}

func MetadataEvent__ReadJSON(item *MetadataEvent, j interface{}) error { return item.readJSON(j) }
func (item *MetadataEvent) readJSON(j interface{}) error {
	_jm, _ok := j.(map[string]interface{})
	if j != nil && !_ok {
		return ErrorInvalidJSON("metadata.event", "expected json object")
	}
	_jFieldMask := _jm["field_mask"]
	delete(_jm, "field_mask")
	if err := JsonReadUint32(_jFieldMask, &item.FieldMask); err != nil {
		return err
	}
	_jId := _jm["id"]
	delete(_jm, "id")
	if err := JsonReadInt64(_jId, &item.Id); err != nil {
		return err
	}
	_jName := _jm["name"]
	delete(_jm, "name")
	if err := JsonReadString(_jName, &item.Name); err != nil {
		return err
	}
	_jEventType := _jm["event_type"]
	delete(_jm, "event_type")
	if err := JsonReadInt32(_jEventType, &item.EventType); err != nil {
		return err
	}
	_jUnused := _jm["unused"]
	delete(_jm, "unused")
	if err := JsonReadUint32(_jUnused, &item.Unused); err != nil {
		return err
	}
	_jVersion := _jm["version"]
	delete(_jm, "version")
	if err := JsonReadInt64(_jVersion, &item.Version); err != nil {
		return err
	}
	_jUpdateTime := _jm["update_time"]
	delete(_jm, "update_time")
	if err := JsonReadUint32(_jUpdateTime, &item.UpdateTime); err != nil {
		return err
	}
	_jData := _jm["data"]
	delete(_jm, "data")
	if err := JsonReadString(_jData, &item.Data); err != nil {
		return err
	}
	for k := range _jm {
		return ErrorInvalidJSONExcessElement("metadata.event", k)
	}
	return nil
}

func (item *MetadataEvent) WriteJSON(w []byte) (_ []byte, err error) {
	w = append(w, '{')
	if item.FieldMask != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"field_mask":`...)
		w = basictl.JSONWriteUint32(w, item.FieldMask)
	}
	if item.Id != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"id":`...)
		w = basictl.JSONWriteInt64(w, item.Id)
	}
	if len(item.Name) != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"name":`...)
		w = basictl.JSONWriteString(w, item.Name)
	}
	if item.EventType != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"event_type":`...)
		w = basictl.JSONWriteInt32(w, item.EventType)
	}
	if item.Unused != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"unused":`...)
		w = basictl.JSONWriteUint32(w, item.Unused)
	}
	if item.Version != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"version":`...)
		w = basictl.JSONWriteInt64(w, item.Version)
	}
	if item.UpdateTime != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"update_time":`...)
		w = basictl.JSONWriteUint32(w, item.UpdateTime)
	}
	if len(item.Data) != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"data":`...)
		w = basictl.JSONWriteString(w, item.Data)
	}
	return append(w, '}'), nil
}

func (item *MetadataEvent) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil)
}

func (item *MetadataEvent) UnmarshalJSON(b []byte) error {
	j, err := JsonBytesToInterface(b)
	if err != nil {
		return ErrorInvalidJSON("metadata.event", err.Error())
	}
	if err = item.readJSON(j); err != nil {
		return ErrorInvalidJSON("metadata.event", err.Error())
	}
	return nil
}

type MetadataEventBytes struct {
	FieldMask  uint32
	Id         int64
	Name       []byte
	EventType  int32
	Unused     uint32
	Version    int64
	UpdateTime uint32
	Data       []byte
}

func (MetadataEventBytes) TLName() string { return "metadata.event" }
func (MetadataEventBytes) TLTag() uint32  { return 0x9286affa }

func (item *MetadataEventBytes) Reset() {
	item.FieldMask = 0
	item.Id = 0
	item.Name = item.Name[:0]
	item.EventType = 0
	item.Unused = 0
	item.Version = 0
	item.UpdateTime = 0
	item.Data = item.Data[:0]
}

func (item *MetadataEventBytes) Read(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatRead(w, &item.FieldMask); err != nil {
		return w, err
	}
	if w, err = basictl.LongRead(w, &item.Id); err != nil {
		return w, err
	}
	if w, err = basictl.StringReadBytes(w, &item.Name); err != nil {
		return w, err
	}
	if w, err = basictl.IntRead(w, &item.EventType); err != nil {
		return w, err
	}
	if w, err = basictl.NatRead(w, &item.Unused); err != nil {
		return w, err
	}
	if w, err = basictl.LongRead(w, &item.Version); err != nil {
		return w, err
	}
	if w, err = basictl.NatRead(w, &item.UpdateTime); err != nil {
		return w, err
	}
	return basictl.StringReadBytes(w, &item.Data)
}

func (item *MetadataEventBytes) Write(w []byte) (_ []byte, err error) {
	w = basictl.NatWrite(w, item.FieldMask)
	w = basictl.LongWrite(w, item.Id)
	if w, err = basictl.StringWriteBytes(w, item.Name); err != nil {
		return w, err
	}
	w = basictl.IntWrite(w, item.EventType)
	w = basictl.NatWrite(w, item.Unused)
	w = basictl.LongWrite(w, item.Version)
	w = basictl.NatWrite(w, item.UpdateTime)
	return basictl.StringWriteBytes(w, item.Data)
}

func (item *MetadataEventBytes) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0x9286affa); err != nil {
		return w, err
	}
	return item.Read(w)
}

func (item *MetadataEventBytes) WriteBoxed(w []byte) ([]byte, error) {
	w = basictl.NatWrite(w, 0x9286affa)
	return item.Write(w)
}

func (item MetadataEventBytes) String() string {
	w, err := item.WriteJSON(nil)
	if err != nil {
		return err.Error()
	}
	return string(w)
}

func MetadataEventBytes__ReadJSON(item *MetadataEventBytes, j interface{}) error {
	return item.readJSON(j)
}
func (item *MetadataEventBytes) readJSON(j interface{}) error {
	_jm, _ok := j.(map[string]interface{})
	if j != nil && !_ok {
		return ErrorInvalidJSON("metadata.event", "expected json object")
	}
	_jFieldMask := _jm["field_mask"]
	delete(_jm, "field_mask")
	if err := JsonReadUint32(_jFieldMask, &item.FieldMask); err != nil {
		return err
	}
	_jId := _jm["id"]
	delete(_jm, "id")
	if err := JsonReadInt64(_jId, &item.Id); err != nil {
		return err
	}
	_jName := _jm["name"]
	delete(_jm, "name")
	if err := JsonReadStringBytes(_jName, &item.Name); err != nil {
		return err
	}
	_jEventType := _jm["event_type"]
	delete(_jm, "event_type")
	if err := JsonReadInt32(_jEventType, &item.EventType); err != nil {
		return err
	}
	_jUnused := _jm["unused"]
	delete(_jm, "unused")
	if err := JsonReadUint32(_jUnused, &item.Unused); err != nil {
		return err
	}
	_jVersion := _jm["version"]
	delete(_jm, "version")
	if err := JsonReadInt64(_jVersion, &item.Version); err != nil {
		return err
	}
	_jUpdateTime := _jm["update_time"]
	delete(_jm, "update_time")
	if err := JsonReadUint32(_jUpdateTime, &item.UpdateTime); err != nil {
		return err
	}
	_jData := _jm["data"]
	delete(_jm, "data")
	if err := JsonReadStringBytes(_jData, &item.Data); err != nil {
		return err
	}
	for k := range _jm {
		return ErrorInvalidJSONExcessElement("metadata.event", k)
	}
	return nil
}

func (item *MetadataEventBytes) WriteJSON(w []byte) (_ []byte, err error) {
	w = append(w, '{')
	if item.FieldMask != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"field_mask":`...)
		w = basictl.JSONWriteUint32(w, item.FieldMask)
	}
	if item.Id != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"id":`...)
		w = basictl.JSONWriteInt64(w, item.Id)
	}
	if len(item.Name) != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"name":`...)
		w = basictl.JSONWriteStringBytes(w, item.Name)
	}
	if item.EventType != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"event_type":`...)
		w = basictl.JSONWriteInt32(w, item.EventType)
	}
	if item.Unused != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"unused":`...)
		w = basictl.JSONWriteUint32(w, item.Unused)
	}
	if item.Version != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"version":`...)
		w = basictl.JSONWriteInt64(w, item.Version)
	}
	if item.UpdateTime != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"update_time":`...)
		w = basictl.JSONWriteUint32(w, item.UpdateTime)
	}
	if len(item.Data) != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"data":`...)
		w = basictl.JSONWriteStringBytes(w, item.Data)
	}
	return append(w, '}'), nil
}

func (item *MetadataEventBytes) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil)
}

func (item *MetadataEventBytes) UnmarshalJSON(b []byte) error {
	j, err := JsonBytesToInterface(b)
	if err != nil {
		return ErrorInvalidJSON("metadata.event", err.Error())
	}
	if err = item.readJSON(j); err != nil {
		return ErrorInvalidJSON("metadata.event", err.Error())
	}
	return nil
}

func VectorMetadataEvent0Read(w []byte, vec *[]MetadataEvent) (_ []byte, err error) {
	var l uint32
	if w, err = basictl.NatRead(w, &l); err != nil {
		return w, err
	}
	if err = basictl.CheckLengthSanity(w, l, 4); err != nil {
		return w, err
	}
	if uint32(cap(*vec)) < l {
		*vec = make([]MetadataEvent, l)
	} else {
		*vec = (*vec)[:l]
	}
	for i := range *vec {
		if w, err = (*vec)[i].Read(w); err != nil {
			return w, err
		}
	}
	return w, nil
}

func VectorMetadataEvent0Write(w []byte, vec []MetadataEvent) (_ []byte, err error) {
	w = basictl.NatWrite(w, uint32(len(vec)))
	for _, elem := range vec {
		if w, err = elem.Write(w); err != nil {
			return w, err
		}
	}
	return w, nil
}

func VectorMetadataEvent0ReadJSON(j interface{}, vec *[]MetadataEvent) error {
	l, _arr, err := JsonReadArray("[]MetadataEvent", j)
	if err != nil {
		return err
	}
	if cap(*vec) < l {
		*vec = make([]MetadataEvent, l)
	} else {
		*vec = (*vec)[:l]
	}
	for i := range *vec {
		if err := MetadataEvent__ReadJSON(&(*vec)[i], _arr[i]); err != nil {
			return err
		}
	}
	return nil
}

func VectorMetadataEvent0WriteJSON(w []byte, vec []MetadataEvent) (_ []byte, err error) {
	w = append(w, '[')
	for _, elem := range vec {
		w = basictl.JSONAddCommaIfNeeded(w)
		if w, err = elem.WriteJSON(w); err != nil {
			return w, err
		}
	}
	return append(w, ']'), nil
}

func VectorMetadataEvent0BytesRead(w []byte, vec *[]MetadataEventBytes) (_ []byte, err error) {
	var l uint32
	if w, err = basictl.NatRead(w, &l); err != nil {
		return w, err
	}
	if err = basictl.CheckLengthSanity(w, l, 4); err != nil {
		return w, err
	}
	if uint32(cap(*vec)) < l {
		*vec = make([]MetadataEventBytes, l)
	} else {
		*vec = (*vec)[:l]
	}
	for i := range *vec {
		if w, err = (*vec)[i].Read(w); err != nil {
			return w, err
		}
	}
	return w, nil
}

func VectorMetadataEvent0BytesWrite(w []byte, vec []MetadataEventBytes) (_ []byte, err error) {
	w = basictl.NatWrite(w, uint32(len(vec)))
	for _, elem := range vec {
		if w, err = elem.Write(w); err != nil {
			return w, err
		}
	}
	return w, nil
}

func VectorMetadataEvent0BytesReadJSON(j interface{}, vec *[]MetadataEventBytes) error {
	l, _arr, err := JsonReadArray("[]MetadataEventBytes", j)
	if err != nil {
		return err
	}
	if cap(*vec) < l {
		*vec = make([]MetadataEventBytes, l)
	} else {
		*vec = (*vec)[:l]
	}
	for i := range *vec {
		if err := MetadataEventBytes__ReadJSON(&(*vec)[i], _arr[i]); err != nil {
			return err
		}
	}
	return nil
}

func VectorMetadataEvent0BytesWriteJSON(w []byte, vec []MetadataEventBytes) (_ []byte, err error) {
	w = append(w, '[')
	for _, elem := range vec {
		w = basictl.JSONAddCommaIfNeeded(w)
		if w, err = elem.WriteJSON(w); err != nil {
			return w, err
		}
	}
	return append(w, ']'), nil
}
