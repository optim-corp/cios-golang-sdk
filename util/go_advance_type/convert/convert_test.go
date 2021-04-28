package convert

import (
	"bytes"
	"errors"
	"math"
	"testing"

	"github.com/optim-corp/cios-golang-sdk/util/go_advance_type/check"

	"github.com/optim-corp/cios-golang-sdk/util/go_advance_type/ref"
)

const (
	DEFAULT_INT   = 99
	DEFAULT_FLOAT = 99.99
	DEFAULT_STR   = "default"
)

type (
	testCase struct {
		convertor interface{}
		inputs    []interface{}
		output    interface{}
	}
	args []interface{}
)

func TestAny2Str(t *testing.T) {
	var (
		temp            = 1
		testAny2StrCase = []struct {
			input  interface{}
			output interface{}
		}{
			{input: "int", output: "int"},
			{input: 100, output: "100"},
			{input: int64(100), output: "100"},
			{input: math.MaxInt64, output: "9223372036854775807"},
			{input: math.MinInt64, output: "-9223372036854775808"},
			{input: math.MaxInt32, output: "2147483647"},
			{input: math.MinInt32, output: "-2147483648"},
			{input: math.MaxInt16, output: "32767"},
			{input: math.MinInt16, output: "-32768"},
			{input: math.MaxInt8, output: "127"},
			{input: math.MinInt8, output: "-128"},
			{input: uint64(math.MaxUint64), output: "18446744073709551615"},
			{input: math.MaxUint32, output: "4294967295"},
			{input: math.MaxUint16, output: "65535"},
			{input: math.MaxUint8, output: "255"},
			{input: int32(100), output: "100"},
			{input: int16(100), output: "100"},
			{input: int8(100), output: "100"},
			{input: 0.1, output: "0.1"},
			{input: float32(0.1), output: "0.1"},
			{input: nil, output: ""},
			{input: uint(11), output: "11"},
			{input: uint8(11), output: "11"},
			{input: uint16(11), output: "11"},
			{input: uint32(11), output: "11"},
			{input: uint64(11), output: "11"},
			{input: true, output: "true"},
			{input: false, output: "false"},
			{input: []byte("test"), output: "test"},
			{input: errors.New("test"), output: "test"},
			{input: &temp, output: "1"},
		}
	)
	for _, v := range testAny2StrCase {
		if str, err := Str(v.input); str != v.output {
			t.Errorf("Fail Str, in: %v, converted: %v, output: %v, err: %v", v.input, str, v.output, err)
		} else {
			t.Logf("Accept Str, in: %v, converted: %v,  output: %v, err: %v", v.input, str, v.output, err)
		}
	}

}

func TestAny2Int(t *testing.T) {
	var temp *bool = nil
	Bool(temp)
	var (
		test  = "1010"
		Cases = []testCase{
			{convertor: Int, inputs: args{"int"}, output: 0},
			{convertor: Int, inputs: args{"101"}, output: 101},
			{convertor: Int, inputs: args{"-101"}, output: -101},
			{convertor: Int, inputs: args{"9223372036854775807"}, output: 9223372036854775807},
			{convertor: Int, inputs: args{"-9223372036854775808"}, output: -9223372036854775808},
			{convertor: Int, inputs: args{"2147483647"}, output: 2147483647},
			{convertor: Int, inputs: args{"-2147483648"}, output: -2147483648},
			{convertor: Int, inputs: args{&test}, output: 1010},

			{convertor: Int32, inputs: args{"int"}, output: 0},
			{convertor: Int32, inputs: args{"101"}, output: 101},
			{convertor: Int32, inputs: args{"-101"}, output: -101},
			{convertor: Int32, inputs: args{"9223372036854775807"}, output: 0},
			{convertor: Int32, inputs: args{"-9223372036854775808"}, output: 0},
			{convertor: Int32, inputs: args{"2147483647"}, output: 2147483647},
			{convertor: Int32, inputs: args{"-2147483648"}, output: -2147483648},
			{convertor: Int32, inputs: args{&test}, output: 1010},
			{convertor: Int64, inputs: args{true}, output: 1},

			{convertor: Int64, inputs: args{"int"}, output: 0},
			{convertor: Int64, inputs: args{"101"}, output: 101},
			{convertor: Int64, inputs: args{"-101"}, output: -101},
			{convertor: Int64, inputs: args{"9223372036854775807"}, output: 9223372036854775807},
			{convertor: Int64, inputs: args{"-9223372036854775808"}, output: -9223372036854775808},
			{convertor: Int64, inputs: args{"2147483647"}, output: 2147483647},
			{convertor: Int64, inputs: args{"-2147483648"}, output: -2147483648},
			{convertor: Int64, inputs: args{&test}, output: 1010},
			{convertor: Int64, inputs: args{true}, output: 1},

			{convertor: Int, inputs: args{"int"}, output: 0},
			{convertor: Int, inputs: args{"101"}, output: 101},
			{convertor: Int, inputs: args{"-101"}, output: -101},
			{convertor: Int, inputs: args{"9223372036854775807"}, output: 9223372036854775807},
			{convertor: Int, inputs: args{"-9223372036854775808"}, output: -9223372036854775808},
			{convertor: Int, inputs: args{"2147483647"}, output: 2147483647},
			{convertor: Int, inputs: args{"-2147483648"}, output: -2147483648},
			{convertor: Int, inputs: args{&test}, output: 1010},
			{convertor: Int, inputs: args{true}, output: 1},

			{convertor: Float32, inputs: args{"int"}, output: 0},
			{convertor: Float32, inputs: args{"101"}, output: 101},
			{convertor: Float32, inputs: args{"-101"}, output: -101},
			{convertor: Float32, inputs: args{"-101"}, output: -101},
			{convertor: Float32, inputs: args{&test}, output: 1010},
			{convertor: Float32, inputs: args{true}, output: 1.0},

			{convertor: Float64, inputs: args{"int"}, output: 0},
			{convertor: Float64, inputs: args{"101"}, output: 101},
			{convertor: Float64, inputs: args{"-101"}, output: -101},
			{convertor: Float64, inputs: args{"2147483647"}, output: 2147483647},
			{convertor: Float64, inputs: args{"-2147483648"}, output: -2147483648},
			{convertor: Float64, inputs: args{MustStr(math.MaxFloat64)}, output: math.MaxFloat64},
			{convertor: Float64, inputs: args{&test}, output: 1010},
		}
	)
	for _, c := range Cases {
		if result := ref.ExecFunc(c.convertor, c.inputs...); MustStr(result[0]) != MustStr(c.output) {
			t.Errorf("Fail convert, in: %v, converted: %v, output: %v", c.inputs, MustStr(result[0]), MustStr(c.output))
		} else {
			t.Logf("Accept convert, in: %v, converted: %v, output: %v", c.inputs, MustStr(result[0]), MustStr(c.output))
		}
	}

}
func TestMapStruct(t *testing.T) {
	type (
		a struct {
			One   string
			Two   int32
			Three bool
			test  string
		}
		b struct {
			One   string
			Two   *int32 `json:"two"`
			Three bool   `json:"three22"`
			Four  float64
			test  string
		}
	)
	out := b{}
	if err := CopyCastFields(a{
		One:   "test",
		Two:   10,
		Three: true,
		test:  "test",
	}, &out); err != nil || out.One != "test" || *out.Two != 10 {
		t.Fatal(out)
	} else {
		t.Log(out)
	}
	out = b{}
	if err := CopyCastFields(a{
		One:   "test",
		Two:   10,
		Three: true,
		test:  "test",
	}, &out); err != nil || out.One != "test" || *out.Two != 10 {
		t.Fatal(out)
	} else {
		t.Log(out, *out.Two)
	}
	if mp := StructJsonTag2Map(out); mp["three22"] != true || MustInt(mp["two"]) != 10 {
		t.Fatal(mp)
	} else {
		t.Log(mp)
	}
}
func TestJson(t *testing.T) {
	type T struct {
		Test   string `json:"test"`
		Sample string `json:"sample"`
	}
	var target T
	if v, err := IndentJson(T{Test: "test", Sample: "sample"}); err != nil || v != "{\n  \"test\": \"test\",\n  \"sample\": \"sample\"\n}" {
		t.Error("Fail IndentJson ", v, err)
	} else {
		t.Log(v)
	}
	if v, err := CompactJson(T{Test: "test", Sample: "sample"}); err != nil || v != "{\"test\":\"test\",\"sample\":\"sample\"}" {
		t.Error("Fail CompactJson ", v, err)
	} else {
		t.Log(v)
	}
	b, _ := MarshalJson(T{Test: "test", Sample: "sample"})
	if err := UnMarshalJson(bytes.NewReader(b), &target); err != nil || MustCompactJson(target) != "{\"test\":\"test\",\"sample\":\"sample\"}" {
		t.Error("Fail UnMarshalJson ", target, err)
	} else {
		t.Log(MustCompactJson(target))
	}
	if err := UnMarshalJson(T{Test: "test", Sample: "sample"}, &target); err != nil || MustCompactJson(target) != "{\"test\":\"test\",\"sample\":\"sample\"}" {
		t.Error("Fail UnMarshalJson ", target, err)
	} else {
		t.Log(MustCompactJson(target))
	}
	if j := MustMarshalJson(target); MustCompactJson(j) != "{\"test\":\"test\",\"sample\":\"sample\"}" {
		t.Error("Fail MustMarshalJson ", target)
	} else {
		t.Log(j, MustStr(j))
	}

	_t := T{Test: "test"}
	if _tmap, _ := Struct2JsonMap(_t); !check.AreEqualJson(_t, _tmap) {
		t.Logf("%#v", _t)
		t.Error("Fail CompareJson ", _t, _tmap)
	}
	if testMap := map[string]interface{}{"test": "test", "sample": ""}; !check.AreEqualJson(_t, testMap) {
		t.Error("Fail CompareJson ", _t, testMap)
	}
	if !check.AreEqualJson(_t, MustJson(_t)) {
		t.Error("Fail CompareJson ", _t, MustJson(_t))
	}
}
func TestExchange(t *testing.T) {
	if v := SafeInt(interface{}(100), 10); v != 100 {
		t.Fatal("Unexpect Value func Any2Int.", v)
	}
	if v := SafeInt32(interface{}(int32(100)), 10); v != 100 {
		t.Fatal("Unexpect Value func Any2Int32.", v)
	}
	if v := SafeInt32("100", 10); v != 100 {
		t.Fatal("Unexpect Value func Any2Int32.", v)
	}
	if v := SafeInt64(interface{}(int64(100)), 10); v != 100 {
		t.Fatal("Unexpect Value func Any2Long.", v)
	}
	if v := SafeFloat32(interface{}(float32(100)), 10); v != 100 {
		t.Fatal("Unexpect Value func Any2Float.", v)
	}
	if v := SafeFloat64(interface{}(float64(100)), 10); v != 100 {
		t.Fatal("Unexpect Value func Any2Double.", v)
	}
}
