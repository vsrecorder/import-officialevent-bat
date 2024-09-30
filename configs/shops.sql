# mysql

CREATE TABLE shops (
    id                            INT UNSIGNED NOT NULL PRIMARY KEY,
    name                          VARCHAR(255) NOT NULL,
    term                          TINYINT UNSIGNED NOT NULL,
    zip_code                      VARCHAR(8) DEFAULT NULL,
    prefecture_id                 TINYINT(2) UNSIGNED NOT NULL,
    address                       VARCHAR(255) DEFAULT NULL,
    tel                           VARCHAR(32) DEFAULT NULL,
    access                        TEXT DEFAULT NULL,
    business_hours                VARCHAR(255) DEFAULT NULL,
    url                           VARCHAR(255) DEFAULT NULL,
    geo_coding                    VARCHAR(63) DEFAULT NULL,
    FOREIGN KEY (`prefecture_id`) REFERENCES prefectures (`id`)
);



# postgres

CREATE TABLE shops (
    id                            INT NOT NULL PRIMARY KEY,
    name                          VARCHAR(255) NOT NULL,
    term                          SMALLINT NOT NULL,
    zip_code                      VARCHAR(8) DEFAULT NULL,
    prefecture_id                 SMALLINT NOT NULL,
    address                       VARCHAR(255) DEFAULT NULL,
    tel                           VARCHAR(32) DEFAULT NULL,
    access                        TEXT DEFAULT NULL,
    business_hours                VARCHAR(255) DEFAULT NULL,
    url                           VARCHAR(255) DEFAULT NULL,
    geo_coding                    VARCHAR(63) DEFAULT NULL,
    FOREIGN KEY (prefecture_id)   REFERENCES prefectures (id)
);



INSERT INTO shops VALUES(
    0,
    '株式会社ポケモン',
    0,
    NULL,
    0,
    NULL,
    NULL,
    NULL,
    NULL,
    NULL,
    NULL
);