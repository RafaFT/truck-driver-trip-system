# Truck-Driver-Trip-System
API for storing truck Drivers and keeping track of their Trips

This API is deployed on Google Cloud Platform and you can access the [base URL here](https://truck-pad.rj.r.appspot.com).

## Entities

### Driver
A Driver represents a Truck driver, all of it's fields are mandatory at creation time.

```
{
  "cpf": "48372162000",
  "name": "Geraldo Benjamin Galvão",
  "birth_date": "1992-02-26T15:00:00Z",
  "gender": "M",
  "has_vehicle": false,
  "cnh_type": "B"
}
```

1. `CPF (string)`: Eleven digits that uniquely identify a Driver
2. `Name (string)`: Driver's name
3. `Birth_Date (string)`: Date on RFC3339 format (hours, minutes and seconds are ignored)
4. `Gender (string)`: A Driver's gender can be defined as `"M"`, `"F"` or `"O"`.
5. `Has_Vehicle (boolean)`: Whehthe the Driver has it's own vehicle or not
6. `CNH_Type (string)`: Can be one of the following values: `"A"`, `"B"`, `"C"`, `"D"`, `"E"`, `"F"`.

OBS: When fetching a Driver, the API will also calculate and return an `Age` field.

### Trip
A Trip represents a Driver's checkpoint on a Terminal, where information about his/her travel is collected.
All fields but `driver_id` are mandatory at creation time.

```
{
  "driver_id": "48372162000",  // this field is ignored on creation if the Driver ID (CPF) comes from the URL path
  "has_load": false,
  "vehicle_type": 1,
  "time": "2020-07-05T15:12:06Z",
  "origin": {
    "latitude": -50,
    "longitude": 89
  },
  "destination": {
    "latitude": -90,
    "longitude": 179
  }
}
```

1. `Driver_ID (string)`: Eleven digits corresponding to a Driver's CPF
2. `Has_Load (boolean)`: Indicates whether the Driver had load when passing through a Terminal
3. `Vehicle_Type (number)`: Integer that represents the Driver's vehicle (valid values below)
4. `Time (string)`: Date in RFC3339 format, that indicates the timestamp of arrival at the Terminal
5. `Origin (object)`: Map containing the latitude and longitude of the Driver's trip origin
6. `Destination (object)`: Map containing the latitude and longitude of the Driver's trip destination

The valid values for a `Vehicle_Type` and it's meanings are:

```
1 -> Caminhão 3/4
2 -> Caminhão Toco
3 -> Caminhão Truck
4 -> Carreta Simples
5 -> Carreta Eixo Extendido
```

## Routes

All routes paths and query strings are **case-sensitive**.

### Drivers

1. `/drivers`

`POST`

Add a Driver.
Returns status `201` if everything went OK (no body content, for now), or a `409` if a Driver of the same `CPF` already existed.

Payload example:
```
{
  "cpf": "48372162000",
  "name": "Geraldo Benjamin Galvão",
  "birth_date": "1992-02-26T15:00:00Z",
  "gender": "M",
  "has_vehicle": false,
  "cnh_type": "B"
}
```

`GET`

Get all Drivers.
GET requests for Drivers support query strings on the URL (defined below).

Example: Get `cpf`, `name` and `cnh_type` from all Drivers that have a vehicle.

Request: `/drivers?fields=cpf,name,cnh_type&has_vehicle=true`

Response Body:
```
[
  {
    "cpf": "48372162000",
    "name": "Geraldo Benjamin Galvão",
    "cnh_type": "B"
  },
  {
    "cpf": "52488334855",
    "name": "Analu Sarah Aparício",
    "cnh_type": "E"
  }
]
```

***

2. `/drivers/<CPF>`

`PATCH`

Update information about a Driver.
You can update any Driver field, except his/her `CPF`.
Successful updates do not return a body (for now).

Example: Update `has_vehicle` and `gender` fields of Driver _52488334855_.

Request URL: `/drivers/52488334855`

Payload:
```
{
  "has_vehicle": true,
  "gender:" "O"
}
```

`GET`

Get data from a specific Driver.
Has support for query string.

Example: Get all information from Driver _48372162000_.

Request: `/drivers/48372162000`

Response:
```
{
  "cpf": "48372162000",
  "name": "Geraldo Benjamin Galvão",
  "birth_date": "1992-02-26T15:00:00Z",
  "age": 27,
  "gender": "M",
  "has_vehicle": false,
  "cnh_type": "B"
}
```

### Trips By Drivers

1. `/drivers/<CPF>/trips`

`POST`

Add a Trip to a Driver.
Every Trip has a programmatically defined ID field, which is only unique per Driver and it's the string concatenation of the Trip's `Time` value.

Example

Request: `/drivers/48372162000/trips`

Payload:

```
{
  "driver_id": "00000000001", // this field is ignored, as the CPF on the URL path dominates
  "has_load": false,
  "vehicle_type": 1,
  "time": "2020-07-05T15:12:09Z",
  "origin": {
    "latitude": -50,
    "longitude": 89
  },
  "destination": {
    "latitude": -90,
    "longitude": 179
  }
}
```

`GET`

Return all Trips from the specified Driver. Has support for query parameters (no support for origin/destination queries yet).

Example: Get specific fields of all Trips from Driver _14912725544_, from Feb 15º 2010 to Jan 1º 2020, that didn't have load.

Request: `/drivers/14912725544/trips?from=2010-02-15&to=2020-01-01&has_load=false&fields=time,destination,id`

Response:

```
[
  {
    "id": "20200214150000",
    "time": "2020-02-14T15:00:00Z",
    "destination": {
      "latitude": -77.02629,
      "longitude": 66.16744
    }
  },
  {
    "id": "20191226150000",
    "time": "2019-12-26T15:00:00Z",
    "destination": {
      "latitude": 37.5985,
      "longitude": 125.81492
    }
  }
]
```

2. `/drivers/<CPF>/trips/<ID>`

`GET`

Return a Trip by it's ID. A Trip ID is only unique per Driver, and is equal to the string concatenation of it's `Time` properties. It exists to prevent duplicated Trips per Driver and to allow direct access by this endpoint.

Query parameters are ignored (except `fields`).

Example: Get Trip trying to filter by `has_load`.

Request: `/drivers/14912725544/trips/20190601150000?has_load=false`

Response:

```
{
  "id": "20190601150000",
  "driver_id": "14912725544",
  "has_load": true,  // got result despite filter
  "vehicle_type": 1,
  "time": "2019-06-01T15:00:00Z",
  "origin": {
    "latitude": 35.70101,
    "longitude": 3.95183
  },
  "destination": {
    "latitude": -77.02684,
    "longitude": 124.07724
  }
}
```

2. `/drivers/<CPF>/trips/latest`

`GET`

Return the latest Trip (by `Time`) of the given Driver.

Query parameters are ignored (except `fields`).

Example:

Request: `/drivers/14912725544/trips/latest?fields=id,vehicle_type,time,has_load`

Response:

```
{
  "id": "20200214150000",
  "has_load": false,
  "vehicle_type": 1,
  "time": "2020-02-14T15:00:00Z"
}
```

### Trips

1. `/trips`

`GET`

Return all Trips. Has support for query parameters (no support for origin/destination queries yet).

Example 1: Get all trips from year 2020.

Request 1: `/trips?from=2020-01-01&to=2021-01-01`

***

Example 2: Get all trips of Carreta Simples that didn't have load, on February of 2020 in ascending order (by time), limiting to the first 10 results.

Request 2: `/trips?vehicle_type=4&has_load=false&from=2020-02-01&to=2020-03-01&order=asc&limit=10`

`POST`

It's also possible to add a Trip to a Driver using the `/trips` route. In this case however, the `driver_id` is passed on the request body.

Payload example:
```
{
	"driver_id": "14912725544",
	"has_load": false,
	"vehicle_type": 1,
	"time": "2020-02-14T15:00:00Z",
	"origin": {
			"latitude": 67.90649,
			"longitude": 113.38823
	},
	"destination": {
			"latitude": -77.02629,
			"longitude": 66.16744
	}
}
```

Trying to add the same Trip twice return a `409` status. The conflict is based on the Trip's `Time` and `driver_id`. A Driver cannot have more t han one trip with the same timestamp.

***

### The Future

1. Add query by `origin` and `destination` values.
2. Return URL for accessing the added resource after a successful `POST` request (for Trip and Driver).
3. Add pagination for Trips
4. Make sure status codes and response body's are intuitive and "RESTfull"
5. Internally, Trips are organized into SubCollections belonging to each Driver document. This seemed OK initially, but since it makes sense to query (frequently) Trips by not specifying a Driver and because all Trip queries use a Collection Group, I should make Trips be a top level collection.
6. Globally unique Trip IDs.
