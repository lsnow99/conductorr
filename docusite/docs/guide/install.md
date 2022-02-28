# Installation

Follow the instructions below depending on your system and preferred installation method. For configuration instructions post-installation, check [here](/guide/configuration.html).

## Manual Installation

Download the latest release for your operating system [here](https://github.com/lsnow99/conductorr/releases).

Put the executable file in a suitable location. Simply double click the executable or run `./conductorr` in your terminal to start Conductorr. On some systems you may need to first run `chmod +x conductorr` in your terminal before launching.

Navigate to [http://localhost:6416/](http://localhost:6416/) and you should see the setup screen prompting you to create the admin user.

> NOTE: To launch Conductorr automatically at startup, you will need to follow instructions provided by your operating system. Alternatively, you can configure Conductorr via Docker to launch automatically.

## Docker Compose

Create a `docker-compose.yml` file with the following contents:

```yml
version: "3.4"
services:
  conductorr:
    image: "lsnow99/conductorr"
    ports:
      - "6416:6416"
    volumes:
      # Replace the paths before the colon with the corresponding paths on your host system
      - ./data:/app/conductorr.db
      - ./library:/app/library
      - ./downloads:/app/downloads
    environment:
      # Add any environment variables for configuration here
```
