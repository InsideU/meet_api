
## The task is to develop a basic version of meeting scheduling API. You are only required to develop the API for the system. Below are the details.

### Note : The name of the database is zoom and the name of the collection is meets.

### Meetings should have the following Attributes. All fields are mandatory unless marked optional:
Id
Title
Participants
Start Time
End Time
Creation Timestamp

Participants should have the following Attributes. All fields are mandatory unless marked optional:
Name
Email
RSVP (i.e. Yes/No/MayBe/Not Answered)

You are required to Design and Develop an HTTP JSON API capable of the following operations,
Schedule a meeting
Should be a POST request
Use JSON request body
URL should be ‘/meetings’
Must return the meeting in JSON format
Get a meeting using id
Should be a GET request
Id should be in the url parameter
URL should be ‘/meeting/<id here>’
Must return the meeting in JSON format
List all meetings within a time frame
Should be a GET request
URL should be ‘/meetings?start=<start time here>&end=<end time here>’
Must return a an array of meetings in JSON format that are within the time range
List all meetings of a participant
Should be a GET request
URL should be ‘/meetings?participant=<email id>’
Must return a an array of meetings in JSON format that have the participant received in the email within the time range
