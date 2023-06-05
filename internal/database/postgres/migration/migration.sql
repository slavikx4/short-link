CREATE TABLE "Link" (
    "id" serial NOT NULL,
    "shortLink" character varying UNIQUE NOT NULL ,
    "originalLink" character varying UNIQUE NOT NULL,
    CONSTRAINT "Link_pk" PRIMARY KEY ("id")
) WITH (
    OIDS=FALSE
);