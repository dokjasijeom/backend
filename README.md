# Backend

독자시점 웹서비스의 백엔드입니다.

## 기술 스택

- **프로그래밍 언어:** Go (v1.23.4)
- **웹 프레임워크:** Fiber (v2.52.8)
- **데이터베이스:** MySQL
- **ORM:** GORM (v1.26.1)
- **인증:** JWT (JSON Web Tokens)
- **컨테이너화:** Docker (Alpine Linux 사용)
- **로깅:**
  - Axiom
  - Slog
  - Logrus
- **설정 관리:** Godotenv
- **이미지 관리:** Cloudinary
- **ID 해싱:** Hashids
- **기타 주요 라이브러리:**
  - `golang.org/x/crypto` (암호화)
  - `golang.org/x/image` (이미지 조작)
  - `github.com/samber/lo` (유틸리티 라이브러리)
  - `github.com/google/uuid` (UUID 생성)

## 시작하기

로컬에서 프로젝트를 실행하려면 다음 단계를 따르세요:

1.  **저장소 복제:**

    ```bash
    git clone <repository-url>
    cd backend
    ```

2.  **Go 설치:**
    Go 버전 1.23.4 이상이 설치되어 있는지 확인하세요. [golang.org](https://golang.org/dl/)에서 다운로드할 수 있습니다.

3.  **의존성 설치:**

    ```bash
    go mod download
    ```

    또는 모든 의존성을 가져오려면:

    ```bash
    go get ./...
    ```

4.  **환경 변수 설정:**
    프로젝트 루트에 `.env.example`을 복사하거나 (존재하는 경우) 수동으로 생성하여 `.env` 파일을 만듭니다. 이 파일에는 데이터베이스 자격 증명, JWT 비밀 키, Cloudinary 세부 정보 등과 같은 필요한 구성이 포함되어야 합니다.
    예제 `.env` 구조:

    ```env
    # 서버 설정
    PORT=8080
    APP_ENV=development

    # 데이터베이스 설정
    DB_HOST=localhost
    DB_PORT=3306
    DB_USER=your_db_user
    DB_PASSWORD=your_db_password
    DB_NAME=your_db_name

    # JWT 설정
    JWT_SECRET_KEY=your_jwt_secret
    JWT_EXPIRATION_HOURS=72

    # Cloudinary 설정
    CLOUDINARY_CLOUD_NAME=your_cloud_name
    CLOUDINARY_API_KEY=your_api_key
    CLOUDINARY_API_SECRET=your_api_secret

    # 필요에 따른 기타 설정
    ```

5.  **애플리케이션 실행:**

    ```bash
    go run main.go
    ```

    이제 애플리케이션이 실행되어야 하며, 일반적으로 `http://localhost:8080` (또는 `.env` 파일에 지정된 포트)에서 실행됩니다.

6.  **Docker 사용 (대안):**
    Docker를 사용하려면:

    ```bash
    # Docker 이미지 빌드
    docker build -t dokjasijeom-backend .

    # Docker 컨테이너 실행 (.env 파일이 구성되었는지 확인하거나 환경 변수 전달)
    docker run -p 8080:8080 --env-file .env dokjasijeom-backend
    ```

    `Dockerfile`이 환경 변수를 처리하고 필요한 포트를 노출하도록 올바르게 설정되었는지 확인하세요.

## API 엔드포인트

이 섹션은 포괄적인 API 문서(예: Swagger, Postman 컬렉션 또는 전용 API 문서 사이트)에 대한 링크를 이상적으로 제공해야 합니다.

주요 API 엔드포인트 그룹은 컨트롤러에서 관리합니다:

- `/users`: 사용자 관리, 인증, 프로필.
- `/series`: 시리즈 정보, 에피소드, 장르 등.
- `/categories`: 카테고리 관리.
- `/genres`: 장르 관리.
- `/providers`: 콘텐츠 제공자 정보.
- `/roles`: 사용자 역할 관리.
- `/search`: 검색 기능.
- `/backoffice`: 관리 작업을 위한 엔드포인트, 다음과 같은 하위 경로를 가집니다:
  - 장르
  - 인물 (작가/배우)
  - 제공자
  - 발행일
  - 출판사
  - 시리즈 관리

요청/응답 형식 및 특정 엔드포인트에 대한 자세한 내용은 API 문서를 참조하거나 소스 코드의 `controller` 디렉토리를 살펴보세요.

## 프로젝트 구조

이 프로젝트는 Go 애플리케이션에서 일반적인 계층형 아키텍처를 따라 관심사를 분리합니다:

```
.
├── Dockerfile                # 애플리케이션 컨테이너화를 위한 Docker 설정
├── go.mod                    # Go 모듈 정의 파일
├── go.sum                    # Go 모듈 체크섬
├── main.go                   # 주 애플리케이션 진입점
├── README.md                 # 이 파일
├── common/                   # 유틸리티 함수 및 공유 코드 (예: 해싱, JWT)
├── configuration/            # 애플리케이션 설정 (데이터베이스, Fiber, Cloudinary)
├── controller/               # 들어오는 HTTP 요청 처리, 요청 유효성 검사 및 응답 형식 지정
│   └── backoffice/           # 백오피스/관리 기능 관련 컨트롤러
├── entity/                   # 데이터베이스 테이블을 나타내는 GORM 모델 정의
├── exception/                # 사용자 정의 오류 처리 및 정의
├── middleware/               # HTTP 미들웨어 (예: JWT 권한 부여)
├── model/                    # 요청 및 응답을 위한 데이터 전송 객체 (DTO)
├── repository/               # 데이터 액세스 계층, 데이터베이스 상호 작용 추상화
│   └── impl/                 # 저장소 인터페이스의 구체적인 구현
├── service/                  # 비즈니스 로직 계층, 컨트롤러와 저장소 간의 작업 조정
│   └── impl/                 # 서비스 인터페이스의 구체적인 구현
└── ... (기타 파일 .gitignore, .dockerignore 등)
```

- **`main.go`**: Fiber 웹 서버를 초기화하고 시작하며, 경로를 설정하고 의존성을 연결합니다.
- **`common/`**: JWT 생성/유효성 검사, 암호 해싱 등과 같은 공유 유틸리티를 포함합니다.
- **`configuration/`**: 애플리케이션 설정, 데이터베이스 연결 및 타사 서비스 초기화를 관리합니다.
- **`controller/`**: HTTP 핸들러를 정의합니다. 각 컨트롤러는 일반적으로 리소스 또는 관련 기능 그룹에 해당합니다.
- **`entity/`**: 데이터베이스 테이블에 매핑되는 구조체 정의(GORM 모델)를 포함합니다.
- **`exception/`**: 사용자 정의 오류 유형 및 전역 오류 처리기를 정의합니다.
- **`middleware/`**: HTTP 요청에 적용되는 사용자 정의 미들웨어 함수(예: 인증, 로깅)를 보관합니다.
- **`model/`**: 데이터베이스 엔티티와 구별되는 API 요청 본문 및 응답 페이로드에 사용되는 구조체를 정의합니다.
- **`repository/`**: 데이터 소스를 추상화하는 데이터베이스 작업을 위한 인터페이스 및 구현입니다.
- **`service/`**: 컨트롤러와 저장소 간을 중재하는 핵심 비즈니스 로직을 포함합니다.
