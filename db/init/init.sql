CREATE DATABASE anomaly;

CREATE TABLE IF NOT EXISTS transmitters(
    session_id CHAR(255),
    frequence DOUBLE PRECISION,
    timestamp INT 
);