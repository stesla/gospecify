package main
import (
    "runtime"
    "path"
    "io"
    "os"
    "os/exec"
)
func main(){
    src := path.Join(runtime.GOROOT(),"src/pkg/github.com/doun/gospecify/cmd/testmain.template")
    cSrc,err := os.Open(src)
    if err!=nil {
        panic("open template failed")
    }

    des,err := os.Create("./__test/testmain.go")
    if err !=nil {
        panic("create template failed")
    }
    io.Copy(des,cSrc)
    cmd := exec.Command("go run ./__test/testmain.go")
    cmd.Run()
}
