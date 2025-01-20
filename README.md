# **Web crawler in GIN** 

This a Web-Crawlar service that extract all the relevant urls that are present in a particular domain.

To start the service : 

Switch to '**main**' branch.
Redirect to /src folder that contains main.go.
Run the '**go run main.go**'.


**Usage of service**

1. Incorrect input

Curl:

```bash
curl -X POST http://localhost:8080/webcrawl \
-H "Content-Type: application/json" \
-d '["ebay.com","walmart.com","indiamart.com","flipkart.com","tradeindia.com","tatanexarc.com","moglix.com","ofbusiness.com]' 
```

```
Sample Output : {"error":"invalid input format"}
```

No output.json will be generated in this case.

2. Correct Input & urls generation

```bash
curl -X POST http://localhost:8080/webcrawl \
-H "Content-Type: application/json" \
-d '["ebay.com","walmart.com","indiamart.com","flipkart.com","tradeindia.com","tatanexarc.com","moglix.com","ofbusiness.com"]' 
```

```
Sample Output : {"error":"invalid input format"}
```

output.json will be generated in this case.