package wrapper

import "testing"

func TestExchange(t *testing.T) {
	str := MakeString("aaaaaaabbbbbcccccc")
	if str.ToUpperCase(); str.ChatAt(0).AsString() != "A" {
		t.Fatal("Unexpect Value  String ToUpperCase.", str.AsPrimitive())
	}
	if str.ToLowerCase(); str.ChatAt(0).AsString() != "a" {
		t.Fatal("Unexpect Value  String ToLowerCamel.", str.AsPrimitive())
	}
	if str.Reverse(); str.ChatAt(0).AsString() != "c" {
		t.Fatal("Unexpect Value String Reverse.", str.AsPrimitive())
	}
	if str.IsEmpty() {
		t.Fatal("Unexpect Value String IsEmpty.", str.AsPrimitive())
	}
	if !str.Contains("a") || str.ContainsSome("d", "z", "e") || !str.ContainsSome("d", "z", "a") || str.ContainsAll("d", "z", "a") || !str.ContainsAll("c", "b", "a") {
		t.Fatal("Unexpect Value String Contains.", str.AsPrimitive())
	}
	if str.DeleteAll("a"); str.Contains("a") {
		t.Fatal("Unexpect Value String Delete.", str.AsPrimitive())
	}
	if str = MakeString("abcd"); str.IsNumeric() {
		t.Fatal("Unexpect Value String IsNumeric.", str.AsPrimitive())
	}
	if str.Reverse(); !str.Equals("dcba") || str.AsPrimitive() != "dcba" {
		t.Fatal("Unexpect Value String Reverse.", str.AsPrimitive())
	}
	if str.IsBlank() || str.IsEmpty() {
		t.Fatal("Unexpect Value String IsBlank.", str.AsPrimitive())
	}
	if str.StartWith("ab") || !str.StartWith("dc") {
		t.Fatal("Unexpect Value String StartWith.", str.AsPrimitive())
	}
	if str.EndWith("cd") || !str.EndWith("ba") {
		t.Fatal("Unexpect Value String EndWith.", str.AsPrimitive())
	}
	if str.MathPattern(`\s+`) {
		t.Fatal("Unexpect Value String MathPattern.", str.AsPrimitive())
	}
	if str.Reverse().Add("  a"); !str.Equals("abcd  a") {
		t.Fatal("Unexpect Value String Add.", str.AsPrimitive())
	}
	if str.AddAll("b", "b", "z"); !str.Equals("abcd  abbz") {
		t.Fatal("Unexpect Value String AddAll.", str.AsPrimitive())
	}
	if str.AddHead("z"); !str.Equals("zabcd  abbz") {
		t.Fatal("Unexpect Value String AddHead.", str.AsPrimitive())
	}
	if str.AddAllHead("z", "y"); !str.Equals("yzzabcd  abbz") {
		t.Fatal("Unexpect Value String AddAllHead.", str.AsPrimitive())
	}
	if str.Slice(0, 0); str.AddAll("aab").AsRunes()[0] != []rune("aab")[0] || str.AsRunes()[1] != []rune("aab")[1] || str.AsRunes()[2] != []rune("aab")[2] {
		t.Fatal("Unexpect Value String AsRunes.", str.AsPrimitive())
	}
	if string(str.AsBytes()) != string([]byte("aab")) {
		t.Fatal("Unexpect Value String AsBytes.", str.AsPrimitive())
	}
	if str.Delete("a", 1); !str.Equals("ab") {
		t.Fatal("Unexpect Value String Delete.", str.AsPrimitive())
	}
	if str.AddAll("a", "a", "a").DeleteAll("a"); !str.Equals("b") {
		t.Fatal("Unexpect Value String DeleteAll.", str.AsPrimitive())

	}
	r, _ := EncodeJson(map[string]interface{}{
		"aaa": "test",
		"bbb": "test",
	})
	if !r.AsString().Equals("{\"aaa\":\"test\",\"bbb\":\"test\"}") {
		t.Fatal("Unexpect Value Resource AsString.", str.AsPrimitive())
	}
	temp := struct {
		Aaa string
		Bbb string
	}{}
	if _ = r.DecodeJson(&temp); temp.Aaa != "test" {
		t.Fatal("Unexpect Value Resource DecodeJson.", str.AsPrimitive())
	}
	if _r, _ := r.Copy(); _r == &r {
		t.Fatal("Unexpect Value Resource Copy.", str.AsPrimitive())
	}
	if text, _ := r.IndentJson(); text != "{\n  \"aaa\": \"test\",\n  \"bbb\": \"test\"\n}" {
		t.Fatal("Unexpect Value Resource IndentJson.", str.AsPrimitive())
	}
}
