insert into users (name, nick, email, password)
values
("Tony Stark", "IronMan", "starkindustries@email.com", "$2a$10$BnO7SGvrn9yqr1EsHssO5OEMf15SIAdDLOj1M.a0ScJXZyQwnlPIC"),
("Natasha Romanoff", "BlackWidow", "avengersblackwidow@email.com", "$2a$10$BnO7SGvrn9yqr1EsHssO5OEMf15SIAdDLOj1M.a0ScJXZyQwnlPIC"),
("Steve Rogers", "CaptainAmerica", "avengerscaptainamerica@email.com", "$2a$10$BnO7SGvrn9yqr1EsHssO5OEMf15SIAdDLOj1M.a0ScJXZyQwnlPIC");

insert into followers (user_id, follower_id)
values
(1, 2),
(1, 3),
(3, 1),
(2, 3);

insert into tweets (title, content, author_id)
values
("Cant copy my armor", "I am Iron Man", 1),
("I escaped the Red Room", "They made me a killing machine", 2),
("Hydra is still alive", "They took Bucky", 3);