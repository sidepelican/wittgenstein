# wittgenstein
insert shell command result in the text file

## development environment
- macOS 10.14.2

```
$ go version
go version go1.11.5 darwin/amd64
```

## build

```
$ go build
```

## example
My usecase is `cmake`.
It is easy to add new files ( `*.cpp`, `*.h` ) into the compile target list by `wittgenstein` .

- firetree:
```
├── CMakeLists.txt
└── Classes
    ├── ClassA.cpp
    └── ClassB.cpp
```

- original file: 

```CMakeLists.txt
...
list(APPEND SOURCE_FILES
# WITTGENSTEIN_BEGIN `find -L Classes -name *.cpp`
Classes/ClassA.cpp
# WITTGENSTEIN_END
)
...
```

- run: 

```
$ wittgenstein CMakeLists.txt
```

- result: 

```CMakeLists.txt
...
list(APPEND SOURCE_FILES
# WITTGENSTEIN_BEGIN `find -L Classes -name *.cpp`
Classes/ClassA.cpp
Classes/ClassB.cpp    # 👈 added by wittgenstein
# WITTGENSTEIN_END
)
...
```


## format

```
# WITTGENSTEIN_BEGIN `<command>`
<command result here>
# WITTGENSTEIN_END
```

or 

```
// WITTGENSTEIN_BEGIN `<command>`
<command result here>
// WITTGENSTEIN_END
```

## license
MIT.
