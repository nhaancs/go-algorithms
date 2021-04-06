# Big O
How many steps the computer has to perfom for each function.

![big-o-complexity-chart.png](./images/big-o-complexity-chart.png)

## O(n)
Linear time, the big O depends on the number n (number of inputs)

```go
func findNemo(arr []string) { // O(n)
	for i:=0; i < len(arr); i++ {
		if arr[i] == "nemo" {
			fmt.Println("Found NEMO")
		}
	}
}
```

![o-n.png](./images/o-n.png)

## O(1)
Constant time, the number of operations is the constant no matter how large the input is.

```go
func logFirstTwoElements(arr []string) {
	fmt.Println(arr[0])
	fmt.Println(arr[1])
}
```

![o-1.png](./images/o-1.png)

## Analysis a simple funciton

```go
func funChallenge(input []int) {
	a := 10 // O(1)
	a = 50 + 3 // O(1)

	for i := 0; i < len(input); i++ { // O(n)
		anotherFunction() // O(n)
		let stranger = true // O(n)
		a++ // O(n)
	}

	return a // O(1)
}

// Big O(3 + 4n) => Big O(n)
```

## Big O rule 1: Worst case
Always think about the worst case.

## Big O rule 2: Remove constants
Remove constants.

## Big O rule 3: Different terms for inputs
```go
func compressBoxes(boxes1, boxes2 []int) {
	for i := 0; i < len(boxes1); i++ {
		fmt.Println(i)
	}
	
	for i := 0; i < len(boxes2); i++ {
		fmt.Println(i)
	}
}

// Big O(a+b)
```

## O(n^2)
Use multiplication for nested loops.

```go
func logAllPairsOfArr(arr []string) {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr); j++ {
			fmt.Println(arr[i], arr[j])
		}
	}
}

// O(n*n) => O(n^2)
```

**Apply rule 3**

```go
func logAllPairsOfArr(arr1, arr2 []string) {
	for i := 0; i < len(arr1); i++ {
		for j := 0; j < len(arr2); j++ {
			fmt.Println(arr1[i], arr2[j])
		}
	}
}

// O(a*b)
```

