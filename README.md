# Declaration Speeds ![Stars](https://img.shields.io/github/stars/Simpson-Computer-Technologies-Research/DeclarationSpeeds?color=brightgreen) ![Watchers](https://img.shields.io/github/watchers/Simpson-Computer-Technologies-Research/DeclarationSpeeds?label=Watchers)
![banner](https://user-images.githubusercontent.com/75189508/194382177-1c11ee03-f048-4a36-995e-2a1673d08afc.png)

When Declaring a Variable in Golang, should you use 'var n string' or 'n :='? Which is Faster?

# Benchmarks
```go

PreDeclaredType() -> 333.664513ms

NoDeclaredType() -> 367.542388ms

```


# Functions

```go
// Pre declare the variables type
func PreDeclaredType() {
	// Declare the s variable 10^8 amount of times
	for i := 0; i < 1000000000; i++ {
		var s string = "this is a string!"
		_ = s
	}
}

// No declared variable type
func NoDeclaredType() {
	// Declare the s variable 10^8 amount of times
	for i := 0; i < 1000000000; i++ {
		s := "this is a string!"
		_ = s
	}
}

```
