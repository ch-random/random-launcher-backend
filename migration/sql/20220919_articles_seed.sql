-- users
DESCRIBE users;

INSERT INTO
  `users` (google_id, role, name)
VALUES
  ('LyC0r1s', 'user', 'Chisato');

SELECT
  *
FROM
  users;

-- articles
DESCRIBE articles;

INSERT INTO
  `articles` (user_id, title, body, public)
VALUES
  (
    1,
    'Vue + TailwindCSS Hands-on',
    '# Intro ...',
    (TRUE)
  );

SELECT
  *
FROM
  articles;

-- article_game_contents
DESCRIBE article_game_contents;

INSERT INTO
  `article_game_contents` (id, exec_path, zip_url)
VALUES
  (
    1,
    'Neatly.exe',
    'https://drive.google.com/uc?export=download&id=1Sn08keQU9eSbGDvYja8_PhBPiegJ99V_'
  );

SELECT
  *
FROM
  article_game_contents;

-- article_owners
DESCRIBE article_owners;

INSERT INTO
  `article_owners` (id, article_id)
VALUES
  (1, 1);

SELECT
  *
FROM
  article_owners;

-- article_tags
DESCRIBE article_tags;

INSERT INTO
  `article_tags` (id, article_id, name)
VALUES
  (1, 1, 'アクション');

SELECT
  *
FROM
  article_tags;

-- article_comments
DESCRIBE article_comments;

INSERT INTO
  `article_comments` (id, article_id, body, rate)
VALUES
  (
    1,
    1,
    "What a fucking game! I'll never play it again!",
    1
  );

SELECT
  *
FROM
  article_comments;

-- article_image_urls
DESCRIBE article_image_urls;

INSERT INTO
  `article_image_urls` (id, article_id, image_url)
VALUES
  (
    1,
    1,
    'https://www.famitsu.com/images/000/260/379/y_6274fb3411310.jpg'
  );

SELECT
  *
FROM
  article_image_urls;
