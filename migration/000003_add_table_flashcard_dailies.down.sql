DROP TABLE flashcard_dailies;

ALTER TABLE user_daily_word_statistics DROP COLUMN date;

ALTER TABLE user_flashcard_logs DROP COLUMN date;

ALTER TABLE vocabularies DROP COLUMN topic;