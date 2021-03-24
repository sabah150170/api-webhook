
##Docker local build and run:
#	sudo docker build -t api .
#	sudo docker run api


##Heroku Docker Deploy:
#	sudo su 									(docker is running with root perm. )
#	heroku login -i 									
#	heroku container:login 
#	heroku create 									(heroku app name will be created)
#	heroku container:push web -a <heroku_app_name>
#	heroku container:release web -a <heroku_app_name>
#	heroku open -a <heroku_app_name>
