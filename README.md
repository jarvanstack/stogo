

# stogo 
多连表查询语句直接生成 go DTO 层结构体代码

> 许多 orm 框架都有能直接单表映射代码的功能, 但是多表却做不到
> **因为多表组合的可能是无限种**
> 每次写数据映射表都很难受, 于是答主就写了一个工具支持**连表查询语句直接生成 golang DTO 层结构体代码**
## 快速开始

### get 

```bash
go get -u github.com/dengjiawen8955/stogo@v0.0.3
```

### 使用示例

```go
package tests

import (
	"testing"

	"github.com/dengjiawen8955/stogo/stogo"
)

func Teststogo(t *testing.T) {
    //查询语句
	ssql := `	SELECT
	community_question.pk_id AS q_id, 
	community_question.title, 
	community_question.info, 
	community_question.currency as q_currency, 
	community_question.comment_num, 
	community_question.review_num, 
	community_question.update_time as q_update_time, 
	community_question.kind1_id, 
	community_question.kind2_id, 
	community_question.questioner_id, 
	user.pk_id as u_id, 
	user.avatar, 
	user.nickname as nickname, 
	user.answer_score_total, 
	user.answer_score_num
FROM
	community_question,
	user
WHERE
	community_question.questioner_id = user.pk_id 
	AND  community_question.kind1_id = 1
	AND community_question.kind2_id = 11`
    //数据Driver
	driver := "root:root.@tcp(localhost:3306)/mydb"
    //多连表查询语句直接生成 go DTO 层结构体代码
	stogo.GenerateStruct(ssql, driver)
}
```
输出效果如下:

```go

type AutoDTO struct {
	QId              int64   `db:"q_id"`
	Title            string  `db:"title"`
	Info             string  `db:"info"`
	QCurrency        float64 `db:"q_currency"`
	CommentNum       int64   `db:"comment_num"`
	ReviewNum        int64   `db:"review_num"`
	QUpdateTime      string  `db:"q_update_time"`
	Kind1Id          int64   `db:"kind1_id"`
	Kind2Id          int64   `db:"kind2_id"`
	QuestionerId     int64   `db:"questioner_id"`
	UId              int64   `db:"u_id"`
	Avatar           string  `db:"avatar"`
	Nickname         string  `db:"nickname"`
	AnswerScoreTotal int64   `db:"answer_score_total"`
	AnswerScoreNum   int64   `db:"answer_score_num"`
}
```