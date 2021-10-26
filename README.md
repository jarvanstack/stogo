# stogo 

多连表查询语句直接生成 go DTO 层结构体代码

## Quick start 

### get

```bash
go get -u github.com/dengjiawen8955/stogo@v0.0.2
```

### usage

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
``