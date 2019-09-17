package ZhihuCrawler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type ZhihuClient struct{
	client *http.Client
}

func NewZhihuClient() *ZhihuClient {
	httpClient := &http.Client{}
	c := &ZhihuClient{
		client: httpClient,
	}
	return c
}


type TopicDesc struct {
	ID      string
	Limit   int
	AfterID string
	Include []string
}

func NewTopicDescr(id string, limit int, afterID string) *TopicDesc {
	topic := &TopicDesc{id, limit, afterID, nil}

	//reference: https://www.twblogs.net/a/5c6e516abd9eee5c86dcee6e
	topic.Include = []string{
		// "data[?(target.type=topic_sticky_module)].target.data[?(target.type=answer)].target.content",
		// "relationship.is_authorized",
		// "is_author",
		// "voting",
		// "is_thanked",
		// "is_nothelp;data[?(target.type=topic_sticky_module)].target.data[?(target.type=answer)].target.is_normal",
		// "comment_count",
		// "voteup_count",
		// "content",
		// "relevant_info",
		// "excerpt.author.badge[?(type=best_answerer)].topics;data[?(target.type=topic_sticky_module)].target.data[?(target.type=article)].target.content",
		// "voteup_count",
		// "comment_count",
		// "voting",
		// "author.badge[?(type=best_answerer)].topics;data[?(target.type=topic_sticky_module)].target.data[?(target.type=people)].target.answer_count",
		// "articles_count",
		// "gender",
		// "follower_count",
		// "is_followed",
		// "is_following",
		// "badge[?(type=best_answerer)].topics;data[?(target.type=answer)].target.annotation_detail",
		// "content",
		// "hermes_label",
		// "is_labeled",
		// "relationship.is_authorized",
		// "is_author",
		// "voting",
		// "is_thanked",
		// "is_nothelp;data[?(target.type=answer)].target.author.badge[?(type=best_answerer)].topics;data[?(target.type=article)].target.annotation_detail",
		// "content",
		// "hermes_label",
		// "is_labeledj",
		// "author.badge[?(type=best_answerer)].topics;data[?(target.type=question)].target.annotation_detail",
		// "comment_count;",
	}
	return topic
}






func (t *ZhihuClient) GetTopicQuestions(topic *TopicDesc) ([] Question, error) {
	urlStr := fmt.Sprintf("https://www.zhihu.com/api/v4/topics/%s/feeds/essence", topic.ID)
	//urlStr := fmt.Sprintf("https://www.zhihu.com/api/v4/topics/%s/feeds/top_activity", topic.ID)

	limit := topic.Limit
	if limit <= 0 {
		limit = 20
	}
	params := url.Values{
		"limit":    {strconv.Itoa(topic.Limit)},
		"include":  {strings.Join(topic.Include, ",")},
		"after_id": {topic.AfterID},
	}
	res, err := t.SendRequest(urlStr, params)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	result := &QuestionList{}
	err = json.Unmarshal(res, result)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	i:=0
	var Questions [] Question
	for !result.Paging.IsEnd {

		if len(Questions) >= 20 {
			break
		}

		for _, data := range result.Data {


			if i < 70 {
				i++
				continue
			}

			question := data.Target.Question
			keywords := []string{"", ""}
			for _, keyword := range keywords {
				if strings.Contains(question.Title, keyword) {
					Questions = append(Questions, question)
					fmt.Println(question.Title)
					fmt.Println(question.URL)
					//t.GetQuestionAnswersByQid(question.ID)
					fmt.Println()
					break
				}
			}
		}

		res, err := t.SendRequest(result.Paging.Next, nil)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}

		result = &QuestionList{}
		err = json.Unmarshal(res, result)
		if err != nil {
			log.Println(err)
			return nil, err
		}
	}
	return Questions, nil
}

type QuestionDesc struct {
	ID      int
	Limit   int
	AfterID string
	Include []string
}



//func (t *ZhihuClient) GetQuestionAnswersByQid(qid int){ //302684598
//	q := Question{
//		ID: qid,
//	}
//	questions := [] Question {q}
//
//	for _, question := range questions {
//
//		qd := &QuestionDesc{
//			ID: question.ID,
//		}
//
//		fmt.Println(question.Title)
//		fmt.Println(question.URL)
//
//		t.GetAnswerList(qd)
//		//for i, ans := range res {
//		//	fmt.Println("No.", i, ":", string(ans.Content))
//		//}
//
//		fmt.Println()
//		fmt.Println()
//		fmt.Println()
//	}
//}

//https://www.zhihu.com/api/v4/answers/719779548/root_comments?order=normal&limit=20&offset=0&status=open

func (t *ZhihuClient) GetAnswerList(question *QuestionDesc) ([]Answer, error) {
	urlStr := fmt.Sprintf("https://www.zhihu.com/api/v4/questions/%d/answers?include=content", question.ID)

	fmt.Println(urlStr)

	limit := question.Limit
	if limit <= 0 {
		limit = 20
	}

	res, err := t.SendRequest(urlStr, nil)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	fmt.Println(string(res))

	result := &AnswerList{}
	err = json.Unmarshal(res, result)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	var Answers [] Answer
	for !result.Paging.IsEnd {

		for _, ans := range result.Answers {
			if strings.Contains(ans.Content, "深圳") && strings.Contains(ans.Content, "") {
				thirtyDaysAgo := time.Now().Unix() - 15 * 24 * 3600
				if ans.UpdatedTime < thirtyDaysAgo {
					continue
				}

				Answers = append(Answers, ans)
				//fmt.Println(ans.URL)
				//fmt.Println(time.Unix(int64(ans.CreatedTime), 0).Format("2006-01-02 15:04:05")) //2018-07-11 15:10:19
				//fmt.Println(time.Unix(int64(ans.UpdatedTime), 0).Format("2006-01-02 15:04:05")) //2018-07-11 15:10:19
				fmt.Println("[", ans.ID,"]", ans.Question.URL)
				fmt.Println("[", ans.ID,"]", ans.URL)
				fmt.Println("[", ans.ID,"]", trimHtml(ans.Content))
				fmt.Println()
			}
		}

		res, err := t.SendRequest(result.Paging.Next, nil)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}

		result = &AnswerList{}
		err = json.Unmarshal(res, result)
		if err != nil {
			log.Println(err)
			return nil, err
		}
	}

	fmt.Println("IsEnd:", result.Paging)
	if len(Answers) < 100 {
		Answers = append(Answers, result.Answers...)
	}
	return Answers, nil
}

func (t *ZhihuClient) SendRequest(url string, params url.Values) ([]byte, error) {
	encParams := params.Encode()
	req, err := http.NewRequest("GET", url, bytes.NewBufferString(encParams))
	if err != nil {
		fmt.Println("[SendRequest]",err)
		return nil, err
	}
	req.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_3) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/73.0.3683.75 Safari/537.36")
	req.Header.Add("Host", "zhuanlan.zhihu.com")
	//req.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8")
	req.Header.Add("Referer", "https://zhuanlan.zhihu.com/")
	//req.Header.Add("Accept-encoding", "gzip, deflate, br")
	//req.Header.Add("Accept-language", "zh-CN,zh;q=0.9,en-US;q=0.8,en;q=0.7")

	res, err := t.client.Do(req)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	defer res.Body.Close()

	bodyByte, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return bodyByte, nil
}


func trimHtml(src string) string {
	//将HTML标签全转换成小写
	re, _ := regexp.Compile("\\<[\\S\\s]+?\\>")
	src = re.ReplaceAllStringFunc(src, strings.ToLower)
	//去除STYLE
	re, _ = regexp.Compile("\\<style[\\S\\s]+?\\</style\\>")
	src = re.ReplaceAllString(src, "")
	//去除SCRIPT
	re, _ = regexp.Compile("\\<script[\\S\\s]+?\\</script\\>")
	src = re.ReplaceAllString(src, "")
	//去除所有尖括号内的HTML代码，并换成换行符
	re, _ = regexp.Compile("\\<[\\S\\s]+?\\>")
	src = re.ReplaceAllString(src, "\n")
	//去除连续的换行符
	re, _ = regexp.Compile("\\s{2,}")
	src = re.ReplaceAllString(src, "\n")
	return strings.TrimSpace(src)
}




