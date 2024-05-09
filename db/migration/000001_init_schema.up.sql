CREATE TABLE "users" (
  "id" bigserial UNIQUE PRIMARY KEY NOT NULL,
  "username" varchar(255) NOT NULL,
  "email" varchar(255) NOT NULL,
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
  "owner_id" bigint NOT NULL
);

CREATE TABLE "pages" (
  "id" bigserial UNIQUE PRIMARY KEY NOT NULL,
  "website_id" bigint NOT NULL,
  "title" varchar(255) NOT NULL,
  "url" varchar(255) NOT NULL,
  "menu_order" bigint NOT NULL
);

CREATE TABLE "page_components" (
  "id" bigserial UNIQUE PRIMARY KEY NOT NULL,
  "page_id" bigint NOT NULL,
  "component_type" varchar(255) NOT NULL,
  "component_value" text NOT NULL,
  "label" varchar(255) NOT NULL,
  "option_id" bigint NOT NULL,
  "option_name" varchar(255),
  "option_value" text,
  "required" boolean NOT NULL
);

CREATE TABLE "posts" (
  "id" bigserial UNIQUE PRIMARY KEY NOT NULL,
  "title" varchar(255) NOT NULL,
  "content" text NOT NULL,
  "author" varchar(255) NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT (now()),
  "updated_at" timestamp,
  "status" varchar(255) NOT NULL,
  "published_at" timestamp,
  "edited_at" timestamp,
  "published_by_id" bigint,
  "website_id" bigint,
  "post_mime_type" varchar(255)
);

CREATE TABLE "meta" (
  "id" bigserial UNIQUE PRIMARY KEY NOT NULL,
  "page_id" bigint,
  "posts_id" bigint,
  "meta_title" varchar(255),
  "meta_description" text,
  "meta_robots" varchar(255),
  "meta_viewport" varchar(255),
  "meta_charset" varchar(255),
  "website_id" bigint NOT NULL,
  "page_amount" bigint,
  "site_language" varchar(255),
  "meta_key" varchar(255),
  "meta_value" varchar(255) NOT NULL
);

CREATE INDEX ON "posts" ("created_at");

CREATE INDEX ON "posts" ("updated_at");

CREATE INDEX ON "posts" ("title");

CREATE INDEX ON "posts" ("created_at", "updated_at");

ALTER TABLE "website" ADD FOREIGN KEY ("owner_id") REFERENCES "users" ("id");

ALTER TABLE "pages" ADD FOREIGN KEY ("website_id") REFERENCES "website" ("id");

ALTER TABLE "page_components" ADD FOREIGN KEY ("page_id") REFERENCES "pages" ("id");

ALTER TABLE "posts" ADD FOREIGN KEY ("published_by_id") REFERENCES "users" ("id");

ALTER TABLE "posts" ADD FOREIGN KEY ("website_id") REFERENCES "website" ("id");

ALTER TABLE "meta" ADD FOREIGN KEY ("page_id") REFERENCES "pages" ("id");

ALTER TABLE "meta" ADD FOREIGN KEY ("posts_id") REFERENCES "posts" ("id");

ALTER TABLE "meta" ADD FOREIGN KEY ("website_id") REFERENCES "website" ("id");
