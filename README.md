# Contact Form Service

A simple and lightweight API for handling contact form submissions.

## Configuration
Create a config.yaml file in the same directory as the executable

```yaml
# Configure multiple recipients
# Specify the 'id' in the API call
destinations:
  - id: personalwebsite
    name: "Joe Bloggs"
    email: "joe@example.com"

  - id: appsupport
    name: "App Support"
    email: "support@example.com"

# Configure API keys for MailJet
mailjet: 
  privatekey: xxx
  publickey: xxx
  from: noreply@example.com
  name: Contact Form
  subject: Contact Form Message
```

## Available as a Docker image
```
docker pull ghcr.io/michaelrausch/contact:latest
``` 

## Usage

```javascript
POST "/"

{
   "Name": "Joe Bloggs"
   "Email": "example@email.com"
   "Message": "Hello, World!"
   "Destination": "personalwebsite"
}
```

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

## License
[BSD-3-Clause](https://opensource.org/licenses/BSD-3-Clause)