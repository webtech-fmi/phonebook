CREATE TABLE IF NOT EXISTS contacts
(
	id UUID
		CONSTRAINT pk_contacts PRIMARY KEY,
	owner_id UUID NOT NULL,
	created_time TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT NOW(),
    modified_time TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT NOW(),
    favourite BOOLEAN DEFAULT FALSE,
    email UUID
		CONSTRAINT fk_contacts_contact_info REFERENCES contact_info (id),
    personal JSONB,
    phone UUID
		CONSTRAINT fk_contacts_contact_info REFERENCES contact_info (id),
   	metadata JSONB
);