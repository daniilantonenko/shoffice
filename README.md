# Shoffice – Easily share a file with the office
The program provides a simple and intuitive way to share files with colleagues in the office.

### Sending
To submit a file, you need to scan a QR code, which will automatically redirect the user to the page with the form.

### Receiving
After successful submission, the file will be automatically delivered, depending on the preset parameters on the server. For example, it is sent by email or uploaded to a public server.


### Configuration
Create the conf.json file according to your settings, below is an example with possible parameters.

#### Required:
- "EmailServer": "smtp.example.com"
- "EmailPort": 587
- "FromEmail": "mail@example.com"
- "FromPass": "example pass"

#### Optional:
- "CompanyName": "My Company" - the company name is added on the main page
- "FileFormats" : [".jpg", ".png"] – can contain an array of formats available for loading through the form
- "MaxUploadSize": 10485760 – maximum file size when uploading through the form (10485760 bytes = 10 megabytes)

### Build and Run 
Creating an application from source files. This assumes that git and go are already installed on the device.
```
git clone https://github.com/daniilantonenko/shoffice.git
go build ./cmd/app/
./app
```

### Launch

1. Go to http://localhost:8080/generate
2. Select an address from the list or enter it manually (item “specify yours”)
3. Save the generated image for posting


### Plan
- [X] Basic file upload function
- [X] Generation of QR code
- [ ] Installing the application via Docker
- [ ] Adding Multiple Files
- [ ] Adding captcha
- [ ] Adding HTTPS
- [ ] Print a beautiful page with a QR code