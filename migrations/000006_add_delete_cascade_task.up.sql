ALTER TABLE todoapp.tasks
ADD CONSTRAINT tasks_author_user_id_fkey
FOREIGN KEY (author_user_id)
REFERENCES todoapp.users(id)
ON DELETE CASCADE;