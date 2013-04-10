package specify
import . "github.com/doun/gospecify"
var defualtRunner Runner = NewRunner();
func DefaultRunner()Runner{
    return defualtRunner
}
func After(block AfterFunc) {
	defualtRunner.After(block);
}

func Before(block BeforeBlock) {
	defualtRunner.Before(block);
}

func Describe(name string, block ExampleGroupBlock) {
	defualtRunner.Describe(name, block);
}

func It(name string, block ExampleBlock) {
	defualtRunner.It(name, block);
}

func getCurrentRunner()Runner{
    return defualtRunner
}
