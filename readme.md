We will be using Google cloud for this project. The very first thing is to register an application to Google Console Dashboard. You can follow one of the online tutorials on how to do so. 
The following are the quick steps to take register an app:
1. Go to Google Cloud Platform (https://console.cloud.google.com/apis/dashboard)
2. From left option, Go to "**APIs & Services**"
3. Create "New Project" by selecting from the dropdown on top of screen. Then give the project name. And click **create**.
4. Make sure you have correct project selected from dropdown in dashboard
5. From left, select **OAuth Consent Screen**, select "External"
6. Give the app name. eg: vickey-go-app-01
7. add supporting email, and at bottom, developer email
8. Select the scope of application as **userinfo.email** and **userinfo.profile** to get data from google resource server. To learn more about the scopes, go to (https://developers.google.com/identity/protocols/oauth2/scopes)
9. Add Test users. Just use your Gmail email id
10. One more thing - select create credentials, OAuth ID client, select application type as Web Application, and name ex. Testing-1.
11. Add the Authorized redirect URIs. Since I am running non-secured connection,I had http://localhost:3000/callback, click create. That should now create OAuth client on Google Cloud. You should get client ID and Secret.
12. Download that **Json** file. You will use this info in the Go backend service.

Under the root of your project, you will create _.env_ file, and have the following fields as env var where you can simply copy the values from the JSON you downloaded:
<br>CLIENT_ID=
<br>CLIENT_SECRET=
<br>REDIRECT_URL=

Setup is complete. Just fire up the main file.

Go to localhost:3000

Click "**Login with Google**"

If it takes you to Google login screen and ask to enter google credentials to access the project you created in step 2 above, you are in right path. Importantly, you should notice the clientID and client Secret you had inside that JSON.

Now use the Google account you added as Test user. I had my own email address. Hence here is my output:

<pre><code>
Redirected URL:
http://localhost:3000/callback?state=random-string&code=<SOME_LONG_CODE>&scope=email+https%3A%2F%2Fwww.googleapis.com%2Fauth%2Fuserinfo.email+openid&authuser=0&prompt=consent

Data : {
  "id": "109345186767983474090",
  "email": "vickey.shrestha.1987@gmail.com",
  "verified_email": true,
  "picture": "https://lh3.googleusercontent.com/a/default-user=s96-c"
}
</code></pre>


