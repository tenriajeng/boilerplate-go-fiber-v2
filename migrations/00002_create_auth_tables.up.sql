-- Create auth_sessions table
CREATE TABLE auth_sessions (
    id BIGSERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    token VARCHAR(500) UNIQUE NOT NULL,
    refresh_token VARCHAR(500) UNIQUE NOT NULL,
    expires_at TIMESTAMP NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Create password_resets table
CREATE TABLE password_resets (
    id BIGSERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    token VARCHAR(255) UNIQUE NOT NULL,
    expires_at TIMESTAMP NOT NULL,
    used BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Create tfa_codes table
CREATE TABLE tfa_codes (
    id BIGSERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    code VARCHAR(10) NOT NULL,
    expires_at TIMESTAMP NOT NULL,
    used BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Create indexes
CREATE INDEX idx_auth_sessions_user_id ON auth_sessions(user_id);

CREATE INDEX idx_auth_sessions_token ON auth_sessions(token);

CREATE INDEX idx_auth_sessions_refresh_token ON auth_sessions(refresh_token);

CREATE INDEX idx_auth_sessions_expires_at ON auth_sessions(expires_at);

CREATE INDEX idx_password_resets_user_id ON password_resets(user_id);

CREATE INDEX idx_password_resets_token ON password_resets(token);

CREATE INDEX idx_password_resets_expires_at ON password_resets(expires_at);

CREATE INDEX idx_tfa_codes_user_id ON tfa_codes(user_id);

CREATE INDEX idx_tfa_codes_code ON tfa_codes(code);

CREATE INDEX idx_tfa_codes_expires_at ON tfa_codes(expires_at);