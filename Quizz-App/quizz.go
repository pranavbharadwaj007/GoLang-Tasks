package main
import (
	"flag"
	"os"
	"fmt"
	"encoding/csv"
	"strings"
	"time"
	
)

type problem struct{
	quest string
	ans string
}

func parselines(lines [][]string) []problem{
	ret:=make([]problem,len(lines))
	for i, line:=range lines{
		ret[i]=problem{
			quest:line[0],
			ans:line[1],
		}
	}
	return ret
	}

func handleErr(err error){
	if err!=nil{
		panic(err)
	}
}

func main() {
	csvFilename:=flag.String("csv","problems.csv","a csv file in the format of 'question,answer'")
	timeLimit:=flag.Int("limit",30,"the time limit for the quiz in seconds")
	flag.Parse()
	file, err:=os.Open(*csvFilename)
	handleErr(err)
	fileContent:=csv.NewReader(file)
	lines, err:=fileContent.ReadAll()
	handleErr(err)
	problems:=parselines(lines)
	timer:=time.NewTimer(time.Duration(*timeLimit)*time.Second)

	fmt.Println(problems)
	points:=0
	for i,p:=range problems{
		fmt.Printf("Question #%d: %s= \n",i+1, p.quest)
		answerCh:=make(chan string)
		go func(){
			var answer string
			fmt.Scanf("%s\n",&answer)
			answerCh<-answer
		}()
		select{
		case <-timer.C:
			fmt.Println("Time's up")
			fmt.Println("You scored",points,"out of",len(problems))
			return
		case answer:=<-answerCh:
			if answer==strings.TrimSpace(p.ans){
				points++
				fmt.Println("Correct")
			}else{
				fmt.Println("Incorrect")
				fmt.Println("Correct answer is:",p.ans)
			}
	}
	}	
	fmt.Println("You scored",points,"out of",len(problems))

}
