# ContainerGO

ContainerGO is a lightweight, minimalistic container runtime written in **Go**.

## 📌 Features

- **Namespace Isolation**: Processes run in isolated environments.
- **Cgroups Implementation**: Limits and tracks resource usage.
- **OverlayFS**: Provides efficient container filesystem layering.
- **Image Management**: Pull and extract images from Docker Hub.
- **Process Management**: Start, pause, resume, and stop containers.
- **Basic CLI Interface**: Simple commands to manage containers easily.

## 🚀 Installation

### Build the Project

```sh
go build -o congo .
```

## 🛠 Usage

Once built, you can use `congo means containerGO` to manage lightweight containers.

### 1️⃣ Run a New Container

```sh
./congo run /bin/bash
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
