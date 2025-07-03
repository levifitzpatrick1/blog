CREATE TABLE blog (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    title TEXT NOT NULL,
    slug TEXT NOT NULL UNIQUE,
    blurb TEXT NOT NULL,
    content TEXT NOT NULL,
    published DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated DATETIME DEFAULT CURRENT_TIMESTAMP
);

CREATE UNIQUE INDEX idx_posts_slug ON blog (slug);

CREATE INDEX idx_posts_published ON blog (published);

CREATE TABLE frc (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL,
    slug TEXT NOT NULL UNIQUE,
    script TEXT NOT NULL,
    presentation TEXT NOT NULL,
    description TEXT NOT NULL,
    video TEXT NOT NULL,
    image TEXT NOT NULL,
    created DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated DATETIME DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_frc_name ON frc (name);

CREATE UNIQUE INDEX idx_frc_slug ON frc (slug);

CREATE TABLE tags (
id INTEGER PRIMARY KEY AUTOINCREMENT,
name TEXT NOT NULL,
color TEXT NOT NULL
);

create TABLE tags_blog (
blog_id INTEGER NOT NULL,
tag_id INTEGER NOT NULL,
PRIMARY KEY (blog_id, tag_id),
FOREIGN KEY (blog_id) REFERENCES blog(id),
FOREIGN KEY (tag_id) REFERENCES tags(id)
);

create TABLE tags_frc (
frc_id INTEGER NOT NULL,
tag_id INTEGER NOT NULL,
PRIMARY KEY (frc_id, tag_id),
FOREIGN KEY (frc_id) REFERENCES frc(id),
FOREIGN KEY (tag_id) REFERENCES tags(id)
);
