

CREATE TABLE reactions (
id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
from_user_id UUID NOT NULL,
to_user_id UUID NOT NULL,
type INT2 NOT NULL,

created_at timestamptz DEFAULT now()
);

CREATE UNIQUE INDEX ux_reactions_from_to ON reactions(from_user_id, to_user_id);
CREATE INDEX idx_reactions_from_user ON reactions(from_user_id);
CREATE INDEX idx_reactions_to_user ON reactions(to_user_id);

CREATE TABLE matches (
id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
user_first UUID NOT NULL,
user_second UUID NOT NULL,
created_at timestamptz DEFAULT now()
);

CREATE UNIQUE INDEX ux_matches_users 
ON matches(LEAST(user_first, user_second), GREATEST(user_first, user_second));
CREATE INDEX idx_matches_user_first ON matches(user_first);
CREATE INDEX idx_matches_user_second ON matches(user_second);
