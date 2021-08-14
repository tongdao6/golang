package main
import (
    "demo/tool"
    "fmt"
)
const STEP = 33*3
func main(){
    rs := tool.InitSource()
    fmt.Printf("%v\n",rs)
    fc:= rs.Map().FillCntByStep(STEP).SortRed().ListRedCnt()
    //.FillCnt()
    // fmt.Printf("%v",rs)
    // rcns := rs.GetRedCnt()
    fmt.Printf("%v",fc)
}