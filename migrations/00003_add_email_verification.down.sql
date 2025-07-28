-- Migration 00003: add_email_verification
-- Down migration
-- Drop email_verifications table
DROP TABLE IF EXISTS email_verifications CASCADE;

-- Remove email verification fields from users table
ALTER TABLE
    users DROP COLUMN IF EXISTS email_verification_token,
    DROP COLUMN IF EXISTS email_verification_sent_at,
    DROP COLUMN IF EXISTS email_verification_attempts;