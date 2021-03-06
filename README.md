# Cloaka

Cloaka is an Automatic Account Creator for Open Tibia. Brings a new level to the AAC scene. Created by **Raggaer**

Cloaka provides great performance and an easy to use scripting environment using **lua**. Faster and less demanding than the PHP AACs out there.

# Features

Cloaka provides a whole new level of features to the AAC scene. Here is a small preview

- LUA bindings. Make your custom pages using your favourite scripting languague
- OTBM loader. Generate images of your server houses and much more
- Load items.xml, monsters.xml, stages.xml, config.lua
- No APACHE / NGINX / UniServer required. Cloaka uses its own HTTP server.
- Only one executable. Cloaka its only one file (apart from your html layout files). Easy to distribute and installed
- Background services aka daemons
- Fast. Cloaka uses Go == Fast == Amazing
- Minimal set-up. Cloaka installs on its own and reads values from your config.lua to provide almost no configuration needed

# Commands

Cloaka support console commnads (handy for development) such as

- reload template (will compile all the templates.html)
- reload monsters (will load all the monsters from monsters.xml file)
- reload config lua (will reload the config.lua file into memory)
- reload config (will reload the config.json file into memory)
- reload stages (will reload the stages.xml file into memory)
- reload items (will reload the items.xml file into memory)
- reload map (will reload the whole .otbm map house images)
- exit (will terminate cloaka)

Feel free to create your own commands

# Compatible OT distributions

Cloaka is only compatible with [The Forgotten Server](https://github.com/otland/forgottenserver). Cloaka only supports its newer version 1.0 or greater.

# Download

Sadly there are currently no releases of Cloaka

# Custom version

To build your custom version of Cloaka you will need to have the following dependencies:

+ Go 1.6 or greater
+ Git

Once everything is installed you can simply do (to get the newer version)

```
go get github.com/Cloakaac/cloak
```

Now you can freely edit Cloaka and build it later with

```
go build
```
