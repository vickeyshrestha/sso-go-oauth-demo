We will be using Google cloud for this project. The very first thing is to register an application to Google Console Dashboard. You can follow one of the online tutorials on how to do so. 
The following are the quick steps to take register an app:
1. Go to Google Cloud Platform
2. Create New Project
3. Create Credentials -> OAuth client ID
4. In OAuth Consent Screen, select "External"
5. Give the app name. eg: vickey-go-app-01
6. add supporting email, and at bottom, developer email
7. Select the scope of application as userinfo.email and userinfo.profile to get data from google resource server

8. Add Test users. Just use your gamil email id
9. One more thing - select create credentials, OAuth ID client, select application type as Web  Application, and name as Testing-1, and Authorized redirect URIs. I had https://localhost:3000/callback, click create. That should now create OAuth client on Google Cloud. You should get client ID and Secret.
10. Download that **Json** file. You will use this info in the Go backend service.

Under the root of your project, you will create _.env_ file, and have the following fields as env var where you can simply copy the values from the JSON you downloaded:
<br>CLIENT_ID=
<br>CLIENT_SECRET=
<br>REDIRECT_URL=

Setup is complete. Just fire up the main file.

Go to localhost:3000

Click "**Login with Google**"

If it takes you to Google login screen and ask to enter google credentials to access the project you created in step 2 above, you are in right path. Importantly, you should notice the clientID and client Secret you had inside that JSON.



