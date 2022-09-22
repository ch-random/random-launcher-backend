-- github.com/google/uuid: 8-4-4-4-12
-- users
DESCRIBE users;

INSERT INTO
  `users` (id, google_id, role, name)
VALUES
  (
    "aid3x0lv-zotm-8kjg-z4bh-ovrq3khodmf5",
    "gidhrldl3o2t0afkj039eeu7xqj2",
    "user",
    "錦木 千束"
  );

SELECT
  *
FROM
  users;

-- articles
DESCRIBE articles;

INSERT INTO
  `articles` (id, user_id, title, body, public)
VALUES
  (
    "aid3x0lv-zotm-8kjg-z4bh-ovrq3khodmf5",
    "uidp9zck-tz6h-ouz6-f4wj-o0k2kedvnv2q",
    "Vue + TailwindCSS Hands-on",
    '# Intro
1. _
2. _
3. _',
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
    "aid3x0lv-zotm-8kjg-z4bh-ovrq3khodmf5",
    "Neatly.exe",
    "https://drive.google.com/uc?export=download&id=1Sn08keQU9eSbGDvYja8_PhBPiegJ99V_"
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
  (
    "uidp9zck-tz6h-ouz6-f4wj-o0k2kedvnv2q",
    "aid3x0lv-zotm-8kjg-z4bh-ovrq3khodmf5"
  );

SELECT
  *
FROM
  article_owners;

-- article_tags
DESCRIBE article_tags;

INSERT INTO
  `article_tags` (id, article_id, name)
VALUES
  (
    "tidl674q-2nqi-wrwr-bc86-xt2rmxkq0s76",
    "aid3x0lv-zotm-8kjg-z4bh-ovrq3khodmf5",
    "アクション"
  );

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
    "cidmo6ep-tjtw-utsb-y3so-8vwi1lw3d3sf",
    "aid3x0lv-zotm-8kjg-z4bh-ovrq3khodmf5",
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
    "iuidczlt-t24m-pgx9-80es-r8sz1r26afa3",
    "aid3x0lv-zotm-8kjg-z4bh-ovrq3khodmf5",
    "https://www.famitsu.com/images/000/260/379/y_6274fb3411310.jpg"
  );

SELECT
  *
FROM
  article_image_urls;
