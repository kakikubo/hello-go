# hello-go

## delveのインストール

<https://github.com/go-delve/delve/tree/master/Documentation/installation>

```bash
go install github.com/go-delve/delve/cmd/dlv@latest
```

## `fmt.Println`のコールグラフ

```mermaid
graph TD


main --> Println

subgraph fmt
    Println --> Fprintln
    Fprintln
end

Fprintln --> File.Write

subgraph os/file
  File.Write --> os/file_posix.go.File.Write
  subgraph os/file_posix.go
    os/file_posix.go.File.Write((File.Write))
  end
end

os/file_posix.go.File.Write --> internal/poll/fd_unix.go.FD.write((FD.write))
os/file_posix.go.File.Write --> internal/poll/fd_windows.go.FD.write((FD.write))

subgraph intenal/poll
  subgraph internal/poll/fd_unix.go
    internal/poll/fd_unix.go.FD.write
  end
  subgraph internal/poll/fd_windows.go
    internal/poll/fd_windows.go.FD.write --> internal/poll/fd_windows.go.FD.writeConsole((FD.writeConsole))
  end
end


subgraph syscall
  internal/poll/fd_unix.go.FD.write((FD.write)) --> syscall.Write((Write))
  internal/poll/fd_windows.go.FD.write((FD.write)) --> syscall.Write
  internal/poll/fd_windows.go.FD.writeConsole((FD.writeConsole)) --> syscall.WriteConsole((WriteConsole))
end
```

## gpコマンドのインストール

<https://github.com/tenntenn/goplayground>

```bash
go install github.com/tenntenn/goplayground/cmd/gp@latest
```

## fmt.Scanln関数

https://go.dev/play/p/F6_TPUeLa7X
標準入力から値をポインタで読み取り、変数に格納する。
```go
package main
import "fmt"
func main() {
  var price int
  fmt.Print("値段>")
  fmt.Scanln(&price)
  fmt.Printf("%d円\n", price)
}
```
