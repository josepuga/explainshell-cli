# An explainshell.com CLI made with Go

This is just a personal project for my own use. Sometimes you need to know an specific command and/or parameter off the shell. Man is too heavy and tldr is too light. explainshell.com goes right to the point.

## Some samples

```code
$ expl tar zxvf example.tar.gz
The GNU version of the tar archiving utility
-z, --gzip, --gunzip --ungzip
-x, --extract, --get
      extract files from an archive
-v, --verbose
      verbosely list files processed
-f, --file ARCHIVE
      use archive file or device ARCHIVE
```

```code
$ expl du -sh
estimate file space usage
-s, --summarize
       display only a total for each argument
-h, --human-readable
       print sizes in human readable format (e.g., 1K 234M 2G)
```

## Install / Use

No installation required. Is only one executable. On bin/ there are compiled versions for linux, win64 and win32 (Useful if you are using Cygwin).

For Linux copy the expl file inside your $HOME/bin or any bin directory of your $PATH

## Compile by your self

Golang must be installed.

Clone the git repository an exec the compile.sh script. 3 files (expl, expl32.exe and expl.exe) will be created again in bin/
