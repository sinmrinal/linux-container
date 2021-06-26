
# Linux Container

Create a lightwight linux container just by providing file system of distribution of your choice.
## Requirements

- [Golang](https://golang.org/)
## Run Locally


```bash
$ git clone https://github.com/sinmrinal/linux-container.git 
```
Go to the project directory

```bash
$ cd linux-container
```
Start the container

```bash
$ go run main.go -fs="path-to-file-system" run bin/bash
```

## How to downlaod ubuntu file system (optional)

To downlaod file system of Ubuntu 20.04 run this command in your terminal.

Note:- This will download file system in the same directory from which you are running this command.

```bash
$ wget https://cloud-images.ubuntu.com/minimal/releases/focal/release/ubuntu-20.04-minimal-cloudimg-amd64-root.tar.xz
```

Extract the file. (You may need to run this comman as sudo)

```bash
$ tar -xf ubuntu-20.04-minimal-cloudimg-amd64-root.tar.xz ubuntu-fs
```
## Feedback

If you have any feedback, you can reach out to on [LinkedIn](https://www.linkedin.com/in/sinmrinal) or [mrinal_singh@outlook.com](mailto:mrinal_singh@outlook.com)