CREATE TABLE IF NOT EXISTS contact_info (
    id UUID CONSTRAINT pk_contact_info PRIMARY KEY,
    owner_id UUID NOT NULL,
    owner_type TEXT NOT NULL CONSTRAINT owner_type_in CHECK (owner_type IN ('profile', 'contact')),
    created_time TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT NOW(),
    modified_time TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT NOW(),
    personal TEXT [],
    office TEXT [],
    home TEXT []
);