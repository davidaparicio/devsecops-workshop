# Protection - Intro

Once your specification has been updated in line with company policies using auditing and the implementation is well respected with scanning. It's time to deploy your API and protect it with the firewall.


# How to run the Firewall

To deploy the API Firewall in front of the Petstore API you have to get a protection token linked to the API in 42Crunch platform. To get this token follow the tutorial (https://docs.42crunch.com/latest/content/tasks/protect_apis.htm
)

TLDR :

- In the UI Platform, select a collection then an API to protect
- Click on the button 'Protect API'
- Click on the button 'View instructions'
- Select the menu 'Protection tokens' to generate a firewall's token
- When you have the protection token you can execute the command below in the current directory

```
PROTECTION_TOKEN=<token> docker-compose -p 42Crunch -f protect.yml up fw-petstore-secured
```

It deploys a FW in 'front' of the API, and listen on the port 4241. All requests are redirected to the petstore API after validation against the OpenAPI specification file.

You can add in the host's file the following lines to make it easier to test requests :

```
127.0.0.1    petstore-secured
127.0.0.1    petstore
```

Now it is time to test the firewall,execute the following command to create a petstore : 

curl --insecure --location 'https://petstore-secured:4241/petstores' --header 'Content-Type: application/json' --header 'apikey: 65496ebe-6544-4e77-bb66-20b97f6994bb' --data '{
    "name":"volcamp"
}'

You should receive the following response from the API:

API Response: 
```
{
    "id": "234bf052-3848-4cc8-acd4-3f7403aae504",
    "name": "volcamp"
}
```

Let's try to execute a request which does not follow the CreatePetstore specification:

curl --insecure --location 'https://petstore-secured:4241/petstores' --header 'Content-Type: application/json' --header 'apikey: 65496ebe-6544-4e77-bb66-20b97f6994bb' --data '{
    "name":"volcamp1"
}'

Instead of receiving an error from the API, the firewall blocked the request:
```
{
    "status": 400,
    "title": "request validation",
    "detail": "Bad Request",
    "uuid": "b9ec2e83-acdf-424e-9c9c-de5262df353d"
}
```


You can see the logs in the 'Transaction Logs' menu in the platform UI and have an overview of all requests in the 'Security dashboard' 

### Documentation

Step by step deployment with docker compose

https://github.com/42Crunch/resources/blob/master/firewall-deployment/DockerCompose.md
