#
## links
- https://github.com/joomcode/errorx
- https://go.dev/blog/go1.13-errors
- https://blog.devgenius.io/wrapping-multiple-errors-in-go-1-20-78163ef5fc2c
## Output
`GOROOT=/usr/local/go #gosetup
GOPATH=/Users/linczy/go #gosetup
/usr/local/go/bin/go build -o /Users/linczy/Library/Caches/JetBrains/GoLand2023.1/tmp/GoLand/___go_build_error_pr_cmd error-pr/cmd #gosetup
/Users/linczy/Library/Caches/JetBrains/GoLand2023.1/tmp/GoLand/___go_build_error_pr_cmd
Application starting
--------Simple error------------
2023/04/24 15:54:00 main.go:19: Error: Level 1 error message Level 2 error message
--------WRAP error------------
2023/04/24 15:54:00 main.go:25: Error: Level 1 error message Level 2 error message
--------ErrorX------------
2023/04/24 15:54:00 main.go:31: Error: Level 1 error message, cause: common.illegal_state: Level 2 error message
at main.throwError2X()
/Users/linczy/code/go-playground/error-pr/cmd/main.go:36
at main.throwError1X()
/Users/linczy/code/go-playground/error-pr/cmd/main.go:40
at main.main()
/Users/linczy/code/go-playground/error-pr/cmd/main.go:29
at runtime.main()
/usr/local/go/src/runtime/proc.go:250
at runtime.goexit()
/usr/local/go/src/runtime/asm_amd64.s:1598

Process finished with the exit code 0


`




