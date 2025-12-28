CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TYPE auth_provider_type AS ENUM ('google', 'github');
CREATE TYPE member_role_type AS ENUM ('admin', 'editor', 'viewer');

CREATE TABLE users (
  id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  email VARCHAR(255) NOT NULL UNIQUE,
  provider auth_provider_type NOT NULL,
  provider_id VARCHAR(255) NOT NULL,
  created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,

  CONSTRAINT unique_provider_user UNIQUE (provider, provider_id)
);

CREATE TABLE vault_rooms (
  id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  owner_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
  name VARCHAR(255) NOT NULL,
  access_code VARCHAR(255),
  expires_at TIMESTAMP WITH TIME ZONE,
  is_active BOOLEAN DEFAULT true,
  created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,

  CONSTRAINT room_name_length CHECK (length(name) >= 3 AND length(name) <= 255)
);

CREATE TABLE room_members (
  room_id UUID NOT NULL REFERENCES vault_rooms(id) ON DELETE CASCADE,
  user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
  role member_role_type NOT NULL DEFAULT 'viewer',
  created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,

  PRIMARY KEY (room_id, user_id),
  CONSTRAINT unique_membership UNIQUE (room_id, user_id)
);

CREATE TABLE secret_items (
  id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  room_id UUID NOT NULL REFERENCES vault_rooms(id) ON DELETE CASCADE,
  creator_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
  encrypted_content BYTEA NOT NULL,
  nonce BYTEA NOT NULL,
  is_burned BOOLEAN DEFAULT false,
  created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
  burned_at TIMESTAMP WITH TIME ZONE,

  CONSTRAINT valid_nonce_length CHECK (length(nonce) >= 12),
  CONSTRAINT content_not_empty CHECK (length(encrypted_content) > 0)
);

CREATE INDEX idx_users_email ON users(email);
CREATE INDEX idx_users_provider ON users(provider, provider_id);
CREATE INDEX idx_users_created_at ON users(created_at);
CREATE INDEX idx_users_id ON users(id);

CREATE INDEX idx_vault_rooms_owner ON vault_rooms(owner_id);
CREATE INDEX idx_vault_rooms_expires ON vault_rooms(expires_at) WHERE expires_at IS NOT NULL;
CREATE INDEX idx_vault_rooms_active ON vault_rooms(is_active) WHERE is_active = true;
CREATE INDEX idx_vault_rooms_created_at ON vault_rooms(created_at);

CREATE INDEX idx_room_members_user ON room_members(user_id);
CREATE INDEX idx_room_members_room ON room_members(room_id);
CREATE INDEX idx_room_members_role ON room_members(role);

CREATE INDEX idx_secret_items_room ON secret_items(room_id);
CREATE INDEX idx_secret_items_creator ON secret_items(creator_id);
CREATE INDEX idx_secret_items_burned ON secret_items(is_burned) WHERE is_burned = true;
CREATE INDEX idx_secret_items_created_at ON secret_items(created_at);
CREATE INDEX idx_secret_items_room_active ON secret_items(room_id, is_burned) WHERE is_burned = false;
