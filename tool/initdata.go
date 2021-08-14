package tool
import (
    "strconv"
    "bufio"
	"fmt"
	"io"
	"os"
    "sort"
	"strings"
)
// 对每一行读入的数据存入指定格式的结构体中
func processLine(w *ArrIssueInf) func(line []byte) {
    return func(line []byte){
        //获取并分离字符串
        //去掉\n\r
        s1:=strings.Replace(string(line[:]), "\r\n", "", -1)
        s := strings.Split(s1, ",")
        
        //期号
        issue,_:= strconv.Atoi(s[0])
        redNums := make([]NumInf,0,6)
        for i:=1;i<7;i++{
            redNums=append(redNums,NumInf{
                Num:s[i],
                ID:0,
                Cnt:0,
                Col:0,
            })
        }
        
        *w = append(*w,IssueInf{
            issue:issue,
            RedNums:redNums,
            Blue:NumInf{
                Num:s[7],
                ID:0,
                Cnt:0,
                Col:0,
                },
            Stat:make(map[string]int,33),
        }) 
    }
}

// 对数据文件进行逐行数据读入并处理
func ReadLine(filePth string, hookfn func(line []byte)) error {
	f, err := os.Open(filePth)
	if err != nil {
		fmt.Println("读取错误")
		return err
	}
	defer f.Close()

	bfRd := bufio.NewReader(f)
	for {
		line, err := bfRd.ReadBytes('\n')
		hookfn(line)    //放在错误处理前面，即使发生错误，也会处理已经读取到的数据。
		if err != nil { //遇到任何错误立即返回，并忽略 EOF 错误信息
			if err == io.EOF {
				return nil
			}
			return err
		}
	}
}


// 从数据文件中读取数据并 初始化格式
func ReadData() *ArrIssueInf {
    hwn :=make(ArrIssueInf,0,128)
    cbfn := processLine(&hwn)
    ReadLine("code.txt",cbfn)
    return &hwn
}


//初始化输入数据为要求格式
func InitSource() *ArrIssueInf{
    hwn:=ReadData()
    sort.Stable(*hwn)
    return hwn
}

// func (a *ArrIssueInf)Map() *ArrNumInf{
//     ret := make(ArrNumInf,0,0)
//     for id,v := range *a{
//         for col,numinf := range v.RedNums{
//             ret = append(ret,NumInf{
//                 Num:numinf.Num,
//                 ID:id,
//                 Cnt:0,
//                 Col:col+1,
//                 })
//         }
//     }
//     return &ret
// }

func (a *ArrNumInf) Filter(cb func (v NumInf) bool)[]NumInf{
    ret := make([]NumInf,0,128)
    for _,v:=range *a{
        if cb(v) {
            ret = append(ret,v)
        }
    }
    fmt.Println("长度：",len(ret))
    return ret
}

func (a *ArrNumInf) FilterRed(cb func (v NumInf) bool) *ArrNumInf{
    ret := make(ArrNumInf,0,128)
    for _,v:=range *a{
        if cb(v) && v.Col<7 {
            ret = append(ret,v)
        }
    }
    fmt.Println("长度：",len(ret))
    return &ret
}
//复制map
func copyMap(s map[string]int) map[string]int{
    ret := make(map[string]int,len(s))
    for key,val := range s{
        ret[key]=val
    }
    return ret
}

//逐行填充出现次数的统计值
func (a *ArrIssueInf) Map()*ArrIssueInf{
    ret := make(map[string]int,33)
    for id,numinf:= range *a{    
        for _,v := range numinf.RedNums {
            if _,ok:= ret[v.Num];ok{
                ret[v.Num]++
            }else{
                ret[v.Num]=1
            }
        }
        (*a)[id].Stat = copyMap(ret)
    }
    return a
}

//将统计的次数填充到每个数对应的cnt属性上
func (a *ArrIssueInf) FillCnt() *ArrIssueInf{
    for i,numinf := range *a{
        for j,reds := range numinf.RedNums{
            (*a)[i].RedNums[j].Cnt = (*a)[i].Stat[reds.Num]
            (*a)[i].RedNums[j].ID = i
        }
    }
    return a
}

//将红球部分排序并填充Col属性
func (a *ArrIssueInf) SortRed() *ArrIssueInf{
    for i,numinf := range *a{
        sort.Stable(ArrNumInf(numinf.RedNums))
        for j := range numinf.RedNums{
            (*a)[i].RedNums[j].Col = j+1
        }
    }
    return a
}


//根据指定的步长STEP 将统计的次数填充到每个数对应的cnt属性上
func (a *ArrIssueInf) FillCntByStep(step int) *ArrIssueInf{
    for i,numinf := range *a{
        if i<=step {
            continue
        }
        for j,reds := range numinf.RedNums{
            b,e := 0,0
            if x,ok :=(*a)[i].Stat[reds.Num];!ok{
                e=0
            } else{
                e = x
            }
            if x,ok := (*a)[i-step].Stat[reds.Num];!ok{
                b = 0
            }else{
                b = x
            }
            (*a)[i].RedNums[j].Cnt = e-b
            (*a)[i].RedNums[j].ID = i
        }
    }
    return a
}

//打印红球信息
func (a *ArrIssueInf) List(){
    for _,v := range *a{
        fmt.Printf("%v\n",v.RedNums)
    }
}

//提取红球的次数信息
func (a *ArrIssueInf)GetRedCnt()[][]int{
    ret := make([][]int,0,0)
    for _,val := range *a{
        ired := make([]int,6,6)
        for i,rinf := range val.RedNums{
            ired[i]=rinf.Cnt
        }
        ret = append(ret,ired)
    }
    return ret
}

func (a *ArrIssueInf)ListRedCnt()[]int{
    ret := make([]int,0,2048)
    for _,val := range *a{
        ired := make([]int,6,6)
        for i,rinf := range val.RedNums{
            ired[i]=rinf.Cnt
        }
        fc := int(FangCha(ired))
        fmt.Printf("%v aver=%v\n",ired,fc)
        ret = append(ret,fc)
    }
    return ret
}

//求平均值
func average(s []int) float32{
    sum:=float32(0)
    for _,v := range s{
        sum+=float32(v)
    }
    return sum/float32(len(s))
}

//方差
func FangCha(s []int) float32{
   av := average(s)
   fc:=float32(0)
   for _,v := range s{
       t := (float32(v)-av)
       fc = fc + t*t
   }
   return fc/float32(len(s))
}