# cuidgen

cuidgen is a simple command-line utility to generate collision-resistant unique identifiers (CUIDs). It leverages the [lucsky/cuid](https://github.com/lucsky/cuid) library to produce identifiers that are excellent for distributed systems, where the risk of collision needs to be minimized.

## Installation

To install `cuidgen`, run the following command:

```bash
go install github.com/shiroemons/cuidgen
```

This will download and install the `cuidgen` executable to your Go bin directory.

## Usage

Once installed, you can generate a CUID by simply running:

```bash
cuidgen
```

This will output a new, unique CUID to the console.

## Features

- Generates collision-resistant unique identifiers.
- Simple and easy-to-use command-line interface.
- Utilizes the [lucsky/cuid](https://github.com/lucsky/cuid) library.

## License

[MIT License](LICENSE)
