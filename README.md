# Go Spinners
A collection of progress spinners for Go.

![SpinnerDemo](https://user-images.githubusercontent.com/41476809/181807553-316d0f69-426a-4559-9268-d50a8ae9d611.gif)

## Installation

```bash
go get github.com/tom-draper/go-spinners
```

## Usage

Import the package into your project.

```go
import spinners "github.com/tom-draper/go-spinners"
```

Create a spinner with the <code>Spinner</code> function, passing in the name of the spinner.

```go
s := spinners.Spinner("line")
s.Start()
time.Sleep(time.Second * 5) // Perform computation
s.Stop()
```

Prefix text can be specified with <code>SetPrefix</code> to appear before the spinner animation. Similarly, a postfix can be set with <code>SetPostfix</code>.

```go
s := spinners.Spinner("flip")
s.Start()
s.SetPrefix("Loading")
time.Sleep(time.Second * 5) // Perform computation
s.Stop()
```

The animation speed can be modified using the <code>SetDelay</code> function. The default delay is 100 milliseconds.

```go
s := spinners.Spinner("dots2")
s.SetDelay(time.Millisecond * 500)
s.Start()
time.Sleep(time.Second * 5) // Perform computation
s.Stop()
```
