## Employee Commute Route Optimizer

Employee Commute Route Optimizer (ECRO) is a web application that helps employees find the optimal route to reach their destination in the shortest possible time. The project uses Google Maps API to determine the user's current location and HERE API for routing from the user's location to the destination. The app allows users to enter multiple stops and provides the shortest route to visit all of them.

ECRO takes input from the user in the form of a list of stops, and returns a map with the optimal route. The application provides multiple modes of transportation, such as walking, driving, cycling, and public transport. The user can choose their preferred mode of transportation, and ECRO will return the optimal route accordingly.

The project uses Golang as the programming language and implements the Goroutines feature of Golang to make concurrent API requests to HERE API. The API responses are sent through channels, which helps to improve the performance of the application.

ECRO also provides a visual representation of the optimal route using diagrams that are stored in the static folder. The project's architecture is documented in the diagrams in the static folder.

The project is configured using the config.json file, which contains the Google Maps API Key, HERE API Key, and the port number. The ECRO_API_Guideline.json file provides a guideline on how to use the web app. 

The project can be run by executing the following command:

```javascript
go run .\src\main.go
```

### Containerized Program Execution

- Build the Docker image using the following command:

```javascript
docker build -t ecro:1.0 .
```

- Push the Docker image to a Docker registry such as Docker Hub or leave it on your local machine

- Run the following command to apply the kubernetes deployment.yaml to a kubernetes cluster:

```javascript
kubectl apply -f deployment.yaml
```


### static folder

The "static" folder of the Employee Commute Route Optimizer project contains detailed diagrams that clearly illustrate the architecture and design of the project. These diagrams serve as an essential resource for developers who want to understand the project's structure, workflow, and logic. They provide a high-level view of the project's components, their relationships, and the data flow between them. By consulting these diagrams, developers can gain a deep understanding of the project's design and quickly identify potential issues or areas for improvement.

### \config\config.json

The config.json file is an essential component for running the Employee Commute Route Optimizer (ECRO) application. It contains three crucial fields: Port, GoogleMapAPIKey, and HEREAPIKey. While the Port field can be assigned any value, obtaining a valid GoogleMapAPIKey and HEREAPIKey is essential for the proper functioning of the application. To obtain a valid GoogleMapAPIKey, please refer to the developer guidelines. Similarly, to generate a HEREAPIKey, you will need to sign up at https://platform.here.com/ and follow the guidelines provided by the HERE developer page.

### Google Map API 

To retrieve the current location of a client, the project uses the Google Maps API. This API allows the application to obtain the precise geographical coordinates of the user's device, which can then be used to optimize the employee commute route. To use this API, developers must obtain a Google Maps API key, which can be done by following the guidelines provided by Google.

### HERE API

The project uses the HERE API to find routing from a geo point A to one or more other points B. To access this service, you will need to generate a HERE API Key by signing up at https://platform.here.com/ and following the developer guidelines.

### ECRO_API_Guideline.json

ECRO_API_Guideline.json is a file that provides a comprehensive guideline on how to use the Employee Commute Route Optimizer (ECRO) Web App. The file outlines the available endpoints and the required input parameters for each endpoint, as well as the expected output format. It also includes examples of how to make requests to the API and interpret the responses. The guideline is also available online at https://documenter.getpostman.com/view/9501436/2s93eSaavu.

### Modes

Both the `/getRouteDataHandler` and `/getTheShortestLocationHandler` endpoints support the following modes:

- car
- truck
- pedestrian
- bicycle
- publicTransport
- any (combines all available modes)

### Waypoints

A waypoint is a location on a route between the start point and the destination point. It can be an address, latitude-longitude coordinate, or a place ID. Waypoints help in finding optimal routes and provide additional context for navigation.

Here are some examples of what can be included as waypoints:

Latitude and longitude coordinates (e.g., "waypoint0=geo!52.5208,13.4050")
Street addresses (e.g., "waypoint0=street!invalidenstrasse%20116,10115%20berlin,germany")
Cities or regions (e.g., "waypoint0=city!berlin,germany")
Postal codes (e.g., "waypoint0=postalCode!90210,usa")
Landmarks or points of interest (e.g., "waypoint0=landmark!brandenburg%20gate,berlin,germany")
The specific format of the waypoint parameter depends on the type of waypoint being used.

### Contact

If you have any questions or problems, please email me at muradshahsuvarov@gmail.com
