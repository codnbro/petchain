# go-multibase 설치 가이드

이 문서는 `github.com/multiformats/go-multibase` 패키지를 Go 프로젝트에 설치하는 방법

## 설치 단계

### 1. Go 모듈 초기화

`go-multibase` 패키지를 설치하기 전에, 프로젝트 디렉토리가 Go 모듈로 초기화되어 있는지 확인하세요. 프로젝트가 아직 Go 모듈이 아닌 경우, 프로젝트의 루트 디렉토리에서 다음 명령어를 실행

```shell
go mod init github.com/multiformats/go-multibase`
```
### 2. 패키지 설치

Go 모듈이 초기화되면, 다음 명령어를 사용하여 `go-multibase` 패키지를 설치

```shell
go get github.com/multiformats/go-multibase
```

### 3. 패키지 사용

go-multibase 패키지를 성공적으로 설치한 후, 프로젝트에서 다음과 같이 임포트하고 사용

```go
import "github.com/multiformats/go-multibase"
```

### 4. 패키지 문제 발생 시

 Go 모듈 캐시를 지우고 다시 시도

```shell
 go clean -modcache
```