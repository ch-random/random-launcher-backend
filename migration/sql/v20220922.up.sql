DROP TABLE IF EXISTS `users`;

DROP TABLE IF EXISTS `articles`;

DROP TABLE IF EXISTS `article_game_contents`;

DROP TABLE IF EXISTS `article_owners`;

DROP TABLE IF EXISTS `article_tags`;

DROP TABLE IF EXISTS `article_comments`;

DROP TABLE IF EXISTS `article_image_urls`;

CREATE TABLE `users` (
  `id` char(36) PRIMARY KEY,
  `created_at` DATETIME(6) DEFAULT NOW(),
  `updated_at` DATETIME(6) DEFAULT NOW() ON UPDATE CURRENT_TIMESTAMP,
  `google_id` text NOT NULL,
  `role` text,
  `name` text
);

CREATE TABLE `articles` (
  `id` char(36) PRIMARY KEY,
  `created_at` DATETIME(6) DEFAULT NOW(),
  `updated_at` DATETIME(6) DEFAULT NOW() ON UPDATE CURRENT_TIMESTAMP,
  `user_id` char(36),
  `title` text,
  `body` text,
  `public` bool
);

CREATE TABLE `article_game_contents` (
  `id` char(36) PRIMARY KEY,
  `created_at` DATETIME(6) DEFAULT NOW(),
  `updated_at` DATETIME(6) DEFAULT NOW() ON UPDATE CURRENT_TIMESTAMP,
  `exec_path` text,
  `zip_url` text
);

CREATE TABLE `article_owners` (
  `id` char(36),
  `article_id` char(36),
  PRIMARY KEY (`id`, `article_id`)
);

CREATE TABLE `article_tags` (
  `id` char(36) AUTO_INCREMENT,
  `article_id` char(36),
  `name` text,
  PRIMARY KEY (`id`, `article_id`)
);

CREATE TABLE `article_comments` (
  `id` char(36) AUTO_INCREMENT,
  `created_at` DATETIME(6) DEFAULT NOW(),
  `updated_at` DATETIME(6) DEFAULT NOW() ON UPDATE CURRENT_TIMESTAMP,
  `article_id` char(36),
  `body` text,
  `rate` int,
  PRIMARY KEY (`id`, `article_id`)
);

CREATE TABLE `article_image_urls` (
  `id` char(36) AUTO_INCREMENT,
  `article_id` char(36),
  `image_url` text,
  PRIMARY KEY (`id`, `article_id`)
);

CREATE INDEX `users_index_0` ON `users` (`id`);

CREATE INDEX `articles_index_1` ON `articles` (`id`);

CREATE INDEX `article_game_contents_index_2` ON `article_game_contents` (`id`);

-- ALTER TABLE
--   `articles`
-- ADD
--   FOREIGN KEY (`user_id`) REFERENCES `users` (`id`);

-- ALTER TABLE
--   `article_game_contents`
-- ADD
--   FOREIGN KEY (`id`) REFERENCES `articles` (`id`);

-- ALTER TABLE
--   `article_owners`
-- ADD
--   FOREIGN KEY (`id`) REFERENCES `users` (`id`);

-- ALTER TABLE
--   `article_owners`
-- ADD
--   FOREIGN KEY (`article_id`) REFERENCES `articles` (`id`);

-- ALTER TABLE
--   `article_tags`
-- ADD
--   FOREIGN KEY (`article_id`) REFERENCES `articles` (`id`);

-- ALTER TABLE
--   `article_comments`
-- ADD
--   FOREIGN KEY (`article_id`) REFERENCES `articles` (`id`);

-- ALTER TABLE
--   `article_image_urls`
-- ADD
--   FOREIGN KEY (`article_id`) REFERENCES `articles` (`id`);
