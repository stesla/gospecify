package specify

import "testing"
import spec "github.com/doun/gospecify"
func TestRunAllSpec(t *testing.T){
    spec.Main(getCurrentRunner())
}
