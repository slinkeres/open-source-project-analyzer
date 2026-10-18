CREATE TABLE IF NOT EXISTS repositories (
    id BIGSERIAL PRIMARY KEY,
    github_id BIGINT NOT NULL UNIQUE,
    owner TEXT NOT NULL,
    name TEXT NOT NULL,
    full_name TEXT NOT NULL UNIQUE,
    description TEXT,
    stars BIGINT NOT NULL DEFAULT 0,
    forks BIGINT NOT NULL DEFAULT 0,
    open_issues BIGINT NOT NULL DEFAULT 0,
    subscribers BIGINT NOT NULL DEFAULT 0,
    language TEXT,
    license TEXT,
    archived BOOLEAN NOT NULL DEFAULT FALSE,
    has_issues BOOLEAN NOT NULL DEFAULT TRUE,
    url TEXT NOT NULL,
    default_branch TEXT NOT NULL,
    topics TEXT[] NOT NULL DEFAULT '{}',
    github_created_at TEXT NOT NULL,
    github_updated_at TEXT NOT NULL,
    pushed_at TEXT NOT NULL,
    readme TEXT NOT NULL DEFAULT '',
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);


CREATE TABLE IF NOT EXISTS repository_scores (
    id BIGSERIAL PRIMARY KEY,
    repository_id BIGINT NOT NULL UNIQUE
    REFERENCES repositories(id)
    ON DELETE CASCADE,
    activity_score INTEGER NOT NULL,
    documentation_score INTEGER NOT NULL,
    community_score INTEGER NOT NULL,
    maintenance_score INTEGER NOT NULL,
    overall_score INTEGER NOT NULL,
    reasons TEXT[] NOT NULL DEFAULT '{}',
    calculated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS repository_ai_analysis(
    id BIGSERIAL PRIMARY KEY,                                              
    repository_id BIGINT NOT NULL UNIQUE
    REFERENCES repositories(id)
    ON DELETE CASCADE,
    summary TEXT NOT NULL ,
    documentation_quality_score INTEGER NOT NULL,
    onboarding_score INTEGER NOT NULL,
    beginner_friendliness_score INTEGER NOT NULL,
    strengths TEXT[] NOT NULL DEFAULT '{}',
    weaknesses TEXT[] NOT NULL DEFAULT '{}',
    suggested_improvements TEXT[] NOT NULL DEFAULT '{}',
    model TEXT NOT NULL,
    analyzed_at timestamptz not null default now()
);

