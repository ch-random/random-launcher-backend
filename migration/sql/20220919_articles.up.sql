DROP TABLE IF EXISTS `users`;

DROP TABLE IF EXISTS `articles`;

DROP TABLE IF EXISTS `article_game_contents`;

DROP TABLE IF EXISTS `article_owners`;

DROP TABLE IF EXISTS `article_tags`;

DROP TABLE IF EXISTS `article_comments`;

DROP TABLE IF EXISTS `article_image_urls`;

CREATE TABLE `users` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `created_at` DATETIME DEFAULT NOW(),
  `updated_at` DATETIME DEFAULT NOW() ON UPDATE CURRENT_TIMESTAMP,
  `google_id` varchar(255) UNIQUE NOT NULL,
  `role` varchar(255),
  `name` varchar(255)
);

CREATE TABLE `articles` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `created_at` DATETIME DEFAULT NOW(),
  `updated_at` DATETIME DEFAULT NOW() ON UPDATE CURRENT_TIMESTAMP,
  `user_id` int,
  `title` varchar(255),
  `body` varchar(255),
  `public` bool
);

CREATE TABLE `article_game_contents` (
  `id` int PRIMARY KEY,
  `created_at` DATETIME DEFAULT NOW(),
  `updated_at` DATETIME DEFAULT NOW() ON UPDATE CURRENT_TIMESTAMP,
  `exec_path` varchar(255),
  `zip_url` varchar(255)
);

CREATE TABLE `article_owners` (
  `id` int,
  `article_id` int,
  PRIMARY KEY (`id`, `article_id`)
);

CREATE TABLE `article_tags` (
  `id` int AUTO_INCREMENT,
  `article_id` int,
  `name` varchar(255),
  PRIMARY KEY (`id`, `article_id`)
);

CREATE TABLE `article_comments` (
  `id` int AUTO_INCREMENT,
  `created_at` DATETIME DEFAULT NOW(),
  `updated_at` DATETIME DEFAULT NOW() ON UPDATE CURRENT_TIMESTAMP,
  `article_id` int,
  `body` varchar(255),
  `rate` int,
  PRIMARY KEY (`id`, `article_id`)
);

CREATE TABLE `article_image_urls` (
  `id` int AUTO_INCREMENT,
  `article_id` int,
  `image_url` varchar(255),
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
