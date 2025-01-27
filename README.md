# Vesperis Proxy

[![GitHub Release](https://img.shields.io/github/v/release/team-vesperis/vesperis-proxy?link=https%3A%2F%2Fgithub.com%2Fteam-vesperis%2Fvesperis-proxy%2Freleases)](https://github.com/team-vesperis/vesperis-proxy/releases)
[![License](https://img.shields.io/github/license/team-vesperis/vesperis-proxy?link=https%3A%2F%2Fgithub.com%2Fteam-vesperis%2Fvesperis-proxy%2Fblob%2Fmain%2FLICENSE)](https://github.com/team-vesperis/vesperis-proxy/blob/main/LICENSE)

## Installation Instructions

1. **Download and Place Executable**  
   - Download the latest release of the software.  
   - Place the `.exe` file in a designated folder of your choice.

2. **Prepare Databases**  
   - Ensure you have a **Redis** database and a **MySQL** database set up and accessible.  
   - Note down the connection details (host, port, username, password).

3. **Configure Gate Proxy**  
   - Create a file named `config.yml` for the **Gate Proxy** from Minekube.  
   - Inside `config.yml`, configure the following:
     - Servers
     - Connection Forwarding
     - Online Mode
     - Any other required settings.  
   - Refer to the Minekube documentation for additional configuration details.

4. **Run the Executable (First Run)**  
   - Execute the `.exe` file.  
   - This will generate a folder named `config` containing a file called `vesperis.yml`.  
   - The program will stop at this point (crash), as `vesperis.yml` requires further configuration.

5. **Edit Configuration**  
   - Open `vesperis.yml` and add the appropriate usernames, passwords, and connection details for the Redis and MySQL databases.  

6. **Run the Executable (Second Run)**  
   - Run the `.exe` file again.  
   - The program should now start successfully.  

---
