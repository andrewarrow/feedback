
CREATE TABLE users (
    id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    email varchar(255),
    phrase varchar(255),
    flavor varchar(255),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    UNIQUE KEY unique_email (email)
) ENGINE InnoDB;

CREATE TABLE inboxes (
    id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    sent_to varchar(255),
    sent_from varchar(255),
    subject varchar(255),
    is_spam int,
    spam_score float,
    body text,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    KEY sent_key (sent_to)
) ENGINE InnoDB;
