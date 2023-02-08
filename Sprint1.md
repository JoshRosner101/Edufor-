# Sprint 1
## User Stories
- As a frequent user, I can log in to my account on the website.
- As a site visitor with the role of forum writer, I can add my own new forum posts, so that I can ask helpful questions about my software project.
- As a forum writer, I can add posts that are classified either as a software question or a software project collaboration.
- As a site visitor with the role of forum writer, I want my forum posts to stay on the website, so even if the website goes down for maintenence my questions and answers will still be there.
- As a frequent user, I want to know when each thread was originally uploaded, so that I can know what helpful information is most up to date.
- As a frequent user, I want to search and sort posts in order to quickly access desired ones.
- As a frequent user, I want to be able to see the most helpful answers first on a forum post with multiple responses. 
- As a collaborator (answers or collaborates on posted questions or projects), I want to post solutions to questions asked by a question or forum writer.
- As a collaborator, I want to join as a team member to projects posted on the forum.
## What issues we planned to address for this sprint
- Creating a basic back-end with a RESTful API using go that could store necessary data for forum posts
- Creating the UI Prototype using angular
- Setting up the connection between the front-end and back-end
## Which issues were successfully completed
- For the backend, the basic backend was implemented and the id of each post, the username of the one who posted it, the title of the post, the main content of the post and the time that it was created can all be stored. This data can either be gotten all at once, or you can get the data for one specific thread based on its unique id. Currently the backend only supports a POST command for adding that information, a GET command for getting all of the threads, and a GET command for getting a thread with a certain id.
- For the frontend, communication with the backend was successfully implemented, so the frontend can access all of the data currently stored in the backend and it can post data by filling out a form and clicking the button. The website automatically updates when new information is added to the database on it, so all of the posts can display properly.
## Which issues were unsuccessful and why
- For the backend, the main issue that was unsuccessful was the implementation of the database. We have been looking into multiple options for a database to implement (mysql, sqlite, mongodb) and decided that it would be better to focus on creating communication between the frontend and backend as it works now and to implement a more functional user interface before making any major backend overhauls. Implementing the database would take a major rewrite of the backend, so we thought that it would be better to stick with what's currently functional so we have something to show for this sprint instead of potentially breaking the backend while trying to add support for a database.
- For the frontend, the UI is very bare bones at the moment because we wanted to focus on getting the initial functionality with angular working before doing anything too crazy. Our next goal is to create a better looking UI that will look more like an actual forum.
