-- +goose Up
-- +goose StatementBegin
CREATE TABLE quotes (
    id BIGSERIAL PRIMARY KEY,      -- Auto-incrementing ID
    text TEXT NOT NULL,         -- Quote text
    author VARCHAR(255) NOT NULL -- Author of the quote
);

INSERT INTO quotes (text, author) VALUES
    ('Life is 10% what happens to us and 90% how we react to it.', 'Charles R. Swindoll'),
    ('It takes courage to grow up and become who you really are.', 'E.E. Cummings'),
    ('Your self-worth is determined by you. You don''t have to depend on someone telling you who you are.', 'Beyoncé'),
    ('Nothing is impossible. The word itself says ''I''m possible!''', 'Audrey Hepburn'),
    ('Keep your face always toward the sunshine, and shadows will fall behind you.', 'Walt Whitman'),
    ('You have brains in your head. You have feet in your shoes. You can steer yourself any direction you choose. You''re on your own. And you know what you know. And you are the guy who''ll decide where to go.', 'Dr. Seuss'),
    ('Attitude is a little thing that makes a big difference.', 'Winston Churchill'),
    ('To bring about change, you must not be afraid to take the first step. We will fail when we fail to try.', 'Rosa Parks'),
    ('All our dreams can come true, if we have the courage to pursue them.', 'Walt Disney'),
    ('Don''t sit down and wait for the opportunities to come. Get up and make them.', 'Madam C.J. Walker'),
    ('Champions keep playing until they get it right.', 'Billie Jean King'),
    ('I am lucky that whatever fear I have inside me, my desire to win is always stronger.', 'Serena Williams'),
    ('You are never too old to set another goal or to dream a new dream.', 'C.S. Lewis'),
    ('It is during our darkest moments that we must focus to see the light.', 'Aristotle'),
    ('Believe you can and you''re halfway there.', 'Theodore Roosevelt'),
    ('Life shrinks or expands in proportion to one’s courage.', 'Anaïs Nin'),
    ('Just don''t give up trying to do what you really want to do. Where there is love and inspiration, I don''t think you can go wrong.', 'Ella Fitzgerald'),
    ('Try to be a rainbow in someone''s cloud.', 'Maya Angelou'),
    ('If you don''t like the road you''re walking, start paving another one.', 'Dolly Parton'),
    ('Real change, enduring change, happens one step at a time.', 'Ruth Bader Ginsburg'),
    ('All dreams are within reach. All you have to do is keep moving towards them.', 'Viola Davis');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS quotes;
-- +goose StatementEnd
