# Download and Installation

Follow the instructions below depending on your system and preferred installation method. For configuration instructions post-installation, check [here]().

## Manual Installation

Download the latest release for your operating system [here](https://github.com/lsnow99/conductorr/releases).

Put the binary file in a suitable location. Simply double click the binary or run `./conductorr` in your terminal to start Conductorr. On some systems you may need to first run `chmod +x conductorr` in your terminal before launching.

To launch automatically at startup, please refer to your operating system's instructions.

Navigate to [http://localhost:6416/](http://localhost:6416/) and you should see the setup screen prompting you to create the admin user.

## Docker Compose

Create a `docker-compose.yml` file with the following contents:

```yml
version: "3.9"
services:
  conductorr:
    image: "conductorr"
    ports:
      - "6416:6416"
    volumes:
      # Replace the . with your preferred location on your system for the database file
      - .:/app/conductorr.db  
    environment:
      TMDB_API_KEY: yourapikeyhere
      # Add any other environment variables for configuration here
```
