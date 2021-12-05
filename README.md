# x-knight-tour

[![CI](https://github.com/caulagi/x-knight-tour/actions/workflows/ci.yaml/badge.svg)](https://github.com/caulagi/x-knight-tour/actions/workflows/ci.yaml)

A ~mutated~ modified knight tour with the following rules:

* Move 3 squares horizontally or
* Move 3 squares vertically or
* Move 2 squares diagonally.
* Moves to squares already visited or outisde of the checkerboard are NOT allowed.

```shell
$ go build

$ ./x-knight-tour --help
NAME:
   x-knight-tour - a modified knight's tour

USAGE:
   x-knight-tour [global options] startX startY

DESCRIPTION:
   uses warnsdorff algorithm to find a modified knight tour

GLOBAL OPTIONS:
   --help             show help (default: false)
   --boardSize value  size of board (default: 10)

$ ./x-knight-tour --boardSize 5
   1  14   7   4  15
   9  22  17  12  23
  19   5  25  20   6
   2  13   8   3  16
  10  21  18  11  24

$ ./x-knight-tour
   1  62  40  16  61  39  75  49  38  74
  54  18   3  53  58   4  52  72   5  51
  41  15  60  67  94  79  66  95  78  48
   2  63  57  17  64  69  76  50  37  73
  55  19  91  82  59  90  99  71   6  98
  42  14  29  68  93  80  65  96  77  47
  23  83  56  24  86  70  25  87  36  10
  30  20  92  81  32  89 100  33   7  97
  43  13  28  44  12  27  45  11  26  46
  22  84  31  21  85  34   8  88  35   9
```

## Tests

```
$ go test -v ./...
?   	github.com/caulagi/x-knight-tour	[no test files]
=== RUN   TestNewBoard
--- PASS: TestNewBoard (0.00s)
=== RUN   TestHorizontalCandidates
--- PASS: TestHorizontalCandidates (0.00s)
=== RUN   TestVerticalCandidates
--- PASS: TestVerticalCandidates (0.00s)
=== RUN   TestDiagonalCandidates
--- PASS: TestDiagonalCandidates (0.00s)
=== RUN   TestSolve
--- PASS: TestSolve (0.00s)
PASS
ok  	github.com/caulagi/x-knight-tour/board	(cached)
```

#### License

This software is licensed under the MIT License. See [LICENSE](LICENSE) for details.
