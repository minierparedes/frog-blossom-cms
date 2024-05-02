CREATE TABLE "websites" (
  "id" bigserial UNIQUE PRIMARY KEY NOT NULL,
  "name" varchar(255) NOT NULL,
  "domain" varchar(255) NOT NULL,
  "owner_id" bigserial NOT NULL,
  "password" varchar(255),
  "template_id" bigserial,
  "builder_enabled" boolean DEFAULT false
);

CREATE TABLE "template" (
  "id" bigserial UNIQUE PRIMARY KEY NOT NULL,
  "name" varchar(255) NOT NULL
);

CREATE TABLE "pages" (
  "id" bigserial UNIQUE PRIMARY KEY NOT NULL,
  "website_id" bigserial NOT NULL,
  "title" varchar(255) NOT NULL,
  "content" text NOT NULL,
  "url" varchar(255) NOT NULL
);

CREATE TABLE "site_meta_tags" (
  "id" bigserial UNIQUE PRIMARY KEY NOT NULL,
  "page_id" bigserial NOT NULL,
  "name" varchar(255) NOT NULL,
  "content" text NOT NULL
);

CREATE TABLE "contact_forms" (
  "id" bigserial UNIQUE PRIMARY KEY NOT NULL,
  "website_id" bigserial NOT NULL,
  "form_id" varchar(255) NOT NULL
);

CREATE TABLE "form_fields" (
  "id" bigserial UNIQUE PRIMARY KEY NOT NULL,
  "form_id" bigserial NOT NULL,
  "label" varchar(255) NOT NULL,
  "type" varchar(255) NOT NULL,
  "required" boolean NOT NULL
);

CREATE TABLE "templates" (
  "id" bigserial UNIQUE PRIMARY KEY NOT NULL,
  "name" varchar(255) NOT NULL,
  "preview_image_url" varchar(255) NOT NULL
);

CREATE TABLE "layout_options" (
  "id" bigserial UNIQUE PRIMARY KEY NOT NULL,
  "name" varchar(255) NOT NULL,
  "description" text NOT NULL,
  "preview_image_url" varchar(255) NOT NULL
);

CREATE TABLE "users" (
  "id" bigserial UNIQUE PRIMARY KEY NOT NULL,
  "username" varchar(255) UNIQUE NOT NULL,
  "email" varchar(255) UNIQUE NOT NULL,
  "password" varchar(255) NOT NULL,
  "role" varchar(255) DEFAULT 'user',
  "owner_id" bigserial NOT NULL,
  "first_name" varchar(255),
  "last_name" varchar(255),
  "avatar_url" varchar(255),
  "bio" text,
  "created_at" timestamptz DEFAULT (now()),
  "updated_at" timestamptz
);

CREATE TABLE "content" (
  "id" bigserial UNIQUE PRIMARY KEY NOT NULL,
  "title" varchar(255) NOT NULL,
  "body" text NOT NULL,
  "author_id" bigserial NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT (now()),
  "updated_at" timestamp,
  "status" varchar(255) NOT NULL,
  "published_at" timestamp,
  "edited_at" timestamp,
  "organization_id" bigserial,
  "published_by_id" bigserial
);

CREATE TABLE "organizations" (
  "id" bigserial UNIQUE PRIMARY KEY NOT NULL,
  "name" varchar(255) NOT NULL
);

CREATE TABLE "content_meta_tags" (
  "id" bigserial UNIQUE PRIMARY KEY NOT NULL,
  "content_id" bigserial,
  "name" varchar(255) NOT NULL,
  "content" text NOT NULL
);

CREATE TABLE "content_categories" (
  "id" bigserial UNIQUE PRIMARY KEY NOT NULL,
  "content_id" bigserial,
  "category_id" bigserial
);

CREATE TABLE "categories" (
  "id" bigserial UNIQUE PRIMARY KEY NOT NULL,
  "name" varchar(255) NOT NULL
);

CREATE INDEX ON "websites" ("owner_id");

CREATE INDEX ON "websites" ("name");

CREATE INDEX ON "websites" ("domain");

CREATE INDEX ON "template" ("name");

CREATE INDEX ON "pages" ("website_id");

CREATE INDEX ON "pages" ("title");

CREATE INDEX ON "site_meta_tags" ("page_id");

CREATE INDEX ON "site_meta_tags" ("name");

CREATE INDEX ON "templates" ("name");

CREATE INDEX ON "users" ("username");

CREATE INDEX ON "users" ("owner_id");

CREATE INDEX ON "content" ("author_id");

CREATE INDEX ON "content" ("created_at");

CREATE INDEX ON "content" ("updated_at");

CREATE INDEX ON "content" ("organization_id");

CREATE INDEX ON "content" ("title");

CREATE INDEX ON "content" ("created_at", "updated_at");

CREATE INDEX ON "content_meta_tags" ("content_id");

CREATE INDEX ON "categories" ("name");

ALTER TABLE "websites" ADD FOREIGN KEY ("owner_id") REFERENCES "users" ("id");

ALTER TABLE "websites" ADD FOREIGN KEY ("template_id") REFERENCES "template" ("id");

ALTER TABLE "pages" ADD FOREIGN KEY ("website_id") REFERENCES "websites" ("id");

ALTER TABLE "site_meta_tags" ADD FOREIGN KEY ("page_id") REFERENCES "pages" ("id");

ALTER TABLE "contact_forms" ADD FOREIGN KEY ("website_id") REFERENCES "websites" ("id");

ALTER TABLE "form_fields" ADD FOREIGN KEY ("form_id") REFERENCES "contact_forms" ("id");

ALTER TABLE "users" ADD FOREIGN KEY ("owner_id") REFERENCES "websites" ("id");

ALTER TABLE "content" ADD FOREIGN KEY ("author_id") REFERENCES "users" ("id");

ALTER TABLE "content" ADD FOREIGN KEY ("organization_id") REFERENCES "organizations" ("id");

ALTER TABLE "content" ADD FOREIGN KEY ("published_by_id") REFERENCES "users" ("id");

ALTER TABLE "content_meta_tags" ADD FOREIGN KEY ("content_id") REFERENCES "content" ("id");

ALTER TABLE "content_categories" ADD FOREIGN KEY ("content_id") REFERENCES "content" ("id");

ALTER TABLE "content_categories" ADD FOREIGN KEY ("category_id") REFERENCES "categories" ("id");
