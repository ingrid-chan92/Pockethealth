# Pockethealth Programming Challenge

This microservice is used for uploading DICOM Files. These files can then be parsed and processed to retrieve the image data and header attributes elements.

# API Endpoints

## Upload DICOM File

`POST /`

Given a DICOM File, store the file in the microservice. This file must be a valid DICOM file. Assigns a file ID to the file and returns it after processing

### Sample

**Request**
```
curl --request POST --header 'content-data' -F dicom=@test/IM000003 --url http://localhost:3333
```

**Response**
```
HTTP/1.1 200 OK

a2f72f93-fde0-46c6-be57-34ebe99b2a9e
```

## Get DICOM File as PNG 

`GET /:id/image`

Given a File ID, convert the file's pixel data into a PNG and return it. Currently only supports PNG conversion by default.

### Sample

**Request**
```
curl --request GET --url http://localhost:3333/a2f72f93-fde0-46c6-be57-34ebe99b2a9e/image
```

**Response**
```
HTTP/1.1 200 OK
```

![sample](./sample.png)

## Get Header Attributes from DICOM File via DICOM Tag

`GET /:id/headerAttribute?tag=(0000,0000)`

Given a File ID and a Tag, retrieve the header attribute element from that file.

Tag must be formatted `(XXXX,YYYY)` where XXXX and YYYY are hexadecimal values. See [DICOM Tags](https://www.dicomlibrary.com/dicom/dicom-tags/) for more details.

### Sample

**Request**
```
curl --request GET --url http://localhost:3333/44707cb0-4fe3-41f8-912e-1491d087738d/headerAttribute?tag=(7FE0,0010)
```

**Response**
```
HTTP/1.1 200 OK

{"TagName":"PixelData","Element":"FramesLength=1 FrameSize rows=1053 cols=967"}
```

## Get DICOM File Metadata 

`GET /:id`

Retrieves File Metadata such as the File Location and the User (Currently only User Id 1 is supported).

Debug function for easy data retrieval. 

### Sample

**Request**
```
curl --request GET --url http://localhost:3333/44707cb0-4fe3-41f8-912e-1491d087738d
```

**Response**
```
HTTP/1.1 200 OK

{"Id":"e2cb45d4-db19-458b-9b71-7a09ae4efd79","FileLocation":"persistence/dicom/e2cb45d4-db19-458b-9b71-7a09ae4efd79","UserId":1}
```

# Components

High level breakdown of the repository

**internal/handlers**

Contains business logic for handling the incoming requests. Files are divided per handler

**internal/model**

Contains models shared between handlers

**internal/router**

Contains the bootstrapping for setting up the endpoints. Primarily used for passing the DB between handlers

**persistence**

Contains the Sqlite database and the filesystem for storing the DICOM files. 

**test**

Contains test DICOM files for manual testing

# Next Steps
* Features
  * Support more image types, toggle-able via query params in the /image endpoint
  * Prevent users from uploading duplicate files by storing Md5 Hashes in Metadata
  * Caching DICOM Files/Caching Images
  * Bulk Upload/Bulk Tag lookup
  * User Permissions E.g. Tie Uploads to users. Do not allow users to access DICOM files they do not own
  * Authorization on endpoints. E.g. Debug API should be internal users only
* Maintenance
  * Replace Sqlite Database
  * Proper Logging for errors
  * Testing
    * Integration Tests
    * E2E Tests