### Description
#

![Screenshot from 2024-04-10 21-05-38](https://github.com/kentlouisetonino/baseshift/assets/69438999/51305be2-4714-4225-9378-e686684acf0d)




<br />

> - A CLI tool for converting number systems.

> - This will support binary, decimal, octal, and hexadecimal systems.

> - Snap : https://snapcraft.io/baseshift

<br />
<br />



### System Design
#

<br />

![System Design](./docs/image-system-design.png)

<br />
<br />



### Local Development
#
> - Note: `The root main.go is a symbolic link`.

> - Build and run the app.

```bash
# Change permission for the build script.
sudo chmod +x run.sh

# Run the app.
./run.sh
```

> - Below are the installation commands.

```bash
# If app is already installed.
sudo snap refresh baseshift

# If app is not yet installed.
sudo snap install baseshift

# Remove the app using snap.
sudo snap remove baseshift

# Running the app.
baseshift
```
