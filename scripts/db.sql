CREATE TABLE IF NOT EXISTS traffic  (
    key TEXT PRIMARY KEY,
    first_seen INT(10),
    last_seen INT(10),
    bytes BIGINT(20)
);

CREATE TABLE IF NOT EXISTS clients (
    key TEXT,
    name TEXT,
    ip VARCHAR(48),
    port UNSIGNED SMALLINT(5),

    FOREIGN KEY (key)
        REFERENCES traffic (key)
        ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS servers (
    key TEXT,
    name TEXT,
    ip VARCHAR(48),
    port UNSIGNED SMALLINT(5),
    is_broadcast_domain BOOLEAN,
    is_dhcp BOOLEAN,
    country VARCHAR(48),

    FOREIGN KEY (key)
        REFERENCES traffic (key)
        ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS protocols (
    key TEXT,
    l4 VARCHAR(15),
    l7 VARCHAR(45),

    FOREIGN KEY (key)
        REFERENCES traffic (key)
        ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS hosts (
    key INTEGER PRIMARY KEY AUTOINCREMENT,
    name VARCHAR(48),
    asname VARCHAR(48),
    privatehost BOOLEAN,
    ip VARCHAR(48),
    mac VARCHAR(48),
    city VARCHAR(48),
    country VARCHAR(48)
);