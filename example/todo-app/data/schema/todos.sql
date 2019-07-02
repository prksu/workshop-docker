CREATE TABLE IF NOT EXISTS todos (
    id integer AUTO_INCREMENT NOT NULL,
    title varchar(50) NOT NULL,
    completed boolean DEFAULT false,
    createtime timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    updatetime timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`)
);

INSERT INTO todos (title) VALUES ("first todo");