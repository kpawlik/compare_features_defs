# Compare feature defs

## Usage

- basic

```bash
go run main.go -d1 ./d1 -d2 ./d2 
```

Result:

```bash
$ go run main.go -d1 ./d1 -d2 ./d2 
2025/03/03 10:18:36 Comparing ./d1 ==> ./d2
2025/03/03 10:18:36 Compared files: d1\address.def ==> d2\address.def
2025/03/03 10:18:36   Field Name                    |Value 1                       |Value 2                       |Description
2025/03/03 10:18:36 ------------------------------------------------------------------------------------------------------------------------
2025/03/03 10:18:36   id2                           |                              |                              |Field does not exists in ./d2
2025/03/03 10:18:36   herkomst                      |string(1)                     |string(11)                    |Different values
2025/03/03 10:18:36   straat                        |Straat                        |Straata                       |Different values
```

- both way

```bash
go run main.go -d1 ./d1 -d2 ./d2 -bw
```

Result

```bash
025/03/03 11:19:14 Comparing ./d1 ==> ./d2
2025/03/03 11:19:14 Compared files: d1\address.def ==> d2\address.def
2025/03/03 11:19:14   Field Name                    |Value 1                       |Value 2                       |Description
2025/03/03 11:19:14 ------------------------------------------------------------------------------------------------------------------------
2025/03/03 11:19:14   id2                           |                              |                              |Field does not exists in ./d2
2025/03/03 11:19:14   herkomst                      |string(1)                     |string(11)                    |Different values
2025/03/03 11:19:14   straat                        |Straat                        |Straata                       |Different values
2025/03/03 11:19:14 Comparing ./d2 ==> ./d1
2025/03/03 11:19:14 Compared files: d2\address.def ==> d1\address.def
2025/03/03 11:19:14   Field Name                    |Value 1                       |Value 2                       |Description
2025/03/03 11:19:14 ------------------------------------------------------------------------------------------------------------------------
2025/03/03 11:19:14   id3                           |                              |                              |Field does not exists in ./d1
2025/03/03 11:19:14   herkomst                      |string(11)                    |string(1)                     |Different values
2025/03/03 11:19:14   straat                        |Straata                       |Straat                        |Different values
```

- display not existing files

```bash
go run main.go -d1 ./d1 -d2 ./d2 -ne
```

Result:

```bash
2025/03/03 11:19:42 Comparing ./d1 ==> ./d2
2025/03/03 11:19:42 Compared files: d1\address.def ==> d2\address.def
2025/03/03 11:19:42   Field Name                    |Value 1                       |Value 2                       |Description
2025/03/03 11:19:42 ------------------------------------------------------------------------------------------------------------------------
2025/03/03 11:19:42   id2                           |                              |                              |Field does not exists in ./d2
2025/03/03 11:19:42   herkomst                      |string(1)                     |string(11)                    |Different values
2025/03/03 11:19:42   straat                        |Straat                        |Straata                       |Different values
2025/03/03 11:19:42 File d2\address2 copy.def not exists
```
