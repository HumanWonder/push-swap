# Push-swap

## Table of Contents

- [About](#about)
- [Getting Started](#getting_started)
- [Usage](#usage)

## About <a name = "about"></a>

Push-swap is a project helping with the learning of using and coming up with sorting algorithms.

We have at our disposal a list of int values, two stacks (a and b) and a set of instructions.

We had to write 2 programs:

- _push-swap_, which calculates and displays on the standard output the smallest program using push-swap instruction language that sorts integer arguments received.

- _checker_, which takes integer arguments and reads instructions on the standard input. Once read, checker executes them and displays OK if integers are sorted. Otherwise, it will display KO.

These are the instructions that you can use to sort the stack :

- **pa** push the top first element of stack b to stack a
- **pb** push the top first element of stack a to stack b
- **sa** swap first 2 elements of stack a
- **sb** swap first 2 elements of stack b
- **ss** execute sa and sb
- **ra** rotate stack a (shift up all elements of stack a by 1, the first element becomes the last one)
- **rb** rotate stack b
- **rr** execute ra and rb
- **rra** reverse rotate a (shift down all elements of stack a by 1, the last element becomes the first one)
- **rrb** reverse rotate b
- **rrr** execute rra and rrb


## Getting Started <a name = "getting_started"></a>

Clone the project : https://zone01normandie.org/git/afouquem/push-swap.git

Inside the directory, use this command to build and use our programs' executable files :
```
go build -o push-swap

go build -o checker
```

## Usage <a name = "usage"></a>

The push-swap program will only receive a list of unique integers separated by a space. Otherwise it will display an error or nothing if there are no arguments. A correct use should look like this :

```
$ ./push-swap "2 1 3 6 5 8"
pb
pb
ra
sa
rrr
pa
pa
```
The checker program takes a list of instructions in a string and applies it to a list of integers (stored in the previoulsy seen stack). If the the stack is sorted after the instructions and the other stack stays empty, checker display a "OK", otherwise it displays "KO". A basic use looks like this :

```
$ echo -e "rra\npb\nsa\nrra\npa" | ./checker "3 2 1 0"
OK
```
If you want to use the program as a whole, use this type of command in your terminal :

```
$ ARG="4 67 3 87 23"; ./push-swap "$ARG" | ./checker "$ARG"
OK
```
In this case, if checker displays "KO", it means that push-swap came up with a list of instructions that did not sort the stack.

Subject : https://zone01normandie.org/git/root/public/src/branch/master/subjects/push-swap

_Project done by **afouquem**, **avaren** and **drodolph**._
