# Cardiology Test Repo
## Getting Started
To run this project, you will need:
* Golang (1.17 or newer) (Install from https://go.dev/dl/go1.19.5.windows-amd64.msi)
* NodeJS (Anything past V14 works) (Install from https://nodejs.org/dist/v18.13.0/node-v18.13.0-x64.msi)

> If you are not running Windows, let me know and I'll help you set it up - Chase

Then, once those are installed you will need install dependencies and run the process for both servers **in parallel**. (Don't kill either process, they communicate with each other)
## Running the Golang Backend
With Golang fully installed, run the following from the root directory:
```bash
# Enter the backend folder
cd backend

# Install dependencies
go install

# Run the backend
go run server.go
```

## Running the NextJS frontend
With NodeJS fully installed, run the following from the root directory in a new terminal:
```bash
# Enter the frontend folder
cd frontend

# Install dependencies
npm install

# Run the server
npm run dev
```

## Enjoy!
Now you should be able to go to `http://localhost:3000` and will see the project running!