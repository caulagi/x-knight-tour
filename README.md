# x-knight-tour

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

$ ./x-knight-tour 2 3
  93   2  23  77  33  22  86  32  21  87
  13  39  64  12  40  65  11  41  66  10
  24  50  94   1  91  97  34  90  55  31
  75   3  43  78  63  42  85  62  20  88
  14  38  71  51   0  70  54  98  67   9
  25  49  95  36  92  96  35  89  56  30
  74   4  44  79  60  45  84  61  19  83
  15  37  72  52 100  69  53  99  68   8
  26  48  59  27  47  58  28  46  57  29
  73   5  16  80   6  17  81   7  18  82
```