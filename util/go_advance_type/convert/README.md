# Convert Utils

## Cast

support

int, int8, int16, int32, int64, float32, float64, bool, string

```go 
package main
import (
    "github.com/optim-corp/cios-golang-sdk/util/go_advance_type/convert"
)

func main() {
	a := convert.MustStr(120)  // Same SafeStr(120, "")
	b := convert.MustInt("120") // Same SafeInt("120", 0)
	c := convert.MustInt64([]byte("120"))  // Same SafeInt64([]byte("120"), 0)
        // var a string = "120" 
        // var b int = 120      
        // var c int64 = 120    


	d := convert.SafeInt32(struct{}{}, 21)
	e := convert.SafeFloat32([]byte(""), 1.2)
        // var d int32  = 21
        // var e float32 = 1.2


	f, err := convert.Float64("hogehoge") 
	g, err := convert.Bool("gegege")      
        // error
        // error

}
```

## Json

Powered by [json](https://github.com/json-iterator/go)
```go
package main
import (
    "github.com/optim-corp/cios-golang-sdk/util/go_advance_type/convert"
)

func main() {
    type T1 struct {
        Test   string
        Sample string
    }
    t1 := T1{Test: "test", Sample: "sample"}
    
    jsonStr, err := IndentJson(t1) 
    //  jsonStr = "{\n  \"Test\": \"test\",\n  \"Sample\": \"sample\"\n}"
    
    jsonStr, err := CompactJson(t1)
    //  jsonStr = "{\"Test\":\"test\",\"Sample\":\"sample\"}"  
    
    err := UnMarshalJson(t1}, &hogehoge)
    err := MarshalJson(t1) 


}
```


## Convert

```go
package main
import (
    "github.com/optim-corp/cios-golang-sdk/util/go_advance_type/convert"
)

func main() {
    var b Hoge
    var c Foga
    a := Hoge {
        a: "a",
    }

    convert.Struct2Map(a) 
    // map[string]interface{}{"jsontag_a": "a"}
    
    convert.Map2Struct(map[string]interface{}{"a": "a"}, &b) 
    // map[string]interface{}{"a": "a"}
    
    convert.StructTag2Map(a, "json")
    //  Hoge {a: "a"}
    
    convert.StructJsonTag2Map(a) 
    //  map[string]interface{}{"jsontag_a": "a"}
   

    convert.GetObjectValues(a)
    convert.GetObjectKeys(a)
    
    convert.DeepCopy(a, &b)
    convert.CopyFields(a, &c)
    convert.CopyCastFields(a, &c)
}
```

## Exchange

```go
package main
import (
	"github.com/optim-corp/cios-golang-sdk/util/go_advance_type/convert"
)

func main() {
	ptrString := convert.StringPtr("test") // &"test"
	ptrInt := convert.IntPtr(2)            // &2
}
```

