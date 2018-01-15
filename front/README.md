### Description
This repository contains a front-end SPA app, its using React to manipulate the DOM, the skelton of the app was generate using : https://github.com/facebookincubator/create-react-app

### Usage
The input is one number per line in the rendered textarea, after clicking Submit the React app will send an HTTP GET request to Factorial Go API and log the result in the console.

### How to run Localy
`npm install` then `npm start` or `yarn start`

### How to build
`npm run build` or `yarn build`
This will build the app files in `/build` directory, ready to be dispacted to production

### Dependencies
* Node >=8.5.0
* npm 5.2+
* Factorial Go API

### FYI
API Endpoint is hardcoded in `src/App.js` Line 5, I don't think this is a good practice, do you ?

### Good to have
The response of each call is logged to browser console, try to log it in a DOM element once it is received from the API server

### Deployment
This application needs to be deployed behind a web server, preferred Nginx
