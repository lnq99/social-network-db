package helper

import (
	"encoding/json"

	"app/internal/model"
)

type CmtBuilder struct {
	id       int
	userId   int
	postId   int
	parentId int
	content  string
	created  string
	children []*model.Comment
}

var cmtBuilder CmtBuilder

func (b *CmtBuilder) New(id, userId, postId int, content string) *CmtBuilder {
	b.id = id
	b.userId = userId
	b.postId = postId
	b.content = content
	return b
}

func (b *CmtBuilder) Parent(parentId int) *CmtBuilder {
	b.parentId = parentId
	return b
}

func (b *CmtBuilder) Created(created string) *CmtBuilder {
	b.created = created
	return b
}

func (b *CmtBuilder) AddChildren(child *model.Comment) *CmtBuilder {
	b.children = append(b.children, child)
	return b
}

func (b *CmtBuilder) Build() model.Comment {
	return model.Comment{
		b.id,
		b.userId,
		b.postId,
		b.parentId,
		b.content,
		b.created,
		b.children,
	}
}

type CmtMother struct {
	data         map[int]model.Comment
	PostEmptyId  int
	PostSingleId int
	PostNestedId int
	PostEmpty    string
	PostSingle   string
	PostNested   string
}

func (m *CmtMother) Init() {
	m.data = map[int]model.Comment{
		1: cmtBuilder.New(1, 1, 1, "cmt 1").Build(),
		2: cmtBuilder.New(2, 1, 2, "cmt 2").Build(),
		3: cmtBuilder.New(3, 2, 2, "cmt 3").Build(),
		4: cmtBuilder.New(4, 3, 2, "cmt 4").Parent(3).Build(),
		5: cmtBuilder.New(5, 1, 2, "cmt 5").Parent(4).Build(),
	}

	m.PostEmptyId = 3
	m.PostSingleId = 1
	m.PostNestedId = 2

	m.PostEmpty = "[]"
	singleCmt, _ := json.Marshal(m.GetByPost(m.PostSingleId))
	m.PostSingle = string(singleCmt)
	m.PostNested = `
[
  {
    "id": 2, "userId": 1, "postId": 2, "parentId": 0, "content": "cmt 2", "created": ""
  },
  {
    "id": 3, "userId": 2, "postId": 2, "parentId": 0, "content": "cmt 3", "created": "",
    "children": [
      {
        "id": 4, "userId": 3, "postId": 2, "parentId": 3, "content": "cmt 4", "created": "",
        "children": [
          {
            "id": 5, "userId": 1, "postId": 2, "parentId": 4, "content": "cmt 5", "created": ""
          }
        ]
      }
    ]
  }
]
`
}

func (m *CmtMother) Get(cmtId int) model.Comment {
	return m.data[cmtId]
}

func (m *CmtMother) GetByPost(postId int) (res []model.Comment) {
	for _, e := range m.data {
		if e.PostId == postId {
			res = append(res, e)
		}
	}
	return
}
