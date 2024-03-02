Modular
===

This is a small CLI project to serve as an initial modular CLI tool where you can build modules, or, plugins 
to extend the capabilities of your terminal.

## How to develop a plugin

- Check [HelloWorld command example](https://github.com/renatoaraujo/modular-helloworld-plugin)

## How to install a plugin

Build the application using Make
```shell
make build
```

Execute the plugin install command
```shell
./bin/modular plugin install -r renatoaraujo/modular-helloworld-plugin
```

## TODO

- Uninstall plugin
- Version management of plugins
- Documentation
- Publish for homebrew