# csvify

This package let's you write .csv files easily with Go.

For more convenience it simply appends a new line to a file.
If the file doesn't exist, it will be created.

If you want to write multiple lines, please use a loop.
In the example below I'm looping over structs but please note it would also works with loops over slices, arrays, maps or primitive types.
```
type test struct {
	Floats  []float64
	Slices  [][]int
	Map     map[string]int
	Arrays  [2][3]int
	Boolean bool
}

func main() {
	test := []test{{
		Floats:  []float64{1.1, 2.2, 3.3},
		Slices:  [][]int{{3, 4}, {6, 7, 8, 9, 10}},
		Map:     map[string]int{"earth": 1, "mars": 2, "jupiter": 3},
		Arrays:  [2][3]int{{1, 2, 3}, {4, 5, 6}},
		Boolean: true,
	}, {
		Floats:  []float64{7.7, 8.8, 9.9},
		Slices:  [][]int{{6, 7}, {10, 11, 12, 13, 14}},
		Map:     map[string]int{"trappist": 1, "sun": 2, "proxima-b": 3},
		Arrays:  [2][3]int{{10, 20, 30}, {40, 50, 60}},
		Boolean: false,
	}}
	for i := 0; i < len(test); i++ {
		err := CsvifyLine("test.csv", test[i])
		if err != nil {
			fmt.Println("Error:", err)
		}
	}
}
```

It will end up in appending those lines to test.csv :

```
1.1;2.2;3.3;3;4;6;7;8;9;10;earth=1;mars=2;jupiter=3;1;2;3;4;5;6;true;
7.7;8.8;9.9;6;7;10;11;12;13;14;trappist=1;sun=2;proxima-b=3;10;20;30;40;50;60;false;
```
