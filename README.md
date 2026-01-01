# GoLang
I have started learning Go programming languages, these are what i learnt

# Hello World

```Golang
package greetings

import "fmt"


func Hello() {
	fmt.Println("Hello world!")
}
```
First i imported fmt library and as you can see i show **Hello World**. The first listen was simple but it was hard to export them another go file. Because i exported files in JavaScript like that: 
```JavaScript
const hello = () => {
  console.log("Hello world");
};

export hello;
```
i showed all my codes in one go gile, this file is **main.go**.

The second code i learnt is **Values**, the code:
```
package values


import "fmt"

func Values () {
	fmt.Println("go" + "lang")

	fmt.Println("1+1 =", 1+1)
    fmt.Println("7.0/3.0 =", 7.0/3.0)

	fmt.Println(true && false)
    fmt.Println(true || false)
    fmt.Println(!true)
}
```
I know *JavaScript* and *PYTHON*, that helps me to understand them. Because i used these values in *JavaScript* and *PYTHON* too many times.

The third code i learnt is how to make **variables**, Making variables is important!
```
package variables

import "fmt"

func Variables() {

	var a = "initial"
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b, c)

	var d = true
	fmt.Println(d)

	var e int
	fmt.Println(e)

	f := "apple"
	fmt.Println(f)

}
```
It is same as **JavaScript**, I usally use for making variables these commands:
```
const name = "Azizbek";
let last_name = "Jumaboyev";
var age = 19;
```

4th code is constants variables:
```
package constants

import (
    "fmt"
    "math"
)

const s string = "constant"

func Constant() {
    fmt.Println(s)

    const n = 500000000

    const d = 3e20 / n
    fmt.Println(d)

    fmt.Println(int64(d))

    fmt.Println(math.Sin(n))
}
```
**You can't not change that kind of variables**

5th is *for* and *loop*:
```
package for_test

import "fmt"

func For () {
	i := 1

	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	for j:= 0; j < 3; j++ {
		fmt.Println(j)
	}

	for i := range 3 {
		fmt.Println("range: ", i)
	}

	for {
		fmt.Println("Loop")
		break
	}

	for n := range 6 {
		if n % 2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}
```

I'm still learning, one day i will make big project with **Golang** or just simple **Telegram bot** then i will uplaod the code to the **GitHub**
