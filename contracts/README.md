# truffle/hdwallet-provider 설치 가이드
```bash   
- npm install @truffle/hdwallet-provider
- npm install -g truffle
```

## 프로젝트 디렉토리 생성 및 초기화 
1. **폴더 생성**
- mkdir contracts
- cd contracts

2. **Truffle 초기화**
- truffle init

### truffle-config.js 설정
#### 프로젝트의 루트 디렉토리에 위치한 truffle-config.js 파일을 열고, 프로젝트의 요구사항에 맞게 설정

## 주요 Truffle 명령어

### 컴파일
```bash
- truffle compile
```

### 배포 (여기서 <network-name>은 truffle-config.js 파일에 설정된 네트워크의 이름입니다.)
```bash
truffle migrate --network <network-name>
```

### 테스트
```bash
truffle test
```

### 콘솔
```bash
truffle console
```

## 주의사항
### - truffle-config.js 파일은 프로젝트의 중심이므로, 이 파일의 설정이 모든 Truffle 명령어의 동작 방식을 결정합니다.
### - 테스트넷이나 메인넷에 배포하기 전에 항상 로컬 환경에서 충분히 테스트하는 것이 좋습니다.
### - 환경 변수를 사용하여 중요한 정보(예: 비밀 키)를 관리하는 것이 좋습니다.