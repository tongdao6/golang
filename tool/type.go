package tool

//号码基本信息
type NumInf struct {
	Num string
	ID  int
	Cnt int
	Col int
}
type ArrNumInf []NumInf

type IssueInf struct {
	issue   int
	RedNums []NumInf
	Blue    NumInf
	Stat    map[string]int
}
type ArrIssueInf []IssueInf
