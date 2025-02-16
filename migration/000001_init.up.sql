CREATE TABLE users (
    id integer PRIMARY KEY AUTO_INCREMENT,
    username varchar(255) NOT NULL,
    password varchar(255) NOT NULL,
    role varchar(255) NOT NULL DEFAULT 'user',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE categories (
    id integer PRIMARY KEY AUTO_INCREMENT,
    name varchar(255) NOT NULL,
    user_id integer DEFAULT -1,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE vocabularies (
    id integer PRIMARY KEY AUTO_INCREMENT,
    word varchar(255) NOT NULL,
    meaning varchar(255) NOT NULL,
    ipa varchar(255),
    type varchar(255),
    description text,
    url varchar(255),
    category_id int DEFAULT -1,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE examples (
    id integer PRIMARY KEY AUTO_INCREMENT,
    vocabulary_id integer DEFAULT -1,
    sentence varchar(255) NOT NULL,
    meaning varchar(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE user_flashcard_logs (
    id integer PRIMARY KEY AUTO_INCREMENT,
    user_id integer DEFAULT -1,
    vocabulary_id integer DEFAULT -1,
    answer varchar(255),
    is_correct bool DEFAULT false,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE user_daily_word_statistics (
    id integer PRIMARY KEY AUTO_INCREMENT,
    user_id integer DEFAULT -1,
    correct_answers integer,
    wrong_answers integer,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);