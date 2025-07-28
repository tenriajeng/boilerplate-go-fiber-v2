-- Migration 00003: add_email_verification
-- Up migration
-- Add email verification fields to users table
ALTER TABLE
    users
ADD
    COLUMN email_verification_token VARCHAR(255),
ADD
    COLUMN email_verification_sent_at TIMESTAMP,
ADD
    COLUMN email_verification_attempts INTEGER DEFAULT 0;

-- Create email_verifications table for tracking verification attempts
CREATE TABLE email_verifications (
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    email VARCHAR(255) NOT NULL,
    token VARCHAR(255) NOT NULL UNIQUE,
    expires_at TIMESTAMP NOT NULL,
    verified_at TIMESTAMP,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Create index for faster lookups
CREATE INDEX idx_email_verifications_user_id ON email_verifications(user_id);

CREATE INDEX idx_email_verifications_token ON email_verifications(token);

CREATE INDEX idx_email_verifications_email ON email_verifications(email);

-- Add comment for documentation
COMMENT ON TABLE email_verifications IS 'Email verification tracking table';

COMMENT ON COLUMN email_verifications.token IS 'Unique verification token';

COMMENT ON COLUMN email_verifications.expires_at IS 'Token expiration time';

COMMENT ON COLUMN email_verifications.verified_at IS 'When email was verified';