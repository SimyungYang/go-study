/************************************************************************************
 *  json unmarshall
 ************************************************************************************/
/*

package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	doc := `
	{
		"name": "maria",
		"age": 10
	}
	`

	var data map[string]interface{} // JSON 문서의 데이터를 저장할 공간을 맵으로 선언

	json.Unmarshal([]byte(doc), &data) // doc를 바이트 슬라이스로 변환하여 넣고,
                                           // data의 포인터를 넣어줌

	fmt.Println(data["name"], data["age"]) // maria 10: 맵에 키를 지정하여 값을 가져옴
}

// */

/************************************************************************************
 *  json marshall
 ************************************************************************************/
/*
package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	data := make(map[string]interface{}) // 문자열을 키로하고 모든 자료형을 저장할 수 있는 맵 생성

	data["name"] = "maria"
	data["age"] = 10

	// doc, _ := json.Marshal(data)                // 맵을 JSON 문서로 변환
	doc, _ := json.MarshalIndent(data, "", "    ") // 맵을 JSON 문서로 변환

	fmt.Println(string(doc)) // {"age":10,"name":"maria"}: 문자열로 변환하여 출력
}

// */

/************************************************************************************
 *  구조체 json (json -> data)
 ************************************************************************************/
/*
package main

import (
	"encoding/json"
	"fmt"
)

type Author struct {
	Name  string
	Email string
}

type Comment struct {
	Id      uint64
	Author  Author // Author 구조체
	Content string
}

type Article struct {
	Id         uint64
	Title      string
	Author     Author // Author 구조체
	Content    string
	Recommends []string  // 문자열 배열
	Comments   []Comment // Comment 구조체 배열
}

func main() {
	doc := `
	[{
		"Id": 1,
		"Title": "Hello, world!",
		"Author": {
			"Name": "Maria",
			"Email": "maria@example.com"
		},
		"Content": "Hello~",
		"Recommends": [
			"John",
			"Andrew"
		],
		"Comments": [{
			"id": 1,
			"Author": {
				"Name": "Andrew",
				"Email": "andrew@hello.com"
			},
			"Content": "Hello Maria"
		}]
	}]
	`

	var data []Article // JSON 문서의 데이터를 저장할 구조체 슬라이스 선언

	json.Unmarshal([]byte(doc), &data) // doc의 내용을 변환하여 data에 저장

	fmt.Println(data) // [{1 Hello, world! {Maria maria@exa... (생략)
}

// */

/************************************************************************************
 *  data -> json
 ************************************************************************************/
/*

package main

import (
	"encoding/json"
	"fmt"
)

type Author struct {
	Name  string
	Email string
}

type Comment struct {
	Id      uint64
	Author  Author // Author 구조체
	Content string
}

type Article struct {
	Id         uint64
	Title      string
	Author     Author // Author 구조체
	Content    string
	Recommends []string  // 문자열 배열
	Comments   []Comment // Comment 구조체 배열
}

func main() {
	data := make([]Article, 1) // 값을 저장할 구조체 슬라이스 생성

	data[0].Id = 1
	data[0].Title = "Hello, world!"
	data[0].Author.Name = "Maria"
	data[0].Author.Email = "maria@example.com"
	data[0].Content = "Hello~"
	data[0].Recommends = []string{"John", "Andrew"}
	data[0].Comments = make([]Comment, 1)
	data[0].Comments[0].Id = 1
	data[0].Comments[0].Author.Name = "Andrew"
	data[0].Comments[0].Author.Email = "andrew@hello.com"
	data[0].Comments[0].Content = "Hello Maria"

	doc, _ := json.Marshal(data) // data를 JSON 문서로 변환

	fmt.Println(string(doc)) // [{"Id":1,"Title":"Hello, world!","Au... (생략)
}

// */

/************************************************************************************
 *  json lowercase
 ************************************************************************************/
/*

package main

import (
	"encoding/json"
	"fmt"
)

type Author struct {
	Name  string `json:"name"` // 구조체 필드에 태그 지정
	Email string `json:"email"`
}

type Comment struct {
	Id      uint64 `json:"id"`
	Author  Author `json:"author"`
	Content string `json:"content"`
}

type Article struct {
	Id         uint64    `json:"id"`
	Title      string    `json:"title"`
	Author     Author    `json:"author"`
	Content    string    `json:"content"`
	Recommends []string  `json:"recommends"`
	Comments   []Comment `json:"comments"`
}

func main() {
	data := make([]Article, 1) // 값을 저장할 구조체 슬라이스 생성

	data[0].Id = 1
	data[0].Title = "Hello, world!"
	data[0].Author.Name = "Maria"
	data[0].Author.Email = "maria@example.com"
	data[0].Content = "Hello~"
	data[0].Recommends = []string{"John", "Andrew"}
	data[0].Comments = make([]Comment, 1)
	data[0].Comments[0].Id = 1
	data[0].Comments[0].Author.Name = "Andrew"
	data[0].Comments[0].Author.Email = "andrew@hello.com"
	data[0].Comments[0].Content = "Hello Maria"

	doc, _ := json.Marshal(data) // data를 JSON 문서로 변환

	fmt.Println(string(doc)) // [{"Id":1,"Title":"Hello, world!","Au... (생략)
}

// */

/************************************************************************************
 *  json + writeFile
 ************************************************************************************/
/*
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Author struct {
	Name  string `json:"name"` // 구조체 필드에 태그 지정
	Email string `json:"email"`
}

type Comment struct {
	Id      uint64 `json:"id"`
	Author  Author `json:"author"`
	Content string `json:"content"`
}

type Article struct {
	Id         uint64    `json:"id"`
	Title      string    `json:"title"`
	Author     Author    `json:"author"`
	Content    string    `json:"content"`
	Recommends []string  `json:"recommends"`
	Comments   []Comment `json:"comments"`
}

func main() {
	data := make([]Article, 1) // 값을 저장할 구조체 슬라이스 생성

	data[0].Id = 1
	data[0].Title = "Hello, world!"
	data[0].Author.Name = "Maria"
	data[0].Author.Email = "maria@example.com"
	data[0].Content = "Hello~"
	data[0].Recommends = []string{"John", "Andrew"}
	data[0].Comments = make([]Comment, 1)
	data[0].Comments[0].Id = 1
	data[0].Comments[0].Author.Name = "Andrew"
	data[0].Comments[0].Author.Email = "andrew@hello.com"
	data[0].Comments[0].Content = "Hello Maria"

	doc, _ := json.Marshal(data) // data를 JSON 문서로 변환

	err := ioutil.WriteFile("./articles.json", doc, os.FileMode(0644)) // articles.json 파일에 JSON 문서 저장
	if err != nil {
		fmt.Println(err)
		return
	}
}

// */

/************************************************************************************
 *  json + readFile
 ************************************************************************************/
/*
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Author struct {
	Name  string `json:"name"` // 값을 저장할 구조체 슬라이스 생성
	Email string `json:"email"`
}

type Comment struct {
	Id      uint64 `json:"id"`
	Author  Author `json:"author"`
	Content string `json:"content"`
}

type Article struct {
	Id         uint64    `json:"id"`
	Title      string    `json:"title"`
	Author     Author    `json:"author"`
	Content    string    `json:"content"`
	Recommends []string  `json:"recommends"`
	Comments   []Comment `json:"comments"`
}

func main() {
	b, err := ioutil.ReadFile("./articles.json") // articles.json 파일의 내용을 읽어서 바이트 슬라이스에 저장
	if err != nil {
		fmt.Println(err)
		return
	}

	var data []Article // JSON 문서의 데이터를 저장할 구조체 슬라이스 선언

	json.Unmarshal(b, &data) // JSON 문서의 내용을 변환하여 data에 저장

	fmt.Println(data) // [{1 Hello, world! {Maria maria@exa... (생략)
}

// */



/************************************************************************************
 *  json marshall
 ************************************************************************************/
/*


// */



/************************************************************************************
 *  json marshall
 ************************************************************************************/
/*


// */



/************************************************************************************
 *  json marshall
 ************************************************************************************/
/*


// */
