CREATE TABLE posts (
    user_id INT NOT NULL,
    id INT NOT NULL UNIQUE,
    title VARCHAR(255) NOT NULL,
    body TEXT NOT NULL
);

CREATE TABLE comments (
    post_id INT NOT NULL,
    id INT NOT NULL UNIQUE,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    body TEXT NOT NULL
);

ALTER TABLE comments ADD FOREIGN KEY (post_id) REFERENCES posts (id) ON DELETE CASCADE;
