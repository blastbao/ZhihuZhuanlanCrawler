package ZhihuCrawler

import (
	"fmt"
	"testing"
)


const PID = 19588255 // https://zhuanlan.zhihu.com/p/41604227

func TestTopicClient_Demo(t *testing.T) {
	c := NewZhihuClient()
	pidStr := fmt.Sprintf("%d", PID)
	desc := NewTopicDescr(pidStr, 10, "1")

	questions, err := c.GetTopicQuestions(desc)
	if err != nil {
		fmt.Println("result empty error")
		return
	}

	for _, question := range questions{
		fmt.Println(question.Title)
		fmt.Println(question.URL)
		fmt.Println()
		Demo(question.ID)
	}

	//skip := len(questions) / 10
	//cnt := len(questions) / skip
	//
	//var questionSlices [][]Question
	//for i:=1;i<cnt;i++{
	//	tmpSlice := questions[(i-1)*skip:i*skip]
	//	questionSlices = append(questionSlices, tmpSlice)
	//}
	//
	//wg := sync.WaitGroup{}
	//wg.Add(len(questionSlices))
	//for _, j := range questionSlices {
	//	go func(a []Question){
	//		for _, q := range a{
	//			Demo(q.ID)
	//		}
	//		wg.Done()
	//	}(j)
	//}
	//wg.Wait()
}


func Demo(qid int){ //302684598

	c := NewZhihuClient()

	//pidStr := fmt.Sprintf("%d", PID)
	//desc := NewTopicDescr(pidStr, 10, "1")

	//questions, err := c.GetTopicQuestions(desc)
	//if err != nil {
	//	fmt.Println("result empty error")
	//	return
	//}

	q := Question{
		ID: qid,
	}

	questions := [] Question {q}

	for _, question := range questions {

		qd := &QuestionDesc{
			ID: question.ID,
		}

		fmt.Println(question.Title)
		fmt.Println(question.URL)

		c.GetAnswerList(qd)
		//for i, ans := range res {
		//	fmt.Println("No.", i, ":", string(ans.Content))
		//}

		fmt.Println()
		fmt.Println()
		fmt.Println()
	}
	//resultStr, _ := json.Marshal(topics)
	//fmt.Println(string(resultStr))
	//fmt.Println("GetPinnedArticlePidAndAuthor ok")
}

func TestQuestionsClient_Demo(t *testing.T) {
	c := NewZhihuClient()

	//pidStr := fmt.Sprintf("%d", PID)
	//desc := NewTopicDescr(pidStr, 10, "1")

	//questions, err := c.GetTopicQuestions(desc)
	//if err != nil {
	//	fmt.Println("result empty error")
	//	return
	//}

	q := Question{
		ID: 295608199,
	}

	questions := [] Question {q}

	for _, question := range questions {

		qd := &QuestionDesc{
			ID: question.ID,
		}

		fmt.Println(question.Title)
		fmt.Println(question.URL)

		c.GetAnswerList(qd)
		//for i, ans := range res {
		//	fmt.Println("No.", i, ":", string(ans.Content))
		//}

		fmt.Println()
		fmt.Println()
		fmt.Println()
	}
	//resultStr, _ := json.Marshal(topics)
	//fmt.Println(string(resultStr))

	t.Log("GetPinnedArticlePidAndAuthor ok")
}