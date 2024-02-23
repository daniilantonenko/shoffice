# Shoffice – Easily share a file with the office
The program provides a simple and intuitive way to share files with colleagues in the office.

### Sending
To submit a file, you need to scan a QR code, which will automatically redirect the user to the page with the form.

### Receiving
After successful submission, the file will be automatically delivered, depending on the preset parameters on the server. For example, it is sent by email or uploaded to a public server.

## Quick start with DockerHub
https://hub.docker.com/r/daniilantonenko/shoffice

## Quick start with Docker
Use this command to set up an Shoffice server:
``` bash
docker run \
    --name shoffice \
    --restart=always \
    -p 8080:8080 \
```

### Access to uploaded files
Add to docker run key:
``` bash
-v <local_path>:/build/web/uploads
```

## Build and run from source
Creating an application from source files. This assumes that git and go are already installed on the device.
``` bash
git clone https://github.com/daniilantonenko/shoffice.git
go build ./cmd/app/
./app
```

## Launch

1. Go to http://localhost:8080/generate
2. Select an address from the list or enter it manually (item “specify yours”)
3. Save the generated image for posting

## Environment variables

**Note:** All the variables to this image are optional, which means you don't have to type in any variable, and you can have an server out of the box! To do that, create an empty `env` file using `touch .env`, and skip to the next section.

This Docker image uses the following variables, that can be declared in an `env` file (see [example](.env.example)):

```
COMPANY_NAME=your_companu_name
FILE_FIRMATS=your_formats
MAX_UPLOAD_SIZE=your_max_size
```

## Plan
- [X] Basic file upload function
- [X] Generation of QR code
- [X] Installing the application via Docker
- [X] Configuration via env
- [ ] Email Queue
- [ ] Adding Multiple Files
- [ ] CAPTCHA verification
- [ ] Adding HTTPS
- [ ] Print a beautiful page with a QR code