CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- CreateEnum
CREATE TYPE "TaskStatus" AS ENUM ('PENDING', 'PROCESSING', 'COMPLETED', 'FAILED', 'CANCELLED');
-- CreateTable
CREATE TABLE users (
    "id" UUID NOT NULL DEFAULT uuid_generate_v4(),
    "username" TEXT NOT NULL UNIQUE,
    "email" TEXT NOT NULL UNIQUE,
    "password" TEXT NOT NULL,
    "name" TEXT NOT NULL,
    "created_at" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,

    CONSTRAINT "users_pkey" PRIMARY KEY ("id")
);

-- CreateTable
CREATE TABLE login_data (
    "id" UUID NOT NULL DEFAULT uuid_generate_v4(),
    "token" TEXT NOT NULL UNIQUE,
    "user_agent" TEXT,
    "ip_address" TEXT,
    "isp" TEXT,
    "state" TEXT,
    "city" TEXT,
    "country" TEXT,
    "device_type" TEXT NOT NULL,
    "created_at" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "last_logged_in" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "userID" UUID NOT NULL,

    CONSTRAINT "login_data_pkey" PRIMARY KEY ("id")
);

-- CreateTable
CREATE TABLE tasks (
    "id" UUID NOT NULL DEFAULT uuid_generate_v4(),
    "file_name" TEXT NOT NULL,
    "status" "TaskStatus" NOT NULL DEFAULT 'PENDING',
    "upload_url" TEXT NOT NULL UNIQUE,
    "result_url" TEXT NOT NULL UNIQUE,
    "created_at" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "task_type" TEXT NOT NULL,
    "userID" UUID NOT NULL,

    CONSTRAINT "tasks_pkey" PRIMARY KEY ("id")
);

-- CreateIndex
CREATE UNIQUE INDEX "users_username_key" ON users("username");

-- CreateIndex
CREATE UNIQUE INDEX "users_email_key" ON users("email");

-- CreateIndex
CREATE UNIQUE INDEX "login_data_token_key" ON login_data("token");

-- CreateIndex
CREATE UNIQUE INDEX "tasks_file_name_key" ON task("file_name");

-- AddForeignKey
ALTER TABLE login_data ADD CONSTRAINT "login_data_userID_fkey" FOREIGN KEY ("userID") REFERENCES users("id") ON DELETE RESTRICT ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE tasks ADD CONSTRAINT "tasks_userID_fkey" FOREIGN KEY ("userID") REFERENCES users("id") ON DELETE RESTRICT ON UPDATE CASCADE;
