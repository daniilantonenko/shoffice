# Shoffice – Easily share a file with the office
The program provides a simple and intuitive way to share files with colleagues in the office.

### Sending
To submit a file, you need to scan a QR code, which will automatically redirect the user to the page with the form.

### Receiving
After successful submission, the file will be automatically delivered, depending on the preset parameters on the server. For example, it is sent by email or uploaded to a public server.


### Configuration
Change the conf.json file according to your settings, below is an example with possible parameters.

#### Required:
- "EmailServer": "smtp.example.com"
- "EmailPort": 587
- "FromEmail": "mail@example.com"
- "FromPass": "example pass"

#### Optional:
- "CompanyName": "My Company" - the company name is added on the main page
- "FileFormats" : [".jpg", ".png"] – can contain an array of formats available for loading through the form
- "MaxUploadSize": 10485760 – maximum file size when uploading through the form (10485760 bytes = 10 megabytes)

## Plan
- [X] Basic file upload function
- [ ] Generation of QR code
