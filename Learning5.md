### 51) Text Templates:

```go
package main

import (
	"os"
	"text/template"
)

// Go offers built-in support for creating dynamic content or showing customized output to the user with the text/template package.
// A sibling package named html/template provides the same API but has additional security features and should be used for generating HTML.

func main() {
	// We can create a new template and parse its body from a string. Templates are a mix of static text and “actions” enclosed in {{...}} that are used to dynamically insert content.
	t1 := template.New("t1")
    t1, err := t1.Parse("Value is {{.}}\n")
    if err != nil {
        panic(err)
    }

	// Alternatively, we can use the template.Must function to panic in case Parse returns an error. This is especially useful for templates initialized in the global scope.
    t1 = template.Must(t1.Parse("Value: {{.}}\n"))

	// By “executing” the template we generate its text with specific values for its actions. The {{.}} action is replaced by the value passed as a parameter to Execute.
    t1.Execute(os.Stdout, "some text")
	// Value: some text
    t1.Execute(os.Stdout, 5)
	// Value: 5
    t1.Execute(os.Stdout, []string{
		"Go",
        "Rust",
        "C++",
        "C#",
    })
	// Value: [Go Rust C++ C#]

	// Helper function we’ll use below.
    Create := func(name, t string) *template.Template {
        return template.Must(template.New(name).Parse(t))
    }

	// If the data is a struct we can use the {{.FieldName}} action to access its fields. The fields should be exported to be accessible when a template is executing.
    t2 := Create("t2", "Name: {{.Name}}\n")

	t2.Execute(os.Stdout, struct {
        Name string
    }{"Jane Doe"})
	// Name: Jane Doe

	// The same applies to maps; with maps there is no restriction on the case of key names.
    t2.Execute(os.Stdout, map[string]string{
		"Name": "Mickey Mouse",
    })
	// Name: Mickey Mouse

	// if/else provide conditional execution for templates. A value is considered false if it’s the default value of a type, such as 0, an empty string, nil pointer, etc.
	// This sample demonstrates another feature of templates: using - in actions to trim whitespace.
    t3 := Create("t3",
	"{{if . -}} yes {{else -}} no {{end}}\n")
    t3.Execute(os.Stdout, "not empty")
	// yes
    t3.Execute(os.Stdout, "")
	// no

	// range blocks let us loop through slices, arrays, maps or channels. Inside the range block {{.}} is set to the current item of the iteration.
    t4 := Create("t4",
	"Range: {{range .}}{{.}} {{end}}\n")
    t4.Execute(os.Stdout,
        []string{
			"Go",
            "Rust",
            "C++",
            "C#",
        },
	)
	// Range: Go Rust C++ C#
}
```

### 52) Regular Expressions:

```go
package main

import (
	"bytes"
	"fmt"
	"regexp"
)

// Go offers built-in support for regular expressions. Here are some examples of common regexp-related tasks in Go.

func main() {
	// This tests whether a pattern matches a string.
	match, _ := regexp.MatchString("p([a-z]+)ch", "peach")
	fmt.Println(match)
	// true

	// Above we used a string pattern directly, but for other regexp tasks you'll need to Compile an optimized Regexp struct.
	r, _ := regexp.Compile("p([a-z]+)ch")

	// Many methods are available on these structs. Here's a match text like we saw earlier.
	fmt.Println(r.MatchString("peach"))
	// true

	// This finds the match of the regexp
	fmt.Println(r.FindString("peach punch"))
	// peach

	// This also finds teh first match but returns the start and end indexes for the match instead of the matching text
	fmt.Println("idx:", r.FindStringIndex("peach punch"))
	// idx: [0 5]

	// The Submatch variants include information about both the whole-pattern matches and the submatches within those matches. For example this will return information for both p([a-z]+)ch and ([a-z]+).
	fmt.Println(r.FindStringSubmatch("peach punch"))
	// [peach ea]

	// Similarly this will return information about the indexes of matches and submatches.
	fmt.Println(r.FindStringSubmatchIndex("peach punch"))
	// [0 5 1 3]

	// The All variants of these functions apply to all matches in the input, not just the first. For example to find all matches for a regexp.
    fmt.Println(r.FindAllString("peach punch pinch", -1))
	// [peach punch pinch]

	// These All variants are available for the other functions we saw above as well.
	fmt.Println("all:", r.FindAllStringSubmatchIndex(
        "peach punch pinch", -1))
	// all: [[0 5 1 3] [6 11 7 9] [12 17 13 15]]

	// Providing a non-negative integer as the second argument to these functions will limit the number of matches.
    fmt.Println(r.FindAllString("peach punch pinch", 2))
	// [peach punch]

	// Our examples above had string arguments and used names like MatchString. We can also provide []byte arguments and drop String from the function name.
    fmt.Println(r.Match([]byte("peach")))
	// true

	// When creating global variables with regular expressions you can use the MustCompile variation of Compile. MustCompile panics instead of returning an error, which makes it safer to use for global variables.
    r = regexp.MustCompile("p([a-z]+)ch")
    fmt.Println("regexp:", r)
	// regexp: p([a-z]+)ch

	// The regexp package can also be used to replace subsets of strings with other values.
	fmt.Println(r.ReplaceAllString("a peach", "<fruit>"))
	// a <fruit>

	// The Func variant allows you to transform matched text with a given function.
    in := []byte("a peach")
    out := r.ReplaceAllFunc(in, bytes.ToUpper)
    fmt.Println(string(out))
	// a PEACH
}
```

### 53) JSON:

```go
package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// Go offers built-in support for JSON encoding and decoding, including to and from built-in and custom data types.

// We'll use these two structs to demonstrate encoding and decoding of custom types below.
type response1 struct {
	Page   int
	Fruits []string
}

// Only exported fields will be encoded/decoded in JSON. Fields must start with capital letters to be exported.
type response2 struct {
	Page   int      `json:"page"`
	Fruits []string `json:"fruits"`
}

func main() {
	// First we'll look at encoding basic data types to JSON strings. Here are some examples for atomic values.
	bolB, _ := json.Marshal(true)
	fmt.Println(string(bolB))
	// true

	intB, _ := json.Marshal(1)
	fmt.Println(string(intB))
	// 1
	fltB, _ := json.Marshal(2.34)
	fmt.Println(string(fltB))
	// 2.34
	strB, _ := json.Marshal("gopher")
	fmt.Println(string(strB))
	// "gopher"

	// And here are some for slices and maps, which encode to JSON arrays and objects as you'd expect.
	slcD := []string{"apple", "peach", "pear"}
	slcB, _ := json.Marshal(slcD)
	fmt.Println(string(slcB))
	// ["apple","peach","pear"]

	// The JSON package can automatically encode your custom data types. It will only include exported fields in the encoded output and will by default use those names as the JSON keys.
	res1D := &response1{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"}}
	res1B, _ := json.Marshal(res1D)
	fmt.Println(string(res1B))
	// {"Page":1,"Fruits":["apple","peach","pear"]}

	mapD := map[string]int{"apple": 5, "lettuce": 7}
	mapB, _ := json.Marshal(mapD)
	fmt.Println(string(mapB))
	// {"apple":5,"lettuce":7}

	// The JSON package can automatically encode your custom data types. It will only include exported fields in the encoded output and will by default use those names as the JSON keys.
	resp1D := &response1{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"}}
	resp1B, _ := json.Marshal(resp1D)
	fmt.Println(string(resp1B))
	// {"Page":1,"Fruits":["apple","peach","pear"]}

	// You can use tags on struct field declarations to customize the encoded JSON key names. Check the definition of response2 above to see an example of such tags.
	resp2D := &response2{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"}}
	resp2B, _ := json.Marshal(resp2D)
	fmt.Println(string(resp2B))
	// {"page":1,"fruits":["apple","peach","pear"]}

	// Now let’s look at decoding JSON data into Go values. Here’s an example for a generic data structure.
	byt := []byte(`{"num":6.13,"strs":["a","b"]}`)

	// We need to provide a variable where the JSON package can put the decoded data. This map[string]interface{} will hold a map of strings to arbitrary data types.
	var dat map[string]interface{}

	// Here’s the actual decoding, and a check for associated errors.
	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}
	fmt.Println(dat)
	// map[num:6.13 strs:[a b]]

	// In order to use the values in the decoded map, we’ll need to convert them to their appropriate type. For example here we convert the value in num to the expected float64 type.
	num := dat["num"].(float64)
	fmt.Println(num)
	// 6.13

	// Accessing nested data requires a series of conversions.
	strs := dat["strs"].([]interface{})
	str1 := strs[0].(string)
	fmt.Println(str1)
	// a

	// We can also decode JSON into custom data types. This has the advantages of adding additional type-safety to our programs and eliminating the need for type assertions when accessing the decoded data.
	str := `{"page": 1, "fruits": ["apple", "peach"]}`
	res := response2{}
	json.Unmarshal([]byte(str), &res)
	fmt.Println(res)
	// {1 [apple peach]}
	fmt.Println(res.Fruits[0])
	// apple

	// In the examples above we always used bytes and strings as intermediates between the data and JSON representation on standard out. We can also stream JSON encodings directly to os.Writers like os.Stdout or even HTTP response bodies.
	enc := json.NewEncoder(os.Stdout)
	d := map[string]int{"apple": 5, "lettuce": 7}
	enc.Encode(d)
	// {"apple":5,"lettuce":7}
}

```

### 54) XML:

```go
package main

import (
	"encoding/xml"
	"fmt"
)

// Go offers built-in support for XML and XML-lke formats with the encoding/xml package

// Plant will be mapped to XML. Similarly to the JSON examples, field tags contain directives for the encoder and decoder.
// Here we use some special features of the XML package: the XMLName field name dictates the name of the XML element representing this struct; id,attr means that the Id field is an XML attribute rather than a nested element.
type Plant struct {
	XMLName xml.Name `xml:"plant"`
	Id      int      `xml:"id,attr"`
	Name    string   `xml:"name"`
	Origin  []string `xml:"origin"`
}

func (p Plant) String() string {
	return fmt.Sprintf("Plant id=%v, name=%v, origin=%v", p.Id, p.Name, p.Origin)
}

func main() {
	coffee := &Plant{Id: 27, Name: "Coffee"}
	coffee.Origin = []string{"Ethiopia", "Brazil"}

	// Emit XML representing our plant; using MarshalIndent to produce a more human-readable output.
	out, _ := xml.MarshalIndent(coffee, " ", " ")
	fmt.Println(string(out))
	/*
	 <plant id="27">
	  <name>Coffee</name>
	  <origin>Ethiopia</origin>
	  <origin>Brazil</origin>
	 </plant>
	*/

	// To add a generic XML header to the output, append it explicitly.
	fmt.Println(xml.Header + string(out))
	// <?xml version="1.0" encoding="UTF-8"?>
	// <plant id="27">
	//  <name>Coffee</name>
	//  <origin>Ethiopia</origin>
	//  <origin>Brazil</origin>
	// </plant>

	// Use Unmarshal to parse a stream of bytes with XML into a data structure.
	// If the XML is malformed or cannot be mapped onto Plant, a descriptive error will be returned.
	var p Plant
	if err := xml.Unmarshal(out, &p); err != nil {
		panic(err)
	}
	fmt.Println(p)
	// Plant id=27, name=Coffee, origin=[Ethiopia Brazil]

	tomato := &Plant{Id: 81, Name: "Tomato"}
	tomato.Origin = []string{"Mexico", "California"}
	// The parent>child>plant field tag tells the encoder to nest all plants under <parent><child>...
	type Nesting struct {
		XMLName xml.Name `xml:"nesting"`
		Plants  []*Plant `xml:"parent>child>plant"`
	}

	nesting := &Nesting{}
	nesting.Plants = []*Plant{coffee, tomato}

	out, _ = xml.MarshalIndent(nesting, " ", " ")
	fmt.Println(string(out))
	// <nesting>
	// <parent>
	// 	<child>
	// 	<plant id="27">
	// 	<name>Coffee</name>
	// 	<origin>Ethiopia</origin>
	// 	<origin>Brazil</origin>
	// 	</plant>
	// 	<plant id="81">
	// 	<name>Tomato</name>
	// 	<origin>Mexico</origin>
	// 	<origin>California</origin>
	// 	</plant>
	// 	</child>
	// </parent>
	// </nesting>
}
```

### 55) Time:

```go
package main

import (
	"fmt"
	"time"
)

// Go offers extensive support for times and durations; here are some examples.

func main() {
	p := fmt.Println

	// We'll start by getting the current time.
	now := time.Now()
	p(now)
	// 2024-08-19 20:56:25.679406 +0530 IST m=+0.000098793

	// You can buld a time struct by providing the year, month, day, etc. Times are always associated with a  Location, i.e. timezone.
	then := time.Date(
		2009, 11, 17, 20, 34, 58, 651387237, time.UTC,
	)
	p(then)
	// 2009-11-17 20:34:58.651387237 +0000 UTC

	// You can extract the various components of the time value as expected.
	p(then.Year())
	// 2009
	p(then.Month())
	// November
	p(then.Day())
	// 17
	p(then.Hour())
	// 20
	p(then.Minute())
	// 34
	p(then.Second())
	// 58
	p(then.Nanosecond())
	// 651387237
	p(then.Location())
	// UTC

	// The Monday-Sunday Weekdays are also available.
	p(then.Weekday())
	// Tuesday

	// These methods compare two times, testing if the first occur before, after, or at the same time as the second, respectively.
	p(then.Before(now))
	// true
	p(then.After(now))
	// false
	p(then.Equal(now))
	// false

	// The Sub methods returns a Duration representing teh interval between two times.
	diff := now.Sub(then)
	p(diff)
	// 129330h54m55.854280763s

	// We can compute the length of the duration in various units.
	p(diff.Hours())
	// 129330.92314752132
	p(diff.Minutes())
	// 7.759855388851279e+06
	p(diff.Seconds())
	// 4.6559132333107674e+08
	p(diff.Nanoseconds())
	// 465591323331076763

	// You can use Add to advance a time by a given duration, or with a - to move backwards by a duration.
	p(then.Add(diff))
	// 2024-08-19 15:32:47.704227 +0000 UTC
	p(then.Add(-diff))
	// 1995-02-16 01:37:09.598547474 +0000 UTC
}
```
