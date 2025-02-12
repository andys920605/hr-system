CREATE TABLE employee (
    id              BIGINT PRIMARY KEY,   -- 使用 Snowflake ID，不自動遞增
    name            VARCHAR(255) NOT NULL,
    email           VARCHAR(255) UNIQUE NOT NULL,
    phone           VARCHAR(20),
    address         TEXT,
    position        TINYINT NOT NULL,  -- 職位 (Engineer, Manager, Admin)
    job_level       TINYINT NOT NULL,  -- 工作等級 (1~5)
    status          TINYINT NOT NULL CHECK (status IN (0, 1)), -- 0: Resigned, 1: Active
    created_at      TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at      TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);


