# yup-comm

```
NAME:
   comm - compare two sorted files line by line

USAGE:
   comm [OPTIONS] FILE1 FILE2

      Compare sorted files FILE1 and FILE2 line by line.

      With no options, produce three-column output. Column one contains
      lines unique to FILE1, column two contains lines unique to FILE2,
      and column three contains lines common to both files.

COMMANDS:
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --suppress-column-1, -1  suppress column 1 (lines unique to FILE1) (default: false)
   --suppress-column-2, -2  suppress column 2 (lines unique to FILE2) (default: false)
   --suppress-column-3, -3  suppress column 3 (lines that appear in both files) (default: false)
   --check-order            check that the input is correctly sorted (default: false)
   --total                  output a summary (default: false)
   --help, -h               show help
```
