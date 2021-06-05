# Roll the Dice

A simple dice-rolling command line program that takes two arguments: the number of dice, and the number of sides per.

```
go get github.com/HowellBP/roll
```

## Usage

```
$ roll
Usage: roll [# dice] [# sides]
Example: "roll 1 6"
```

## Output

```
$ roll 3 6
Rolling 3d6
Results: [4 3 2]
Total: 9
```

## Compiling

Simply run `go install` on any platform.

Built with [Golang](https://golang.org/) 1.16.
