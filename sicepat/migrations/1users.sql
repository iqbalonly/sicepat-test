use db_sicepat;

DROP TABLE IF EXISTS `users`;

CREATE TABLE if not exists `users` (
	id int NOT NULL auto_increment,
	name varchar(100) NOT NULL,
	email varchar(100) NOT NULL,
	date_of_birth DATE DEFAULT NULL NULL,
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
	deleted_at TIMESTAMP NULL DEFAULT NULL,
	primary key (id)
)
ENGINE=InnoDB
DEFAULT CHARSET=utf8mb4
COLLATE=utf8mb4_unicode_ci;

-- SEED:
INSERT INTO `users` (name, email, date_of_birth) 
	VALUES 
		('iqbal', 'iqbal@gmail.com', '1998-06-28'),
		('test', 'test@gmail.com', NULL);