-- Create a `snippets` table.
CREATE TABLE snippets (
    id SERIAL, 
    title VARCHAR(100) NOT NULL,
    content TEXT NOT NULL,
    created TIMESTAMP NOT NULL,
    expires TIMESTAMP NOT NULL
);
-- Add an index on the created column.
CREATE INDEX idx_snippets_created ON snippets(created);

INSERT INTO public.snippets (title, content, created, expires) VALUES (
'An old silent pond',
'An old silent pond...\nA frog jumps into the pond,\nsplash! Silence again.\n\n– Matsuo Bashō', 
timezone('utc+2', now()),
timezone('utc+2', now()) + INTERVAL '365 DAYS'
);

INSERT INTO public.snippets (title, content, created, expires) VALUES (
'Over the wintry forest',
'Over the wintry\nforest, winds howl in rage\nwith no leaves to blow.\n\n– Natsume Soseki', 
timezone('utc+2', now()),
timezone('utc+2', now()) + INTERVAL '365 DAYS'
);

INSERT INTO public.snippets (title, content, created, expires) VALUES (
'First autumn morning',
'First autumn morning\nthe mirror I stare into\nshows my father''s face.\n\n– Murakami Kijo', 
timezone('utc+2', now()),
timezone('utc+2', now()) + INTERVAL '365 DAYS'
);