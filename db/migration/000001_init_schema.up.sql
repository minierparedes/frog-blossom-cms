CREATE TABLE "users" (
  "id" bigserial UNIQUE PRIMARY KEY NOT NULL,
  "username" varchar(255) UNIQUE NOT NULL,
  "email" varchar(255) UNIQUE NOT NULL,
  "password" varchar(255) NOT NULL,
  "role" varchar(255) DEFAULT 'user',
  "first_name" varchar(255),
  "last_name" varchar(255),
  "avatar_url" varchar(255),
  "bio" text,
  "created_at" timestamptz DEFAULT (now()),
  "updated_at" timestamptz
);

CREATE TABLE "website" (
  "id" bigserial UNIQUE PRIMARY KEY NOT NULL,
  "name" varchar(255) NOT NULL,
  "domain" varchar(255) NOT NULL,
  "owner_id" int NOT NULL,
  "selected_template" int NOT NULL
);

CREATE TABLE "template" (
  "id" bigserial UNIQUE PRIMARY KEY NOT NULL,
  "name" varchar(255) NOT NULL
);

CREATE TABLE "website_template" (
  "website_id" int,
  "template_id" int
);

CREATE TABLE "pages" (
  "id" bigserial UNIQUE PRIMARY KEY NOT NULL,
  "website_id" int NOT NULL,
  "title" varchar(255) NOT NULL,
  "url" varchar(255) NOT NULL
);

CREATE TABLE "page_components" (
  "id" bigserial UNIQUE PRIMARY KEY NOT NULL,
  "page_id" int NOT NULL,
  "component_type" varchar(255) NOT NULL,
  "component_data" jsonb NOT NULL
);

CREATE TABLE "content" (
  "id" bigserial UNIQUE PRIMARY KEY NOT NULL,
  "title" varchar(255) NOT NULL,
  "body" text NOT NULL,
  "author_id" int NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT (now()),
  "updated_at" timestamp,
  "status" varchar(255) NOT NULL,
  "published_at" timestamp,
  "edited_at" timestamp,
  "published_by_id" int,
  "component_id" int NOT NULL
);

CREATE TABLE "content_images" (
  "id" bigserial UNIQUE PRIMARY KEY NOT NULL,
  "content_id" int,
  "file_path" varchar(255) NOT NULL,
  "title" varchar(255) NOT NULL,
  "description" varchar(255) NOT NULL,
  "alt_text" VARCHAR(255)
);

CREATE TABLE "form_fields" (
  "id" bigserial UNIQUE PRIMARY KEY NOT NULL,
  "component_id" int NOT NULL,
  "label" varchar(255) NOT NULL,
  "type" varchar(255) NOT NULL,
  "required" boolean NOT NULL
);

CREATE INDEX ON "users" ("username");

CREATE INDEX ON "website" ("owner_id");

CREATE INDEX ON "website" ("name");

CREATE INDEX ON "website" ("domain");

CREATE INDEX ON "template" ("name");

CREATE INDEX ON "pages" ("website_id");

CREATE INDEX ON "pages" ("title");

CREATE INDEX ON "page_components" ("page_id");

CREATE INDEX ON "content" ("author_id");

CREATE INDEX ON "content" ("created_at");

CREATE INDEX ON "content" ("updated_at");

CREATE INDEX ON "content" ("title");

CREATE INDEX ON "content" ("created_at", "updated_at");

ALTER TABLE "website" ADD FOREIGN KEY ("owner_id") REFERENCES "users" ("id");

ALTER TABLE "website" ADD FOREIGN KEY ("selected_template") REFERENCES "template" ("id");

ALTER TABLE "website_template" ADD FOREIGN KEY ("website_id", "template_id") REFERENCES "website" ("id", "selected_template");

ALTER TABLE "website_template" ADD FOREIGN KEY ("website_id") REFERENCES "website" ("id");

ALTER TABLE "website_template" ADD FOREIGN KEY ("template_id") REFERENCES "template" ("id");

ALTER TABLE "pages" ADD FOREIGN KEY ("website_id") REFERENCES "website" ("id");

ALTER TABLE "page_components" ADD FOREIGN KEY ("page_id") REFERENCES "pages" ("id");

ALTER TABLE "content" ADD FOREIGN KEY ("author_id") REFERENCES "users" ("id");

ALTER TABLE "content" ADD FOREIGN KEY ("published_by_id") REFERENCES "users" ("id");

ALTER TABLE "content" ADD FOREIGN KEY ("component_id") REFERENCES "page_components" ("id");

ALTER TABLE "content_images" ADD FOREIGN KEY ("content_id") REFERENCES "content" ("id");

ALTER TABLE "form_fields" ADD FOREIGN KEY ("component_id") REFERENCES "page_components" ("id");
