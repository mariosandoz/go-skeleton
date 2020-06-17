# Go Skeleton
A simple golang project structure for creating API

## Structure
The project structure follow Hexagonal architecture from [GopherCon 2018](https://youtu.be/oL6JBUk6tj0).

This project structure only caters to the simplest project.
You can add the folder/package that needed.

## Folders
### /cmd
Add the http handler to invoke the function using http or command.

### /configs
All the configuration files will be put here.

### /internal
All the code that private and not for shared.

### /pkg
All the code that can be accessed by other.


#### These next folder naming are following the youtube video with some change.
Use Case: CRUD API for Creator

#### /pkg/adding
struct for adding creator and service to call the add to repository

#### /pkg/http/rest
handler for http request

#### /pkg/storage
#### /pkg/storage/mongodb
package mongodb for storage data in mongodb
#### /pkg/storage/json
package mongodb for storage data in json

In this case, there are 2 types of storage to cater to different use cases, if your project store the data using SQL, NoSQL, or write to a file. No matter the storage type, the changes should be minor.

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.
Please make sure to update tests as appropriate.

## License
[MIT](https://choosealicense.com/licenses/mit/)
