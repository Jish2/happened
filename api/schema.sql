-- Add new schema named "public"
CREATE SCHEMA IF NOT EXISTS "public";
-- Set comment to schema: "public"
COMMENT ON SCHEMA "public" IS 'standard public schema';
-- Create "events" table
CREATE TABLE "public"."events" ("id" text NOT NULL, "title" text NOT NULL, "description" text NULL, "created_at" timestamp NOT NULL, "updated_at" timestamp NOT NULL, "thumbnail" text NULL, PRIMARY KEY ("id"));
