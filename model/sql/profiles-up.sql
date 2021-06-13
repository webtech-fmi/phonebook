CREATE TABLE IF NOT EXISTS profiles
(
	id UUID
		CONSTRAINT pk_profiles PRIMARY KEY
		CONSTRAINT fk_profiles_contacts REFERENCES contacts (owner_id),
	user_id UUID
		CONSTRAINT fk_profiles_users REFERENCES users (id),
	created_time TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT NOW(),
    modified_time TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT NOW(),
    email JSONB,
    personal JSONB,
    phone JSONB,
   	metadata JSONB
);