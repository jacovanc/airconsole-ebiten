# AirConsole + Ebiten Boilerplate

## Overview

This boilerplate project provides a basic setup for developing games that integrate AirConsole for input handling and Ebiten for game logic. It demonstrates how to set up a web-based controller interface and a Go-based game engine that communicates with it.

## Structure

The project is divided into two main parts:

- **web/** */: This directory contains the web frontend. It includes the HTML (controller.html, screen.html), JavaScript (script.js), and CSS (style.css) files necessary for both the controller interface and the html that will render the WASM
- **game/** */: This houses the Go codebase for the game. The Go code is compiled to WebAssembly, enabling it to run in the browser and interact with the web-based controller.

## Hosting

This project can be hosted in various ways, but below are instructions for using PHP as a local server for development and testing.

### General Hosting Guidelines

- **WebAssembly Compilation**: Compile the Go code in the game/ directory into WebAssembly and place the compiled file in the web/build directory.
- **Static File Server**: Any static file server can be used to serve the files in the web/ directory. This can range from a simple server in Python or Node.js to more robust servers like Apache or Nginx.

### Hosting with PHP

For local development, PHP's built-in server is a convenient option. Follow these steps to use PHP:

- **Check PHP Installation**: Ensure PHP is installed on your machine. Verify this by running php -v in your command line.
- **Start PHP Server**: In the command line, navigate to the web/ directory and start the server:
```
php -S localhost:8000
```
- **Access The Project**: With the server running, you can access the game and controller interfaces at:
Game interface: http://localhost:8000/screen.html
Controller interface: http://localhost:8000/controller.html.
Depending on how you are hosting, you may want to setup a redirect at the route of the web folder to the screen.html

- The PHP server will serve your static files, and any changes will be reflected upon refreshing the browser.

## Playing the game
- Find out the local IP of where you're hosting it (assuming that you are hosted it on your local network). Public IP is not needed as everything is done inside your local network.
- You can now test your game using the airconsole infrastructure as explained in [https://developers.airconsole.com/#!/getting_started]
- You can access it in two ways *(for both, make sure you replace the IP with the local IP that it's being hosted on)*
 - http://www.airconsole.com/#http://192.168.0.1:8080/ (this will let you play the game via the airconsole controller app)
 - http://www.airconsole.com/simulator/#http://192.168.0.1:8080/ (this will simulate the controllers in a split screen)

