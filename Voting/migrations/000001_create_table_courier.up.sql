
CREATE TABLE election (                             -- выборы 
    id UUID PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    date TIMESTAMP DEFAULT NOW()
);

CREATE TABLE candidate (                             -- кандидаты                   
    id UUID PRIMARY KEY,
    election_id UUID REFERENCES election(id),
    party_id UUID NOT NULL,
    public_id UUID NOT NULL, 
    UNIQUE (election_id, public_id)
);

CREATE TABLE public_vote (                           -- голосование
    id UUID PRIMARY KEY,
    election_id UUID REFERENCES election(id),
    public_id UUID NOT NULL,
    UNIQUE (election_id, public_id)
);

CREATE TABLE votes (                                  -- голоса
    id UUID PRIMARY KEY,
    candidate_id UUID REFERENCES candidate(id)
);
