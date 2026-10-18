# Open source project analyzer 
Сервис для анализа GitHub-репозиториев с помощью AI

Собирает метаданные репозитория через GitHub API, делает оценку по 4 критериям (activity, documentation, community, maintenance) и проводит AI-анализ через Groq

## Стек

- **Go 1.26** + Gin 
- **PostgreSQL** 
- **Groq API** AI-анализ (модель `openai/gpt-oss-120b`)
- **Docker** + docker-compose
- **golang-migrate** 

## Возможности
* Парсинг URL GitHub-репозитория
* Fetch метаданных через GitHub API
* Оценка по 4 критериям (activity, documentation, community, maintenance)
* AI-анализ (summary, scores, strengths, weaknesses, improvements)
* Комбинированный отчёт (среднее из AI оценки + строго прописанной оценки)

## Архитектура
Handler -> Service -> Repository

## Быстрый старт 
### Клонирование и настройка
1. git clone https://github.com/slinkeres/open-source-project-analyzer
2. cd open-source-project-analyzer
3. Скопируй env.example в .env (cp env.example .env)
Открой .env и заполни:
GROQ_API_KEY — получи на console.groq.com/keys
Остальные значения можно оставить по умолчанию
4. Запускаем контейнеры
```bash
docker-compose up --build
```


## API Endpoints
### GET
* /health - проверка работоспособности
* /api/repositories - получить список репозиториев
* /api/repositories/:id - получить репозиторий по id
* /api/repositories/:id/score - получить оценку репозитория по id
* /api/repositories/:id/ai-analysis - получить сохраненный AI анализ через GROQ
* /api/repositories/:id/report - получить отчет

### POST
* /api/repositories/preview - парсинг URL
* /api/repositories/fetch - загрузка репозитория с GitHub
* /api/repositories/:id/score - рассчитать оценку
* /api/repositories/:id/ai-analysis - запустить AI анализ через GROQ

## Тесты 
```bash
go test ./... -v
```

## Примеры запросов
* curl http://localhost:8080/health
Ответ : {"status":"ok"}

* curl -X POST http://localhost:8080/api/repositories/preview \
  -H "Content-Type: application/json" \
  -d '{"url":"https://github.com/gin-gonic/gin"}'

Ответ: 
```
{
    "owner": "gin-gonic",
    "name": "gin",
    "normalized_url": "https://github.com/gin-gonic/gin"
}
```
* curl http://localhost:8080/api/repositories
Ответ:
```
[
    {
        "database_id": 1,
        "full_name": "gin-gonic/gin",
        "description": "Gin is a high-performance HTTP web framework written in Go. It provides a Martini-like API but with significantly better performance—up to 40 times faster—thanks to httprouter. Gin is designed for building REST APIs, web applications, and microservices.",
        "stars": 89291,
        "forks": 8734,
        "open_issues": 791,
        "language": "Go",
        "license": "MIT",
        "archived": false,
        "default_branch": "master"
    }
]
```

* curl -X POST http://localhost:8080/api/repositories/1/ai-analysis
Ответ:
```
{
    "repository_id": 1,
    "summary": "Gin is a high-performance Go web framework offering a Martini-like API with a zero-allocation router, extensive middleware support, and strong benchmarking results. It provides clear installation steps, a concise first‑app example, multilingual documentation, and a vibrant ecosystem of community middleware and production adopters.",
    "documentation_quality_score": 85,
    "onboarding_score": 80,
    "beginner_friendliness_score": 75,
    "strengths": [
        "Exceptional request routing performance with zero heap allocations",
        "Rich middleware ecosystem via gin-contrib and community packages",
        "Comprehensive README with quick‑start example and links to full docs",
        "Multilingual official documentation and tutorials",
        "Active community and many high‑traffic production users",
        "Clear versioning and release notes"
    ],
    "weaknesses": [
        "Limited out‑of‑the‑box features compared to full‑stack frameworks",
        "Requires solid Go knowledge, which can be a barrier for newcomers",
        "Occasional breaking API changes across major releases",
        "Documentation, while extensive, can be fragmented across multiple sites",
        "Test coverage could be higher for core components"
    ],
    "suggested_improvements": [
        "Add more step‑by‑step beginner tutorials and video guides",
        "Provide a detailed migration guide for major version upgrades",
        "Include additional built‑in utilities (e.g., request validation, structured error responses)",
        "Consolidate documentation into a single, searchable portal",
        "Increase core library test coverage and publish coverage metrics"
    ],
    "model": "openai/gpt-oss-120b",
    "analyzed_at": "2026-10-05T18:20:23.353345Z"
}
```