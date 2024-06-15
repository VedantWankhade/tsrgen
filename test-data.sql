DROP TABLE IF EXISTS entries;

CREATE TABLE entries (
  	entry_id SERIAL PRIMARY KEY,
  	page_id VARCHAR(30),
  	page_title VARCHAR(50),
  	page_link VARCHAR(100)
  );
  
 INSERT INTO entries (page_id, page_title, page_link) VALUES
 ('1234', 'Test', 'https://google.com'),
 ('35462', 'Page', 'http://vedaant.dev'),
 ('23424', 'Hello World', 'http://vedantwankhade.atlassian.net')
