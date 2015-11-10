/************************************************************************************
 *  io.Reader, io.Write 인터페이스 활용
 ************************************************************************************/
/*

type Reader interface {
        Read(p []byte) (n int, err error)
}

type Writer interface {
        Write(p []byte) (n int, err error)
}

// */

/************************************************************************************
 *  bufio methods
 ************************************************************************************/
/*

func NewReader(rd io.Reader) *Reader //: io.Reader 인터페이스로 io.Reader 인터페이스 따르는 읽기 인스턴스 생성
func NewWriter(w io.Writer) *Writer //: io.Writer 인터페이스로 io.Writer 인터페이스를 따르는 쓰기 인스턴스 생성
func (b *Writer) WriteString(s string) (int, error) //: 문자열을 버퍼에 저장
func (b *Writer) Flush() error //: 버퍼의 데이터를 파일에 저장

// */

/************************************************************************************
 *  bufio 예제
 ************************************************************************************/
//*
package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
)

func main() {
	file, err := os.OpenFile(
		"text01.txt",
		os.O_CREATE|os.O_RDWR|os.O_TRUNC, // 파일이 없으면 생성,
		// 읽기/쓰기, 파일을 연 뒤 내용 삭제
		os.FileMode(0644), // 파일 권한은 644
	)

	fmt.Println(reflect.TypeOf(file))
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close() // main 함수가 끝나기 직전에 파일을 닫음

	w := bufio.NewWriter(file) // io.Writer 인터페이스를 따르는 file로
	// io.Writer 인터페이스를 따르는 쓰기 인스턴스 w 생성

	fmt.Println(reflect.TypeOf(w))

	w.WriteString("Hello, world!") // 쓰기 인스턴스로 버퍼에 Hello, world! 쓰기
	w.Flush()                      // 버퍼의 내용을 파일에 저장

	r := bufio.NewReader(file) // io.Reader 인터페이스를 따르는 file로
	// io.Reader 인터페이스를 따르는 읽기 인스턴스 r 생성
	fi, _ := file.Stat()         // 파일 정보 구하기
	b := make([]byte, fi.Size()) // 파일 크기만큼 바이트 슬라이스 생성

	file.Seek(0, os.SEEK_SET) // 파일 읽기 위치를 파일의 맨 처음(0)으로 이동
	r.Read(b)                 // 읽기 인스턴스로 파일의 내용을 읽어서 b에 저장

	fmt.Println(string(b)) // 문자열로 변환하여 b의 내용 출력
}

// */

/************************************************************************************
 *	문자열 읽고 파일로 (ReadFrom)
 ************************************************************************************/
/*
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	_ "io"
)

func main() {
	file, err := os.OpenFile(
		"text02.txt",
		os.O_CREATE|os.O_RDWR|os.O_TRUNC, // 파일이 없으면 생성,
		// 읽기/쓰기, 파일을 연 뒤 내용 삭제
		os.FileMode(0644)) // 파일 권한은 644
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close() // main 함수가 끝나기 직전에 파일을 닫음

	s := "Hello, world!"
	r := strings.NewReader(s) // 문자열로 io.Reader 인터페이스를 따르는
	// 읽기 인스턴스 r 생성

	w := bufio.NewWriter(file) // io.Writer 인터페이스를 따르는 file로
	// io.Writer 인터페이스를 따르는 쓰기 인스턴스 w 생성
	w.ReadFrom(r) // 읽기 인스턴스 r에서 데이터를 읽어서 w의 버퍼에 저장
	// io.Copy(w, r) // **** 위와 동일 ****

	w.Flush()     // 버퍼의 내용을 파일에 저장
}

// */

/************************************************************************************
 *	문자열을 화면에... io.Reader를 os.Stdout으로
 ************************************************************************************/
/*
package main

import (
	"io"
	"os"
	"strings"
)

func main() {
	s := "Hello, world!"
	r := strings.NewReader(s) // 문자열로 io.Reader 인터페이스를 따르는
                                  // 읽기 인스턴스 r 생성

	io.Copy(os.Stdout, r) // os.Stdout에 io.Reader를 복사하면 화면에 그대로 출력됨
}
// */

/************************************************************************************
 *	기본 입출력 함수 + 입출력 인터페이스
 ************************************************************************************/
/*
package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "Hello, world!"
	r := strings.NewReader(s) // 문자열로 io.Reader 인터페이스를 따르는
	// 읽기 인스턴스 r 생성

	var s1, s2 string
	n, _ := fmt.Fscanf(r, "%s %s", &s1, &s2) // 형식을 지정하여 읽기 인스턴스 r에서
	// 문자열 읽기

	fmt.Println("입력 개수:", n) // 입력 개수: 2
	fmt.Println(s1)          // Hello,
	fmt.Println(s2)          // world!
}

// */

/************************************************************************************
 *	bufio + 기본입출력함수
 ************************************************************************************/
/*
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.OpenFile(
		"hello.txt",
		os.O_CREATE|os.O_RDWR|os.O_TRUNC, // 파일이 없으면 생성,
		// 읽기/쓰기, 파일을 연 뒤 내용 삭제
		os.FileMode(0644)) // 파일 권한은 644
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close() // main 함수가 끝나기 직전에 파일을 닫음

	w := bufio.NewWriter(file) // file로 io.Writer 인터페이스를 따르는
	// 쓰기 인스턴스 w 생성
	fmt.Fprintf(w, "%d,%f,%s", 1, 1.1, "Hello") // 형식을 지정하여 쓰기 인스턴스 w의 버퍼에
	// 1, 1.1, Hello 저장
	w.Flush() // 버퍼의 내용을 파일에 저장
}

// */

/************************************************************************************
 *	io.ReadWriter + bufio.ReadLine
 ************************************************************************************/
/*
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.OpenFile(
		"hello.txt",
		os.O_CREATE|os.O_RDWR|os.O_TRUNC, // 파일이 없으면 생성,
		// 읽기/쓰기, 파일을 연 뒤 내용 삭제
		os.FileMode(0644)) // 파일 권한은 644
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close() // main 함수가 끝나기 직전에 파일을 닫음

	r := bufio.NewReader(file) // file로 io.Reader 인터페이스를 따르는 읽기 인스턴스 r 생성
	w := bufio.NewWriter(file) // file로 io.Writer 인터페이스를 따르는 쓰기 인스턴스 w 생성

	rw := bufio.NewReadWriter(r, w) // r, w를 사용하여 io.ReadWriter 인터페이스를 따르는
	// 읽기/쓰기 인스턴스 생성
	rw.WriteString("Hello, world!") // 읽기/쓰기 인스턴스로 버퍼에 Hello, world! 쓰기
	rw.Flush()                      // 버퍼의 내용을 파일에 저장

	file.Seek(0, os.SEEK_SET)   // 파일 읽기 위치를 파일의 맨 처음(0)으로 이동
	data, _, _ := rw.ReadLine() // 파일에서 문자열 한 줄을 읽어서 data에 저장
	fmt.Println(string(data))   // 문자열로 변환하여 data의 내용 출력
}

// */
