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

type EngineMetafilesStat struct {
	Data []EngineMetafilesOneMemoryStat
}

func (EngineMetafilesStat) TLName() string { return "engine.metafilesStatData" }
func (EngineMetafilesStat) TLTag() uint32  { return 0xb673669b }

func (item *EngineMetafilesStat) Reset() {
	item.Data = item.Data[:0]
}

func (item *EngineMetafilesStat) Read(w []byte) (_ []byte, err error) {
	return VectorEngineMetafilesOneMemoryStat0Read(w, &item.Data)
}

func (item *EngineMetafilesStat) Write(w []byte) (_ []byte, err error) {
	return VectorEngineMetafilesOneMemoryStat0Write(w, item.Data)
}

func (item *EngineMetafilesStat) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0xb673669b); err != nil {
		return w, err
	}
	return item.Read(w)
}

func (item *EngineMetafilesStat) WriteBoxed(w []byte) ([]byte, error) {
	w = basictl.NatWrite(w, 0xb673669b)
	return item.Write(w)
}

func (item EngineMetafilesStat) String() string {
	w, err := item.WriteJSON(nil)
	if err != nil {
		return err.Error()
	}
	return string(w)
}

func EngineMetafilesStat__ReadJSON(item *EngineMetafilesStat, j interface{}) error {
	return item.readJSON(j)
}
func (item *EngineMetafilesStat) readJSON(j interface{}) error {
	_jm, _ok := j.(map[string]interface{})
	if j != nil && !_ok {
		return ErrorInvalidJSON("engine.metafilesStatData", "expected json object")
	}
	_jData := _jm["data"]
	delete(_jm, "data")
	for k := range _jm {
		return ErrorInvalidJSONExcessElement("engine.metafilesStatData", k)
	}
	if err := VectorEngineMetafilesOneMemoryStat0ReadJSON(_jData, &item.Data); err != nil {
		return err
	}
	return nil
}

func (item *EngineMetafilesStat) WriteJSON(w []byte) (_ []byte, err error) {
	w = append(w, '{')
	if len(item.Data) != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"data":`...)
		if w, err = VectorEngineMetafilesOneMemoryStat0WriteJSON(w, item.Data); err != nil {
			return w, err
		}
	}
	return append(w, '}'), nil
}

func (item *EngineMetafilesStat) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil)
}

func (item *EngineMetafilesStat) UnmarshalJSON(b []byte) error {
	j, err := JsonBytesToInterface(b)
	if err != nil {
		return ErrorInvalidJSON("engine.metafilesStatData", err.Error())
	}
	if err = item.readJSON(j); err != nil {
		return ErrorInvalidJSON("engine.metafilesStatData", err.Error())
	}
	return nil
}
