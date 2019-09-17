package ZhihuCrawler

type Zhuanlan struct {
	Slug string `json:"slug"`
	Name string `json:"name"`
}

type Author struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Gender      int    `json:"gender"`
	Headline    string `json:"headline"`
	Description string `json:"description"`
	AvatarUrl   string `json:"avatar_url"`
	Type        string `json:"type"`
	UID         string `json:"uid"`
	URL         string `json:"url"`
	URLToken    string `json:"url_token"`
	UserType    string `json:"user_type"`
	AvatarURL   string `json:"avatar_url"`
}

type PinnedArticleAndAuthor struct {
	Type     string `json:"type"`
	ID       int    `json:"id"`
	Updated  int64  `json:"updated"`
	Created  int64  `json:"created"`
	Title    string `json:"title"`
	ImageURL string `json:"image_url"`
	URL      string `json:"url"`
	Excerpt  string `json:"excerpt"`
	Author   Author
}

type Topic struct {
	Url  string `json:"url"`
	Type string `json:"type"`
	Id   string `json:"id"`
	Name string `json:"name"`
}

type Article struct {
	ID       int     `json:"id"`
	Type     string  `json:"type"`
	Title    string  `json:"title"`
	URL      string  `json:"url"`
	Updated  int64   `json:"updated"`
	Created  int64   `json:"created"`
	Excerpt  string  `json:"excerpt"`
	Content  string  `json:"content"`
	ImageURL string  `json:"image_url"`
	Topics   []Topic `json:"topics"`
}

type ArticleList struct {
	Paging struct {
		IsEnd   bool `json:"is_end"`
		Totals  int  `json:"totals"`
		IsStart bool `json:"is_start"`
	} `json:"paging"`
	Data []struct {
		ID int `json:"id"`
	} `json:"data"`
}

type QuestionList struct {
	Paging struct {
		IsEnd   bool `json:"is_end"`
		IsStart bool `json:"is_start"`
		Previous string `json:"previous"`
		Next string `json:"next"`
	} `json:"paging"`
	Data []struct {
		Target struct {
			Question Question `json:"question"`
		} `json:"target"`
	} `json:"data"`
}

type Question struct {
	ID int `json:"id"`
	Title string `json:"title"`
	URL string `json:"url"`
}

type Answer struct {
	ID int `json:"id"`
	Author Author `json:"author"`
	URL string `json:"url"`
	Question Question `json:"question"`
	Content string `json:"content"`
	CreatedTime int64 `json:"created_time"`
	UpdatedTime int64 `json:"updated_time"`
}

type AnswerList struct {
	Paging struct {
		IsEnd   bool `json:"is_end"`
		IsStart bool `json:"is_start"`
		Previous string `json:"previous"`
		Next string `json:"next"`
		Totals int `json:"totals"`
	} `json:"paging"`
	Answers [] Answer `json:"data"`
}