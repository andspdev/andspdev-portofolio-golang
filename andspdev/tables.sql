CREATE TABLE IF NOT EXISTS `pengguna` (
    `id` varchar(36) NOT NULL DEFAULT UUID(),
    `email` varchar(50) NOT NULL,
    `kata_sandi` varchar(255) NOT NULL,
    `username` varchar(25) NOT NULL,
    `login_terakhir` timestamp DEFAULT current_timestamp(),
    `tentang_saya` text NOT NULL,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS `pengguna_session` (
    `id` varchar(36) NOT NULL DEFAULT UUID(),
    `session_login` varchar(120) NOT NULL UNIQUE,
    `ip_address` varchar(128) NOT NULL,
    `ua` varchar(500) NOT NULL,
    `waktu_login` timestamp DEFAULT current_timestamp(),
    `is_logout` tinyint(1) NOT NULL DEFAULT 0,
    PRIMARY KEY(`id`)
) ENGINE=InnoDB CHARSET=utf8mb4;