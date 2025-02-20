# ContainerGO

ContainerGO is a lightweight, minimalistic container runtime written in **Go**.

# What I Learned

- **Namespace Isolation**: Processes run in isolated environments using namespaces for resource segregation.
- **Cgroups Implementation**: Resource usage (CPU, memory, etc.) is controlled and tracked via cgroups.
- **OverlayFS**: Efficient container filesystem layering is achieved using OverlayFS.
- **System Calls**: Utilized low-level system calls to manage container processes and resources.
- **OCI Layer**: Complied with the Open Container Initiative (OCI) standards for runtime and image specs.
- **UnionFS**: Combined multiple filesystem layers into a single unified view using UnionFS.
- **Container Networking**: Configured network interfaces and isolation using virtual Ethernet devices and bridges.
- **Image Management**: Pulled, stored, and managed container images using registries and local storage.
- **Logging and Monitoring**: Set up logging and monitoring for container processes and resource usage.
- **CLI Development**: Built a command-line interface (CLI) to interact with the container runtime.

# Features

- **Basic CLI Interface**: Intuitive command-line interface for managing containers with simple, user-friendly commands.
- **Process Management**: Start, pause, resume, and stop container processes with ease.
- **Image Management**: Pull container images from Docker Hub and extract them for local use.
- **Resource Isolation**: Leverage Linux namespaces and cgroups for process and resource isolation.
- **Filesystem Layering**: Use OverlayFS for efficient and lightweight container filesystem management.
- **Networking Support**: Set up isolated network interfaces and bridges for container communication.
- **Logging and Monitoring**: Built-in logging for tracking container activity and resource usage.
- **Security Features**: Implement capabilities, seccomp profiles, and user namespaces for secure container execution.

## 🚀 Installation

### Build the Project

```sh
go build -o congo .
```

## 🛠 Usage

Once built, you can use **_congo_** means **_containerGO_** to manage lightweight containers.

### 1️⃣ Run a New Container

```sh
./congo run --name <container_name> <image_name> <command>
```

like

```sh
./congo run --name ubuntu-container ubuntu /bin/bash
```

This starts a new container with `/bin/bash` as the entry point.

### 2️⃣ Pull an Image from Docker Hub

```sh
./congo pull ubuntu
```

Downloads an `ubuntu` image from Docker Hub and stores it locally.

### 3️⃣ Extract an Image

```sh
./congo extract ubuntu
```

Extracts the pulled image to the appropriate filesystem directory.

### 4️⃣ Pause a Running Container

```sh
./congo pause <PID>
```

Suspends the container process.

### 5️⃣ Resume a Paused Container

```sh
./congo resume <PID>
```

Resumes a paused container process.

### 6️⃣ Stop a Running Container

```sh
./congo stop <PID>
```

Terminates the container process.

## 💡 Contributing

Feel free to fork the repository and submit **pull requests** to improve ContainerGO.
