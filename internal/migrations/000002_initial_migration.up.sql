CREATE TABLE likes (
id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
from_user_id UUID NOT NULL,
to_user_id UUID NOT NULL,
created_at timestamptz DEFAULT now()
);

CREATE UNIQUE INDEX ux_likes_from_to ON likes(from_user_id, to_user_id);
CREATE INDEX idx_likes_from_user ON likes(from_user_id);
CREATE INDEX idx_likes_to_user ON likes(to_user_id);

CREATE TABLE dislikes (
id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
from_user_id UUID NOT NULL,
to_user_id UUID NOT NULL,
created_at timestamptz DEFAULT now()
);

CREATE UNIQUE INDEX ux_dislikes_from_to ON dislikes(from_user_id, to_user_id);
CREATE INDEX idx_dislikes_from_user ON dislikes(from_user_id);
CREATE INDEX idx_dislikes_to_user ON dislikes(to_user_id);

CREATE TABLE matches (
id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
user_first UUID NOT NULL,
user_second UUID NOT NULL,
created_at timestamptz DEFAULT now()
);

CREATE UNIQUE INDEX ux_matches_users 
ON matches(LEAST(user_first, user_second), GREATEST(user_first, user_second));
CREATE INDEX idx_matches_user_first ON matches(user_first);
CREATE INDEX idx_matches_user_second ON matches(user_second);