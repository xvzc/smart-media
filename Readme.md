# go mod init <module_name>
현재 디렉터리는 모듈이 아닌 패키지를 포함하고 있다. 그 이유는 아직 go.mod 파일이 존재하지 않기 때문이다.
go mod init 명령어를 실행하여, 현재 디렉터리를 모듈의 루트로 만들고, go test를 실행한다.
# go mod vendor 
패키지를 모두 vendor 폴더에 설치한다.