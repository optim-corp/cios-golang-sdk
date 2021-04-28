package collection

import "testing"

func TestCollection(t *testing.T) {
	stringStream := StringStreamOf("0", "1", "2", "3", "4", "5")
	intStream := IntegerStreamOf(0, 1, 2, 3, 4, 5)
	int32Stream := Int32StreamOf(0, 1, 2, 3, 4, 5)
	int64Stream := Int64StreamOf(0, 1, 2, 3, 4, 5)
	float32Stream := Float32StreamOf(0, 1, 2, 3, 4, 5)
	float64Stream := Float64StreamOf(0, 1, 2, 3, 4, 5)
	if v := stringStream.Clone().Sum(); v != "012345" {
		t.Fatal("Unexpect Value stringStream Sum.", v)
	}
	if v := intStream.Clone().Sum(); v != 15 {
		t.Fatal("Unexpect Value intStream Sum.", v)
	}
	if v := intStream.Min(); v != 0 {
		t.Fatal("Unexpect Value intStream Min.", v)
	}
	if v := intStream.Max(); v != 5 {
		t.Fatal("Unexpect Value intStream Max.", v)
	}
	if v := intStream.Average(); v != 2 {
		t.Fatal("Unexpect Value intStream Average.", v)
	}
	if v := int32Stream.Clone().Sum(); v != 15 {
		t.Fatal("Unexpect Value int32Stream Sum.", v)
	}
	if v := int32Stream.Min(); v != 0 {
		t.Fatal("Unexpect Value int32Stream Min.", v)
	}
	if v := int32Stream.Max(); v != 5 {
		t.Fatal("Unexpect Value int32Stream Max.", v)
	}
	if v := int32Stream.Average(); v != 2 {
		t.Fatal("Unexpect Value int32Stream Average.", v)
	}
	if v := int64Stream.Clone().Sum(); v != 15 {
		t.Fatal("Unexpect Value int64Stream Sum.", v)
	}
	if v := int64Stream.Min(); v != 0 {
		t.Fatal("Unexpect Value int64Stream Min.", v)
	}
	if v := int64Stream.Max(); v != 5 {
		t.Fatal("Unexpect Value int64Stream Max.", v)
	}
	if v := int64Stream.Average(); v != 2 {
		t.Fatal("Unexpect Value int64Stream Average.", v)
	}
	if v := float32Stream.Clone().Sum(); v != 15 {
		t.Fatal("Unexpect Value float32Stream Sum.", v)
	}
	if v := float32Stream.Min(); v != 0 {
		t.Fatal("Unexpect Value float32Stream Min.", v)
	}
	if v := float32Stream.Max(); v != 5 {
		t.Fatal("Unexpect Value float32Stream Max.", v)
	}
	if v := float32Stream.Average(); v != 2.5 {
		t.Fatal("Unexpect Value float32Stream Average.", v)
	}
	if v := float64Stream.Clone().Sum(); v != 15 {
		t.Fatal("Unexpect Value float64Stream Sum.", v)
	}
	stream := StringStreamOf("0") //[0]
	if stream.Get(0) == nil || *stream.Get(0) != "0" {
		t.Fatal("Unexpect Value StringStream StringStreamOf.", stream)
	}
	if stream.Add("1"); stream.Len() != 2 || *stream.Get(1) != "1" {
		t.Fatal("Unexpect Value StringStream Add.", stream)
	}
	if stream.AddAll("2", "3", "4", "5"); stream.Len() != 6 || *stream.Last() != "5" {
		t.Fatal("Unexpect Value StringStream AddAll.", stream)
	}
	if stream.Concat([]string{"4", "5"}); stream.Len() != 8 || *stream.Last() != "5" {
		t.Fatal("Unexpect Value StringStream Concat.", stream)
	}
	if stream.Set(0, "Hogehoge"); *stream.Get(0) != "Hogehoge" || *stream.Get(1) != "1" {
		t.Fatal("Unexpect Value StringStream Set.", stream)
	}
	if a := "a"; *stream.AddSafe(&a).AddSafe(nil).Last() != "a" || *stream.First() != "Hogehoge" {
		t.Fatal("Unexpect Value StringStream AddSafe.", stream)
	}

	stream = StringStreamOf("0", "1", "2", "3", "4", "5", "4", "5") //[0]
	stream.DeleteRange(1, 3)                                        //[0 4 5 4 5]
	stream.Delete(0)                                                //[4 5 4 5]
	stream.Distinct()                                               //[4 5]
	stream.Skip(1)                                                  // [5]
	stream.Limit(0)                                                 //[]

	if stream.IsPreset() {
		t.Fatal("Unexpect Value float64Stream Limit.", stream)
	}
	//stream = StringStreamOf("0", "1", "2", "3", "4", "5", "4", "5")          //[0]
	//result := stream.Get(-1)                                                 //result = nil
	//result = stream.Get(2)                                                   //result = *"2"
	//result = stream.First()                                                  //result = *"0"
	//result = stream.Last()                                                   //result = *"5"
	//result = stream.Find(func(arg string, _ int) bool { return arg == "2" }) //result = *"2"
	//
	//resultMatch := stream.AllMatch(func(_ string, _ int) bool { return true })       //resultMatch = true
	//resultMatch = stream.AnyMatch(func(arg string, _ int) bool { return arg == "" }) //resultMatch = false
	//resultMatch = stream.NoneMatch(func(_ string, _ int) bool { return false })      //resultMatch = true
	//resultMatch = stream.Contains("5")                                               //resultMatch = true
	//resultMatch = stream.Equals([]string{"0", "2", "3"})                             //resultMatch = false
	//resultMatch = stream.IsEmpty()                                                   //resultMatch = false
	//resultMatch = stream.IsPreset()                                                  //resultMatch = false
	//stream.FindIndex(func(arg string, _ int) bool { return arg == "3" })             //3
	//
	//stream.ForEach(func(arg string, index int) {})
	//stream.ForEachRight(func(arg string, index int) {})
	//stream.Peek(func(arg *string, index int) {})
	//stream.Filter(func(arg string, index int) bool { return index == 2 })
	//stream.GroupBy(func(arg string, index int) string { return "sample" })
	//stream.Reduce(func(result, current string, index int) string { return "" })
	//stream.ReduceInit(func(result, current string, index int) string { return "" }, "")
	//stream.Slice(0, 5)
	//stream.Sort(func(i, j int) bool { return true })
	//stream.SortStable(func(i, j int) bool { return true })
	//stream.Map(func(arg string, index int) string { return "" })
	//stream.While(func(arg string, index int) bool { return true })
	//stream.Reverse()
	//stream.ToList()
	//stream.Val() //same ToList()
	//stream.Len()
	//stream.IndexOf("1")
	//stream.Clone()
	//stream.Copy() //same Clone()

	stream = StringStreamOf("a", "b", "c")
	if !stream.ContainsAll("a", "b", "c") || stream.ContainsAll("a", "d") {
		t.Fatal("Unexpect Value ContainsAll.", stream)
	}
	if !stream.ContainsSome("a", "e", "d") || stream.ContainsSome("e", "d") {
		t.Fatal("Unexpect Value ContainsAll.", stream)
	}
}
