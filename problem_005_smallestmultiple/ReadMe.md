# Project 5 - Smallest Multiple

## Selected Problem: Problem 5 - Smallest Multiple

* Source: https://projecteuler.net/problem=5

* Title: Smallest multiple - Problem 5

* Problem Description

```
2520 is the smallest number that can be divided by each of the
numbers from 1 to 10 without any remainder. What is the smallest
positive number that is evenly divisible by all of the numbers
from 1 to 20?
```

## Solutions
* Solution - 1 Slowest brute force solution

    ```
    ./01/main.go
    ```

* Solution - 2 Modified brute force solution. Very slow for upper limits greater than 30

    ```
    ./02/main.go
    ```

* Solution - 3 Uses Prime Factors and uint64 numeric types. Very fast but fails when upper limits exceed 40. Failure is due to generated values greater than the bounds of uint64

    ```
    ./03/main.go
    ```

* Solution - 4 Uses Prime Factors and big.Int (math/big) numeric types.  Very fast and capable of handling very large upper limit values.  This is the preferred solution.

    ```
    ./04/main.go
    ```
