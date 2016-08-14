Build Updated Selenium Based Image
==================================


# What's the Selenium Image

Selenium images include many parts and versions, like `Standalone`, `hub` etc.
And also nodes can run different browser-based cases, like `Firefox`, `Chrome` etc.


## How to Build the Image?

1. git clone https://github.com/HackToday/docker-selenium.git

2. cd docker-selenium

3. change the Makefile version like below:

   ```
   VERSION := $(or $(VERSION),$(VERSION),'3.0.0-beta2')
   ```

4. change the version in download file of [Dockerfile of Base](https://github.com/HackToday/docker-selenium/blob/master/Base/Dockerfile#L30)
   changes like below:

   ```
   wget --no-verbose https://selenium-release.storage.googleapis.com/3.0-beta2/selenium-server-standalone-3.0.0-beta2.jar -O /opt/selenium/selenium-server-standalone.jar
   ```

5. build the image with following command:

   ```
   make standalone_firefox
   ```

## How to run the Image?

Refer the same as before [README.md](https://github.com/HackToday/docker-selenium)
