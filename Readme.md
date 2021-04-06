## Go Modules 

#### 모듈 초기화
```bash
 $ go mod init <module_name>
```
&nbsp;현재 디렉터리는 모듈이 아닌 패키지를 포함하고 있다. 그 이유는 아직 go.mod 파일이 존재하지 않기 때문이다.
go mod init 명령어를 실행하여, 현재 디렉터리를 모듈의 루트로 만들고, go test를 실행한다.

- go.mod 파일
    의존하는 모듈과 버전을 명시하는 파일
- go.sum 파일
    모듈의 파일이 바뀌지 않았음을 확인하기 위해 저장하는 첵썸 파일
    go 명령이 go.mod 에 정의된 모듈을 다운로드할 때 각 모듈의 첵썸값을 go.sum 파일에 자동 기록
    이를 첵썸 디비(아래 설명)와 대조, 모듈 내용의 변경 여부를 확인

*go.mod 파일과 go.sum 파일은 소스관리 대상*


#### 
```bash
 $ go mod vendor 
```
&nbsp;go mod vendor 명령어로 의존성 패키지 vendor 디렉토리 만들수 있으며, 해당 디렉토리에는 go.mod 에 명시된 패키지들이 다운로드 되어 저장된다. vendor 폴더에는 modules.txt 라는 파일이 생성되고, 상세한 하위 패키지와 디렉토리가 기술되어 저장 된다.

