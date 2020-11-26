CREATE TABLE "users" (
  "id" bigserial PRIMARY KEY,
  "username" varchar NOT NULL,
  "password" varchar NOT NULL,
  "dob" date,
  "phone" varchar
);

CREATE TABLE "addresses" (
  "id" bigserial PRIMARY KEY,
  "user_id" bigint,
  "Line1" varchar,
  "Line2" varchar,
  "Line3" varchar,
  "City" varchar,
  "State" varchar,
  "Postal" varchar
);

CREATE TABLE "products" (
  "id" bigserial PRIMARY KEY,
  "product_name" varchar NOT NULL,
  "category" varchar NOT NULL,
  "availability" boolean NOT NULL DEFAULT (true),
  "price" decimal NOT NULL,
  "discount" decimal DEFAULT 0,
  "description" varchar,
  "rating" decimal
);

CREATE TABLE "orders" (
  "id" bigserial[pk],
  "user_id" bigint,
  "order_date" datetime NOT NULL DEFAULT (now())
);

CREATE TABLE "order_details" (
  "id" bigserial[pk],
  "order_id" bigint,
  "product_id" bigint,
  "quantity" int,
  "shipping_date" datetime,
  "delivery_status" varchar,
  "address_id" bigint
);

ALTER TABLE "addresses" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "orders" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "order_details" ADD FOREIGN KEY ("order_id") REFERENCES "orders" ("id");

ALTER TABLE "order_details" ADD FOREIGN KEY ("product_id") REFERENCES "products" ("id");

ALTER TABLE "order_details" ADD FOREIGN KEY ("address_id") REFERENCES "addresses" ("id");
