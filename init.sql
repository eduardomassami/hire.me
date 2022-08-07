CREATE DATABASE IF NOT EXISTS bemobi;

use bemobi;

CREATE TABLE `URLS` (
	`ID` int NOT NULL AUTO_INCREMENT,
    `Alias` varchar(255) NOT NULL,
    `URL` varchar(255) NOT NULL,
    `RETRIEVAL_COUNT` int,
    PRIMARY KEY (`ID`),
    UNIQUE KEY `Alias` (`Alias`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci