
Shopify Coding Challenge
Justin Tran

Chosen requirements:
- Add/Remove images and users
- Buy/Sell images
- public/private images

This project was built in Golang using the ent framework. I was able to build the API endpoints for users allowing creation, querying and updating of user objects. The same endpoints were completed for images.

Given more time, I would have set up all endpoints dealing with Buy/Sell from an ImageRepo. I would also setup a more dynamic env for others to use their own MySQL database. The images were purposely set up so that their fileLocation attribute would point to a resource in memory instead of having the static images saved inefficiently on the database. Basic Auth would be implemented to check if users have authorization for an image otherwise the image would not be shown. I had a lot of fun building this Repo in Golang but my flow was cut short when ent started to become confusing to upkeep. ent had a large library of files/configurations and I'd run in major blocks trying to understand the documentation. I should have maybe used a more widely/commonly used framework like gorm. I was only able to understand the basic implementations of ent like how to build and manipulate schemas. 

Attached is a completed basic UML Diagram for the structure of the objects of User, Image and ImageRepos.

Thanks for taking the time to read through this project!

