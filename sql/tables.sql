DROP SCHEMA public CASCADE;

CREATE SCHEMA public;

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE EXTENSION IF NOT EXISTS "pgcrypto";

SELECT * FROM pg_extension;

CREATE TABLE meeting(
  meeting_id	          UUID        DEFAULT gen_random_uuid(),
  meeting_name			    CHAR(50) 		not null,
  PRIMARY KEY (meeting_id)
);

CREATE TABLE race(
  race_id              UUID           DEFAULT gen_random_uuid(),
  meeting_id	         UUID,
  closing_time			   TIMESTAMP 		  not null,
  race_type            char(50)       not null,
  PRIMARY KEY (race_id),
  FOREIGN KEY (meeting_id) references meeting (meeting_id)
);

CREATE TABLE competitor(
  competitor_id       UUID          DEFAULT gen_random_uuid(),
  race_id	            UUID,
  position            INT           not null,
  type                char(50)      not null,
  name                char(50),
  PRIMARY KEY (competitor_id),
  FOREIGN KEY (race_id) references race (race_id)
);

INSERT INTO meeting (meeting_name) VALUES ('Gold Coast');

SELECT * FROM meeting;


