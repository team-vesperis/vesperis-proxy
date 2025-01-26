# Vesperis Proxy

[![GitHub Release](https://img.shields.io/github/v/release/team-vesperis/vesperis-proxy?link=https%3A%2F%2Fgithub.com%2Fteam-vesperis%2Fvesperis-proxy%2Freleases)](https://github.com/team-vesperis/vesperis-proxy/releases)
[![License](https://img.shields.io/github/license/team-vesperis/vesperis-proxy?link=https%3A%2F%2Fgithub.com%2Fteam-vesperis%2Fvesperis-proxy%2Fblob%2Fmain%2FLICENSE)](https://github.com/team-vesperis/vesperis-proxy/blob/main/LICENSE)

## Installation
1. Download the latest release and add the .exe file in a folder.
2. Create a file called "config.yml" in the main folder with the correct contents from Gate Minekube.
3. Add a folder called: "config" with a file inside called "vesperis.yml". Inside the file put the following with the correct usernames and passwords for the databases:

 ```
databases:
  mysql:
    username: ""
    password: ""
    host: "localhost"
    port: 3306
    database: "vesperis"
  
  redis:
    host: "localhost"
    port: 6379
    database: 0
    username: ""
    password: ""
 ```

4. Run the .exe file.
