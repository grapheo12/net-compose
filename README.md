# Net-compose

Brother of docker-compose. Handles only Network Namespaces

## Capabilities

- [x] Creating and removing network namespaces
- [x] Support for Veth pairs
- [ ] Support for Bridges and other devices
- [x] Adding routing entries to devices, with nexthop
- [ ] Support for default gateway
- [x] Shell using a particular namespace


## Installing

This project works in Linux.
Make sure you have the following dependencies installed:

- golang
- iproute2
- make


In the project folder, run:

```bash
make && sudo make install
```

The executable `net-compose` will be installed in `/usr/bin`.


## Usage

See sample configs in the `test` folder

### Establish a virtual network

Create a config file following the format in the given examples.
(Let's say its name is `compose.yaml`)

To see what commands will be run to create the namespace, run:

```bash
net-compose compose.yaml dry-run
```

Currently the `up` command is defunct. :sweat-smile:

To really create the network, for now, you can do:

```bash
net-compose compose.yaml dry-run | sudo sh -
```


### Deleting the network

For now, anything created on the default namespace stays.
The cleanup operation deletes all the namespace.

To see the commands that will be used to delete the network, run:

```bash
net-compose compose.yaml down
```

To execute these commands, run:

```bash
net-compose compose.yaml down | sudo sh -
```


### Shell in a namespace

To open a shell where you can execute commands from a particular namespace, run:

```bash
sudo net-compose compose.yaml shell <name-of-the-namespace>
```

The namespace should be present in `compose.yaml` and the network must be established beforehand.

To exit from the shell, enter the `exit` command or press `Ctrl+C`.