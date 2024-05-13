CREATE TABLE "users" (
  "id" bigserial UNIQUE PRIMARY KEY NOT NULL,
  "username" varchar(255) UNIQUE NOT NULL,
  "email" varchar(255) NOT NULL,
  "password" varchar(255) NOT NULL,
  "role" varchar(255) NOT NULL,
  "first_name" varchar(255) NOT NULL,
  "last_name" varchar(255) NOT NULL,
  "user_url" text,
  "description" text,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL
);

CREATE TABLE "pages" (
  "id" bigserial UNIQUE PRIMARY KEY NOT NULL,
  "domain" varchar(255) NOT NULL,
  "author_id" bigint NOT NULL,
  "page_author" varchar(255) NOT NULL,
  "title" varchar(255) NOT NULL,
  "url" varchar(255) NOT NULL,
  "menu_order" bigint NOT NULL,
  "component_type" varchar(255) NOT NULL,
  "component_value" text NOT NULL,
  "page_identifier" varchar(255) NOT NULL,
  "option_id" bigint NOT NULL,
  "option_name" varchar(255) NOT NULL,
  "option_value" text NOT NULL,
  "option_required" boolean NOT NULL
);

CREATE TABLE "posts" (
  "id" bigserial UNIQUE PRIMARY KEY NOT NULL,
  "title" varchar(255) NOT NULL,
  "content" text NOT NULL,
  "author_id" bigint NOT NULL,
  "url" varchar(255) NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT (now()),
  "updated_at" timestamp NOT NULL,
  "status" varchar(255) NOT NULL,
  "published_at" timestamp NOT NULL,
  "edited_at" timestamp NOT NULL,
  "post_author" bigint NOT NULL,
  "post_mime_type" varchar(255) NOT NULL,
  "published_by" varchar(255) NOT NULL,
  "updated_by" varchar(255) NOT NULL
);

CREATE TABLE "meta" (
  "id" bigserial UNIQUE PRIMARY KEY NOT NULL,
  "page_id" bigint,
  "posts_id" bigint,
  "meta_title" varchar(255),
  "meta_description" text,
  "meta_robots" varchar(255),
  "meta_og_image" varchar(255),
  "locale" varchar(10),
  "page_amount" bigint NOT NULL,
  "site_language" varchar(255),
  "meta_key" varchar(255) NOT NULL,
  "meta_value" varchar(255) NOT NULL
);

CREATE INDEX ON "users" ("username");

CREATE INDEX ON "users" ("first_name");

CREATE INDEX ON "users" ("last_name");

CREATE INDEX ON "users" ("created_at", "updated_at");

CREATE INDEX ON "pages" ("page_author");

CREATE INDEX ON "pages" ("domain");

CREATE INDEX ON "pages" ("url");

CREATE INDEX ON "posts" ("created_at");

CREATE INDEX ON "posts" ("updated_at");

CREATE INDEX ON "posts" ("title");

CREATE INDEX ON "posts" ("created_at", "updated_at");

ALTER TABLE "pages" ADD FOREIGN KEY ("author_id") REFERENCES "users" ("id");

ALTER TABLE "pages" ADD FOREIGN KEY ("page_author") REFERENCES "users" ("username");

ALTER TABLE "posts" ADD FOREIGN KEY ("author_id") REFERENCES "users" ("id");

ALTER TABLE "posts" ADD FOREIGN KEY ("post_author") REFERENCES "users" ("username");

ALTER TABLE "posts" ADD FOREIGN KEY ("published_by") REFERENCES "users" ("username");

ALTER TABLE "posts" ADD FOREIGN KEY ("updated_by") REFERENCES "users" ("username");

ALTER TABLE "meta" ADD FOREIGN KEY ("page_id") REFERENCES "pages" ("id");

ALTER TABLE "meta" ADD FOREIGN KEY ("posts_id") REFERENCES "posts" ("id");
