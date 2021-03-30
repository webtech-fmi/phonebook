CREATE TABLE IF NOT EXISTS profiles
(
	id UUID
		CONSTRAINT pk_profiles PRIMARY KEY
		CONSTRAINT fk_profiles_contacts REFERENCES contacts (owner_id),
	user_id UUID
		CONSTRAINT fk_profiles_users REFERENCES users (id),
	created_time TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT NOW(),
    modified_time TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT NOW(),
    email UUID
		CONSTRAINT fk_profiles_contact_info REFERENCES contact_info (id),
    personal JSONB,
    phone UUID
		CONSTRAINT fk_profiles_contact_info REFERENCES contact_info (id),
   	metadata JSONB
);