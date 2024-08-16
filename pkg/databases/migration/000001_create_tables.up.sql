BEGIN;

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
SELECT uuid_generate_v4();

CREATE OR REPLACE FUNCTION set_updated_at_column()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = CURRENT_TIMESTAMP;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION generate_unique_string()
RETURNS TEXT LANGUAGE plpgsql AS $$
DECLARE
    new_string TEXT;
BEGIN
    LOOP
        new_string := substring(md5(random()::text), 1, 6);
        
        IF NOT EXISTS (SELECT id FROM codes WHERE id = new_string) THEN
            EXIT;
        END IF;
    END LOOP;
    
    RETURN new_string;
END;
$$;

CREATE TABLE "users" (
  "id" uuid PRIMARY KEY DEFAULT uuid_generate_v4() not null,
  "username" varchar not null,
  "email" varchar UNIQUE not null,
  "password" varchar not null,
  "created_at" timestamp DEFAULT CURRENT_TIMESTAMP not null,
  "updated_at" timestamp DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE "notes" (
  "id" uuid PRIMARY KEY DEFAULT uuid_generate_v4() not null,
  "user_id" uuid,
  "blocks" varchar not null,
  "created_at" timestamp DEFAULT CURRENT_TIMESTAMP not null,
  "updated_at" timestamp DEFAULT CURRENT_TIMESTAMP not null
);

CREATE TABLE "codes" (
  "id" varchar PRIMARY KEY DEFAULT generate_unique_string() not null,
  "note_id" uuid not null,
  "expired_at" timestamp DEFAULT CURRENT_TIMESTAMP not null,
  "created_at" timestamp DEFAULT CURRENT_TIMESTAMP not null
);

CREATE TABLE "images" (
  "id" uuid PRIMARY KEY DEFAULT uuid_generate_v4() not null,
  "note_id" uuid not null,
  "url" varchar not null,
  "created_at" timestamp DEFAULT CURRENT_TIMESTAMP not null
);

CREATE TABLE "auth" (
  "id" serial PRIMARY KEY not null,
  "access_token" varchar not null,
  "refresh_token" varchar not null,
  "created_at" timestamp DEFAULT CURRENT_TIMESTAMP not null
);

ALTER TABLE "notes" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");
ALTER TABLE "images" ADD FOREIGN KEY ("note_id") REFERENCES "notes" ("id");
ALTER TABLE "codes" ADD FOREIGN KEY ("note_id") REFERENCES "notes" ("id");

CREATE TRIGGER set_updated_at BEFORE UPDATE ON users FOR EACH ROW EXECUTE FUNCTION set_updated_at_column();
CREATE TRIGGER set_updated_at BEFORE UPDATE ON notes FOR EACH ROW EXECUTE FUNCTION set_updated_at_column();

COMMIT;