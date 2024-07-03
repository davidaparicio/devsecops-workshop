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
PROTECTION_TOKEN=<token> docker compose -p 42Crunch -f protect.yml up fw-petstore-secured
```

It deploys a FW in 'front' of the API, and listen on the port 4241. All requests are redirected to the petstore API after validation against the OpenAPI specification file.

You can add in the host's file the following lines to make it easier to test requests :

```
127.0.0.1    petstore-secured
127.0.0.1    petstore
```

Now it is time to test the firewall in the following section

## Conformance 

We will execute some curl script to assess the firewall's compliance with the API Specification.

```bash
curl --insecure --location 'https://petstore-secured:4241/petstores' --header 'Content-Type: application/json' --header 'apikey: 65496ebe-6544-4e77-bb66-20b97f6994bb' --data '{"name":"sunnytech"}'
```

You should receive the following response from the API:

API Response: 
```
{
    "id": "234bf052-3848-4cc8-acd4-3f7403aae504",
    "name": "sunnytech"
}
```

Let's attempt a request that deviates from the CreatePetstore operation's specifications


```bash
curl --insecure --location 'https://petstore-secured:4241/petstores' --header 'Content-Type: application/json' --header 'apikey: 65496ebe-6544-4e77-bb66-20b97f6994bb' --data '{"name":"sunnytech1"}'
```

The API didn't generate an error; instead, the firewall blocked the request.

```
{
    "status": 400,
    "title": "request validation",
    "detail": "Bad Request",
    "uuid": "b9ec2e83-acdf-424e-9c9c-de5262df353d"
}
```

You can access the blocked request in the 'Protection' tab, under the 'Transaction Logs' menu, and filter it by the UUID provided in the Firewall response.

## Rate limiting 

Now, we will assess the firewall's rate-limiting capability.
In the OpenAPI specification, specifically within the 'UserLogin' operation, we have configured the following strategy: 
:

```
"x-42c-local-strategy": {
    "x-42c-strategy": {
        "protections": [
            {
                "x-42c-request-limiter_0.1": {
                    "key": "req_limiter_login",
                    "window": 10,
                    "hits": 10,
                    "burst.enabled": true,
                    "burst.window": 2,
                    "burst.hits": 5
                }
            }
        ]
    }
},
```

[More info about strategy](https://docs.42crunch.com/latest/content/extras/protection_rate_limiting.htm)

To assess the burst rate, you can execute the following command:

```bash
for i in {1..6}; do
    echo "Execution $i:"
    curl --insecure --location 'https://petstore-secured:4241/login' --header 'Content-Type: application/json' --header 'Cookie: session=b54cc9f9-4006-44df-9c6e-325f9532987b' --data-raw '{
        "user":"admin@42crunch.com",
        "password":"HelloWorld"
    }'
done
```

Response: 

```
Execution 1:
{"id":"d8a679c0-06ae-4c0e-8045-36832953230a"}
Execution 2:
{"id":"a90ddb9d-2ed9-448e-8f78-d83812c2beb7"}
Execution 3:
{"id":"385051c9-5501-48f4-87a5-de49fa5d2efe"}
Execution 4:
{"id":"79192a85-36ce-42f2-8a71-711ad9813654"}
Execution 5:
{"id":"4fc37f87-7cad-4d53-bc59-b78966afef3a"}
Execution 6:
{"title":"Too Many Requests","detail":"Too Many Requests","instance":"https://www.42crunch.com/too-many-requests","status":429,"uuid":"839461a0-7471-430c-bdc8-22b6428a598f"}
```

To evaluate window rate limiting, please execute the following command: 

```bash
for i in {1..11}; do
    echo "Execution $i:"
    curl --insecure --location 'https://petstore-secured:4241/login' --header 'Content-Type: application/json' --header 'Cookie: session=b54cc9f9-4006-44df-9c6e-325f9532987b' --data-raw '{
        "user":"admin@42crunch.com",
        "password":"HelloWorld"
    }'
    sleep 0.5
done
```

Response:

```
Execution 1:
{"id":"ca43df1e-48f7-4e81-9d74-9ae19812be00"}
Execution 2:
{"id":"ac6c070a-8b45-4d4b-b7c7-932406729e23"}
Execution 3:
{"id":"29ff07c6-c348-4626-a0f9-8c025a09c0b0"}
Execution 4:
{"id":"0f0292c6-106f-47e7-b162-ece6bd7b39a9"}
Execution 5:
{"id":"3e3d28a6-6e94-4f0e-a70f-510dea06b269"}
Execution 6:
{"id":"daf12ddb-8fd1-4c13-bfb9-01a5470518c5"}
Execution 7:
{"id":"ad951fec-5843-4abf-8330-e89eae9f522a"}
Execution 8:
{"id":"85f668bd-5fea-4fdf-b374-2ab42af298b9"}
Execution 9:
{"id":"9e9e8ad9-82d7-4694-8caf-3527d2bbda7d"}
Execution 10:
{"id":"ca3706ee-7540-4e31-93d1-8645aa1d46a7"}
Execution 11:
{"title":"Too Many Requests","detail":"Too Many Requests","instance":"https://www.42crunch.com/too-many-requests","status":429,"uuid":"6fc9b968-ddf2-4b5f-a38b-2579ca28a28b"}
```

### Documentation

Step by step deployment with docker compose

https://github.com/42Crunch/resources/blob/master/firewall-deployment/DockerCompose.md
