/***********************************************************************************
 *        File Write 예제
 ************************************************************************************/
//*

package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Create("text01.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer file.Close() // main 함수가 끝나기 직전에 파일을 닫음

	s := "Hello, world!"

	n, err := file.Write([]byte(s)) // s를 []byte 바이트 슬라이스로 변환, s를 파일에 저장
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(n, "바이트 저장 완료")
}

// */

/************************************************************************************
 *        이전 예제 : file.Write 대신 fmt.Fprint 함수 사용 예제.
 ************************************************************************************/
/*
//
package main

import (
	"fmt"
	"os"
)

func main() {
	file1, _ := os.Create("old1.txt")        // hello1.txt 파일 생성
	defer file1.Close()                        // main 함수가 끝나기 직전에 파일을 닫음
	fmt.Fprint(file1, 1, 1.1, "Hello, world!") // 값을 그대로 문자열로 만든 뒤 파일에 저장

	file2, _ := os.Create("old2.txt")          // hello2.txt 파일 생성
	defer file2.Close()                          // main 함수가 끝나기 직전에 파일을 닫음
	fmt.Fprintln(file2, 1, 1.1, "Hello, world!") // 값을 그대로 문자열로 만든 뒤
	                                             // 문자열 끝에 개행 문자(\n)를 붙이고 파일에 저장

	file3, _ := os.Create("old3.txt")                     // hello3.txt 파일 생성
	defer file3.Close()                                     // main 함수가 끝나기 직전에 파일을 닫음
	fmt.Fprintf(file3, "%d,%f,%s", 1, 1.1, "Hello, world!") // 형식을 지정하여 파일에 저장
}
// */

/************************************************************************************
 *        Open & Read & Stat & Size
 ************************************************************************************/
/*
package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("text01.txt") // hello.txt 파일을 열기
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close() // main 함수가 끝나기 직전에 파일을 닫음

	fi, err := file.Stat() // 파일 정보를 가져오기
	if err != nil {
		fmt.Println(err)
		return
	}

	var data = make([]byte, fi.Size()) // 파일 크기만큼 바이트 슬라이스 생성
	// var data = make([]byte, 10) // 10만큼?

	n, err := file.Read(data) // 파일의 내용을 읽어서 바이트 슬라이스에 저장
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(n, "바이트 읽기 완료")
	fmt.Println(string(data)) // 문자열로 변환하여 data의 내용 출력
}

// */

/************************************************************************************
 *        이전 Read 예제
 ************************************************************************************/
/*
package main

import (
	"fmt"
	"os"
)

func main() {
	var num1 int
	var num2 float32
	var s string

	file1, _ := os.Open("old1.txt")          // hello1.txt 파일 열기
	defer file1.Close()                        // main 함수가 끝나기 직전에 파일을 닫음
	n, _ := fmt.Fscan(file1, &num1, &num2, &s) // 파일을 읽은 뒤 공백, 개행 문자로
	                                           // 구분된 문자열에서 입력을 받음
	fmt.Println("입력 개수:", n)                // 입력 개수: 3
	fmt.Println(num1, num2, s)                 // 1 1.1 Hello

	file2, _ := os.Open("old2.txt")    // hello2.txt 파일 열기
	defer file2.Close()                  // main 함수가 끝나기 직전에 파일을 닫음
	fmt.Fscanln(file2, &num1, &num2, &s) // 파일을 읽은 뒤 공백으로
	                                     // 구분된 문자열에서 입력을 받음
	fmt.Println("입력 개수:", n)          // 입력 개수: 3
	fmt.Println(num1, num2, s)           // 1 1.1 Hello

	file3, _ := os.Open("old3.txt")               // hello3.txt 파일 열기
	defer file3.Close()                             // main 함수가 끝나기 직전에 파일을 닫음
	fmt.Fscanf(file3, "%d,%f,%s", &num1, &num2, &s) // 파일을 읽은 뒤 문자열에서
	                                                // 형식을 지정하여 입력을 받음
	fmt.Println("입력 개수:", n)                     // 입력 개수: 3
	fmt.Println(num1, num2, s)                      // 1 1.1 Hello
}
// */

/************************************************************************************
 *        OpenFile & FileMode & Stat & Seek
 ************************************************************************************/
/*

package main

import (
	"fmt"
	"os"
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

	n := 0
	s := "안녕하세요"
	n, err = file.Write([]byte(s)) // s를 []byte 바이트 슬라이스로 변환, s를 파일에 저장
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(n, "바이트 저장 완료")

	fi, err := file.Stat() // 파일 정보 가져오기
	if err != nil {
		fmt.Println(err)
		return
	}

	var data = make([]byte, fi.Size()) // 파일 크기만큼 바이트 슬라이스 생성

	// file.Seek(-12, os.SEEK_END)
	file.Seek(0, os.SEEK_SET) // Write 함수로 인해 파일 읽기/쓰기 위치가
	// 이동했으므로 file.Seek 함수로 읽기/쓰기 위치를
	// 파일의 맨 처음(0)으로 이동
	n, err = file.Read(data) // 파일의 내용을 읽어서 바이트 슬라이스에 저장
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(n, "바이트 읽기 완료")
	fmt.Println(string(data)) // 문자열로 변환하여 data의 내용 출력
}

// */

/************************************************************************************
 *        ioUtil & WriteFile & ReadFile
 ************************************************************************************/
/*

package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	s := "Hello, world!"

	err := ioutil.WriteFile("text03.txt", []byte(s), os.FileMode(644))
                                                  // s를 []byte 바이트 슬라이스로 변환, s를 hello.txt 파일에 저장
	if err != nil {
		fmt.Println(err)
		return
	}

	data, err := ioutil.ReadFile("hello.txt") // hello.txt의 내용을 읽어서 바이트 슬라이스 리턴
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(data)) // 문자열로 변환하여 data의 내용 출력
}
// */
