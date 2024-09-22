# Electronic Leash (Web)

Electronic Leash is a web-based pet monitoring system that allows pet owners to
track their pets' status. This project uses QR codes attached to pet collars,
which can be scanned to view the pet's information.

Obs.: you can find the embedded part
[here](https://github.com/leakedmemory/prototyping-class-project-embedded).

## Features

- QR code generation for pet collars
- Pet status monitoring
- Responsive web interface
- Direct access to pet information via QR code scan

## Getting Started

### Prerequisites

- Running on a Linux machine
- Go 1.23 or higher
- [templ](https://templ.guide/quick-start/installation)
- [Air](https://templ.guide/quick-start/installation)

Air and templ can be automatically installed by running the `make watch` command
and accepting the request.

### Developing

- Create a `.env` file and set up your environment variables based on the
  [example.env](example.env) file
- Building

```bash
make build
```

- Running

```bash
make run
```

- Live reload

```bash
make watch
```

- Cleaning the binary output

```bash
make clean
```

- Deploying on [Fly.io](https://fly.io) (adjust [fly.toml](fly.toml))

```bash
fly deploy
```

### Accessing the Web Interface

Once running, access `localhost:8080`.

### QR Code Generation

QR codes are generated for each pet and can be accessed at
`/pet/qrcode?leash-id=<leash_id>`. When scanned, these QR codes redirect to the
pet's information page.

### Pet Monitoring

Pet status is updated every 15 seconds and can be viewed on the main dashboard.

## Built With

- [Go](https://go.dev/)
- [HTMX](https://htmx.org/)
- [Pico CSS](https://picocss.com/)
- [templ](https://templ.guide/)
- [Docker](https://www.docker.com/)
- [Fly.io](http://fly.io)
- [Twilio](https://www.twilio.com/en-us)

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file
for details.

## Disclaimers

- This was a project built for my prototyping class at university
- It's my first project with this stack, so keep in mind that a lot of things
  are most likely made in an unusual way
