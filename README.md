# Modular CLI

#### Overview

Modular is a compact and efficient CLI tool designed to supercharge your terminal experience with extendable modules or plugins. It provides a foundational framework for creating and integrating bespoke functionalities into your command-line interface, catering to a wide array of tasks and workflows. Whether you're looking to automate routine tasks, enhance productivity, or add novel capabilities, Modular offers a versatile platform for innovation and customization.

#### Features

- **Extensibility:** Easily extend the core functionalities with custom modules or plugins.
- **Simplicity:** Designed with simplicity in mind, enabling developers to create and share their plugins.
- **Customization:** Tailor the tool to meet your specific needs by integrating only the plugins you need.

#### Getting Started

1. **Prerequisites**
    - Ensure you have Go installed on your system to build and run Modular.
    - Basic knowledge of Go and CLI operations will be beneficial.

2. **Installation**
    - Clone the repository: `git clone https://github.com/renatoaraujo/modular.git`
    - Navigate to the project directory: `cd modular`
    - Build the application: `make build`

3. **Running Modular**
    - After building, you can run Modular directly from the `bin` directory: `./bin/modular`

#### Developing a Plugin

To create your own plugin for Modular:

1. Start by checking out the [HelloWorld command example](https://github.com/renatoaraujo/modular-helloworld-plugin), which provides a basic template.
2. Create your plugin following the guidelines and structure demonstrated in the HelloWorld example.
3. Build your plugin and integrate it with Modular to extend its capabilities.

#### Installing a Plugin

To install a new plugin into Modular:

```shell
./bin/modular plugin install -r [GitHub Repo URL of Plugin]
```

Replace `[GitHub Repo URL of Plugin]` with the actual repository URL of your plugin.

#### Contributing

We welcome contributions from the community! If you have an idea for a new feature, a bug fix, or a new plugin, please feel free to fork the repository, make your changes, and submit a pull request.

#### Future Plans

- Plugin uninstallation feature.
- Version management for plugins.
- Comprehensive documentation for developers and users.
- Distribution via package managers like Homebrew.