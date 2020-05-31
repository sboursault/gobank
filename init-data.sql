-- Table: gobank.t_event

DROP TABLE gobank.t_event;

CREATE TABLE gobank.t_event
(
    id             SERIAL PRIMARY KEY,
	aggregate_type VARCHAR NOT NULL,
	stream_id      VARCHAR NOT NULL,
	event_type     VARCHAR NOT NULL,
	payload        VARCHAR NOT NULL
)

TABLESPACE pg_default;

ALTER TABLE gobank.t_event
    OWNER to postgres;

/*
INSERT INTO gobank.event(
	aggregate_type, stream_id, event_type, payload)
	VALUES ('test', 'test', 'test', 'test');
*/
