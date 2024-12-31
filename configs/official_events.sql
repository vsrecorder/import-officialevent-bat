# mysql

CREATE TABLE official_events (
    id                      INT UNSIGNED NOT NULL PRIMARY KEY,
    title                   VARCHAR(255) NOT NULL,
    address                 VARCHAR(255) NOT NULL,
    venue                   VARCHAR(255) DEFAULT NULL, 
    date                    DATE NOT NULL,
    started_at              DATETIME DEFAULT NULL,
    ended_at                DATETIME DEFAULT NULL,
    deck_count              VARCHAR(255) DEFAULT NULL,
    type_id                 INT UNSIGNED DEFAULT NULL,
    type_name               VARCHAR(255) DEFAULT NULL,
    csp_flg                 TINYINT(2) UNSIGNED DEFAULT NULL,
    league_id               INT UNSIGNED DEFAULT NULL,
    league_title            VARCHAR(255) DEFAULT NULL,
    regulation_id           INT UNSIGNED DEFAULT NULL,
    regulation_title        VARCHAR(255) DEFAULT NULL,
    capacity                INT UNSIGNED DEFAULT NULL,
    attr_id                 INT UNSIGNED DEFAULT NULL,
    shop_id                 INT UNSIGNED DEFAULT NULL,
    shop_name               VARCHAR(255) DEFAULT NULL,
    FOREIGN KEY (`shop_id`) REFERENCES shops (`id`)
);



# postgres

CREATE TABLE official_events (
    id                      INT NOT NULL PRIMARY KEY,
    title                   VARCHAR(255) NOT NULL,
    address                 VARCHAR(255) NOT NULL,
    venue                   VARCHAR(255) DEFAULT NULL, 
    date                    DATE NOT NULL,
    started_at              TIMESTAMP DEFAULT NULL,
    ended_at                TIMESTAMP DEFAULT NULL,
    deck_count              VARCHAR(255) DEFAULT NULL,
    type_id                 INT DEFAULT NULL,
    type_name               VARCHAR(255) DEFAULT NULL,
    csp_flg                 BOOLEAN DEFAULT NULL,
    league_id               INT DEFAULT NULL,
    league_title            VARCHAR(255) DEFAULT NULL,
    regulation_id           INT DEFAULT NULL,
    regulation_title        VARCHAR(255) DEFAULT NULL,
    capacity                INT DEFAULT NULL,
    attr_id                 INT DEFAULT NULL,
    shop_id                 INT  DEFAULT NULL,
    shop_name               VARCHAR(255) DEFAULT NULL,
    FOREIGN KEY (shop_id)   REFERENCES shops (id)
);
