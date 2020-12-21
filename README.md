# numbarb_backend
Build docker image:
```docker build -t numbarb_backend .```
Then run:
```docker run -p 3000:3000 numbarb_backend```
And check in browser: http://localhost:3000/convert?number=1234