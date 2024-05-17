
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE posts (
    post_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    title VARCHAR(255) NOT NULL,
    content TEXT NOT NULL,
    author VARCHAR(100) NOT NULL,
    publication_date DATE NOT NULL,
    tags VARCHAR(100) NOT NULL
);
