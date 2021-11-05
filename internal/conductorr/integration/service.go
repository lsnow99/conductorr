package integration

/*
All services that Conductorr communicates with must implement the Service interface.
The Service interface provides a means for initializing the service from a configuration
string, testing the connection to the service, as well as generating the configuration
string for when it needs to be written to the database.
*/
type Service interface {
	TestConnection() error
	GetConfig() (string, error)
	GetName() string
}