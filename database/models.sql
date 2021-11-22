CREATE TABLE Elections (
    id              TEXT NOT NULL,
    name            TEXT NOT NULL,
    published       BOOLEAN NOT NULL,
    finalized       BOOLEAN NOT NULL,
    openTime        TIMESTAMP,
    closeTime       TIMESTAMP,
    PRIMARY KEY(id)
);

CREATE TABLE Valid_Voters (
    email           TEXT NOT NULL,
    PRIMARY KEY(email)
);

CREATE TABLE Candidates (
    id              TEXT NOT NULL,
    name            TEXT NOT NULL,
    presentation    TEXT,
    PRIMARY KEY(id)
);

CREATE TABLE Casted_Votes (
    email           TEXT NOT NULL,
    electionID      TEXT NOT NULL,
    PRIMARY KEY(email, electionID),

    CONSTRAINT fk_voters
        FOREIGN KEY(email)
        REFERENCES Valid_Voters(email),
    CONSTRAINT fk_elections
        FOREIGN KEY(electionID)
        REFERENCES Elections(id)
);

CREATE TABLE Votes (
    hash            TEXT NOT NULL,
    electionID      TEXT NOT NULL,
    isBlank         BOOLEAN NOT NULL,
    PRIMARY KEY(hash),

    CONSTRAINT fk_elections
        FOREIGN KEY(electionID)
        REFERENCES Elections(id)
);

CREATE TABLE Candidates_In_Elections (
    electionID      TEXT NOT NULL,
    candidateID     TEXT NOT NULL,
    PRIMARY KEY (electionID, candidateID),

    CONSTRAINT fk_elections
        FOREIGN KEY(electionID)
        REFERENCES Elections(id),
    CONSTRAINT fk_candidates
        FOREIGN KEY(candidateID)
        REFERENCES Candidates(id)
);

CREATE TABLE Vote_Log (
    voteHash        TEXT NOT NULL,
    voteTime        TIMESTAMP NOT NULL,
    PRIMARY KEY(voteHash),

    CONSTRAINT fk_votes
        FOREIGN KEY(voteHash)
        REFERENCES Votes(hash)
)
