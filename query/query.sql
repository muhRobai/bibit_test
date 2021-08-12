select 
usr.id,
usr.user_name,
(
	SELECT user_name FROM users ug where ug.id = usr.parent 
) as parent
from users usr