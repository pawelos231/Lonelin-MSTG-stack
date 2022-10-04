# Lonelin

This application is my attempt of creating a fully functional social media website, with backend written in go, this social media app will be created in the MSTG stack.

I know the name of this stack does not sound very appeling but i think it will be very interesting to build something in it !
this tech stack contains:

- S - SvelteKit (for rendering frontEnd)
- T - tailwind (to deal with stupid css)
- G - Golang (to deal with backend)
- M - mongodb as a database  


Project will include CRUD funcionalites, it will cover both frontend and backend, project will resemble reddit, instead of rest API it will feature grapghql, it will also contain 2D games as well as 3D games, created in three.js and ammo.js, if any idea pops into my head it will be featured in this project i also hope to create some rust based functionalities and many more.

For some time the .env file will be present in the folder structure, beacuse it does not contain any vulnerable data, as soon as the project will require this values to be completly private, .env file will not longer be present and all of values inside of it will be replaced

As time passes the technology stack of this application will grow in size, at some point it will require migration to SQL database for better organization of data, it will also need a caching system to which i plan to use redis.

if you want to open it on your own machine you will need to copy this repository and then time <code> npm run dev</code> in the client folder and <code>go run main.go</code> in the server folder. I hope to have alot of good time building this application and learn A LOT of new things ! 
