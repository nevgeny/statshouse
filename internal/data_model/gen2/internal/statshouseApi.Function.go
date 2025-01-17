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

func StatshouseApiFnAvg() StatshouseApiFunction { return StatshouseApiFunction__MakeEnum(5) }

func StatshouseApiFnCount() StatshouseApiFunction { return StatshouseApiFunction__MakeEnum(0) }

func StatshouseApiFnCountNorm() StatshouseApiFunction { return StatshouseApiFunction__MakeEnum(1) }

func StatshouseApiFnCumulAvg() StatshouseApiFunction { return StatshouseApiFunction__MakeEnum(6) }

func StatshouseApiFnCumulCount() StatshouseApiFunction { return StatshouseApiFunction__MakeEnum(2) }

func StatshouseApiFnCumulSum() StatshouseApiFunction { return StatshouseApiFunction__MakeEnum(9) }

func StatshouseApiFnDerivativeAvg() StatshouseApiFunction { return StatshouseApiFunction__MakeEnum(24) }

func StatshouseApiFnDerivativeCount() StatshouseApiFunction {
	return StatshouseApiFunction__MakeEnum(25)
}

func StatshouseApiFnDerivativeCountNorm() StatshouseApiFunction {
	return StatshouseApiFunction__MakeEnum(26)
}

func StatshouseApiFnDerivativeMax() StatshouseApiFunction { return StatshouseApiFunction__MakeEnum(23) }

func StatshouseApiFnDerivativeMin() StatshouseApiFunction { return StatshouseApiFunction__MakeEnum(22) }

func StatshouseApiFnDerivativeSum() StatshouseApiFunction { return StatshouseApiFunction__MakeEnum(27) }

func StatshouseApiFnDerivativeSumNorm() StatshouseApiFunction {
	return StatshouseApiFunction__MakeEnum(28)
}

func StatshouseApiFnDerivativeUnique() StatshouseApiFunction {
	return StatshouseApiFunction__MakeEnum(29)
}

func StatshouseApiFnDerivativeUniqueNorm() StatshouseApiFunction {
	return StatshouseApiFunction__MakeEnum(30)
}

func StatshouseApiFnMax() StatshouseApiFunction { return StatshouseApiFunction__MakeEnum(4) }

func StatshouseApiFnMaxCountHost() StatshouseApiFunction { return StatshouseApiFunction__MakeEnum(21) }

func StatshouseApiFnMaxHost() StatshouseApiFunction { return StatshouseApiFunction__MakeEnum(20) }

func StatshouseApiFnMin() StatshouseApiFunction { return StatshouseApiFunction__MakeEnum(3) }

func StatshouseApiFnP25() StatshouseApiFunction { return StatshouseApiFunction__MakeEnum(11) }

func StatshouseApiFnP50() StatshouseApiFunction { return StatshouseApiFunction__MakeEnum(12) }

func StatshouseApiFnP75() StatshouseApiFunction { return StatshouseApiFunction__MakeEnum(13) }

func StatshouseApiFnP90() StatshouseApiFunction { return StatshouseApiFunction__MakeEnum(14) }

func StatshouseApiFnP95() StatshouseApiFunction { return StatshouseApiFunction__MakeEnum(15) }

func StatshouseApiFnP99() StatshouseApiFunction { return StatshouseApiFunction__MakeEnum(16) }

func StatshouseApiFnP999() StatshouseApiFunction { return StatshouseApiFunction__MakeEnum(17) }

func StatshouseApiFnStddev() StatshouseApiFunction { return StatshouseApiFunction__MakeEnum(10) }

func StatshouseApiFnSum() StatshouseApiFunction { return StatshouseApiFunction__MakeEnum(7) }

func StatshouseApiFnSumNorm() StatshouseApiFunction { return StatshouseApiFunction__MakeEnum(8) }

func StatshouseApiFnUnique() StatshouseApiFunction { return StatshouseApiFunction__MakeEnum(18) }

func StatshouseApiFnUniqueNorm() StatshouseApiFunction { return StatshouseApiFunction__MakeEnum(19) }

var _StatshouseApiFunction = [31]UnionElement{
	{TLTag: 0x89689775, TLName: "statshouseApi.fnCount", TLString: "statshouseApi.fnCount#89689775"},
	{TLTag: 0x60e68b5c, TLName: "statshouseApi.fnCountNorm", TLString: "statshouseApi.fnCountNorm#60e68b5c"},
	{TLTag: 0x871201c4, TLName: "statshouseApi.fnCumulCount", TLString: "statshouseApi.fnCumulCount#871201c4"},
	{TLTag: 0xb4cb2644, TLName: "statshouseApi.fnMin", TLString: "statshouseApi.fnMin#b4cb2644"},
	{TLTag: 0xf90de384, TLName: "statshouseApi.fnMax", TLString: "statshouseApi.fnMax#f90de384"},
	{TLTag: 0x6323c2f6, TLName: "statshouseApi.fnAvg", TLString: "statshouseApi.fnAvg#6323c2f6"},
	{TLTag: 0xf4d9ad09, TLName: "statshouseApi.fnCumulAvg", TLString: "statshouseApi.fnCumulAvg#f4d9ad09"},
	{TLTag: 0x80ce3cf1, TLName: "statshouseApi.fnSum", TLString: "statshouseApi.fnSum#80ce3cf1"},
	{TLTag: 0x361963d5, TLName: "statshouseApi.fnSumNorm", TLString: "statshouseApi.fnSumNorm#361963d5"},
	{TLTag: 0x42fc39b6, TLName: "statshouseApi.fnCumulSum", TLString: "statshouseApi.fnCumulSum#42fc39b6"},
	{TLTag: 0x2043e480, TLName: "statshouseApi.fnStddev", TLString: "statshouseApi.fnStddev#2043e480"},
	{TLTag: 0xcf9ad7bf, TLName: "statshouseApi.fnP25", TLString: "statshouseApi.fnP25#cf9ad7bf"},
	{TLTag: 0x77c5de5c, TLName: "statshouseApi.fnP50", TLString: "statshouseApi.fnP50#77c5de5c"},
	{TLTag: 0xe674272, TLName: "statshouseApi.fnP75", TLString: "statshouseApi.fnP75#0e674272"},
	{TLTag: 0xd4c8c793, TLName: "statshouseApi.fnP90", TLString: "statshouseApi.fnP90#d4c8c793"},
	{TLTag: 0x9a92b76f, TLName: "statshouseApi.fnP95", TLString: "statshouseApi.fnP95#9a92b76f"},
	{TLTag: 0x71992e9a, TLName: "statshouseApi.fnP99", TLString: "statshouseApi.fnP99#71992e9a"},
	{TLTag: 0xa3434c26, TLName: "statshouseApi.fnP999", TLString: "statshouseApi.fnP999#a3434c26"},
	{TLTag: 0xf20fb854, TLName: "statshouseApi.fnUnique", TLString: "statshouseApi.fnUnique#f20fb854"},
	{TLTag: 0x9ceb6f68, TLName: "statshouseApi.fnUniqueNorm", TLString: "statshouseApi.fnUniqueNorm#9ceb6f68"},
	{TLTag: 0xb4790064, TLName: "statshouseApi.fnMaxHost", TLString: "statshouseApi.fnMaxHost#b4790064"},
	{TLTag: 0x885e665b, TLName: "statshouseApi.fnMaxCountHost", TLString: "statshouseApi.fnMaxCountHost#885e665b"},
	{TLTag: 0x4817df2b, TLName: "statshouseApi.fnDerivativeMin", TLString: "statshouseApi.fnDerivativeMin#4817df2b"},
	{TLTag: 0x43eeb810, TLName: "statshouseApi.fnDerivativeMax", TLString: "statshouseApi.fnDerivativeMax#43eeb810"},
	{TLTag: 0x60d2b603, TLName: "statshouseApi.fnDerivativeAvg", TLString: "statshouseApi.fnDerivativeAvg#60d2b603"},
	{TLTag: 0xe617771c, TLName: "statshouseApi.fnDerivativeCount", TLString: "statshouseApi.fnDerivativeCount#e617771c"},
	{TLTag: 0xbfb5f7fc, TLName: "statshouseApi.fnDerivativeCountNorm", TLString: "statshouseApi.fnDerivativeCountNorm#bfb5f7fc"},
	{TLTag: 0xa3a43781, TLName: "statshouseApi.fnDerivativeSum", TLString: "statshouseApi.fnDerivativeSum#a3a43781"},
	{TLTag: 0x96683390, TLName: "statshouseApi.fnDerivativeSumNorm", TLString: "statshouseApi.fnDerivativeSumNorm#96683390"},
	{TLTag: 0x5745a0a3, TLName: "statshouseApi.fnDerivativeUnique", TLString: "statshouseApi.fnDerivativeUnique#5745a0a3"},
	{TLTag: 0x4bd4f327, TLName: "statshouseApi.fnDerivativeUniqueNorm", TLString: "statshouseApi.fnDerivativeUniqueNorm#4bd4f327"},
}

// TODO - deconflict name
func StatshouseApiFunction__MakeEnum(i int) StatshouseApiFunction {
	return StatshouseApiFunction{index: i}
}

type StatshouseApiFunction struct {
	index int
}

func (item StatshouseApiFunction) TLName() string { return _StatshouseApiFunction[item.index].TLName }
func (item StatshouseApiFunction) TLTag() uint32  { return _StatshouseApiFunction[item.index].TLTag }

func (item *StatshouseApiFunction) Reset() { item.index = 0 }

func (item *StatshouseApiFunction) IsFnCount() bool { return item.index == 0 }
func (item *StatshouseApiFunction) SetFnCount()     { item.index = 0 }

func (item *StatshouseApiFunction) IsFnCountNorm() bool { return item.index == 1 }
func (item *StatshouseApiFunction) SetFnCountNorm()     { item.index = 1 }

func (item *StatshouseApiFunction) IsFnCumulCount() bool { return item.index == 2 }
func (item *StatshouseApiFunction) SetFnCumulCount()     { item.index = 2 }

func (item *StatshouseApiFunction) IsFnMin() bool { return item.index == 3 }
func (item *StatshouseApiFunction) SetFnMin()     { item.index = 3 }

func (item *StatshouseApiFunction) IsFnMax() bool { return item.index == 4 }
func (item *StatshouseApiFunction) SetFnMax()     { item.index = 4 }

func (item *StatshouseApiFunction) IsFnAvg() bool { return item.index == 5 }
func (item *StatshouseApiFunction) SetFnAvg()     { item.index = 5 }

func (item *StatshouseApiFunction) IsFnCumulAvg() bool { return item.index == 6 }
func (item *StatshouseApiFunction) SetFnCumulAvg()     { item.index = 6 }

func (item *StatshouseApiFunction) IsFnSum() bool { return item.index == 7 }
func (item *StatshouseApiFunction) SetFnSum()     { item.index = 7 }

func (item *StatshouseApiFunction) IsFnSumNorm() bool { return item.index == 8 }
func (item *StatshouseApiFunction) SetFnSumNorm()     { item.index = 8 }

func (item *StatshouseApiFunction) IsFnCumulSum() bool { return item.index == 9 }
func (item *StatshouseApiFunction) SetFnCumulSum()     { item.index = 9 }

func (item *StatshouseApiFunction) IsFnStddev() bool { return item.index == 10 }
func (item *StatshouseApiFunction) SetFnStddev()     { item.index = 10 }

func (item *StatshouseApiFunction) IsFnP25() bool { return item.index == 11 }
func (item *StatshouseApiFunction) SetFnP25()     { item.index = 11 }

func (item *StatshouseApiFunction) IsFnP50() bool { return item.index == 12 }
func (item *StatshouseApiFunction) SetFnP50()     { item.index = 12 }

func (item *StatshouseApiFunction) IsFnP75() bool { return item.index == 13 }
func (item *StatshouseApiFunction) SetFnP75()     { item.index = 13 }

func (item *StatshouseApiFunction) IsFnP90() bool { return item.index == 14 }
func (item *StatshouseApiFunction) SetFnP90()     { item.index = 14 }

func (item *StatshouseApiFunction) IsFnP95() bool { return item.index == 15 }
func (item *StatshouseApiFunction) SetFnP95()     { item.index = 15 }

func (item *StatshouseApiFunction) IsFnP99() bool { return item.index == 16 }
func (item *StatshouseApiFunction) SetFnP99()     { item.index = 16 }

func (item *StatshouseApiFunction) IsFnP999() bool { return item.index == 17 }
func (item *StatshouseApiFunction) SetFnP999()     { item.index = 17 }

func (item *StatshouseApiFunction) IsFnUnique() bool { return item.index == 18 }
func (item *StatshouseApiFunction) SetFnUnique()     { item.index = 18 }

func (item *StatshouseApiFunction) IsFnUniqueNorm() bool { return item.index == 19 }
func (item *StatshouseApiFunction) SetFnUniqueNorm()     { item.index = 19 }

func (item *StatshouseApiFunction) IsFnMaxHost() bool { return item.index == 20 }
func (item *StatshouseApiFunction) SetFnMaxHost()     { item.index = 20 }

func (item *StatshouseApiFunction) IsFnMaxCountHost() bool { return item.index == 21 }
func (item *StatshouseApiFunction) SetFnMaxCountHost()     { item.index = 21 }

func (item *StatshouseApiFunction) IsFnDerivativeMin() bool { return item.index == 22 }
func (item *StatshouseApiFunction) SetFnDerivativeMin()     { item.index = 22 }

func (item *StatshouseApiFunction) IsFnDerivativeMax() bool { return item.index == 23 }
func (item *StatshouseApiFunction) SetFnDerivativeMax()     { item.index = 23 }

func (item *StatshouseApiFunction) IsFnDerivativeAvg() bool { return item.index == 24 }
func (item *StatshouseApiFunction) SetFnDerivativeAvg()     { item.index = 24 }

func (item *StatshouseApiFunction) IsFnDerivativeCount() bool { return item.index == 25 }
func (item *StatshouseApiFunction) SetFnDerivativeCount()     { item.index = 25 }

func (item *StatshouseApiFunction) IsFnDerivativeCountNorm() bool { return item.index == 26 }
func (item *StatshouseApiFunction) SetFnDerivativeCountNorm()     { item.index = 26 }

func (item *StatshouseApiFunction) IsFnDerivativeSum() bool { return item.index == 27 }
func (item *StatshouseApiFunction) SetFnDerivativeSum()     { item.index = 27 }

func (item *StatshouseApiFunction) IsFnDerivativeSumNorm() bool { return item.index == 28 }
func (item *StatshouseApiFunction) SetFnDerivativeSumNorm()     { item.index = 28 }

func (item *StatshouseApiFunction) IsFnDerivativeUnique() bool { return item.index == 29 }
func (item *StatshouseApiFunction) SetFnDerivativeUnique()     { item.index = 29 }

func (item *StatshouseApiFunction) IsFnDerivativeUniqueNorm() bool { return item.index == 30 }
func (item *StatshouseApiFunction) SetFnDerivativeUniqueNorm()     { item.index = 30 }

func (item *StatshouseApiFunction) ReadBoxed(w []byte) (_ []byte, err error) {
	var tag uint32
	if w, err = basictl.NatRead(w, &tag); err != nil {
		return w, err
	}
	switch tag {
	case 0x89689775:
		item.index = 0
		return w, nil
	case 0x60e68b5c:
		item.index = 1
		return w, nil
	case 0x871201c4:
		item.index = 2
		return w, nil
	case 0xb4cb2644:
		item.index = 3
		return w, nil
	case 0xf90de384:
		item.index = 4
		return w, nil
	case 0x6323c2f6:
		item.index = 5
		return w, nil
	case 0xf4d9ad09:
		item.index = 6
		return w, nil
	case 0x80ce3cf1:
		item.index = 7
		return w, nil
	case 0x361963d5:
		item.index = 8
		return w, nil
	case 0x42fc39b6:
		item.index = 9
		return w, nil
	case 0x2043e480:
		item.index = 10
		return w, nil
	case 0xcf9ad7bf:
		item.index = 11
		return w, nil
	case 0x77c5de5c:
		item.index = 12
		return w, nil
	case 0xe674272:
		item.index = 13
		return w, nil
	case 0xd4c8c793:
		item.index = 14
		return w, nil
	case 0x9a92b76f:
		item.index = 15
		return w, nil
	case 0x71992e9a:
		item.index = 16
		return w, nil
	case 0xa3434c26:
		item.index = 17
		return w, nil
	case 0xf20fb854:
		item.index = 18
		return w, nil
	case 0x9ceb6f68:
		item.index = 19
		return w, nil
	case 0xb4790064:
		item.index = 20
		return w, nil
	case 0x885e665b:
		item.index = 21
		return w, nil
	case 0x4817df2b:
		item.index = 22
		return w, nil
	case 0x43eeb810:
		item.index = 23
		return w, nil
	case 0x60d2b603:
		item.index = 24
		return w, nil
	case 0xe617771c:
		item.index = 25
		return w, nil
	case 0xbfb5f7fc:
		item.index = 26
		return w, nil
	case 0xa3a43781:
		item.index = 27
		return w, nil
	case 0x96683390:
		item.index = 28
		return w, nil
	case 0x5745a0a3:
		item.index = 29
		return w, nil
	case 0x4bd4f327:
		item.index = 30
		return w, nil
	default:
		return w, ErrorInvalidUnionTag("statshouseApi.Function", tag)
	}
}

func (item StatshouseApiFunction) WriteBoxed(w []byte) (_ []byte, err error) {
	w = basictl.NatWrite(w, _StatshouseApiFunction[item.index].TLTag)
	return w, nil
}

func StatshouseApiFunction__ReadJSON(item *StatshouseApiFunction, j interface{}) error {
	return item.readJSON(j)
}
func (item *StatshouseApiFunction) readJSON(j interface{}) error {
	if j == nil {
		return ErrorInvalidJSON("statshouseApi.Function", "expected string")
	}
	_jtype, _ok := j.(string)
	if !_ok {
		return ErrorInvalidJSON("statshouseApi.Function", "expected string")
	}
	switch _jtype {
	case "statshouseApi.fnCount#89689775", "statshouseApi.fnCount", "#89689775":
		item.index = 0
		return nil
	case "statshouseApi.fnCountNorm#60e68b5c", "statshouseApi.fnCountNorm", "#60e68b5c":
		item.index = 1
		return nil
	case "statshouseApi.fnCumulCount#871201c4", "statshouseApi.fnCumulCount", "#871201c4":
		item.index = 2
		return nil
	case "statshouseApi.fnMin#b4cb2644", "statshouseApi.fnMin", "#b4cb2644":
		item.index = 3
		return nil
	case "statshouseApi.fnMax#f90de384", "statshouseApi.fnMax", "#f90de384":
		item.index = 4
		return nil
	case "statshouseApi.fnAvg#6323c2f6", "statshouseApi.fnAvg", "#6323c2f6":
		item.index = 5
		return nil
	case "statshouseApi.fnCumulAvg#f4d9ad09", "statshouseApi.fnCumulAvg", "#f4d9ad09":
		item.index = 6
		return nil
	case "statshouseApi.fnSum#80ce3cf1", "statshouseApi.fnSum", "#80ce3cf1":
		item.index = 7
		return nil
	case "statshouseApi.fnSumNorm#361963d5", "statshouseApi.fnSumNorm", "#361963d5":
		item.index = 8
		return nil
	case "statshouseApi.fnCumulSum#42fc39b6", "statshouseApi.fnCumulSum", "#42fc39b6":
		item.index = 9
		return nil
	case "statshouseApi.fnStddev#2043e480", "statshouseApi.fnStddev", "#2043e480":
		item.index = 10
		return nil
	case "statshouseApi.fnP25#cf9ad7bf", "statshouseApi.fnP25", "#cf9ad7bf":
		item.index = 11
		return nil
	case "statshouseApi.fnP50#77c5de5c", "statshouseApi.fnP50", "#77c5de5c":
		item.index = 12
		return nil
	case "statshouseApi.fnP75#0e674272", "statshouseApi.fnP75", "#0e674272":
		item.index = 13
		return nil
	case "statshouseApi.fnP90#d4c8c793", "statshouseApi.fnP90", "#d4c8c793":
		item.index = 14
		return nil
	case "statshouseApi.fnP95#9a92b76f", "statshouseApi.fnP95", "#9a92b76f":
		item.index = 15
		return nil
	case "statshouseApi.fnP99#71992e9a", "statshouseApi.fnP99", "#71992e9a":
		item.index = 16
		return nil
	case "statshouseApi.fnP999#a3434c26", "statshouseApi.fnP999", "#a3434c26":
		item.index = 17
		return nil
	case "statshouseApi.fnUnique#f20fb854", "statshouseApi.fnUnique", "#f20fb854":
		item.index = 18
		return nil
	case "statshouseApi.fnUniqueNorm#9ceb6f68", "statshouseApi.fnUniqueNorm", "#9ceb6f68":
		item.index = 19
		return nil
	case "statshouseApi.fnMaxHost#b4790064", "statshouseApi.fnMaxHost", "#b4790064":
		item.index = 20
		return nil
	case "statshouseApi.fnMaxCountHost#885e665b", "statshouseApi.fnMaxCountHost", "#885e665b":
		item.index = 21
		return nil
	case "statshouseApi.fnDerivativeMin#4817df2b", "statshouseApi.fnDerivativeMin", "#4817df2b":
		item.index = 22
		return nil
	case "statshouseApi.fnDerivativeMax#43eeb810", "statshouseApi.fnDerivativeMax", "#43eeb810":
		item.index = 23
		return nil
	case "statshouseApi.fnDerivativeAvg#60d2b603", "statshouseApi.fnDerivativeAvg", "#60d2b603":
		item.index = 24
		return nil
	case "statshouseApi.fnDerivativeCount#e617771c", "statshouseApi.fnDerivativeCount", "#e617771c":
		item.index = 25
		return nil
	case "statshouseApi.fnDerivativeCountNorm#bfb5f7fc", "statshouseApi.fnDerivativeCountNorm", "#bfb5f7fc":
		item.index = 26
		return nil
	case "statshouseApi.fnDerivativeSum#a3a43781", "statshouseApi.fnDerivativeSum", "#a3a43781":
		item.index = 27
		return nil
	case "statshouseApi.fnDerivativeSumNorm#96683390", "statshouseApi.fnDerivativeSumNorm", "#96683390":
		item.index = 28
		return nil
	case "statshouseApi.fnDerivativeUnique#5745a0a3", "statshouseApi.fnDerivativeUnique", "#5745a0a3":
		item.index = 29
		return nil
	case "statshouseApi.fnDerivativeUniqueNorm#4bd4f327", "statshouseApi.fnDerivativeUniqueNorm", "#4bd4f327":
		item.index = 30
		return nil
	default:
		return ErrorInvalidEnumTagJSON("statshouseApi.Function", _jtype)
	}
}

func (item StatshouseApiFunction) WriteJSON(w []byte) (_ []byte, err error) {
	w = append(w, '"')
	w = append(w, _StatshouseApiFunction[item.index].TLString...)
	return append(w, '"'), nil

}

func (item StatshouseApiFunction) String() string {
	w, err := item.WriteJSON(nil)
	if err != nil {
		return err.Error()
	}
	return string(w)
}

func VectorStatshouseApiFunctionBoxed0Read(w []byte, vec *[]StatshouseApiFunction) (_ []byte, err error) {
	var l uint32
	if w, err = basictl.NatRead(w, &l); err != nil {
		return w, err
	}
	if err = basictl.CheckLengthSanity(w, l, 4); err != nil {
		return w, err
	}
	if uint32(cap(*vec)) < l {
		*vec = make([]StatshouseApiFunction, l)
	} else {
		*vec = (*vec)[:l]
	}
	for i := range *vec {
		if w, err = (*vec)[i].ReadBoxed(w); err != nil {
			return w, err
		}
	}
	return w, nil
}

func VectorStatshouseApiFunctionBoxed0Write(w []byte, vec []StatshouseApiFunction) (_ []byte, err error) {
	w = basictl.NatWrite(w, uint32(len(vec)))
	for _, elem := range vec {
		if w, err = elem.WriteBoxed(w); err != nil {
			return w, err
		}
	}
	return w, nil
}

func VectorStatshouseApiFunctionBoxed0ReadJSON(j interface{}, vec *[]StatshouseApiFunction) error {
	l, _arr, err := JsonReadArray("[]StatshouseApiFunction", j)
	if err != nil {
		return err
	}
	if cap(*vec) < l {
		*vec = make([]StatshouseApiFunction, l)
	} else {
		*vec = (*vec)[:l]
	}
	for i := range *vec {
		if err := StatshouseApiFunction__ReadJSON(&(*vec)[i], _arr[i]); err != nil {
			return err
		}
	}
	return nil
}

func VectorStatshouseApiFunctionBoxed0WriteJSON(w []byte, vec []StatshouseApiFunction) (_ []byte, err error) {
	w = append(w, '[')
	for _, elem := range vec {
		w = basictl.JSONAddCommaIfNeeded(w)
		if w, err = elem.WriteJSON(w); err != nil {
			return w, err
		}
	}
	return append(w, ']'), nil
}
