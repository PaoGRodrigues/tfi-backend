CREATE TABLE IF NOT EXISTS traffic  (
    key INTEGER PRIMARY KEY,
    first_seen INT(10),
    last_seen INT(10),
    bytes BIGINT(20)
);

CREATE TABLE IF NOT EXISTS clients (
    key INTEGER,
    name TEXT,
    ip VARCHAR(48),
    port UNSIGNED SMALLINT(5),

    FOREIGN KEY (key)
        REFERENCES traffic (key)
        ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS servers (
    key INTEGER,
    name TEXT,
    ip VARCHAR(48),
    port UNSIGNED SMALLINT(5),
    is_broadcast_domain BOOLEAN,
    is_dhcp BOOLEAN,

    FOREIGN KEY (key)
        REFERENCES traffic (key)
        ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS protocols (
    key INTEGER,
    l4 VARCHAR(15),
    l7 VARCHAR(45),

    FOREIGN KEY (key)
        REFERENCES traffic (key)
        ON DELETE CASCADE
);