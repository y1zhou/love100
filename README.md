# love100
A simple website powered by [Vue.js](https://vuejs.org/) and [Gin](https://github.com/gin-gonic/gin#model-binding-and-validation).

## Deployment
### Prerequisites
To run the website you'll need `nginx` and `MySQL`, both of which have countless tutorials for setting up. You also need a `config.json` file with the following keys:

```json
{
  "BackendPort": "9423",
  "MysqlDatabase": "love100",
  "MysqlUser": "username",
  "MysqlPassword": "UsuallyMoreComplicated"
}
```

The json file should be in the same directory with the built `backend` executable file.

To build from source, you'll also need an environment with `Vue.js` and `golang`, and of course the code here:

```bash
git clone https://github.com/y1zhou/love100
```

### Building the Backend
```bash
cd /path/to/love100/backend
go build -o love100_backend

./love100_backend
```

Note that there should be a `config.json` file in the same directory.

### Building the Frontend

```bash
cd /path/to/love100
yarn build
```

Use `nginx` to serve the generated `dist/` folder.

## Todo
- Better handling for config files
- Dockerize the entire application
