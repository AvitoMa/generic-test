### Description
This repository contains a single page app, its React based, the skelton of the app was generate using : https://github.com/facebookincubator/create-react-app

### Usage
The input is one number per line in the rendered textarea, after clicking Submit the React app will send an HTTP GET request to Factorial Go API and log the result in the console.

### How to run Localy
`npm install` then `npm start` or `yarn start`

### How to build
`npm run build` or `yarn build`
This will build the app files in `/build` directory, ready to be dispacted to production

### How to test
Run `mocha` command in CLI, this will run a single scenario test located in `test` folder

### Dependencies
* Node >=8.5.0
* npm 5.2+
* Factorial Go API

### FYI
API Endpoint is hardcoded in `src/App.js` Line 5, I don't think this is a good practice, do you ?

### Test in Pipeline
Run the tests using Mocha on the pre-build process, refer to "How to test section" above

### Deployment
This application needs to be deployed behind a web server, preferred Nginx
